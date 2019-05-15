package main

import "testing"

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func BenchmarkArea(b *testing.B) {
	rectangle := Rectangle{12, 6}
	want := 72.0

	for i := 0; i < b.N; i++ {
		b.Run("rectangles", func(b *testing.B) {
			got := rectangle.Area()
			if got != want {
				b.Errorf("got %.2f want %.2f", got, want)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					rectangle.Area()
				}
			},
		)
	}
}
