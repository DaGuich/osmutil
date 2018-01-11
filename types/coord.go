package geotypes

import (
	"math"
)

// Coord represents a Point as latitude and longitude
type Coord struct {
	Lat    float64
	Lon    float64
	Radian bool
}

// NewCoord creates a new coord with latitude and longitude
// in degrees
func NewCoord(lat, lon float64) *Coord {
	return &Coord{Lat: lat, Lon: lon, Radian: false}
}

// NewCoordFromCartesian returns lat/lon coordinates
// converted from cartesian coordinates with the
// assumption of a spherical earth
func NewCoordFromCartesian(c *CartesianCoord) *Coord {
	var lat, lon, hyp float64

	lon = math.Atan2(c.Y, c.X)
	hyp = math.Sqrt(c.X*c.X + c.Y*c.Y)
	lat = math.Atan2(c.Z, hyp)

	return &Coord{Lat: lat, Lon: lon}
}

// NewCoordRad creates a new coord with latitude and longitude
// in radians
func NewCoordRad(lat, lon float64) *Coord {
	return &Coord{Lat: lat, Lon: lon, Radian: true}
}

// ToRadian returns a copy of Coord converted ro radian
func (c *Coord) ToRadian() *Coord {
	if c.Radian {
		return &Coord{
			Lat:    c.Lat,
			Lon:    c.Lon,
			Radian: true,
		}
	}
	return &Coord{
		Lat:    (c.Lat * math.Pi) / 180,
		Lon:    (c.Lon * math.Pi) / 180,
		Radian: true,
	}
}

// Distance computes the distance between two coordinates in metres
// with the assumption of a spherical earth
func (c *Coord) Distance(dest *Coord) (d float64) {
	var (
		radStart    *Coord
		radDest     *Coord
		earthRadius float64
		a           float64
		b           float64
	)
	earthRadius = 6371e+3

	radStart = c.ToRadian()
	radDest = dest.ToRadian()

	a = math.Sin(radDest.Lat-radStart.Lat) *
		math.Sin(radDest.Lon-radStart.Lon) *
		math.Cos(radStart.Lat) *
		math.Cos(radDest.Lat) *
		math.Sin((radDest.Lon-radStart.Lon)/2) *
		math.Sin((radDest.Lon-radStart.Lon)/2)

	b = 2 *
		math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d = earthRadius * b

	return d
}

// CartesianCoord represents a point with cartesian
// coordinates
type CartesianCoord struct {
	X float64
	Y float64
	Z float64
}

// NewCartesianCoord creates a new Point in cartesian
// coordinate system
func NewCartesianCoord(x, y, z float64) *CartesianCoord {
	return &CartesianCoord{X: x, Y: y, Z: z}
}

// NewCartesianCoordFromCoord returns a point in the
// cartesian coordinate system from lat/lon coordinate
func NewCartesianCoordFromCoord(c *Coord) *CartesianCoord {
	var (
		x, y, z  float64
		radCoord *Coord
	)
	radCoord = c.ToRadian()

	x = math.Cos(radCoord.Lat) * math.Cos(radCoord.Lon)
	y = math.Cos(radCoord.Lat) * math.Sin(radCoord.Lon)
	z = math.Sin(radCoord.Lat)

	return &CartesianCoord{X: x, Y: y, Z: z}
}

// Distance calculates the distance between two cartesian coordinates
func (c *CartesianCoord) Distance(dest *CartesianCoord) float64 {
	return math.Sqrt(c.X*dest.X + c.Y*dest.Y + c.Z*dest.Z)
}
