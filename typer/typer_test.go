package typer

import "testing"

func BenchmarkTestTypeInside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestTypeInside("inside")
	}
}

func BenchmarkTestTypeOutside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestTypeOutside("outside")
	}
}
