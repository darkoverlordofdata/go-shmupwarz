package main

func NewCollisionSystem(game *ShmupWarz) (this *CollisionSystem) {
	this = new(CollisionSystem)
	this.Game = game
	return this
}

type CollisionSystem struct {
	Game *ShmupWarz
}

func (this *CollisionSystem) Update(delta float64) {

}
