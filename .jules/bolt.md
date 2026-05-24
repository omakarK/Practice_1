## 2025-03-17 - Read-Write Lock Optimization in Go
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` for read-heavy endpoints significantly improves throughput in concurrent environments by allowing multiple simultaneous readers. Moving non-shared state modifications (like setting response headers) outside the critical section further reduces lock contention.
**Action:** Always check if a protected resource is predominantly read, and prefer `RWMutex` with `RLock` for those paths. Minimize the duration of the lock by moving independent setup code outside.
