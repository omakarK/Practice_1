# Bolt's Journal - Critical Learnings

## 2026-08-19 - Concurrent HTTP Reads in Go
**Learning:** Using `sync.Mutex` in read-heavy HTTP handlers serializes requests across concurrent clients, causing unnecessary contention. Upgrading to `sync.RWMutex` with `RLock()` and moving non-thread-state operations (like setting HTTP response headers) outside critical lock blocks significantly reduces parallel latency.
**Action:** Always check if locks on read handlers can use `sync.RWMutex` and minimize critical section scope by keeping HTTP header setup outside the lock.
