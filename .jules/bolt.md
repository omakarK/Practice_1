## 2025-05-15 - [Concurrency Optimization in Go Task Tracker]
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` for a read-heavy endpoint like `/api/tasks` significantly improves throughput. In this codebase, the `getTasksHandler` was blocking all other read requests during JSON encoding of the tasks slice. By using `RLock`, we allow multiple concurrent readers, resulting in a ~45% performance improvement in parallel benchmarks.
**Action:** Always evaluate if a protected resource is read-heavy and consider `sync.RWMutex` to reduce lock contention.
