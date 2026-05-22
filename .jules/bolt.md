## 2025-05-15 - Optimizing concurrent reads in Go web server
**Learning:** In read-heavy endpoints, switching from `sync.Mutex` to `sync.RWMutex` significantly improves concurrent throughput by allowing multiple readers to hold the lock simultaneously. Additionally, moving non-shared resource modifications (like setting response headers) outside the critical section reduces lock contention.
**Action:** Always check if a lock is protecting read-only operations and prefer `RWMutex`. Minimize the scope of the lock to only the necessary data access.
