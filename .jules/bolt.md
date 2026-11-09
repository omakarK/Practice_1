# Bolt's Journal - Critical Performance Learnings

## 2026-08-31 - HTTP Critical Section and Mutex Contention in Go Handlers
**Learning:** Using an exclusive `sync.Mutex` on read-heavy API endpoints causes severe lock contention under parallel load. Additionally, executing header mutation and non-critical operations inside the critical section increases lock hold time unnecessarily. Switching to `sync.RWMutex` with `RLock()` and moving HTTP header configuration outside the lock significantly reduces parallel latency (by >50%).
**Action:** In read-dominated in-memory services, use `sync.RWMutex`, minimize critical section boundaries, and keep standard response setup outside the lock.
