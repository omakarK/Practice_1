## 2026-05-16 - Optimize getTasksHandler concurrency
**Learning:** Using sync.RWMutex instead of sync.Mutex for read-heavy endpoints significantly improves parallel performance. Moving non-critical operations (like setting headers) outside the lock also reduces the time spent in the critical section.
**Action:** Always prefer RWMutex for Go handlers that primarily serve data, and minimize the scope of the lock by moving non-state-dependent logic outside of it.
