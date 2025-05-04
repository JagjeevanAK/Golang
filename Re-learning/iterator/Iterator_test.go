package iterator

import "testing"

func BenchmarkLoop(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		loop("a")
	}
}