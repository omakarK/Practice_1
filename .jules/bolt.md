# Bolt's Performance Journal

## 2025-07-22 - Upgrading sync.Mutex to sync.RWMutex and reducing critical lock duration
**Learning:** In Go, handlers that perform read-heavy tasks on shared data blocks can become highly bottle-necked under parallel workloads due to exclusive write locks (`sync.Mutex`). Upgrading the lock to a read-write mutex (`sync.RWMutex`), using `RLock()` / `RUnlock()` for the read path, and moving response setups (like header fields and allocations) out of the critical locked section significantly decreases parallel read latency.
**Action:** Always inspect HTTP handlers dealing with read paths. When safe, apply `sync.RWMutex` instead of `sync.Mutex`, minimize the critical locked section, and perform header setting outside of the lock.
