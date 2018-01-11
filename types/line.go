package geotypes

// Line is a Line
type Line struct {
	Coordinates []*Coord
	Info        map[string]string
}

// NewLine creates a new line
func NewLine(coords []*Coord) *Line {
	return &Line{
		Coordinates: coords,
		Info:        make(map[string]string),
	}
}
