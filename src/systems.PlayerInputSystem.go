package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayerInputSystem(game *ShmupWarz) (this *PlayerInputSystem) {
	this = new(PlayerInputSystem)
	this.Game = game
	this.Keys = sdl.GetKeyboardState()
	this.Player = game.Entities[0]

	return this
}

type PlayerInputSystem struct {
	Game   *ShmupWarz
	Keys   []uint8
	Player *Entity
}

func (this *PlayerInputSystem) Update(delta float64) {

	x, y, _ := sdl.GetMouseState()

	this.Player.Position.X = int32(x)
	this.Player.Position.Y = int32(y)

	if this.Keys[sdl.SCANCODE_ESCAPE] != 0 {
		this.Game.Quit()
	}

}
