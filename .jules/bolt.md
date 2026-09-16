## 2026-08-31 - Use sync.RWMutex for Read-Heavy HTTP Handlers
**Learning:** In Go HTTP applications, using exclusive `sync.Mutex` for read operations creates thread contention under concurrent requests. Switching to `sync.RWMutex` with `RLock()` and moving non-critical operations (like header configuration) outside the lock section reduced latency by over 50%.
**Action:** Always check if global/shared state mutexes in read-heavy HTTP endpoints can use `sync.RWMutex` and minimize the critical section.
