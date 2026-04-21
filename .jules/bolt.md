## 2025-05-15 - [RWMutex for Concurrent Reads]
**Learning:** In read-heavy Go applications with global state protected by a mutex, switching from `sync.Mutex` to `sync.RWMutex` provides a significant performance boost for read operations by allowing them to run concurrently.
**Action:** Always evaluate if a protected resource is frequently read but infrequently modified, and use `sync.RWMutex` to prevent unnecessary serialization of read requests.
