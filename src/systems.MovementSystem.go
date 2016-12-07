package main

func NewMovementSystem(game *ShmupWarz) (this *MovementSystem) {
	this = new(MovementSystem)
	this.Game = game
	return this
}

type MovementSystem struct {
	Game *ShmupWarz
}

func (this *MovementSystem) Update(delta float64) {

}
