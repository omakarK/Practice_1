# Bolt's Performance Journal

## 2026-08-31 - Read Lock Bottleneck in Go HTTP Handlers
**Learning:** Using an exclusive `sync.Mutex` in read-heavy HTTP handlers serializes requests and causes severe thread contention under parallel loads. Moving non-critical state setup (like setting HTTP headers) outside the mutex and switching to `sync.RWMutex` with `RLock` allows concurrent read execution and significantly reduces latency.
**Action:** Always prefer `sync.RWMutex` for read-heavy state and keep header writes and unnecessary logic out of critical sections.
