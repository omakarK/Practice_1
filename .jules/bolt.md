# Bolt's Performance Journal

## 2026-08-31 - Read-Lock Optimization for HTTP Handler in Go
**Learning:** Upgrading `sync.Mutex` to `sync.RWMutex` with `RLock` and moving non-thread-state operations (such as HTTP header setup) outside the lock critical section reduces lock contention and improves parallel read throughput significantly for Go HTTP handlers.
**Action:** When handling read-dominated shared resources in Go, always prefer `sync.RWMutex` with `RLock()` and keep response header setup outside the lock block.
