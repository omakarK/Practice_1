# Bolt's Performance Journal

This is Bolt's performance journal for logging critical learnings and insights regarding application optimizations.

## 2026-08-10 - Optimizing Parallel Reads via sync.RWMutex and Reduced Lock Duration
**Learning:** In HTTP handlers serving static/slow-changing resources, using `sync.Mutex` leads to significant lock contention during high concurrent read loads. Upgrading to `sync.RWMutex` with `RLock`/`RUnlock` enables multiple concurrent readers to proceed in parallel without blocking each other. Additionally, removing HTTP header configurations (like `Content-Type`) from the critical locked section reduces lock duration further, optimizing overall response time.
**Action:** Always favor `sync.RWMutex` for read-heavy resources and ensure that non-critical code (e.g., HTTP headers setting, static value generation) resides outside the mutex scope.
