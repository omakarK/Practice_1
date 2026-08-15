# Bolt's Journal - Critical Learnings

This journal contains critical learnings, architectural bottlenecks, and performance insights discovered in this codebase.

## 2026-08-16 - RWMutex and Scope Reduction for HTTP Handlers
**Learning:** In Go HTTP read handlers holding a full `sync.Mutex`, read operations contend on the lock needlessly. Furthermore, performing response header setup (`w.Header().Set(...)`) inside the lock extends lock contention duration. Upgrading to `sync.RWMutex` with `RLock()` and setting HTTP headers outside the lock allows parallel read execution and minimizes critical section duration.
**Action:** Always check if locks in HTTP read handlers can be converted to `sync.RWMutex` and verify that non-shared operations (like header setup or JSON encoding outside shared slice mutability) are kept outside the lock.
