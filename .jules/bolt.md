## 2025-05-15 - Using sync.RWMutex for concurrent read access
**Learning:** In Go, replacing `sync.Mutex` with `sync.RWMutex` can significantly improve throughput for read-heavy workloads (like a GET API) by allowing multiple goroutines to hold a read lock simultaneously. Benchmarks showed a reduction from ~1688 ns/op to ~884 ns/op (approx. 48% improvement) when accessed in parallel.
**Action:** Always evaluate if a shared resource is primarily read-heavy and consider `sync.RWMutex` for concurrent read optimization.
