# Bolt's Journal

## 2026-08-31 - Read vs Write Locks in Go In-Memory Store
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` with `RLock` for read-only endpoints (like `getTasksHandler`), and moving HTTP header configuration outside of the critical section, drastically reduces parallel read lock contention and latency.
**Action:** Always audit lock scope and use RWMutex for read-heavy handlers.
