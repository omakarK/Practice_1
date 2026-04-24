## 2025-05-15 - RWMutex for Task List
**Learning:** In a read-heavy application, using `sync.Mutex` for all operations can create a bottleneck. Switching to `sync.RWMutex` allows multiple concurrent readers, significantly improving throughput for read endpoints. In this case, `BenchmarkGetTasks` saw a performance improvement from ~2023 ns/op to ~1049 ns/op when run in parallel.
**Action:** Always evaluate if a resource is read-heavy and consider `sync.RWMutex` to improve concurrency.
