package main

import (
	"math/rand"
)

func generateDungeon(gm *gameMap) {
	var roomCount int
	var rooms []mapRectangle
	var currentRoom mapRectangle
	features := []string{"regularRoom"}
	roomQty := 10
	// Fill the map with solid blocks.
	fillMap(gm)

	// Dig out central starting room
	currentRoom = startingRect(gm)
	gm.createRoom(currentRoom)
	rooms = append(rooms, currentRoom)

roomGen:
	for roomCount < roomQty {

		// Pick a random room
		currentRoom = rooms[rand.Intn(len(rooms))]

		// Pick a random wall and tile
		randomWall, dir := currentRoom.GetRandomSide()
		currentTile := randomWall[rand.Intn(len(randomWall))]

		// Pick a random feature to build
		featureChoice := features[rand.Intn(len(features))]

		switch featureChoice {
		case "regularRoom":
			width := rand.Intn(MAX_ROOM_SIZE-MIN_ROOM_SIZE) + MIN_ROOM_SIZE
			height := rand.Intn(MAX_ROOM_SIZE-MIN_ROOM_SIZE) + MIN_ROOM_SIZE
			x := currentTile.x - (width / 2 * dir.y) + (2 * dir.x)
			y := currentTile.y - (height / 2 * dir.x) + (2 * dir.y)
			if dir.y < 0 {
				y = currentTile.y - height - 1
				x -= width - 1
			}
			if dir.x < 0 {
				x = currentTile.x - width - 1
				y -= height - 1
			}
			if x+width+1 > gm.w || y+height+1 > gm.h || x < 1 || y < 1 {
				continue
			}
			newRoom := CreateMapRectangle(x, y, width, height)

			for _, otherRoom := range rooms {
				roomsIntersect := newRoom.DoesIntersect(otherRoom)
				if roomsIntersect {
					continue roomGen
				}
			}

			doorTile := currentTile.Add(dir)
			door := CreateMapRectangle(doorTile.x, doorTile.y, 1, 1)

			gm.createRoom(newRoom)
			gm.createRoom(door)
			rooms = append(rooms, newRoom)
		}

		roomCount++

	}

	placePlayer(gm.player, rooms)

}

func fillMap(gm *gameMap) {
	for i := range gm.tiles {
		for j := range gm.tiles[i] {
			gm.tiles[i][j] = tile{blocksMovement: true}
		}
	}
}

func startingRect(gm *gameMap) mapRectangle {
	w := rand.Intn(7) + 5
	h := rand.Intn(7) + 5

	// Ensure odd numbers
	if !(w%2 == 0) {
		w += 1
	}
	if !(h%2 == 0) {
		h += 1
	}

	x := gm.GetCentre().x - (w / 2)
	y := gm.GetCentre().y - (h / 2)

	return CreateMapRectangle(x, y, w, h)

}

func placePlayer(player *Entity, rooms []mapRectangle) {
	randomRoom := rooms[rand.Intn(len(rooms))]
	player.x = randomRoom.GetCentre().x
	player.y = randomRoom.GetCentre().y
	player.updateGridPos()
}

// func randomRoom(x, y int) mapRectangle {
// 	width := rand.Intn(MAX_ROOM_SIZE-MIN_ROOM_SIZE) + MIN_ROOM_SIZE
// 	height := rand.Intn(MAX_ROOM_SIZE-MIN_ROOM_SIZE) + MIN_ROOM_SIZE

// 	return CreateMapRectangle(x, y, width, height)
// }
