## 2025-05-14 - sync.RWMutex for Read-Heavy Handlers
**Learning:** In simple Go web services like 'go-dynamic-project', using a global `sync.Mutex` for both reads and writes creates unnecessary contention on 'get' endpoints. Switching to `sync.RWMutex` and moving header assignments outside the lock significantly reduces latency.
**Action:** Always check if a shared state is primarily read from and prefer `sync.RWMutex` to allow concurrent access.
