## 2025-05-14 - Concurrent Read Optimization with RWMutex
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` for read-heavy endpoints (like GET /api/tasks) significantly improves parallel throughput. Additionally, moving non-shared state operations like `w.Header().Set` outside the lock further reduces the critical section duration.
**Action:** Always prefer `sync.RWMutex` for shared data structures that are frequently read but infrequently modified. Audit handlers to ensure only the absolute minimum code is inside the lock's scope.

## 2025-05-14 - Parallel Benchmarking in Go
**Learning:** When using `b.RunParallel` in Go benchmarks with `httptest`, the `httptest.NewRecorder` must be instantiated *inside* the worker loop. Using a shared recorder across goroutines leads to data races and incorrect results.
**Action:** Ensure per-iteration objects are correctly scoped within parallel benchmark functions.
