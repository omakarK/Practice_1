# Bolt Journal

## 2025-07-22 - Upgrading sync.Mutex to sync.RWMutex and minimizing critical sections in read-heavy handlers
**Learning:** In read-heavy HTTP APIs in Go, protecting shared state with a standard `sync.Mutex` forces all read requests to block each other, drastically reducing parallel performance. Upgrading to `sync.RWMutex` with `RLock` allows concurrent reads. Furthermore, keeping lock scopes as small as possible by moving non-shared resource logic (like setting response headers) outside of the lock maximizes parallelism.
**Action:** Always prefer `sync.RWMutex` for endpoints with high read/write ratio and minimize the locked critical section.
