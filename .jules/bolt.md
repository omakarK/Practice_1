## 2025-05-22 - Use RWMutex for Concurrent Reads
**Learning:** In read-heavy Go applications, `sync.RWMutex` significantly improves throughput over `sync.Mutex` by allowing multiple concurrent readers. In this task tracker, the `GET /api/tasks` endpoint was previously bottlenecked by a standard mutex, even though it only performs read operations. Switching to `RLock` reduced the benchmark time per operation by approximately 50%.
**Action:** Always prefer `sync.RWMutex` for shared resources where read operations are significantly more frequent than write operations.
