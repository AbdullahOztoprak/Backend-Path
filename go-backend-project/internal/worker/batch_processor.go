package worker

import (
    "sync"
    "sync/atomic"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/rs/zerolog/log"
)

type BatchProcessor struct {
    batchSize     int
    items         []interface{}
    mu            sync.RWMutex
    processedBatches int64
    totalItems    int64
}

func NewBatchProcessor(batchSize int) *BatchProcessor {
    return &BatchProcessor{
        batchSize: batchSize,
        items:     make([]interface{}, 0, batchSize),
    }
}

func (bp *BatchProcessor) AddItem(item interface{}) {
    bp.mu.Lock()
    defer bp.mu.Unlock()
    
    bp.items = append(bp.items, item)
    atomic.AddInt64(&bp.totalItems, 1)
    
    if len(bp.items) >= bp.batchSize {
        go bp.processBatch(bp.items)
        bp.items = make([]interface{}, 0, bp.batchSize)
    }
}

func (bp *BatchProcessor) processBatch(batch []interface{}) {
    log.Info().Msgf("Processing batch of %d items", len(batch))
    
    // Process each item in the batch concurrently
    var wg sync.WaitGroup
    for _, item := range batch {
        wg.Add(1)
        go func(item interface{}) {
            defer wg.Done()
            bp.processItem(item)
        }(item)
    }
    
    wg.Wait()
    atomic.AddInt64(&bp.processedBatches, 1)
    log.Info().Msgf("Completed batch processing")
}

func (bp *BatchProcessor) processItem(item interface{}) {
    // Process individual item based on type
    switch v := item.(type) {
    case *models.Transaction:
        log.Debug().Msgf("Processing transaction %d in batch", v.ID)
    case *models.User:
        log.Debug().Msgf("Processing user %s in batch", v.Username)
    default:
        log.Debug().Msg("Processing unknown item type in batch")
    }
}

func (bp *BatchProcessor) GetStats() (batches, items int64) {
    return atomic.LoadInt64(&bp.processedBatches), atomic.LoadInt64(&bp.totalItems)
}

func (bp *BatchProcessor) Flush() {
    bp.mu.Lock()
    defer bp.mu.Unlock()
    
    if len(bp.items) > 0 {
        go bp.processBatch(bp.items)
        bp.items = make([]interface{}, 0, bp.batchSize)
    }
}
