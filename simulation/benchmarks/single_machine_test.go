package benchmarks

import (
	"testing"
)

func BenchmarkSingleMachine(b *testing.B) {
	runPubSubBenchmark(b, []string{":7476"})
}
