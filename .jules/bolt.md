## 2026-05-28 - Optimize getTasksHandler with RWMutex
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` significantly improves parallel read performance in Go by allowing multiple readers to access shared data concurrently. Reducing the duration of the critical section by moving non-critical operations (like setting HTTP headers) outside the lock further reduces contention.
**Action:** Always check if a shared resource is read-heavy and consider `sync.RWMutex`. Keep critical sections as small as possible.
