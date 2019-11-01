package pow_test

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/Konstantin8105/pow"
)

func TestE2(t *testing.T) {
	tcs := []struct {
		x      float64
		result float64
	}{
		{1.0, 1.0},
		{2.0, 4.0},
		{3.0, 9.0},
		{math.NaN(), math.NaN()},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.Inf(1)},
	}
	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a := pow.E2(tcs[i].x)
			e := math.Pow(tcs[i].x, 2.0)
			if a != e && math.IsNaN(a) != math.IsNaN(e) {
				t.Errorf("got %14e , want %.14e", a, e)
			}
		})
	}
}

func TestE3(t *testing.T) {
	tcs := []struct {
		x      float64
		result float64
	}{
		{1.0, 1.0},
		{2.0, 8.0},
		{3.0, 27.0},
		{math.NaN(), math.NaN()},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.Inf(-1)},
	}
	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a := pow.E3(tcs[i].x)
			e := math.Pow(tcs[i].x, 3.0)
			if a != e && math.IsNaN(a) != math.IsNaN(e) {
				t.Errorf("got %14e , want %.14e", a, e)
			}
		})
	}
}

func TestE4(t *testing.T) {
	tcs := []struct {
		x      float64
		result float64
	}{
		{1.0, 1.0},
		{2.0, 16.0},
		{3.0, 81.0},
		{math.NaN(), math.NaN()},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.Inf(-1)},
	}
	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a := pow.E4(tcs[i].x)
			e := math.Pow(tcs[i].x, 4.0)
			if a != e && math.IsNaN(a) != math.IsNaN(e) {
				t.Errorf("got %14e , want %.14e", a, e)
			}
		})
	}
}

func TestEn(t *testing.T) {
	tcs := []struct {
		x float64
		e int
	}{
		{1.0, 1.0},
		{2.0, 8.0},
		{3.0, 27.0},
		{1.0, -1.0},
		{2.0, -8.0},
		{3.0, -27.0},
	}
	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a := pow.En(tcs[i].x, tcs[i].e)
			e := math.Pow(tcs[i].x, float64(tcs[i].e))
			if a != e && math.IsNaN(a) != math.IsNaN(e) {
				t.Errorf("got %14e , want %.14e", a, e)
			}
		})
	}
}

func Example() {
	x := 2.0
	r2 := pow.E2(x) // math.Pow(x, 2.0)
	r3 := pow.E3(x) // math.Pow(x, 3.0)
	fmt.Fprintf(os.Stdout, "%.4f\n", r2)
	fmt.Fprintf(os.Stdout, "%.4f\n", r3)
	// Output:
	// 4.0000
	// 8.0000
}

func Benchmark(b *testing.B) {
	x := math.Pi
	var y float64

	b.Run("math.Pow2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = math.Pow(x, 2.0)
		}
	})
	b.Run("pow.E2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.E2(x)
		}
	})
	b.Run("pow.En2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.En(x, 2)
		}
	})

	b.Run("math.Pow3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = math.Pow(x, 3.0)
		}
	})
	b.Run("pow.E3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.E3(x)
		}
	})
	b.Run("pow.En3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.En(x, 3)
		}
	})

	b.Run("math.Pow4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = math.Pow(x, 4.0)
		}
	})
	b.Run("pow.E4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.E4(x)
		}
	})
	b.Run("pow.En4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.En(x, 4)
		}
	})

	b.Run("math.Pow(x, 51)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = math.Pow(x, 51.0)
		}
	})
	b.Run("pow.En(x, 51)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			y = pow.En(x, 51)
		}
	})
	_ = y
}
