# fastrand

`fastrand` is the fastest pseudo-random number generator in Go. Support most common APIs of `math/rand`.

This generator comes from Go runtime per-M structure, and the init-seed is provided by Go runtime, which means you can't add your seed, but these methods scale very well on multiple cores.

The generator passes the SmallCrush suite, part of TestU01 framework: http://simul.iro.umontreal.ca/testu01/tu01.html

## Compare to math/rand

- **2 ~ 200x faster**
- Scales well on multiple cores
- **Not** provide a stable value stream (can't inject init-seed)
- Fix bugs in math/rand `Float64` and `Float32`  (since no need to preserve the value stream)

## Benchmark

Go version: go1.18 linux/amd64

CPU: AMD 3700x(8C16T), running at 3.6GHz

OS: ubuntu 18.04

MEMORY: 16G x 2 (3200MHz)


```
name                        old time/op  new time/op  delta
SingleCore/Uint32()-16      10.7ns ± 1%   3.6ns ± 0%  -66.02%  (p=0.016 n=5+4)
SingleCore/Uint64()-16      11.3ns ± 0%   6.0ns ±24%  -47.25%  (p=0.008 n=5+5)
SingleCore/Int()-16         10.8ns ± 0%   5.6ns ± 1%  -48.00%  (p=0.000 n=5+4)
SingleCore/Intn(32)-16      11.9ns ± 0%   4.2ns ± 0%  -64.75%  (p=0.008 n=5+5)
SingleCore/Read/1024-16      700ns ± 0%   157ns ± 1%  -77.52%  (p=0.008 n=5+5)
SingleCore/Read/10240-16    6.84µs ± 0%  1.44µs ± 0%  -78.89%  (p=0.008 n=5+5)
MultipleCore/Uint32()-16     126ns ± 5%     0ns ± 1%  -99.78%  (p=0.008 n=5+5)
MultipleCore/Uint64()-16     132ns ± 3%     0ns ± 0%  -99.69%  (p=0.016 n=5+4)
MultipleCore/Int()-16        127ns ± 4%     1ns ± 0%  -99.57%  (p=0.008 n=5+5)
MultipleCore/Intn(32)-16     129ns ± 4%     0ns ± 1%  -99.67%  (p=0.008 n=5+5)
MultipleCore/Read/1024-16   1.13µs ± 5%  0.10µs ± 0%  -91.29%  (p=0.008 n=5+5)
MultipleCore/Read/10240-16  10.0µs ± 8%   0.6µs ± 1%  -94.09%  (p=0.008 n=5+5)
```
