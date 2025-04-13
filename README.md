# Clin : A smart cli code runner

---

A stupid project to try out a new language (go btw)

It is not that impractical if you are a student learning c, cpp, or java and don't use vscode with code-runner\|  
this is a simple simple tool that does it for you in just one step

## Supported Languages

- C  
- Cpp  
- Fortran  
- Rust  
- Haskel  
- Java  
- Python  
- \<more to be added>

## help

**Usage:**  
`clin [options] [build flags / options] <source-file>`

### Options

- `-v` or `--version` — print version  
- `-h` or `--help` — Show this help message  
- `-o <file>` — Set the output binary file path  
- `-ot <file>` — Set path and not run the binary after building  
- `--build` — Everything after this is considered build flags  

This is not necessary but is there if u need it

### Examples

clin -o bin/myapp test.c
clin -ot bin/myapp hello.cpp
clin script.py
clin --build "
clin -u myscript.py --input data.txt --verbose -n 5
clin myscript.py --build -u--input data.txt --verbose -n 5


### Issues

+ Java no "noexec" support  clin ll always try  to run java class after compilation 
+ compilers with spaces in compilation command  like "go build" and "zig build-exe" are bugged
+ zig  is currently  not working as intended
+ debug thigs are hardcoded and are ON

