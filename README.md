# clin: A Universal Compiler and Interpreter Wrapper

`clin` is a command-line tool that simplifies the process of compiling and running code across various programming languages. It intelligently detects the language based on the file extension and acts as a wrapper around the appropriate compiler or interpreter.

**Key Features:**

* **Automatic Language Detection:** Runs the correct compiler or interpreter based on the input file's extension.
* **Universal `-o` Flag:** Specifies the output path for compiled binaries.
* **Pass-through Flags:** Unknown flags are directly passed to the underlying compiler or interpreter.
* **Default Execution:** By default, `clin` compiles (if necessary) and runs the resulting binary located in a temporary OS-dependent directory (`/tmp`).
* **Customizable Behavior:**
    * `-o <path>`: Specify the output path for the binary.
    * `-ot <path>`: Specify the output path but do not execute the binary.
    * `-t`: Do not execute the binary after compilation.

**Supported Languages:**

`clin` currently supports the following programming languages:

* C (`.c`)
* C++ (`.cpp`)
* Python (`.py`)
* Go (`.go`)
* Zig (`.zig`)
* Rust (`.rs`)
* Java (`.java`)
* Swift (`.swift`)
* Fortran 90 (`.f90`)
* Fortran 95 (`.f95`)
* Fortran 77 (`.f`)
* Fortran 2003 (`.f03`)
* Fortran 2008 (`.f08`)
* Fortran (other extensions) (`.for`)
* Haskell (`.hs`)
* Ruby (`.rb`)
* Perl (`.pl`)
* PHP (`.php`)
* Lua (`.lua`)
* JavaScript (`.js`)
* TypeScript (`.ts`)

With `clin`, you can seamlessly compile and run your code without needing to remember the specific commands for each language.

## Installations


## help

**Usage:**
`clin [options] <source-file> [build flags]`

### Options

* `-v` or `--version` — print the clin version
* `-h` or `--help` — Show this help message
* `--verbose` — enable verbose output (useful for debugging)
* `-o <file>` — Set the output binary file path
* `-ot <file>` — Set the output path for the binary but do not run it after building
* `-t` — do not run the binary after compilation (useful for not running Java classes)
* `--build` — Everything after this option is considered build flags and will be passed directly to the underlying compiler or interpreter. This option is not strictly necessary but can be used for clarity.

### Examples

* `clin -o bin/myapp test.c`
    Compile `test.c` and save the executable as `bin/myapp`.
* `clin -ot bin/myapp hello.cpp`
    Compile `hello.cpp` and save the executable as `bin/myapp`, but do not run it.
* `clin script.py`
    Run the Python script `script.py`.
* `clin --build "-Wall -O2" mycode.c`
    Compile `mycode.c` with the build flags `-Wall` and `-O2`. The output binary will be created in the default temporary location and run.
* `clin -u myscript.py --input data.txt --verbose -n 5`
    Run the Python script `myscript.py`, passing the arguments `-u --input data.txt --verbose -n 5` directly to the interpreter.
* `clin myscript.py --build "-u --input data.txt --verbose -n 5"`
    Run the Python script `myscript.py`, treating `-u --input data.txt --verbose -n 5` as build flags passed to the interpreter. Note the quoting to treat them as a single argument to `--build`.


## Issues

+ Does  not support flags like -d in java and similar, it ll still work but you will get error during runtime use -t & -ot flag for such cases
+ compilers/interpreters for specific languages are hoardcoded. I will add support for clinrc file to solve that issue 
+ No support for runtime arguments or flags, same with runner flags like flag for java somethin.class
+ Could be optimised for efficient logic.

