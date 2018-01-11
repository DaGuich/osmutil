package geotypes

import (
	"testing"
)

func BenchmarkVector2D_Angle(b *testing.B) {
	var x, y *Vector2D
	x = &Vector2D{5, 1}
	y = &Vector2D{-12, 8}

	for n := 0; n < b.N; n++ {
		x.Angle(y)
	}
}

func BenchmarkVector3D_Angle(b *testing.B) {
	var x, y *Vector3D
	x = &Vector3D{5, 1, 2}
	y = &Vector3D{-12, 8, 9}

	for n := 0; n < b.N; n++ {
		x.Angle(y)
	}
}
