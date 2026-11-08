# Bolt's Performance Journal

Critical learnings only. Format:
`## YYYY-MM-DD - [Title]`
`**Learning:** [Insight]`
`**Action:** [How to apply next time]`

## 2026-08-31 - Concurrency Contention in In-Memory HTTP Handlers
**Learning:** Using `sync.Mutex` across high-frequency HTTP read handlers serializes concurrent client requests even when reads do not mutate state. Moving non-critical operations (like `w.Header().Set(...)`) outside lock sections and using `sync.RWMutex.RLock()` dropped parallel read latency from ~2200 ns/op to ~912 ns/op (~58% speedup).
**Action:** When handling read-heavy in-memory state in Go services, use `sync.RWMutex` with `RLock`/`RUnlock` and ensure response header mutations or encoding preparations happen outside the lock where possible.
