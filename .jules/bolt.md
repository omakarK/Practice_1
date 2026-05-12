## 2026-05-12 - Use sync.RWMutex for concurrent reads
**Learning:** In Go, replacing sync.Mutex with sync.RWMutex for read-heavy workloads can significantly reduce lock contention and improve parallel read performance. Parallel benchmarks (b.RunParallel) are essential to observe this benefit, as they simulate concurrent access.
**Action:** Always evaluate if a protected resource is read more often than written, and use sync.RWMutex if it is. Verify the improvement using parallel benchmarks.
