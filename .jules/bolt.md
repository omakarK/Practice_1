## 2025-05-15 - RWMutex for Read-Heavy Endpoints
**Learning:** In Go web services with shared state, using `sync.RWMutex` instead of `sync.Mutex` for read-heavy endpoints significantly improves throughput by allowing concurrent reads. Minimizing the critical section by moving non-shared resource operations (like setting HTTP headers) outside the lock further reduces contention.
**Action:** Always prefer `sync.RWMutex` when a shared resource is read more often than written, and audit critical sections to ensure only essential operations are performed while holding the lock.
