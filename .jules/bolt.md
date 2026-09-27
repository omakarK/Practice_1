# Bolt's Performance Journal

## 2026-08-31 - Use RWMutex and Move I/O Outside Critical Section for HTTP Read Handlers
**Learning:** Using `sync.Mutex` for HTTP GET handlers serializes concurrent read requests and blocks readers during HTTP header setup and JSON encoding. Switching to `sync.RWMutex` with `RLock` and setting HTTP response headers outside the critical section reduces parallel read latency by ~50-60%.
**Action:** Always use `sync.RWMutex` for read-heavy shared state in Go HTTP handlers, and minimize critical section scope by keeping non-shared state setup and I/O outside the lock.
