# Bolt's Journal - Critical Learnings

## 2026-08-31 - HTTP Handlers Concurrent Read Performance
**Learning:** Using `sync.Mutex` in read-heavy HTTP handlers serializes requests unnecessarily. Switching to `sync.RWMutex` with `RLock`/`RUnlock` allows concurrent reads across multiple goroutines, and keeping non-state operations (like setting HTTP headers) outside the critical lock section reduces lock contention significantly (~50%+ speedup in parallel reads).
**Action:** Always check if locks in HTTP read handlers can use `sync.RWMutex` and minimize the scope of the critical section.
