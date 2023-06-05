// run `go test -bench=.``

package main

import (
	"math"
	"testing"
)

func calculateSquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func BenchmarkSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateSquareRoot(100)
	}
}
