## 2025-05-15 - [Concurrency] Use RWMutex for task list reads
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` significantly improves parallel read performance for shared data structures in Go, especially when the workload is read-heavy.
**Action:** Always evaluate if a shared resource is read more often than it's written and use `sync.RWMutex` if appropriate to reduce lock contention.
