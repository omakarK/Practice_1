# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock Concurrency Optimization in Go HTTP Handlers
**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` using `RLock()` for read-only HTTP endpoints (`getTasksHandler`), along with moving header mutations outside the critical section, substantially reduces latency under concurrent workloads by removing write lock contention across parallel requests.
**Action:** Always check if global/shared state in Go HTTP handlers is read-dominated and use `sync.RWMutex` with minimal critical sections for read paths.
