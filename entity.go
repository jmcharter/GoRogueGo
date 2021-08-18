package main

import (
	"image/color"

	"github.com/faiface/pixel"
)

type Entity struct {
	x      int
	y      int
	grid_x float64
	grid_y float64
	rect   pixel.Rect
	colour color.RGBA
}

func CreateEntity(x, y int, colour color.RGBA) Entity {
	F_TILE_SIZE := float64(TILE_SIZE)
	fx := float64(x)
	fy := float64(y)

	return Entity{
		x:      x,
		y:      y,
		grid_x: fx * F_TILE_SIZE,
		grid_y: fy * F_TILE_SIZE,
		rect: pixel.Rect{
			Min: pixel.V(fx*F_TILE_SIZE, fy*F_TILE_SIZE),
			Max: pixel.V(fx*F_TILE_SIZE+F_TILE_SIZE, fy*F_TILE_SIZE+F_TILE_SIZE),
		},
		colour: colour,
	}
}

func (e *Entity) Move(dx, dy int) {
	e.x += dx
	e.y += dy
	e.updateGridPos()
}

func (e *Entity) updateGridPos() {
	e.grid_x = float64(e.x * TILE_SIZE)
	e.grid_y = float64(e.y * TILE_SIZE)
}
