package main

type MapVec struct {
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
func (rec mapRectangle) GetCentre() MapVec {
	centre_x := (rec.minX + rec.maxX) / 2
	centre_y := (rec.minY + rec.maxY) / 2

	return MapVec{centre_x, centre_y}
}

// DoesIntersect returns a boolean.
// If this rectangle interesects with another given rectangle, returns true
func (rec mapRectangle) DoesIntersect(other mapRectangle) bool {
	return rec.minX <= other.maxX &&
		rec.minY <= other.maxY &&
		rec.maxX >= other.minX &&
		rec.maxY >= other.minY
}
