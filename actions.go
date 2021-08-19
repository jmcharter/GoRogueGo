package main

// MoveAction will check that a move is valid for a given entity and if so, call
// that entity's Move method.
func MoveAction(dx, dy int, entity *Entity) {
	newX := entity.x + dx
	newY := entity.y + dy
	if !posInBound(newX, newY) {
		return
	}
	if entity.CheckCollision(dx, dy, entities) {
		return
	}
	entity.Move(dx, dy)

}

func posInBound(x, y int) bool {
	if 0 <= x && x < GRID_WIDTH &&
		0 <= y && y < GRID_HEIGHT {
		return true
	}
	return false
}
