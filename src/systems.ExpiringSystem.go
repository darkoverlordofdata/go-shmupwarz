package main

func NewExpiringSystem(game *ShmupWarz) (this *ExpiringSystem) {
	this = new(ExpiringSystem)
	this.Game = game
	return this
}

type ExpiringSystem struct {
	Game *ShmupWarz
}

func (this *ExpiringSystem) Update(delta float64) {

}
