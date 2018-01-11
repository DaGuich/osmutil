package geotypes

import "math"

// Vector is a interface for all Vector structs
type Vector interface {
	Length() float64
	Normalize() float64
	Add(v *Vector) *Vector
	Sub(v *Vector) *Vector
	Mult(v *Vector) *Vector
	Angle(v *Vector) float64
}

// Vector2D is a vector in 2D room
type Vector2D struct {
	X float64
	Y float64
}

// Length ...
func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a normalized copy of the vector
func (v *Vector2D) Normalize() *Vector2D {
	var (
		norm *Vector2D
		l    float64
	)

	l = v.Length()
	norm = &Vector2D{
		X: v.X / l,
		Y: v.Y / l,
	}

	return norm
}

// Add two vectors
func (v *Vector2D) Add(other *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Sub tract other vector from vector
func (v *Vector2D) Sub(other *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Mult iply other vector with vector
func (v *Vector2D) Mult(other *Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Angle between vector and other vector
func (v *Vector2D) Angle(other *Vector2D) float64 {
	return math.Acos(v.Mult(other)/(v.Length()*other.Length())) * (180 / math.Pi)
}

// Vector3D is a vector in 3D room
type Vector3D struct {
	X float64
	Y float64
	Z float64
}

// Length ...
func (v *Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a normalized copy of the vector
func (v *Vector3D) Normalize() *Vector3D {
	var (
		norm *Vector3D
		l    float64
	)

	l = v.Length()
	norm = &Vector3D{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}

	return norm
}

// Add two vectors
func (v *Vector3D) Add(other *Vector3D) *Vector3D {
	return &Vector3D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Sub tract other vector from vector
func (v *Vector3D) Sub(other *Vector3D) *Vector3D {
	return &Vector3D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// Mult iply other vector with vector
func (v *Vector3D) Mult(other *Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Angle between vector and other vector
func (v *Vector3D) Angle(other *Vector3D) float64 {
	return math.Acos(v.Mult(other)/(v.Length()*other.Length())) * (180 / math.Pi)
}
