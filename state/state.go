package state

type Player struct {
	X int
	Y int
}

type World struct {
	Player *Player
}

func NewWorld(screenWidth, screenHeight int) *World {
	w := &World{&Player{screenWidth / 2, screenHeight / 2}}
	return w
}
