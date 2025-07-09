# 📦 go-ds-implementations

This repository contains manual implementations of classic **data structures** and **caching algorithms** using the Go programming language — built entirely from scratch.

No high-level built-in containers like `container/list` or `map` magic unless explicitly required. The goal is to deepen understanding of how these structures work under the hood, their time/space complexities, and real-world use cases.

---

## 🚀 Why This Exists

- ✅ Strengthen core computer science knowledge  
- ✅ Prepare for coding interviews and system design rounds  
- ✅ Explore internals of data structures and their trade-offs  
- ✅ Apply Go’s low-level concurrency primitives in advanced scenarios

---

## 📚 Implementations So Far

### 🧠 Caching Algorithms

| Algorithm   | Description                            | Status  |
|------------|----------------------------------------|---------|
| LRU Cache  | Least Recently Used cache strategy     | ✅ Done |
| LFU Cache  | Least Frequently Used strategy          | 🔜 Soon |
| FIFO Cache | First In First Out strategy             | 🔜 Soon |
| TTL Cache  | Time-to-live expiring cache             | 🔜 Soon |

### 🧱 Data Structures

| Structure          | Description                                  | Status            |
|-------------------|----------------------------------------------|-------------------|
| Doubly Linked List| Core of LRU, supports both ends              | ✅ Done (used in LRU) |
| Stack             | LIFO implementation                          | 🔜 Planned         |
| Queue             | FIFO implementation                          | 🔜 Planned         |
| Hash Map          | Custom hash table (open addressing / chaining) | 🔜 Planned         |
| Min/Max Heap      | Priority queue support                        | 🔜 Planned         |

--- 

## 📂 Folder Structure

```text
go-ds-implementations/
├── caching/                        # Caching strategies like LRU, LFU, FIFO
│   ├── lru.go                      # LRU Cache implementation
│
├── data-structures/                 # Core data structures implemented from scratch
│   └── queue/                      # Queue-related implementations
│       └── queue.go                # Queue (FIFO) implementation
│
├── benchmarks/                     # (Optional) Benchmark tests (planned)
│
├── README.md                       # Project overview and documentation
├── go.mod                          # Go module metadata
└── go.sum                          # Go module dependency checksums
```

---

## ▶️ Usage

To use any of the data structures or caching algorithms, simply import the relevant package and initialize the structure.  
Make sure you have Go installed and run the project like this:

```bash
go run path/to/your_file.go
```

---

### ✅ Contributions
```markdown
## 🤝 Contributions

Contributions are welcome! Feel free to fork this repo, open issues, or submit pull requests.

If you'd like to add a new data structure, caching strategy, or improve tests and documentation, just follow the existing structure and naming conventions.
```

---

✌️ Peace