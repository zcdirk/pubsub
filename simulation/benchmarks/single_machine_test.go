package benchmarks

import (
	"testing"
)

func BenchmarkSingleMachine(b *testing.B) {
	createPubSubBenchmark(b, []string{":7476"})
}
