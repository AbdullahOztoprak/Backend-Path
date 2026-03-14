//go:build legacy

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

func TestValidateTransaction(t *testing.T) {
	tests := []struct {
		name     string
		tx       entity.Transaction
		expected bool
	}{
		{
			name: "valid transaction",
			tx: entity.Transaction{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     100.50,
				Description: "Payment for services",
			},
			expected: true,
		},
		{
			name: "invalid transaction - negative amount",
			tx: entity.Transaction{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     -50.00,
				Description: "Payment for services",
			},
			expected: false,
		},
		{
			name: "invalid transaction - zero amount",
			tx: entity.Transaction{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     0.00,
				Description: "Payment for services",
			},
			expected: false,
		},
		{
			name: "invalid transaction - missing from user",
			tx: entity.Transaction{
				ToUserID:   2,
				Amount:     100.50,
				Description: "Payment for services",
			},
			expected: false,
		},
		{
			name: "invalid transaction - missing to user",
			tx: entity.Transaction{
				FromUserID: 1,
				Amount:     100.50,
				Description: "Payment for services",
			},
			expected: false,
		},
		{
			name: "invalid transaction - missing description",
			tx: entity.Transaction{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     100.50,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateTransaction(tt.tx)
			assert.Equal(t, tt.expected, result)
		})
	}
}