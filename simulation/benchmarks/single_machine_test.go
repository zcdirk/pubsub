package benchmarks

import (
	"testing"
)

func BenchmarkSingleMachine(b *testing.B) {
	createBenchmark(b, []string{":7476"})
}
