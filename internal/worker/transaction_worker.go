package worker

import (
    "context"
    "log"
    "time"

    "github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
    "github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
    "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/observability"
)

type TransactionWorker struct {
    transactionService service.TransactionService
    logger             observability.Logger
}

func NewTransactionWorker(transactionService service.TransactionService, logger observability.Logger) *TransactionWorker {
    return &TransactionWorker{
        transactionService: transactionService,
        logger:             logger,
    }
}

func (tw *TransactionWorker) ProcessTransaction(ctx context.Context, transaction entity.Transaction) error {
    tw.logger.Info("Processing transaction", "transaction_id", transaction.ID)

    err := tw.transactionService.ExecuteTransaction(ctx, transaction)
    if err != nil {
        tw.logger.Error("Failed to process transaction", "transaction_id", transaction.ID, "error", err)
        return err
    }

    tw.logger.Info("Transaction processed successfully", "transaction_id", transaction.ID)
    return nil
}

func (tw *TransactionWorker) StartWorker(ctx context.Context, transactionChannel <-chan entity.Transaction) {
    for {
        select {
        case transaction := <-transactionChannel:
            if err := tw.ProcessTransaction(ctx, transaction); err != nil {
                tw.logger.Error("Error processing transaction", "transaction_id", transaction.ID, "error", err)
            }
        case <-ctx.Done():
            tw.logger.Info("Transaction worker shutting down")
            return
        }
    }
}

func (tw *TransactionWorker) RetryFailedTransaction(ctx context.Context, transaction entity.Transaction, retryCount int) {
    for i := 0; i < retryCount; i++ {
        err := tw.ProcessTransaction(ctx, transaction)
        if err == nil {
            return
        }
        tw.logger.Warn("Retrying transaction", "transaction_id", transaction.ID, "attempt", i+1)
        time.Sleep(2 * time.Second) // Exponential backoff can be implemented here
    }
    tw.logger.Error("Max retries reached for transaction", "transaction_id", transaction.ID)
}