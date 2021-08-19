package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func drawGrid(surface *pixelgl.Window) error {
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
	return nil
}

func drawWall(GameMap gameMap, surface *pixelgl.Window) {
	F_TILE_SIZE := float64(TILE_SIZE)
	imd := imdraw.New(nil)
	imd.Color = WALL_COLOUR

	for y := range GameMap.tiles {
		for x := range GameMap.tiles[y] {
			if GameMap.tiles[y][x].blocksMovement {
				fy := float64(y) * F_TILE_SIZE
				fx := float64(x) * F_TILE_SIZE
				square := pixel.R(fx, fy, fx+F_TILE_SIZE, fy+F_TILE_SIZE)
				imd.Push(square.Min)
				imd.Push(square.Max)
				imd.Rectangle(0)
			}
		}
	}

	imd.Draw(surface)

}

func drawEntity(entity Entity, surface *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = entity.colour
	imd.Push(entity.GetRect().Min)
	imd.Push(entity.GetRect().Max)
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
		3,
		3,
		colornames.Orangered,
		false,
	)

	var GameMap gameMap
	GameMap.init()
	entities = append(entities, &player)

	for !win.Closed() {
		checkKeyPress(&player, GameMap, win)
		win.Clear(BG_COLOUR)
		drawWall(GameMap, win)
		for _, entity := range entities {
			drawEntity(*entity, win)
		}
		if GRID_ON {
			err := drawGrid(win)
			if err != nil {
				fmt.Print(err)
			}
		}
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
