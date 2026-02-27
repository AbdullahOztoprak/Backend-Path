package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/pkg/idempotency"
)

type TransferFundsUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	idempotency     idempotency.IdempotencyService
}

func NewTransferFundsUseCase(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository, idempotency idempotency.IdempotencyService) *TransferFundsUseCase {
	return &TransferFundsUseCase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		idempotency:     idempotency,
	}
}

type TransferFundsRequest struct {
	FromUserID   int64   `json:"from_user_id"`
	ToUserID     int64   `json:"to_user_id"`
	Amount       float64 `json:"amount"`
	IdempotencyKey string `json:"idempotency_key"`
}

func (uc *TransferFundsUseCase) Execute(ctx context.Context, request TransferFundsRequest) error {
	if err := uc.idempotency.Check(request.IdempotencyKey); err != nil {
		return fmt.Errorf("idempotency check failed: %w", err)
	}

	fromUser, err := uc.userRepo.FindByID(ctx, request.FromUserID)
	if err != nil {
		return fmt.Errorf("failed to find from user: %w", err)
	}

	toUser, err := uc.userRepo.FindByID(ctx, request.ToUserID)
	if err != nil {
		return fmt.Errorf("failed to find to user: %w", err)
	}

	if fromUser.Balance < request.Amount {
		return errors.New("insufficient funds")
	}

	transaction := &entity.Transaction{
		FromUserID: request.FromUserID,
		ToUserID:   request.ToUserID,
		Amount:     request.Amount,
		Timestamp:  time.Now(),
	}

	if err := uc.transactionRepo.Create(ctx, transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	fromUser.Balance -= request.Amount
	toUser.Balance += request.Amount

	if err := uc.userRepo.Update(ctx, fromUser); err != nil {
		return fmt.Errorf("failed to update from user balance: %w", err)
	}

	if err := uc.userRepo.Update(ctx, toUser); err != nil {
		return fmt.Errorf("failed to update to user balance: %w", err)
	}

	if err := uc.idempotency.MarkAsProcessed(request.IdempotencyKey); err != nil {
		return fmt.Errorf("failed to mark idempotency key as processed: %w", err)
	}

	return nil
}