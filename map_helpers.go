package main

import "math/rand"

type mapVec struct {
	x, y int
}

// Add sums the x and y values of two mapVecs and returns a new mapVec using
// these values
func (mv mapVec) Add(other mapVec) mapVec {
	x := mv.x + other.x
	y := mv.y + other.y
	return mapVec{x, y}
}

type mapRectangle struct {
	minX, minY, maxX, maxY   int
	top, right, bottom, left []mapVec
}

// CreateMapRectangle returns a new mapRectangle based on coordinates x,y and
// given width and height
func CreateMapRectangle(x, y, w, h int) mapRectangle {
	maxX := x + w - 1
	maxY := y + h - 1
	top := getSide(x, maxX, maxY, true)
	right := getSide(y, maxY, maxX, false)
	bottom := getSide(x, maxX, y, true)
	left := getSide(y, maxY, x, false)

	return mapRectangle{x, y, maxX, maxY, top, right, bottom, left}
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
	return (rec.minX-1 <= other.maxX &&
		rec.minY-1 <= other.maxY &&
		rec.maxX+1 >= other.minX &&
		rec.maxY+1 >= other.minY)
}

// Return a random side of the rectangle and it's direction expressed as:
// north: +1
// east: +1
// south: -1
// west: -1
func (rec mapRectangle) GetRandomSide() (side []mapVec, dir mapVec) {
	sideChoice := rand.Intn(4)
	switch sideChoice {
	case 0:
		return rec.top, mapVec{0, 1}
	case 1:
		return rec.right, mapVec{1, 0}
	case 2:
		return rec.bottom, mapVec{0, -1}
	case 3:
		return rec.left, mapVec{-1, 0}
	}
	return nil, mapVec{}
}

// start and finish are the co-ordinates of the first and last position to get
// plane is the horizontal or vertical line to be iterated over
func getSide(start, finish, plane int, horizontal bool) (vecs []mapVec) {
	for i := start; i <= finish; i++ {
		var vec mapVec
		if horizontal {
			vec = mapVec{i, plane}
		} else {
			vec = mapVec{plane, i}
		}
		vecs = append(vecs, vec)
	}
	return
}
