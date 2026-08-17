# Bolt's Performance Journal

## 2026-08-17 - Upgrading Mutex to RWMutex for Read-Heavy HTTP Handlers
**Learning:** Exclusive `sync.Mutex` locks on read-heavy HTTP endpoints (like `GET /api/tasks`) create severe contention under parallel requests. Moving non-shared operations (such as setting HTTP headers) outside the lock and switching to `sync.RWMutex.RLock()` allows concurrent reads and drastically reduces handler latency under load.
**Action:** Always prefer `sync.RWMutex` when reads outnumber writes, and ensure header manipulation / request setup occurs outside the lock boundary.
