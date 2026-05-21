## 2024-05-21 - Parallel Benchmarking of Go HTTP Handlers
**Learning:** When using `b.RunParallel` to benchmark Go HTTP handlers, `httptest.NewRecorder` must be instantiated *inside* the worker loop. Shared recorders are not thread-safe for concurrent writes and will cause data races.
**Action:** Always create a new `httptest.NewRecorder` for each iteration within `b.RunParallel` to ensure thread safety and accurate results.
