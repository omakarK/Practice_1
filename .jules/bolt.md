# Bolt's Journal - Critical Learnings

## 2026-08-16 - RWMutex & Unlocking Work Outside Critical Section in Go HTTP Handlers

**Learning:** In Go HTTP handlers reading shared slice state, using exclusive `sync.Mutex` causes unnecessary contention during concurrent GET requests. Upgrading to `sync.RWMutex` with `RLock` allows concurrent reads. Furthermore, setting HTTP headers before/outside the lock reduces lock hold time.
**Action:** Replace `sync.Mutex` with `sync.RWMutex` for read-heavy shared state and move non-state modifications outside the critical lock section.
