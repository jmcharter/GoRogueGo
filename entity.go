package main

import (
	"image/color"

	"github.com/faiface/pixel"
)

// Slice to hold all active entities
var entities = []*Entity{}

// Entity is the base struct that all other entities will be based on
// e.g Player, NPCs, MOBs, Items
type Entity struct {
	x              int
	y              int
	gridX          float64
	gridY          float64
	rect           pixel.Rect
	colour         color.RGBA
	blocksMovement bool
}

// CreateEntity returns a new Entity with give positional co-ordianates
// and colour.
//
// rect is calculated based on x, y ints and TILE_SIZE const.
func CreateEntity(x, y int, colour color.RGBA, blocksMovement bool) Entity {
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
		colour:         colour,
		blocksMovement: blocksMovement,
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

// CheckCollision will compare an entities expected new position with the current
// position of all other existing entities and return a bool.
// Return true if there would be two entities in on position.
func (e Entity) CheckCollision(dx, dy int, others []*Entity) bool {
	for _, entity := range others {
		if !entity.blocksMovement {
			continue
		}
		if entity.x == e.x+dx && entity.y == e.y+dy {
			return true
		}
	}
	return false
}
