import ctypes

lib = ctypes.CDLL("libfib.so")

def main():
    for i in range(1, 41):
        lib.Fib(i)

if __name__ == '__main__':
    main()
