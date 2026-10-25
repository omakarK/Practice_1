# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read vs Write Locks in In-Memory Go Handlers
**Learning:** Using `sync.Mutex` on read-heavy HTTP endpoints like `getTasksHandler` creates unnecessary lock contention under concurrent load. Switching to `sync.RWMutex` with `RLock()` and keeping header modifications outside the locked section significantly reduces read latency without sacrificing safety.
**Action:** Always check if shared state in Go HTTP handlers is read-dominated and use `sync.RWMutex` while minimizing critical section scope.
