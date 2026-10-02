## 2026-08-31 - Use sync.RWMutex for HTTP read handlers in Go

**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` and `RLock` for read-heavy HTTP endpoints (like `getTasksHandler`) allows multiple concurrent read requests to execute without thread contention. Moving non-critical operations (such as `w.Header().Set(...)`) outside the read lock further reduces critical section holding time, improving parallel read latency from ~1817 ns/op to ~923 ns/op (~49% speedup).

**Action:** For read-heavy shared state in Go HTTP handlers, prefer `sync.RWMutex` with `RLock()`, and ensure header manipulations or non-state operations are executed outside the lock.
