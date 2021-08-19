package main

import (
	"golang.org/x/image/colornames"
)

// Settings for window and grid.
// TILE_SIZE should be a factor of both width and height to result in a
// grid that fits flush to the window.
const (
	WIDTH       = 1024               // Window width
	HEIGHT      = 768                // Window height
	TILE_SIZE   = 48                 // Size of individual grid tiles
	GRID_WIDTH  = WIDTH / TILE_SIZE  // Grid width in tiles
	GRID_HEIGHT = HEIGHT / TILE_SIZE // Grid height in tiles
	TITLE       = "GoRL"             // Title displayed for window
)

// Various other settings such as colours and booleans.
var (
	BG_COLOUR   = colornames.Steelblue      // Window background colour
	LINE_COLOUR = colornames.Lightsteelblue // Grid line colour
	GRID_ON     = true                      // true to display grid
)
