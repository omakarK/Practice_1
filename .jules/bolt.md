# Bolt's Performance Journal

## 2026-08-31 - RWMutex and Critical Section Optimization in Read-Heavy HTTP Handlers
**Learning:** Using `sync.Mutex` on read-heavy HTTP endpoints serializes concurrent client requests even when only reading shared in-memory state. Furthermore, setting HTTP headers inside the locked critical section unnecessarily inflates lock hold time. Switching to `sync.RWMutex` with `RLock`/`RUnlock` and preparing HTTP response headers prior to acquiring read locks dramatically improves concurrent throughput and lowers parallel request latency under multi-core load.
**Action:** In read-dominated in-memory Go endpoints, always prefer `sync.RWMutex` and minimize lock scope by performing all header and pre-serialization setup outside the lock.
