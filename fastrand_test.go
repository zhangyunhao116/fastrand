package fastrand

import (
	"math/rand"
	"testing"
)

func TestAll(t *testing.T) {
	_ = Uint32()
}

func BenchmarkSingleCore(b *testing.B) {
	b.Run("math/rand-Int31n(5)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rand.Int31n(5)
		}
	})
	b.Run("fast-rand-Int31n(5)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int31n(5)
		}
	})

	b.Run("math/rand-Int63n(5)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rand.Int63n(5)
		}
	})
	b.Run("fast-rand-Int63n(5)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int63n(5)
		}
	})

	b.Run("math/rand-Uint32()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rand.Uint32()
		}
	})
	b.Run("fast-rand-Uint32()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Uint32()
		}
	})
	b.Run("math/rand-Uint64()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rand.Uint64()
		}
	})

	b.Run("fast-rand-Uint64()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Uint64()
		}
	})

}

func BenchmarkMultipleCore(b *testing.B) {
	b.Run("math/rand-Int31n(5)", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Int31n(5)
			}
		})
	})
	b.Run("fast-rand-Int31n(5)", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Int31n(5)
			}
		})
	})

	b.Run("math/rand-Int63n(5)", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Int63n(5)
			}
		})
	})
	b.Run("fast-rand-Int63n(5)", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Int63n(5)
			}
		})
	})

	b.Run("math/rand-Float32()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Float32()
			}
		})
	})
	b.Run("fast-rand-Float32()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Float32()
			}
		})
	})

	b.Run("math/rand-Uint32()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Uint32()
			}
		})
	})
	b.Run("fast-rand-Uint32()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Uint32()
			}
		})
	})

	b.Run("math/rand-Uint64()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Uint64()
			}
		})
	})
	b.Run("fast-rand-Uint64()", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Uint64()
			}
		})
	})
}
