package main

func NewSoundEffectSystem(game *ShmupWarz) (this *SoundEffectSystem) {
	this = new(SoundEffectSystem)
	this.Game = game
	return this
}

type SoundEffectSystem struct {
	Game *ShmupWarz
}

func (this *SoundEffectSystem) Update(delta float64) {

}
