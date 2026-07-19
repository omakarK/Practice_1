# Bolt Journal

## 2025-05-22 - Upgrading to RWMutex and Minimizing Lock Duration in Go
**Learning:** In read-heavy Go HTTP services, a global `sync.Mutex` serializes all reads, resulting in unnecessary lock contention and latency under parallel workloads. Upgrading to `sync.RWMutex` with `RLock`/`RUnlock` enables safe concurrent reads. Additionally, operations like setting HTTP headers do not touch the shared state, so moving them outside of the critical section minimizes lock duration even further.
**Action:** When optimizing handlers, check if a global lock can be split into read/write locks, and ensure all non-shared operations (like header initialization or pre-processing request/response data) are performed outside the locked block.
