package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	mix "github.com/veandco/go-sdl2/sdl_mixer"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// ShmupWarz SDL game structure
type ShmupWarz struct {
	Game
	Font       *ttf.Font
	Music      *mix.Music
	Sound      *mix.Chunk
	Background *sdl.Texture
	Entities   []*Entity
	Systems    []System
}

// NewShmupWarz Returns new shmupwarz
func NewShmupWarz(width int, height int, title string) (this *ShmupWarz) {
	this = new(ShmupWarz)
	this.Width = width
	this.Height = height
	this.Title = title
	this.Mode = mix.INIT_OGG
	return
}

// Initialize the game
// overrides the base class Initialize
// called by the game engine prior to start
// use to initialize SDL submodules
// func (this *ShmupWarz) Initialize() {
// 	this.Game.Initialize()
// 	return
// }

// LoadContent
// overrides the base class LoadContent
// called by the game engine prior to the game loop
// use to load resources
func (this *ShmupWarz) LoadContent() {

	// Load resources
	this.Font = this.LoadFont("assets/fonts/skranji.regular.ttf", 24)
	this.Music = this.LoadMusic("assets/music/frantic-gameplay.ogg")
	this.Sound = this.LoadSound("assets/sounds/click.wav")
	this.Background = this.LoadTexture("assets/images/BackdropBlackLittleSparkBlack.png")
	this.Music.Play(-1)

	this.Entities = []*Entity{this.CreatePlayer(0, 0)}

	this.Systems = []System{
		NewPlayerInputSystem(this),
		NewMovementSystem(this),
		NewSoundEffectSystem(this),
		NewCollisionSystem(this),
		NewEntitySpawningTimerSystem(this),
		NewColorTweenSystem(this),
		NewScaleTweenSystem(this),
		NewRemoveOffscreenShipsSystem(this),
		NewViewManagerSystem(this),
		NewExpiringSystem(this),
		NewDestroySystem(this),
		NewRenderSystem(this)}

	this.Game.Start()
}

// Update
// Implenents the abstract method Update
// game logic, physics, etc goes here
func (this *ShmupWarz) Update(delta float64) {

	t1 := float64(time.Now().UnixNano()) / 1000000.0

	l := len(this.Systems)
	for i := 0; i < l; i++ {
		this.Systems[i].Update(delta)
	}
	t2 := float64(time.Now().UnixNano()) / 1000000.0

	fmt.Printf("update %f\n", (t2 - t1))
}

// Draw
// Implenents the abstract method Draw
// do all the rendering
func (this *ShmupWarz) Draw(delta float64) {
	var e *Entity
	var tex *sdl.Texture
	var w, h int32
	t1 := float64(time.Now().UnixNano()) / 1000000.0

	this.Renderer.Clear()
	this.Renderer.Copy(this.Background, nil, nil)
	for i := 0; i < len(this.Entities); i++ {
		e = this.Entities[i]
		tex = e.Sprite.Texture

		_, _, w, h, _ = tex.Query()
		x := e.Position.X - w/2
		y := e.Position.Y - h/2
		this.Renderer.Copy(tex, nil, &sdl.Rect{X: x, Y: y, W: w, H: h})
	}
	this.Renderer.Present()
	t2 := float64(time.Now().UnixNano()) / 1000000.0

	fmt.Printf("draw %f\n", (t2 - t1))

}

// Destroy Destroys SDL and releases the memory
// overrides the base class Destroy
func (this *ShmupWarz) Destroy() {
	this.Font.Close()
	this.Music.Free()
	this.Sound.Free()

	ttf.Quit()
	mix.CloseAudio()
	mix.Quit()

	this.Game.Destroy()

}
