## 2026-05-13 - Optimized concurrent reads with RWMutex
**Learning:** In read-heavy Go web services, switching from `sync.Mutex` to `sync.RWMutex` provides a significant performance boost for parallel requests. In this case, it improved parallel read latency from ~1985 ns/op to ~1088 ns/op (approx. 45% improvement).
**Action:** Always check if a shared resource is mostly read, and consider `sync.RWMutex` for better concurrency.
