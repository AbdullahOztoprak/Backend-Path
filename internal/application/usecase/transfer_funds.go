package usecase

import (
	"context"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type IdempotencyChecker interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type TransferFundsUseCase struct {
	transactionRepo repository.TransactionRepository
	idempotency     IdempotencyChecker
}

func NewTransferFundsUseCase(transactionRepo repository.TransactionRepository, idempotency IdempotencyChecker) *TransferFundsUseCase {
	return &TransferFundsUseCase{
		transactionRepo: transactionRepo,
		idempotency:     idempotency,
	}
}

type TransferFundsRequest struct {
	FromUserID     string  `json:"from_user_id"`
	ToUserID       string  `json:"to_user_id"`
	Amount         float64 `json:"amount"`
	IdempotencyKey string `json:"idempotency_key"`
	Description    string  `json:"description"`
}

func (uc *TransferFundsUseCase) Execute(ctx context.Context, request TransferFundsRequest) error {
	if request.Amount <= 0 {
		return entity.ErrInvalidTransactionAmount
	}

	if request.IdempotencyKey != "" && uc.idempotency != nil {
		existing, err := uc.idempotency.Get(request.IdempotencyKey)
		if err != nil {
			return err
		}
		if existing != "" {
			return nil
		}
	}

	transaction, err := entity.NewTransaction(request.FromUserID, request.ToUserID, request.Amount, request.Description, request.IdempotencyKey)
	if err != nil {
		return err
	}

	if err := uc.transactionRepo.Create(ctx, transaction); err != nil {
		return err
	}

	if request.IdempotencyKey != "" && uc.idempotency != nil {
		if err := uc.idempotency.Set(request.IdempotencyKey, transaction.ID); err != nil {
			return err
		}
	}

	return nil
}