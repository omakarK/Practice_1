# Bolt's Journal - Critical Learnings

## 2026-08-31 - Use RWMutex RLock for Concurrent Read Handlers
**Learning:** In Go HTTP services where read requests (GET) heavily outnumber writes (POST), replacing `sync.Mutex` with `sync.RWMutex` and using `RLock()` allows concurrent read handlers to run in parallel without contention. Moving non-critical operations like setting HTTP response headers outside the lock further reduces critical section hold time.
**Action:** Use `sync.RWMutex` with `RLock()` for read operations in Go in-memory state handlers and keep critical sections as small as possible.
