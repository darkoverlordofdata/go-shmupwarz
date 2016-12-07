package main

func NewScaleTweenSystem(game *ShmupWarz) (this *ScaleTweenSystem) {
	this = new(ScaleTweenSystem)
	this.Game = game
	return this
}

type ScaleTweenSystem struct {
	Game *ShmupWarz
}

func (this *ScaleTweenSystem) Update(delta float64) {

}
