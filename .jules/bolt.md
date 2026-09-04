# Bolt Journal - Critical Learnings

## 2026-08-31 - Use RWMutex and narrow critical sections for read-heavy HTTP handlers
**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` (`RLock`) for read-heavy HTTP endpoints and keeping HTTP response header setting outside the lock dramatically reduces lock contention under parallel loads.
**Action:** Always check if global/shared state access in HTTP handlers is read-dominated and reduce lock scope by moving non-state operations (like header setting) outside critical sections.
