# Bolt Performance Journal - Critical Learnings Only

## 2026-08-31 - HTTP Handlers and Mutex Contention in Go Dynamic Project
**Learning:** In Go HTTP handlers serving read operations on shared state, using an exclusive `sync.Mutex` causes significant contention during parallel requests. Upgrading to `sync.RWMutex` with `RLock` and keeping header assignments outside the critical lock reduces parallel execution latency by ~50%.
**Action:** Use `sync.RWMutex` for read-heavy shared data and keep non-state setup code outside critical lock regions.
