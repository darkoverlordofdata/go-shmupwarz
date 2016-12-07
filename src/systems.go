package main

type System interface {
	Update(delta float64)
}

func NewMovementSystem(game *ShmupWarz) (this *MovementSystem) {
	this = new(MovementSystem)
	this.Game = game
	return this
}
func NewPlayerInputSystem(game *ShmupWarz) (this *PlayerInputSystem) {
	this = new(PlayerInputSystem)
	this.Game = game
	return this
}
func NewSoundEffectSystem(game *ShmupWarz) (this *SoundEffectSystem) {
	this = new(SoundEffectSystem)
	this.Game = game
	return this
}
func NewCollisionSystem(game *ShmupWarz) (this *CollisionSystem) {
	this = new(CollisionSystem)
	this.Game = game
	return this
}
func NewEntitySpawningTimerSystem(game *ShmupWarz) (this *EntitySpawningTimerSystem) {
	this = new(EntitySpawningTimerSystem)
	this.Game = game
	return this
}
func NewColorTweenSystem(game *ShmupWarz) (this *ColorTweenSystem) {
	this = new(ColorTweenSystem)
	this.Game = game
	return this
}
func NewScaleTweenSystem(game *ShmupWarz) (this *ScaleTweenSystem) {
	this = new(ScaleTweenSystem)
	this.Game = game
	return this
}
func NewRemoveOffscreenShipsSystem(game *ShmupWarz) (this *RemoveOffscreenShipsSystem) {
	this = new(RemoveOffscreenShipsSystem)
	this.Game = game
	return this
}
func NewViewManagerSystem(game *ShmupWarz) (this *ViewManagerSystem) {
	this = new(ViewManagerSystem)
	this.Game = game
	return this
}
func NewExpiringSystem(game *ShmupWarz) (this *ExpiringSystem) {
	this = new(ExpiringSystem)
	this.Game = game
	return this
}
func NewDestroySystem(game *ShmupWarz) (this *DestroySystem) {
	this = new(DestroySystem)
	this.Game = game
	return this
}
func NewRenderSystem(game *ShmupWarz) (this *RenderSystem) {
	this = new(RenderSystem)
	this.Game = game
	return this
}

type MovementSystem struct {
	Game *ShmupWarz
}

func (this *MovementSystem) Update(delta float64) {

}

type PlayerInputSystem struct {
	Game *ShmupWarz
}

func (this *PlayerInputSystem) Update(delta float64) {

}

type SoundEffectSystem struct {
	Game *ShmupWarz
}

func (this *SoundEffectSystem) Update(delta float64) {

}

type CollisionSystem struct {
	Game *ShmupWarz
}

func (this *CollisionSystem) Update(delta float64) {

}

type EntitySpawningTimerSystem struct {
	Game *ShmupWarz
}

func (this *EntitySpawningTimerSystem) Update(delta float64) {

}

type ColorTweenSystem struct {
	Game *ShmupWarz
}

func (this *ColorTweenSystem) Update(delta float64) {

}

type ScaleTweenSystem struct {
	Game *ShmupWarz
}

func (this *ScaleTweenSystem) Update(delta float64) {

}

type RemoveOffscreenShipsSystem struct {
	Game *ShmupWarz
}

func (this *RemoveOffscreenShipsSystem) Update(delta float64) {

}

type ViewManagerSystem struct {
	Game *ShmupWarz
}

func (this *ViewManagerSystem) Update(delta float64) {

}

type ExpiringSystem struct {
	Game *ShmupWarz
}

func (this *ExpiringSystem) Update(delta float64) {

}

type DestroySystem struct {
	Game *ShmupWarz
}

func (this *DestroySystem) Update(delta float64) {

}

type RenderSystem struct {
	Game *ShmupWarz
}

func (this *RenderSystem) Update(delta float64) {

}
