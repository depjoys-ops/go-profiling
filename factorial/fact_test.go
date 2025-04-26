package factorial

import (
	"testing"
)

const (
	factNumber = 3628800
)

// go test -v ./... -bench=. -benchtime 100x (-count=2)
func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(factNumber)
	}
}

func BenchmarkDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(factNumber)
	}
}

// go test -v ./... -bench="BenchmarkCalculate" -run=^# -count=10 | tee CalcRecursive.txt
// go test -v ./... -bench="BenchmarkCalculate" -run=^# -count=10 | tee CalcDynamic.txt
// benchstat CalcRecursive.txt CalcDynamic.txt
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate()
	}	
}
