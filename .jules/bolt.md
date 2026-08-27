# Bolt's Performance Journal

## 2026-08-27 - RWMutex & Critical Section Minimization in Read-Heavy HTTP Handlers
**Learning:** Replacing `sync.Mutex` with `sync.RWMutex` (`RLock`) and moving non-thread-state operations (such as HTTP header initialization) outside the critical lock section significantly reduces lock contention in parallel read endpoints.
**Action:** Always prefer `sync.RWMutex` for read-heavy handlers and minimize the critical section by hoisting HTTP response initialization outside of mutex locks.
