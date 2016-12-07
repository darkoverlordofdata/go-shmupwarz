package main

func NewEntitySpawningTimerSystem(game *ShmupWarz) (this *EntitySpawningTimerSystem) {
	this = new(EntitySpawningTimerSystem)
	this.Game = game
	return this
}

type EntitySpawningTimerSystem struct {
	Game *ShmupWarz
}

func (this *EntitySpawningTimerSystem) Update(delta float64) {

}
