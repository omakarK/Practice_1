## 2025-05-04 - Parallel Benchmarking Go Handlers
**Learning:** In Go parallel benchmarks using `b.RunParallel`, `httptest.NewRecorder` must be instantiated *inside* the worker loop. Shared recorders are not thread-safe for concurrent writes and will cause race conditions or incorrect results during benchmarking.
**Action:** Always create a fresh `httptest.NewRecorder` within the `pb.Next()` loop when benchmarking HTTP handlers in parallel.
