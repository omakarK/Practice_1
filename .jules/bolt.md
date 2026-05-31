## 2025-05-15 - RWMutex and Header Setting Optimization
**Learning:** Switching from sync.Mutex to sync.RWMutex and moving non-shared resource operations (like setting HTTP headers) outside of critical sections reduced parallel read latency by ~55% in the Go task service.
**Action:** Always audit Go handlers for operations that can be performed outside the locked region and prefer RWMutex for read-heavy shared data.
