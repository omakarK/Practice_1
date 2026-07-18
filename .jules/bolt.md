# Bolt Journal
⚡ Lightning-fast codebase optimizations.

## 2025-05-22 - Upgrading sync.Mutex to sync.RWMutex and reducing critical sections
**Learning:** In read-heavy handlers (like getTasksHandler in Go), using `sync.Mutex` blocks other read operations completely. Upgrading to `sync.RWMutex` with `RLock()` allows concurrent reads. Additionally, setting HTTP headers is non-shared and should be done outside the lock to minimize the critical section.
**Action:** Always inspect if shared states are read-only in some handlers, and upgrade those from `sync.Mutex` to `sync.RWMutex`. Ensure HTTP header writes and parsing operations are kept outside the critical section.
