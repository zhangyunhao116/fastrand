package fastrand

import (
	"math/rand"
	"os"
	"testing"
)

type testFuncs struct {
	fu32     func() uint32
	fu64     func() uint64
	fint     func() int
	fintn    func(_ int) int
	fread    func(_ []byte) (int, error)
	fperm    func(_ int) []int
	fshuffle func(_ int, swap func(i, j int))
}

func init() {
	if os.Getenv("BENCHMARK_TARGET") == "std" {
		targetfs = stdrandfs
	} else {
		targetfs = fastrandfs
	}
}

var (
	targetfs  testFuncs // benchmark this target
	stdrandfs = testFuncs{
		fu32:     rand.Uint32,
		fu64:     rand.Uint64,
		fint:     rand.Int,
		fintn:    rand.Intn,
		fread:    rand.Read,
		fperm:    rand.Perm,
		fshuffle: rand.Shuffle,
	}
	fastrandfs = testFuncs{
		fu32:     Uint32,
		fu64:     Uint64,
		fint:     Int,
		fintn:    Intn,
		fread:    Read,
		fperm:    Perm,
		fshuffle: Shuffle,
	}
)

func benchmarkSingleCore(b *testing.B, fs testFuncs) {
	b.Run("Uint32()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fs.fu32()
		}
	})
	b.Run("Uint64()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fs.fu64()
		}
	})
	b.Run("Int()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fs.fint()
		}
	})
	b.Run("Intn(32)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fs.fintn(32)
		}
	})
	b.Run("Read/1024", func(b *testing.B) {
		p := make([]byte, 1024)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fs.fread(p)
		}
	})
	b.Run("Read/10240", func(b *testing.B) {
		p := make([]byte, 10240)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fs.fread(p)
		}
	})
	b.Run("Perm/1024", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fs.fperm(1024)
		}
	})
	b.Run("Shuffle/1024", func(b *testing.B) {
		x := make([]int, 1024)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fs.fshuffle(1024, func(i, j int) {
				x[i], x[j] = x[j], x[i]
			})
		}
	})
}

func benchmarkMultipleCore(b *testing.B, fs testFuncs) {
	b.Run("Uint32()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fs.fu32()
			}
		})
	})
	b.Run("Uint64()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fs.fu64()
			}
		})
	})
	b.Run("Int()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fs.fint()
			}
		})
	})
	b.Run("Intn(32)", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fs.fintn(32)
			}
		})
	})
	b.Run("Read/1024", func(b *testing.B) {
		p := make([]byte, 1024)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fs.fread(p)
			}
		})
	})
	b.Run("Read/10240", func(b *testing.B) {
		p := make([]byte, 10240)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fs.fread(p)
			}
		})
	})
	b.Run("Perm/1024", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fs.fperm(1024)
			}
		})
	})
}

func BenchmarkSingleCore(b *testing.B) {
	benchmarkSingleCore(b, targetfs)
}

func BenchmarkMultipleCore(b *testing.B) {
	benchmarkMultipleCore(b, targetfs)
}
