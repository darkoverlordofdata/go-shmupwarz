package main

import (
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
func (this *ShmupWarz) Initialize() {
	this.Game.Initialize()

	return
}

// Start
// overrides the base class start
// called by the game engine prior to the game loop
// use to load resources
func (this *ShmupWarz) Start() {

	// Load resources
	this.Font = this.LoadFont("assets/fonts/skranji.regular.ttf", 24)
	this.Music = this.LoadMusic("assets/music/frantic-gameplay.ogg")
	this.Sound = this.LoadSound("assets/sounds/click.wav")
	this.Background = this.LoadTexture("assets/images/BackdropBlackLittleSparkBlack.png")
	this.Music.Play(-1)

	this.Systems = append(this.Systems, NewMovementSystem(this))
	this.Systems = append(this.Systems, NewPlayerInputSystem(this))
	this.Systems = append(this.Systems, NewSoundEffectSystem(this))
	this.Systems = append(this.Systems, NewCollisionSystem(this))
	this.Systems = append(this.Systems, NewEntitySpawningTimerSystem(this))
	this.Systems = append(this.Systems, NewColorTweenSystem(this))
	this.Systems = append(this.Systems, NewScaleTweenSystem(this))
	this.Systems = append(this.Systems, NewRemoveOffscreenShipsSystem(this))
	this.Systems = append(this.Systems, NewViewManagerSystem(this))
	this.Systems = append(this.Systems, NewExpiringSystem(this))
	this.Systems = append(this.Systems, NewDestroySystem(this))
	this.Systems = append(this.Systems, NewRenderSystem(this))

	this.Game.Start()
}

// OnEvent
func (this *ShmupWarz) OnEvent(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		this.Game.Quit()

	case *sdl.MouseButtonEvent:
		this.Sound.Play(2, 0)
		// if t.Type == sdl.MOUSEBUTTONDOWN && t.Button == sdl.BUTTON_LEFT {
		// }

	case *sdl.KeyDownEvent:
		if t.Keysym.Scancode == sdl.SCANCODE_ESCAPE || t.Keysym.Scancode == sdl.SCANCODE_AC_BACK {
			this.Game.Quit()
		}
	}

}

// Update
// Implenents the abstract method Update
// game logic, physics, etc goes here
func (this *ShmupWarz) Update(delta float64) {
	for i := 0; i < len(this.Systems); i++ {
		this.Systems[i].Update(delta)
	}
}

// Draw
// Implenents the abstract method Draw
// do all the rendering
func (this *ShmupWarz) Draw(delta float64) {
	this.Renderer.Clear()
	this.Renderer.Copy(this.Background, nil, nil)
	this.Renderer.Present()

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
