## 2025-05-15 - Optimizing `getTasksHandler` with `sync.RWMutex`
**Learning:** In a read-heavy Go web service, using `sync.Mutex` for simple data retrieval creates unnecessary contention among concurrent requests. Moving side effects like setting HTTP headers outside the critical section further reduces lock hold time.
**Action:** Always prefer `sync.RWMutex` for read-intensive shared state and keep critical sections as small as possible by moving non-shared-state operations (like header manipulation) outside the lock.
