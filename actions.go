package main

func MoveAction(dx, dy int, entity *Entity) {
	new_x := entity.x + dx
	new_y := entity.y + dy
	if !posInBound(new_x, new_y) {
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
