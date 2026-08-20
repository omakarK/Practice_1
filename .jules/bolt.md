## 2026-08-20 - RWMutex for Concurrent Read HTTP Handlers
**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` and calling `RLock()` in Go read-only HTTP handlers dramatically improves parallel throughput by allowing concurrent readers. Moving non-shared state operations like setting response headers (`w.Header().Set(...)`) outside the lock further reduces critical section contention.
**Action:** Always check if locks in Go HTTP handlers protect read-only operations, and prefer `sync.RWMutex` with `RLock` after hoisting header configuration outside the critical section.
