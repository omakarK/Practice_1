## 2026-08-31 - HTTP GET concurrency optimization with sync.RWMutex
**Learning:** Exclusive `sync.Mutex` locks on read-heavy HTTP handlers cause unnecessary thread contention under parallel load. Switching to `sync.RWMutex` with `RLock` and moving non-critical header logic outside the lock reduces lock wait times for concurrent GET requests.
**Action:** Always check if global shared mutexes in read-heavy HTTP endpoints can use `sync.RWMutex` with `RLock`.
