package main


func NewRenderSystem(game *ShmupWarz) (this *RenderSystem) {
	this = new(RenderSystem)
	this.Game = game
	return this
}



type RenderSystem struct {
	Game *ShmupWarz
}

func (this *RenderSystem) Update(delta float64) {

}
