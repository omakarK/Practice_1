## 2025-05-15 - [Race condition in Go parallel benchmarks]
**Learning:** When using `b.RunParallel` in Go benchmarks, sharing a single `httptest.NewRecorder` across multiple goroutines causes a race condition because the recorder is not thread-safe for concurrent writes.
**Action:** Always instantiate `httptest.NewRecorder` inside the `RunParallel` loop to ensure each goroutine has its own recorder.
