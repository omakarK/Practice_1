# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock Optimization in Concurrent Go Handlers
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` with `RLock` for read-heavy endpoint handlers (`getTasksHandler`), along with moving non-critical operations like HTTP header creation outside of lock critical sections, reduces parallel read latency by ~47% (from ~1850 ns/op down to ~970 ns/op) under high parallel load.
**Action:** Always prefer `sync.RWMutex` with `RLock`/`RUnlock` for read-dominated shared resources in Go web servers, and keep the critical section strictly minimal by moving header initialization or response preparation outside the lock.
