package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/test/mocks"
)

func TestTransferFunds(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockUserRepo := new(mocks.UserRepository)
	transactionService := service.NewTransactionService(mockTransactionRepo, mockUserRepo)

	tests := []struct {
		name          string
		fromUserID   int
		toUserID     int
		amount        float64
		expectedError error
	}{
		{
			name:        "Successful transfer",
			fromUserID: 1,
			toUserID:   2,
			amount:      100.00,
			expectedError: nil,
		},
		{
			name:        "Insufficient funds",
			fromUserID: 1,
			toUserID:   2,
			amount:      1000.00,
			expectedError: entity.ErrInsufficientFunds,
		},
		{
			name:        "Invalid user ID",
			fromUserID: 0,
			toUserID:   2,
			amount:      100.00,
			expectedError: entity.ErrInvalidUserID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError == nil {
				mockUserRepo.On("GetBalance", tt.fromUserID).Return(200.00, nil)
				mockTransactionRepo.On("CreateTransaction", mock.AnythingOfType("*entity.Transaction")).Return(nil)
			} else if tt.expectedError == entity.ErrInsufficientFunds {
				mockUserRepo.On("GetBalance", tt.fromUserID).Return(50.00, nil)
			} else if tt.expectedError == entity.ErrInvalidUserID {
				// No need to set up mocks for invalid user ID
			}

			err := transactionService.TransferFunds(tt.fromUserID, tt.toUserID, tt.amount)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			mockUserRepo.AssertExpectations(t)
			mockTransactionRepo.AssertExpectations(t)
		})
	}
}