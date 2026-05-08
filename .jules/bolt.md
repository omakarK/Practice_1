## 2026-05-08 - Use RWMutex for Concurrent Task Reads
**Learning:** In Go, replacing sync.Mutex with sync.RWMutex for read-heavy operations significantly improves parallel throughput and reduces latency by allowing multiple concurrent readers.
**Action:** Always evaluate if a protected resource is read-heavy and consider sync.RWMutex to avoid unnecessary serialization of read requests.
