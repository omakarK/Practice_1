# Bolt's Journal - Critical Learnings

## 2026-08-31 - Optimize `getTasksHandler` in Go backend
**Learning:** In read-heavy HTTP endpoints with concurrent traffic, using `sync.RWMutex` instead of `sync.Mutex` and moving HTTP header configuration outside of the critical section reduces mutex contention and latency from ~1885 ns/op to ~840-1000 ns/op (~50% improvement).
**Action:** Use `RWMutex` with `RLock`/`RUnlock` for read-only slice reads and minimize critical section scope.
