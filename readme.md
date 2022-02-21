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
SingleCore/Uint32()-16      10.7ns ± 0%   3.4ns ±49%  -68.34%  (p=0.008 n=5+5)
SingleCore/Uint64()-16      11.3ns ± 0%   5.0ns ± 0%  -55.36%  (p=0.000 n=5+4)
SingleCore/Int()-16         10.8ns ± 0%   6.4ns ±14%  -40.18%  (p=0.008 n=5+5)
SingleCore/Intn(32)-16      11.9ns ± 0%   4.2ns ± 0%  -64.81%  (p=0.029 n=4+4)
SingleCore/Read/1024-16      701ns ± 0%   159ns ± 1%  -77.30%  (p=0.008 n=5+5)
SingleCore/Read/10240-16    6.84µs ± 0%  1.45µs ± 0%  -78.87%  (p=0.016 n=4+5)
SingleCore/Perm/1024-16     24.7µs ± 2%  11.1µs ± 1%  -54.93%  (p=0.008 n=5+5)
SingleCore/Shuffle/1024-16  12.1µs ± 0%   5.3µs ± 1%  -56.40%  (p=0.008 n=5+5)
MultipleCore/Uint32()-16     124ns ± 4%     0ns ± 1%  -99.85%  (p=0.008 n=5+5)
MultipleCore/Uint64()-16     127ns ± 6%     0ns ± 1%  -99.67%  (p=0.008 n=5+5)
MultipleCore/Int()-16        128ns ± 3%     1ns ± 1%  -99.57%  (p=0.008 n=5+5)
MultipleCore/Intn(32)-16     127ns ± 3%     0ns ± 1%  -99.65%  (p=0.008 n=5+5)
MultipleCore/Read/1024-16   1.13µs ± 2%  0.10µs ± 0%  -91.29%  (p=0.008 n=5+5)
MultipleCore/Read/10240-16  10.3µs ± 6%   0.6µs ± 0%  -94.25%  (p=0.008 n=5+5)
MultipleCore/Perm/1024-16    208µs ± 1%     3µs ± 3%  -98.60%  (p=0.008 n=5+5)
```
