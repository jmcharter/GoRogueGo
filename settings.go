package main

import (
	"golang.org/x/image/colornames"
)

const (
	WIDTH       = 1024
	HEIGHT      = 768
	TILE_SIZE   = 16
	GRID_WIDTH  = WIDTH / TILE_SIZE
	GRID_HEIGHT = HEIGHT / TILE_SIZE
	TITLE       = "GoRL"
)

var (
	BG_COLOUR   = colornames.Steelblue
	LINE_COLOUR = colornames.Lightsteelblue
	GRID_ON     = true
)
