package worker

import (
    "sync"
    "sync/atomic"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/AbdullahOztoprak/go-backend-project/internal/service"
    "github.com/rs/zerolog/log"
)

type TransactionWorkerPool struct {
    workerCount    int
    taskQueue      chan *models.Transaction
    quit           chan bool
    wg             sync.WaitGroup
    txService      service.TransactionService
    // Atomic counters for statistics
    processedCount int64
    failedCount    int64
}

func NewTransactionWorkerPool(workerCount int, queueSize int, txService service.TransactionService) *TransactionWorkerPool {
    return &TransactionWorkerPool{
        workerCount: workerCount,
        taskQueue:   make(chan *models.Transaction, queueSize),
        quit:        make(chan bool),
        txService:   txService,
    }
}

func (p *TransactionWorkerPool) Start() {
    for i := 0; i < p.workerCount; i++ {
        p.wg.Add(1)
        go p.worker(i)
    }
    log.Info().Msgf("Started %d transaction workers", p.workerCount)
}

func (p *TransactionWorkerPool) Stop() {
    close(p.quit)
    p.wg.Wait()
    close(p.taskQueue)
    log.Info().Msg("Stopped transaction worker pool")
}

func (p *TransactionWorkerPool) Submit(tx *models.Transaction) {
    select {
    case p.taskQueue <- tx:
        log.Debug().Msgf("Transaction %d queued", tx.ID)
    default:
        log.Warn().Msg("Transaction queue is full")
        atomic.AddInt64(&p.failedCount, 1)
    }
}

func (p *TransactionWorkerPool) worker(id int) {
    defer p.wg.Done()
    
    for {
        select {
        case tx := <-p.taskQueue:
            if tx != nil {
                p.processTransaction(tx, id)
            }
        case <-p.quit:
            return
        }
    }
}

func (p *TransactionWorkerPool) processTransaction(tx *models.Transaction, workerID int) {
    log.Debug().Msgf("Worker %d processing transaction %d", workerID, tx.ID)
    
    // Process the transaction
    err := p.txService.Create(tx)
    if err != nil {
        log.Error().Err(err).Msgf("Worker %d failed to process transaction %d", workerID, tx.ID)
        atomic.AddInt64(&p.failedCount, 1)
    } else {
        atomic.AddInt64(&p.processedCount, 1)
        log.Debug().Msgf("Worker %d successfully processed transaction %d", workerID, tx.ID)
    }
}

// GetStats returns processing statistics
func (p *TransactionWorkerPool) GetStats() (processed, failed int64) {
    return atomic.LoadInt64(&p.processedCount), atomic.LoadInt64(&p.failedCount)
}
