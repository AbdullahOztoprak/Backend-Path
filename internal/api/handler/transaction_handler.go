package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-backend-project/internal/application/usecase"
    "go-backend-project/internal/domain/entity"
    "go-backend-project/pkg/apperror"
)

type TransactionHandler struct {
    transferFundsUseCase usecase.TransferFunds
    listTransactionsUseCase usecase.ListTransactions
}

func NewTransactionHandler(transferFundsUseCase usecase.TransferFunds, listTransactionsUseCase usecase.ListTransactions) *TransactionHandler {
    return &TransactionHandler{
        transferFundsUseCase: transferFundsUseCase,
        listTransactionsUseCase: listTransactionsUseCase,
    }
}

func (h *TransactionHandler) TransferFunds(c *gin.Context) {
    var request entity.Transaction
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, apperror.NewBadRequestError("Invalid request payload"))
        return
    }

    err := h.transferFundsUseCase.Execute(request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, apperror.NewInternalServerError("Failed to transfer funds"))
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Funds transferred successfully"})
}

func (h *TransactionHandler) ListTransactions(c *gin.Context) {
    transactions, err := h.listTransactionsUseCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, apperror.NewInternalServerError("Failed to retrieve transactions"))
        return
    }

    c.JSON(http.StatusOK, transactions)
}