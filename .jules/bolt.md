## 2026-06-19 - Tooling Truncation Hazard
**Learning:** Standard output tools (read_file, cat) in this environment truncate at 1000 characters, which can lead to incomplete architectural understanding and rejected optimization plans if a file is larger.
**Action:** Always check file size with 'wc -c' before reading, and use 'sed' or 'tail' to read files in chunks if they exceed 1000 bytes.
