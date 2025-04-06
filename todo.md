# Clin
 Cli runner


## 🧪 Project Plan: Minimal CLI Build Runner

### 🗭 Goal:
Create a lightweight CLI tool in **Go** that compiles and runs source code in different languages, cross-platform.

---

### ✅ Phase 1: **Basic Prototype (Go version)**  
- [ ] Write the CLI using Go  
- [ ] Accept a file path as an argument  
- [ ] Detect file extension (`.c`, `.cpp`, `.py`, etc.)  
- [ ] If `.c`, compile it with `gcc`, run the binary  
- [ ] Option to build binary in `./` or `/tmp`  
- [ ] Clean error handling + messages

---

### 🛠️ Phase 2: **Add Support for Other Languages**
- [ ] Python: run with `python3 filename.py`
- [ ] C++: compile with `g++`
- [ ] Java: compile + run `.java` files
- [ ] Rust (optional): use `rustc`
- [ ] Add fallback/default: show “unsupported language”

---

### 🌐 Phase 3: **Cross-Platform Support**
- [ ] Make it work on Linux, Windows, macOS  
- [ ] Handle path differences (`\\` vs `/`)  
- [ ] Detect default compilers for each OS  
- [ ] Compile using Go’s cross-compilation:
  ```bash
  GOOS=windows GOARCH=amd64 go build -o bin.exe
## 🧪 Project Plan: Minimal CLI Build Runner

### 🗭 Goal:
Create a lightweight CLI tool in **Go** that compiles and runs source code in different languages, cross-platform.

---

### ✅ Phase 1: **Basic Prototype (Go version)**  
- [ ] Write the CLI using Go  
- [ ] Accept a file path as an argument  
- [ ] Detect file extension (`.c`, `.cpp`, `.py`, etc.)  
- [ ] If `.c`, compile it with `gcc`, run the binary  
- [ ] Option to build binary in `./` or `/tmp`  
- [ ] Clean error handling + messages

---

### 🛠️ Phase 2: **Add Support for Other Languages**
- [ ] Python: run with `python3 filename.py`
- [ ] C++: compile with `g++`
- [ ] Java: compile + run `.java` files
- [ ] Rust (optional): use `rustc`
- [ ] Add fallback/default: show “unsupported language”

---

### 🌐 Phase 3: **Cross-Platform Support**
- [ ] Make it work on Linux, Windows, macOS  
- [ ] Handle path differences (`\\` vs `/`)  
- [ ] Detect default compilers for each OS  
- [ ] Compile using Go’s cross-compilation:
  ```bash
  GOOS=windows GOARCH=amd64 go build -o bin.exe

