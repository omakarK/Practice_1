# Bolt's Performance Journal

## 2026-08-31 - Use sync.RWMutex for concurrent task reads
**Learning:** `getTasksHandler` only performs read operations on `tasks`. Using `sync.Mutex` causes unnecessary write-locks and contention on high-concurrency read requests. Switching `mu` to `sync.RWMutex` and moving HTTP header settings outside the critical section reduced parallel read latency from ~1904 ns/op to ~841 ns/op (~55% speedup).
**Action:** Use `sync.RWMutex` with `RLock`/`RUnlock` for read-heavy shared slice access in Go HTTP handlers, and keep non-thread-sensitive operations like header setup outside the lock scope.
