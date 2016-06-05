# libfib

An experiment to run Go code directly from Python.

### How it works

`libfib.go` is a very naive implementation of Fibonacci sequence. Using Go's build modes, we build it into a shared library. The library is imported in Python and used there.

Basically, we execute the Go code from Python code even though the latter doesn't even know what Go is.

Run `go help buildmodes` to learn more about Go's build modes.

### Why?

1. It's cool.
2. Go is much faster than Python. You can rewrite slow parts of your code in Go, and call them from your Python code. See [benchmarks](#benchmarks).
3. It's cool.

### Benchmarks

Benchmarks are the best way to compare the performance and really show the **coolness** of Go's build modes.

*The function to benchmark:* Find the first 40 Fibonacci numbers. That is `Fib` function will be called with `n` from `1...40`. The execution time of the full loop is then measured.

*Results*: `libfib` is 76X faster!!!1111

#### Pure Python implementation

`benchmark.py` implements the same Fibonacci function in pure Python.Here's the result of the Linux `time` command:

```
$ time python benchmark.py

real  2m26.726s
user  2m6.232s
sys 0m2.280s
```

#### Shared library implementation

`benchmark_libfib.py` calls the Fibonacci function in the shared library from Python. Here's the result of the Linux `time` command:

```
$ time python benchmark_libfib.py

real  0m1.899s
user  0m1.593s
sys 0m0.039s
```

### Creating the shared objects file

```
go build -buildmode=c-shared -o libfib.so libfib.go
```

This will produce `fib.h` and `fib.so`.

### Using `.so` in Python

Now we can load the library in Python and start using it.

```
import ctypes
lib = ctypes.CDLL("libfib.so")
print(lib.Fib(20))
```

### License

MIT (c) Karan Goel
