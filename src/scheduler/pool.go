package scheduler

import (
  "runtime"
  "sync"
)

type WorkerPool struct {
  jobs      chan func()
  wg        sync.WaitGroup
  numWorker int
}

func NewPool(maxWorkers int) *WorkerPool {
  if maxWorkers <= 0 {
    maxWorkers = runtime.NumCPU()
  }
  pool := &WorkerPool{
    jobs:      make(chan func(), maxWorkers*4),
    numWorker: maxWorkers,
  }
  for i := 0; i < maxWorkers; i++ {
    go func() {
      for job := range pool.jobs {
        job()
        pool.wg.Done()
      }
    }()
  }
  return pool
}

func (p *WorkerPool) Submit(job func()) {
  p.wg.Add(1)
  p.jobs <- job
}

func (p *WorkerPool) Shutdown() {
  p.wg.Wait()
  close(p.jobs)
} 