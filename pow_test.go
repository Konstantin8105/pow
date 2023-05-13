package pow_test

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Konstantin8105/binaryexpr"
	"github.com/Konstantin8105/pow"
)

func Test(t *testing.T) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		return binaryexpr.Test(path)
	})
	if err != nil {
		t.Error(err)
	}
}

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

// cpu: Intel(R) Xeon(R) CPU           X5550  @ 2.67GHz
// Benchmark/math.Pow2-16         	35098486	        34.05 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.E2-16            	1000000000	         1.028 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.En2-16           	142912671	         8.224 ns/op	       0 B/op	       0 allocs/op
// Benchmark/math.Pow3-16         	35733325	        35.76 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.E3-16            	1000000000	         1.056 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.En3-16           	126317251	         9.516 ns/op	       0 B/op	       0 allocs/op
// Benchmark/math.Pow4-16         	33246009	        36.10 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.E4-16            	1000000000	         1.021 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.En4-16           	126656416	         9.524 ns/op	       0 B/op	       0 allocs/op
// Benchmark/math.Pow(x,_51)-16   	27771769	        43.08 ns/op	       0 B/op	       0 allocs/op
// Benchmark/pow.En(x,_51)-16     	48447409	        24.23 ns/op	       0 B/op	       0 allocs/op
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
