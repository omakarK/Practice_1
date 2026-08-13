# Bolt's Performance Journal

This performance journal is maintained by Bolt to track critical learnings, unexpected results, and codebase-specific performance patterns/anti-patterns.

## Critical Learnings

## 2026-08-14 - Optimized `getTasksHandler` with `sync.RWMutex`
**Learning:** In a highly parallel read-heavy Go HTTP API, using a standard `sync.Mutex` and setting HTTP header parameters within the critical section causes lock contention. Upgrading to `sync.RWMutex` with `RLock`/`RUnlock`, and moving the header modification (`w.Header().Set`) out of the lock significantly reduces parallel latency.
**Action:** Always favor `sync.RWMutex` for endpoints that only perform reads, and minimize the critical section by moving any header settings or preparation code outside the lock block.
