## 2026-08-31 - Use sync.RWMutex and move headers outside critical section for HTTP read handlers

**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` and calling `RLock()`/`RUnlock()` in read-heavy HTTP handlers (e.g. `GET /api/tasks`) allows parallel execution of read requests. Additionally, moving non-shared header setup outside the lock section reduces critical section contention, resulting in ~50% latency reduction under parallel load (~2027 ns/op to ~1011 ns/op).

**Action:** Identify read-heavy HTTP handlers in Go applications using global/shared mutexes, switch shared locks to `sync.RWMutex`, and move request-scoped work outside critical sections.
