package main

func NewDestroySystem(game *ShmupWarz) (this *DestroySystem) {
	this = new(DestroySystem)
	this.Game = game
	return this
}

type DestroySystem struct {
	Game *ShmupWarz
}

func (this *DestroySystem) Update(delta float64) {

}
