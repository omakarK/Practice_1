## 2025-05-11 - Thread-safe Benchmarking in Go
**Learning:** When using `b.RunParallel` in Go to benchmark HTTP handlers, `httptest.NewRecorder` must be instantiated inside the worker loop. Shared recorders are not thread-safe for concurrent writes and will lead to race conditions or incorrect results.
**Action:** Always ensure request-scoped objects like `httptest.NewRecorder` are local to the parallel worker loop in benchmarks.
