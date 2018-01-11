package geotypes

// Polygon is a polygon
type Polygon struct {
	Coordinates []*Coord
	Info        map[string]string
}

// NewPolygon creates a new polygon
func NewPolygon(coords []*Coord) *Polygon {
	return &Polygon{
		Coordinates: coords,
		Info:        make(map[string]string),
	}
}
