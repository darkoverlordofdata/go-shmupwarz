package main

func NewViewManagerSystem(game *ShmupWarz) (this *ViewManagerSystem) {
	this = new(ViewManagerSystem)
	this.Game = game
	return this
}

type ViewManagerSystem struct {
	Game *ShmupWarz
}

func (this *ViewManagerSystem) Update(delta float64) {

}
