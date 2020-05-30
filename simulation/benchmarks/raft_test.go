package benchmarks

import (
	"testing"
)

func BenchmarkRaft(b *testing.B) {
	runPubSubBenchmark(b, []string{
		":7476",
		":7477",
		":7478",
		":7479",
	})
}
