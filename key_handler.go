package main

import "github.com/faiface/pixel/pixelgl"

func check_key_press(player *Entity, window *pixelgl.Window) {
	switch {
	case window.JustPressed(pixelgl.KeyLeft):
		player.Move(-1, 0)
	case window.JustPressed(pixelgl.KeyRight):
		player.Move(1, 0)
	case window.JustPressed(pixelgl.KeyUp):
		player.Move(0, 1)
	case window.JustPressed(pixelgl.KeyDown):
		player.Move(0, -1)
	}
}
