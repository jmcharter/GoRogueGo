package main

// Map room settings
const (
	MIN_ROOM_SIZE = 5
	MAX_ROOM_SIZE = 12
	MIN_ROOM_QTY  = 5
	MAX_ROOM_QTY  = 10
)

type tile struct {
	blocksMovement bool
	blocksSight    bool
	highlight      bool
}

type gameMap struct {
	w      int
	h      int
	player *Entity
	tiles  [][]tile
}

func (gm *gameMap) init(pc *Entity) {
	gm.w = GRID_WIDTH
	gm.h = GRID_HEIGHT
	gm.player = pc

	gm.tiles = make([][]tile, GRID_HEIGHT)
	for i := range gm.tiles {
		gm.tiles[i] = make([]tile, GRID_WIDTH)

	}

	generateDungeon(gm)
	// Test room to be removed
	// room1 := CreateMapRectangle(2, 3, 10, 10)
	// room2 := CreateMapRectangle(22, 3, 10, 10)
	// gm.createRoom(room1)
	// gm.createRoom(room2)
	// gm.createHorizontalTunnel(room1.GetCentre().x, room2.GetCentre().x, room1.GetCentre().y)

}

func (gm *gameMap) createRoom(r mapRectangle) {
	for y := r.minY; y <= r.maxY; y++ {
		for x := r.minX; x <= r.maxX; x++ {
			gm.tiles[y][x].blocksMovement = false
		}
	}
}

func (gm *gameMap) createHorizontalTunnel(minX, maxX, y int) {
	for x := minX; x <= maxX; x++ {
		gm.tiles[y][x].blocksMovement = false
	}
}

func (gm *gameMap) createVerticalTunnel(minY, maxY, x int) {
	for y := minY; y <= maxY; y++ {
		gm.tiles[y][x].blocksMovement = false
	}
}

func (gm *gameMap) GetCentre() mapVec {
	centreX := gm.w / 2
	centreY := gm.h / 2

	return mapVec{centreX, centreY}
}
