## 2025-05-18 - Optimized concurrent reads in Go
**Learning:** In read-heavy endpoints, switching from `sync.Mutex` to `sync.RWMutex` significantly improves performance under concurrent load by allowing multiple readers. Additionally, minimizing the critical section by moving non-shared state operations (like setting headers) outside the lock further reduces contention.
**Action:** Always prefer `sync.RWMutex` for data structures that are frequently read but occasionally written, and keep the locked section as small as possible.
