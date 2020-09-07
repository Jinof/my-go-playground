package conster

import "testing"

func BenchmarkConstInsideFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstInsideFunction()
	}
}

func BenchmarkConstOutsideFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstOutsideFunction()
	}
}
