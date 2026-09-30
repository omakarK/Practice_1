# Bolt's Performance Journal - Critical Learnings

## 2026-08-31 - [Go Mutex Contention in Parallel Read Handlers]
**Learning:** Using an exclusive `sync.Mutex` in Go HTTP read handlers creates heavy lock contention under parallel read traffic. Swapping to `sync.RWMutex` with `RLock()` and moving header setting operations outside the critical section nearly cuts request latency in half.
**Action:** Always check if mutexes protecting state in Go HTTP GET handlers can be replaced with `sync.RWMutex` and minimize critical section duration.
