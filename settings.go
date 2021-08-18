package main

import (
	"golang.org/x/image/colornames"
)

const (
	WIDTH       = 1024               // Window width
	HEIGHT      = 768                // Window height
	TILE_SIZE   = 16                 // Size of individual grid tiles
	GRID_WIDTH  = WIDTH / TILE_SIZE  // Grid width in tiles
	GRID_HEIGHT = HEIGHT / TILE_SIZE // Grid height in tiles
	TITLE       = "GoRL"             // Title displayed for window
)

var (
	BG_COLOUR   = colornames.Steelblue      // Window background colour
	LINE_COLOUR = colornames.Lightsteelblue // Grid line colour
	GRID_ON     = true                      // true to display grid
)
