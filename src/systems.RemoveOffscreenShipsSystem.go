package main

func NewRemoveOffscreenShipsSystem(game *ShmupWarz) (this *RemoveOffscreenShipsSystem) {
	this = new(RemoveOffscreenShipsSystem)
	this.Game = game
	return this
}

type RemoveOffscreenShipsSystem struct {
	Game *ShmupWarz
}

func (this *RemoveOffscreenShipsSystem) Update(delta float64) {

}
