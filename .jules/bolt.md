# Bolt's Journal - Critical Learnings

This journal records critical performance learnings specific to this codebase.

## 2026-08-31 - RWMutex for Concurrent HTTP Reads
**Learning:** Using `sync.Mutex` in `getTasksHandler` caused concurrent GET requests to block each other. Switching to `sync.RWMutex` with `RLock()` and moving HTTP response header configuration outside the lock reduced parallel read latency from ~1910 ns/op to ~909 ns/op (~52% speedup).
**Action:** Always prefer `sync.RWMutex` for read-heavy handlers and minimize operations performed inside locked sections.
