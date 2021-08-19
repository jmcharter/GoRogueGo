package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func drawGrid(surface *pixelgl.Window) {
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

func drawEntity(entity Entity, surface *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = entity.colour
	imd.Push(pixel.V(entity.gridX, entity.gridY))
	imd.Push(pixel.V(entity.gridX+TILE_SIZE, entity.gridY+TILE_SIZE))
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
		10,
		5,
		colornames.Orangered,
		false,
	)

	entities = append(entities, &player)

	for x := 10; x < 16; x++ {
		newEntity := CreateEntity(x, 15, colornames.Navajowhite, true)
		entities = append(entities, &newEntity)
	}

	for !win.Closed() {
		checkKeyPress(&player, win)
		win.Clear(BG_COLOUR)
		for _, entity := range entities {
			drawEntity(*entity, win)
		}
		if GRID_ON {
			drawGrid(win)
		}
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
