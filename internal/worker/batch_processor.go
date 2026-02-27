package worker

import (
    "context"
    "log"
    "time"

    "github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
)

type BatchProcessor struct {
    transactionService service.TransactionService
    batchSize          int
    processInterval    time.Duration
}

func NewBatchProcessor(transactionService service.TransactionService, batchSize int, processInterval time.Duration) *BatchProcessor {
    return &BatchProcessor{
        transactionService: transactionService,
        batchSize:          batchSize,
        processInterval:    processInterval,
    }
}

func (bp *BatchProcessor) Start(ctx context.Context) {
    ticker := time.NewTicker(bp.processInterval)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            log.Println("Batch processor stopped")
            return
        case <-ticker.C:
            if err := bp.processBatch(); err != nil {
                log.Printf("Error processing batch: %v", err)
            }
        }
    }
}

func (bp *BatchProcessor) processBatch() error {
    transactions, err := bp.transactionService.FetchPendingTransactions(bp.batchSize)
    if err != nil {
        return err
    }

    for _, transaction := range transactions {
        if err := bp.transactionService.ProcessTransaction(transaction); err != nil {
            log.Printf("Failed to process transaction %d: %v", transaction.ID, err)
            continue
        }
        log.Printf("Successfully processed transaction %d", transaction.ID)
    }

    return nil
}