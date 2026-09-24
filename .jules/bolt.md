# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock & Header Order Optimization in getTasksHandler
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` (`RLock`) and moving response header setup outside the mutex critical section reduces benchmark latency from ~1900 ns/op to ~840 ns/op under parallel read loads.
**Action:** Use `sync.RWMutex` for read-heavy shared data structures and keep non-state I/O operations outside critical lock sections.
