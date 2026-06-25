## 2025-05-15 - Optimizing Go HTTP Read Handlers with RWMutex
**Learning:** In read-intensive Go services, upgrading from `sync.Mutex` to `sync.RWMutex` and minimizing the critical section by moving non-shared resource operations (like HTTP header settings) outside of the lock significantly improves concurrent throughput.
**Action:** Always check if a read-only path is using a full Mutex and look for opportunities to shrink the locked section in Go web handlers.
