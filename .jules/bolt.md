## 2025-05-15 - Concurrent Read Optimization with RWMutex
**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` for read-heavy resources in Go significantly improves performance under concurrent load. In `go-dynamic-project`, parallel read latency for the `/api/tasks` endpoint dropped from ~2025 ns/op to ~984.5 ns/op (a ~51% improvement).
**Action:** Always prefer `sync.RWMutex` over `sync.Mutex` when the protected resource is frequently read and only occasionally modified, ensuring `RLock()` is used for read paths.
