package fixtures

import (
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

var TransactionFixtures = []entity.Transaction{
	{
		ID:          "1",
		FromUserID: "1",
		ToUserID:   "2",
		Amount:     100.00,
		Description: "Payment for services",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
	{
		ID:          "2",
		FromUserID: "2",
		ToUserID:   "3",
		Amount:     50.50,
		Description: "Refund for overpayment",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
	{
		ID:          "3",
		FromUserID: "1",
		ToUserID:   "3",
		Amount:     75.25,
		Description: "Payment for goods",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}