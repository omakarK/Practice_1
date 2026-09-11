## 2026-08-31 - Use sync.RWMutex and Move Non-Critical Work Outside Locks in Read-Heavy HTTP Handlers

**Learning:** In Go HTTP handlers serving read-heavy workloads (such as fetching task lists), using exclusive `sync.Mutex` creates severe lock contention across concurrent requests. Switching to `sync.RWMutex` with `RLock()` and moving response header configuration outside the lock section reduces parallel read latency by ~45% (from ~2058 ns/op to ~1118 ns/op).

**Action:** Identify shared read state in Go web services, replace `sync.Mutex` with `sync.RWMutex` where reads dominate writes, and ensure non-state-dependent setup (e.g. `w.Header().Set(...)`) happens outside the mutex critical section.
