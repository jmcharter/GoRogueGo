package main

import (
	"image/color"

	"github.com/faiface/pixel"
)

// Entity is the base struct that all other entities will be based on
// e.g Player, NPCs, MOBs, Items
type Entity struct {
	x      int
	y      int
	gridX  float64
	gridY  float64
	rect   pixel.Rect
	colour color.RGBA
}

// CreateEntity returns a new Entity with give positional co-ordianates
// and colour.
//
// rect is calculated based on x, y ints and TILE_SIZE const.
func CreateEntity(x, y int, colour color.RGBA) Entity {
	F_TILE_SIZE := float64(TILE_SIZE)
	fx := float64(x)
	fy := float64(y)

	return Entity{
		x:     x,
		y:     y,
		gridX: fx * F_TILE_SIZE,
		gridY: fy * F_TILE_SIZE,
		rect: pixel.Rect{
			Min: pixel.V(fx*F_TILE_SIZE, fy*F_TILE_SIZE),
			Max: pixel.V(fx*F_TILE_SIZE+F_TILE_SIZE, fy*F_TILE_SIZE+F_TILE_SIZE),
		},
		colour: colour,
	}
}

// Move updates the x and y values of Entity by dx and dy
func (e *Entity) Move(dx, dy int) {
	e.x += dx
	e.y += dy
	e.updateGridPos()
}

// updateGridPos updates the values of grid_x and grid_y in respect to x and y
func (e *Entity) updateGridPos() {
	e.gridX = float64(e.x * TILE_SIZE)
	e.gridY = float64(e.y * TILE_SIZE)
}
