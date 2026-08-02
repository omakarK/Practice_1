# Bolt's Journal

## 2025-05-22 - Upgrading getTasksHandler to sync.RWMutex
**Learning:** Upgrading `sync.Mutex` to `sync.RWMutex` and moving HTTP header settings outside the locked region significantly improves read throughput and reduces lock contention during parallel read operations.
**Action:** Always prefer `sync.RWMutex` for read-heavy operations where write operations are infrequent, and ensure HTTP headers/non-critical preparations are set outside the critical section.
