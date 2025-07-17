package scheduler

// RunLexParse runs lexing and parsing jobs in parallel, respecting dependencies.
func RunLexParse(files []string, graph map[string][]string, pool *WorkerPool) map[string]interface{} {
  return map[string]interface{}{}
}

// RunCompile runs compile jobs in parallel, respecting dependencies.
func RunCompile(files []string, graph map[string][]string, pool *WorkerPool) map[string]interface{} {
  return map[string]interface{}{}
}

// RunBatch runs a batch of jobs in parallel using the worker pool.
func RunBatch(jobs []func(), pool *WorkerPool) {
}

func RunBatches(jobs map[string]func(), graph map[string][]string, pool *WorkerPool) {
  completed := make(map[string]bool)
  scheduled := make(map[string]bool)
  total := len(jobs)
  for len(completed) < total {
    batch := NextBatch(graph, completed, scheduled)
    if len(batch) == 0 {
      break // cycle or error
    }
    done := make(chan string, len(batch))
    for _, file := range batch {
      scheduled[file] = true
      f := file
      pool.Submit(func() {
        jobs[f]()
        done <- f
      })
    }
    for i := 0; i < len(batch); i++ {
      completed[<-done] = true
    }
    close(done)
  }
  pool.Shutdown()
} 