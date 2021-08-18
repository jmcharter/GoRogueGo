package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func draw_grid(surface *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = LINE_COLOUR
	for x := 0.0; x < WIDTH; x += TILE_SIZE {
		imd.Push(pixel.V(x, 0.0), pixel.V(x, HEIGHT))
		imd.Line(1)
		imd.Draw(surface)
	}
	for y := 0.0; y < HEIGHT; y += TILE_SIZE {
		imd.Push(pixel.V(0.0, y), pixel.V(WIDTH, y))
		imd.Line(1)
		imd.Draw(surface)
	}
}

func draw_entity(entity Entity, surface *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = entity.colour
	imd.Push(pixel.V(entity.grid_x, entity.grid_y))
	imd.Push(pixel.V(entity.grid_x+TILE_SIZE, entity.grid_y+TILE_SIZE))
	imd.Rectangle(0)
	imd.Draw(surface)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  TITLE,
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	player := CreateEntity(
		rand.Intn(GRID_WIDTH),
		rand.Intn(GRID_HEIGHT),
		colornames.Orangered,
	)

	for !win.Closed() {
		check_key_press(&player, win)
		win.Clear(BG_COLOUR)
		draw_entity(player, win)
		if GRID_ON {
			draw_grid(win)
		}
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
