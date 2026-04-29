## 2025-05-15 - Optimize concurrent read latency in tasks API
**Learning:** In Go web servers with high read-to-write ratios, using `sync.RWMutex` instead of a standard `sync.Mutex` significantly improves throughput and reduces latency for read-only endpoints by allowing concurrent access.
**Action:** Always evaluate the read/write patterns of shared resources. If reads are frequent and writes are rare, prefer `sync.RWMutex` with `RLock`/`RUnlock` for read operations.
