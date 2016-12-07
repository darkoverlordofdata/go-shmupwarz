package main

func NewPlayerInputSystem(game *ShmupWarz) (this *PlayerInputSystem) {
	this = new(PlayerInputSystem)
	this.Game = game
	return this
}

type PlayerInputSystem struct {
	Game *ShmupWarz
}

func (this *PlayerInputSystem) Update(delta float64) {

}
