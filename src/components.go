package main

import "github.com/veandco/go-sdl2/sdl"

const (
	C_BOUNDS      = 1      //2**0
	C_BULLET      = 2      //2**1
	C_COLORTWEEN  = 4      //2**2
	C_DESTROY     = 8      //2**3
	C_ENEMY       = 16     //2**4
	C_EXPIRES     = 32     //2**5
	C_HEALTH      = 64     //2**6
	C_LAYER       = 128    //2**7
	C_PLAYER      = 256    //2**8
	C_POSITION    = 512    //2**9
	C_SCALE       = 1024   //2**10
	C_RESOURCE    = 2048   //2**11
	C_SCALETWEEN  = 4096   //2**12
	C_SCORE       = 8192   //2**13
	C_SOUNDEFFECT = 16384  //2**14
	C_SPRITE      = 32769  //2**15
	C_TINT        = 65536  //2**16
	C_VELOCITY    = 131072 //2**17
)

const (
	ComponentBounds = iota
	ComponentBullet
	ComponentColorTween
	ComponentDestroy
	ComponentEnemy
	ComponentExpires
	ComponentHealth
	ComponentLayer
	ComponentPlayer
	ComponentPosition
	ComponentScale
	ComponentResource
	ComponentScaleTween
	ComponentScore
	ComponentSoundEffect
	ComponentSprite
	ComponentTint
	ComponentVelocity
	ComponentTotal
)

var ComponentType = map[int]string{
	ComponentBounds:      "Bounds",
	ComponentBullet:      "Bullet",
	ComponentColorTween:  "ColorTween",
	ComponentDestroy:     "Destroy",
	ComponentEnemy:       "Enemy",
	ComponentExpires:     "Expires",
	ComponentHealth:      "Health",
	ComponentLayer:       "Layer",
	ComponentPlayer:      "Player",
	ComponentPosition:    "Position",
	ComponentScale:       "Scale",
	ComponentResource:    "Resource",
	ComponentScaleTween:  "ScaleTween",
	ComponentScore:       "Score",
	ComponentSoundEffect: "SoundEffect",
	ComponentSprite:      "Sprite",
	ComponentTint:        "Tint",
	ComponentVelocity:    "Velocity",
}

type AbstractComponent struct{}

type BoundsComponent struct {
	AbstractComponent
	Radius float64
}
type BulletComponent struct {
	AbstractComponent
	Active bool
}
type ColorTweenComponent struct {
	AbstractComponent
	RedMin       float64
	RedMax       float64
	RedSpeed     float64
	GreenMin     float64
	GreenMax     float64
	GreenSpeed   float64
	BlueMin      float64
	BlueMax      float64
	BlueSpeed    float64
	AlphaMin     float64
	AlphaMax     float64
	AlphaSpeed   float64
	RedAnimate   bool
	GreenAnimate bool
	BlueAnimate  bool
	AlphaAnimate bool
	Repeat       bool
}
type DestroyComponent struct {
	AbstractComponent
	Active bool
}
type EnemyComponent struct {
	AbstractComponent
	Active bool
}
type ExpiresComponent struct {
	AbstractComponent
	Delay float64
}
type HealthComponent struct {
	AbstractComponent
	MinHealth float64
	MaxHealth float64
}
type LayerComponent struct {
	AbstractComponent
	Ordinal int
}
type PlayerComponent struct {
	AbstractComponent
	Active bool
}
type PositionComponent struct {
	AbstractComponent
	X float64
	Y float64
}
type ResourceComponent struct {
	AbstractComponent
	Name string
}
type ScaleComponent struct {
	AbstractComponent
	X float64
	Y float64
}
type ScaleTweenComponent struct {
	AbstractComponent
	Min    float64
	Max    float64
	Speed  float64
	Repeat bool
	Active bool
}
type ScoreComponent struct {
	AbstractComponent
	Value int
}
type SoundEffectComponent struct {
	AbstractComponent
	Effect int
}
type SpriteComponent struct {
	AbstractComponent
	Texture *sdl.Texture
}
type TintComponent struct {
	AbstractComponent
	R byte
	G byte
	B byte
	A byte
}
type VelocityComponent struct {
	AbstractComponent
	X float64
	Y float64
}
