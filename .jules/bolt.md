## 2026-08-31 - Read Lock Optimization for HTTP Get Tasks Handler
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` and using `RLock()` for read-heavy HTTP handlers significantly improves throughput under concurrent reads by preventing read contention. Moving non-critical work like setting response headers outside the lock further reduces lock hold times.
**Action:** Use `sync.RWMutex` for read operations and keep critical sections as small as possible.
