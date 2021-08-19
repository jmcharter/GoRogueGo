package main

type mapVec struct {
	x, y int
}

type mapRectangle struct {
	minX, minY, maxX, maxY int
}

// CreateMapRectangle returns a new mapRectangle based on coordinates x,y and
// given width and height
func CreateMapRectangle(x, y, w, h int) mapRectangle {
	maxX := x + w
	maxY := y + h

	return mapRectangle{x, y, maxX, maxY}
}

// GetCentre returns the centre coordinates of the rectangle
func (rec mapRectangle) GetCentre() mapVec {
	centreX := (rec.minX + rec.maxX) / 2
	centreY := (rec.minY + rec.maxY) / 2

	return mapVec{centreX, centreY}
}

// DoesIntersect returns a boolean.
// If this rectangle interesects with another given rectangle, returns true
func (rec mapRectangle) DoesIntersect(other mapRectangle) bool {
	return rec.minX <= other.maxX &&
		rec.minY <= other.maxY &&
		rec.maxX >= other.minX &&
		rec.maxY >= other.minY
}
