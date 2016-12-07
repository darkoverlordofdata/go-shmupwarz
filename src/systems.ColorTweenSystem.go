package main

func NewColorTweenSystem(game *ShmupWarz) (this *ColorTweenSystem) {
	this = new(ColorTweenSystem)
	this.Game = game
	return this
}

type ColorTweenSystem struct {
	Game *ShmupWarz
}

func (this *ColorTweenSystem) Update(delta float64) {

}
