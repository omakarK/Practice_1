## 2026-08-31 - Read concurrency with sync.RWMutex and minimal lock scope

**Learning:** Using `sync.Mutex` on read-dominated HTTP handlers causes contention under parallel requests. Switching to `sync.RWMutex` with `RLock` and moving non-critical setup (e.g., `w.Header().Set`) outside the lock reduces lock contention significantly in Go HTTP handlers.
**Action:** Use `sync.RWMutex` with `RLock()`/`RUnlock()` for read-heavy handlers and keep the locked critical section as minimal as possible.
