package geotypes

// Point is a point
type Point struct {
	Coordinates *Coord
	Info        map[string]string
}

// NewPoint creates a Point
func NewPoint(coord *Coord) *Point {
	return &Point{
		Coordinates: coord,
		Info:        make(map[string]string),
	}
}
