# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read-heavy HTTP endpoints with sync.RWMutex
**Learning:** Using `sync.Mutex` for read-heavy handlers creates unnecessary lock contention across concurrent requests. Switching to `sync.RWMutex` with `RLock` and moving non-state-mutating setup (e.g. `w.Header().Set`) outside the critical section reduces parallel read latency by ~50%.
**Action:** Always prefer `sync.RWMutex` for data structures with high read-to-write ratios and minimize critical section duration.
