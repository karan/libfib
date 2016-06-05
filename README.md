# libfib

An experiment to run Go code directly from Python.

### How it works

The Go (Golang) module is built as a library (`.so` and a free `.h`) that is then imported in Python. This lets Python code execute the Go code even though it doesn't even know what Go is.

### Why?

1. It's cool.
2. Go is much faster than Python. You can rewrite slow parts of your code in Go, and call them from your Python code.
3. It's cool.

### Creating the shared objects file

```
go build -buildmode=c-shared -o fib.so fib.go
```

Run `go help buildmodes` to learn more about Go's build modes.

