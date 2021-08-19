package main

type tile struct {
	blocksMovement bool
	blocksSight    bool
}

type gameMap struct {
	w     int
	h     int
	tiles [][]tile
}

func (gm *gameMap) init() {
	gm.w = GRID_WIDTH
	gm.h = GRID_HEIGHT

	gm.tiles = make([][]tile, GRID_HEIGHT)
	for i := range gm.tiles {
		gm.tiles[i] = make([]tile, GRID_WIDTH)
		for j := range gm.tiles[i] {
			gm.tiles[i][j] = tile{blocksMovement: true}
		}
	}

	// Test room to be removed
	room1 := CreateMapRectangle(2, 3, 10, 10)
	room2 := CreateMapRectangle(22, 3, 10, 10)
	gm.createRoom(room1)
	gm.createRoom(room2)
	gm.createHorizontalTunnel(room1.GetCentre().x, room2.GetCentre().x, room1.GetCentre().y)

}

func (gm *gameMap) createRoom(r mapRectangle) {
	for y := r.minY; y < r.maxY; y++ {
		for x := r.minX; x < r.maxX; x++ {
			gm.tiles[y][x].blocksMovement = false
		}
	}
}

func (gm *gameMap) createHorizontalTunnel(minX, maxX, y int) {
	for x := minX; x < maxX; x++ {
		gm.tiles[y][x].blocksMovement = false
	}
}

func (gm *gameMap) createVerticalTunnel(minY, maxY, x int) {
	for y := minY; y < maxY; y++ {
		gm.tiles[y][x].blocksMovement = false
	}
}

func (gm *gameMap) generateDungeon() {

}
