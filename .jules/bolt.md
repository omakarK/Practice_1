# Bolt's Performance Journal

## 2026-08-31 - Reduce lock contention on GET handler with RWMutex
**Learning:** Using `sync.Mutex` on read-heavy HTTP handlers blocks concurrent GET requests unnecessarily. Switching to `sync.RWMutex` with `RLock()` and moving static response header configuration outside the critical section allows parallel read handling and reduces parallel request latency significantly.
**Action:** When handling read-dominant shared state in Go HTTP APIs, prefer `sync.RWMutex` with `RLock` over `sync.Mutex` and keep non-state header initialization outside the lock.
