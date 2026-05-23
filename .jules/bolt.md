## 2025-05-23 - Optimizing Parallel Read Latency in Go

**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` for read-heavy endpoints significantly reduces contention in parallel execution. Moving non-critical operations (like setting HTTP headers) outside the mutex lock further minimizes the time spent in the critical section.

**Action:** Always prefer `sync.RWMutex` when an endpoint primarily performs read operations and identify work that can be moved out of the lock.
