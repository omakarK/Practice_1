# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock Optimization for HTTP Handlers
**Learning:** Using `sync.RWMutex` with `RLock()` and moving HTTP header setups outside critical sections significantly reduces parallel read latency in Go web handlers (~2470 ns/op down to ~988 ns/op).
**Action:** When handling read-heavy slice or map access in HTTP endpoints, upgrade `sync.Mutex` to `sync.RWMutex` and keep non-shared operations (like header modifications) out of lock boundaries.
