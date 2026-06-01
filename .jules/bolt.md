## 2026-06-01 - Optimizing getTasksHandler with RWMutex
**Learning:** In read-intensive Go web services, switching from `sync.Mutex` to `sync.RWMutex` for GET handlers can lead to significant performance improvements (over 50% in parallel benchmarks). Additionally, moving non-shared resource operations like `w.Header().Set` outside the lock further reduces contention.
**Action:** Always check if a handler is read-only and can benefit from `RLock`. Ensure critical sections are as small as possible by moving header settings and body decoding outside the lock.
