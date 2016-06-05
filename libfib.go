// package name: main
package main

import "C"

// A very crappy implementation of Fibonacci
// to get the point across.
//export Fib
func Fib(n int) int {
  if n <= 2 {
    return 1
  }
  return Fib(n-1) + Fib(n-2)
}

// Required to allow CGo to compile to a shared library.
func main() {}
