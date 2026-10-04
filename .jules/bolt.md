# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read-Locking vs Mutex Lock for HTTP Handlers
**Learning:** Holding an exclusive `sync.Mutex` during HTTP headers setting and JSON encoding blocks all concurrent read requests. Switching to `sync.RWMutex` with `RLock` and moving HTTP header configuration and JSON encoding outside or with minimal read lock time drastically improves parallel throughput.
**Action:** Use `sync.RWMutex` for read-heavy shared data structures and keep lock holding times strictly focused on accessing shared memory.
