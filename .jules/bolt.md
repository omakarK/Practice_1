## 2026-05-05 - Optimize concurrent read performance with sync.RWMutex
**Learning:** In Go web servers with a shared data structure, using `sync.Mutex` for both reads and writes can lead to significant contention under concurrent load. `sync.RWMutex` allows multiple concurrent readers, which drastically improves parallel read latency when the workload is read-heavy.
**Action:** Use `sync.RWMutex` instead of `sync.Mutex` for data structures that are frequently read but infrequently modified.
