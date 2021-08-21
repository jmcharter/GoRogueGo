package main

import "github.com/faiface/pixel/pixelgl"

func checkKeyPress(player *Entity, gm gameMap, window *pixelgl.Window) {
	switch {
	case window.JustPressed(pixelgl.KeyLeft):
		MoveAction(-1, 0, player, gm)
	case window.JustPressed(pixelgl.KeyRight):
		MoveAction(1, 0, player, gm)
	case window.JustPressed(pixelgl.KeyUp):
		MoveAction(0, 1, player, gm)
	case window.JustPressed(pixelgl.KeyDown):
		MoveAction(0, -1, player, gm)
	case window.JustPressed(pixelgl.KeySpace):
		generateDungeon(&gm)
	}
}
