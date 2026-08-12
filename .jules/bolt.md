# Bolt's Performance Journal

This journal contains critical learnings from Bolt's performance optimization sessions.

## 2026-08-11 - Upgrade Mutex to RWMutex in Go HTTP Handlers
**Learning:** In Go, replacing sync.Mutex with sync.RWMutex for read-heavy handlers (like `getTasksHandler`) and moving header configurations outside the lock reduces lock contention significantly in parallel read benchmarks.
**Action:** Always upgrade read-only or read-heavy locks to sync.RWMutex and minimize the critical section by performing setup (like headers) outside the locked block.
