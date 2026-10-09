# Bolt's Performance Journal

## 2026-08-31 - Use RWMutex and narrow lock scope for concurrent HTTP read handlers in Go
**Learning:** For read-heavy Go HTTP handlers protected by a mutex, `sync.RWMutex` with `RLock()`/`RUnlock()` combined with moving static/header configurations outside the critical section reduces lock contention and significantly drops parallel request latency.
**Action:** Use `sync.RWMutex` for shared in-memory read state in Go web handlers and keep HTTP header mutations outside the mutex.
