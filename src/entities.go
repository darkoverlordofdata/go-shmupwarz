package main

var nextEntityId = 0

const (
	LAYER_DEFAULT = iota
	LAYER_BACKGROUND
	LAYER_TEXT
	LAYER_LIVES
	LAYER_ENEMY1
	LAYER_ENEMY2
	LAYER_ENEMY3
	LAYER_PLAYER
	LAYER_BULLET
	LAYER_EXPLOSION
	LAYER_BANG
	LAYER_PARTICLE
	LAYER_HUD
)

const (
	EFFECT_PEW = iota
	EFFECT_ASPLODE
	EFFECT_SMALLASPLODE
)

type System interface {
	Update(delta float64)
}

// GetEntityId
func GetEntityId() int {
	nextEntityId++
	return nextEntityId
}

// Entity
type Entity struct {
	Id          int
	Name        string
	Active      bool
	Mask        uint64
	Bounds      BoundsComponent
	Bullet      BulletComponent
	ColorTween  ColorTweenComponent
	Destroy     DestroyComponent
	Enemy       EnemyComponent
	Expires     ExpiresComponent
	Health      HealthComponent
	Layer       LayerComponent
	Player      PlayerComponent
	Position    PositionComponent
	Scale       ScaleComponent
	Resource    ResourceComponent
	ScaleTween  ScaleTweenComponent
	Score       ScoreComponent
	SoundEffect SoundEffectComponent
	Sprite      SpriteComponent
	Tint        TintComponent
	Velocity    VelocityComponent
}

// CreateEntity
func CreateEntity(mask uint64) (e *Entity) {
	e = new(Entity)
	e.Id = GetEntityId()
	e.Mask = mask | C_LAYER | C_POSITION | C_RESOURCE
	return e
}

// CreatePlayer
func (this *ShmupWarz) CreatePlayer(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_PLAYER | C_BOUNDS | C_HEALTH)

	e.Layer = LayerComponent{Ordinal: LAYER_PLAYER}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/fighter.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Bounds = BoundsComponent{Radius: 45}
	e.Health = HealthComponent{MaxHealth: 100, MinHealth: 100}
	e.Player = PlayerComponent{Active: true}
	return e
}

// CreateBullet
func (this *ShmupWarz) CreateBullet(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_BULLET | C_BOUNDS | C_EXPIRES | C_TINT | C_VELOCITY)

	e.Layer = LayerComponent{Ordinal: LAYER_BULLET}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/bullet.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Bounds = BoundsComponent{Radius: 3}
	e.Expires = ExpiresComponent{Delay: 2}
	e.Health = HealthComponent{MaxHealth: 100, MinHealth: 100}
	e.Velocity = VelocityComponent{X: 0, Y: -800}
	e.Tint = TintComponent{R: 0xad, G: 0xff, B: 0x2f, A: 255}
	e.Bullet = BulletComponent{Active: true}
	return e
}

// CreateEnemy1
func (this *ShmupWarz) CreateEnemy1(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_ENEMY | C_BOUNDS | C_HEALTH | C_VELOCITY)

	e.Layer = LayerComponent{Ordinal: LAYER_ENEMY1}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/enemy1.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Bounds = BoundsComponent{Radius: 35}
	e.Health = HealthComponent{MaxHealth: 10, MinHealth: 10}
	e.Velocity = VelocityComponent{X: 0, Y: 40}
	e.Enemy.Active = true
	return e
}

// CreateEnemy2
func (this *ShmupWarz) CreateEnemy2(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_ENEMY | C_BOUNDS | C_HEALTH | C_VELOCITY)

	e.Layer = LayerComponent{Ordinal: LAYER_ENEMY2}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/enemy2.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Bounds = BoundsComponent{Radius: 86}
	e.Health = HealthComponent{MaxHealth: 20, MinHealth: 20}
	e.Velocity = VelocityComponent{X: 0, Y: 30}
	e.Enemy.Active = true
	return e
}

// CreateEnemy3
func (this *ShmupWarz) CreateEnemy3(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_ENEMY | C_BOUNDS | C_HEALTH | C_VELOCITY)

	e.Layer = LayerComponent{Ordinal: LAYER_ENEMY3}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/enemy3.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Bounds = BoundsComponent{Radius: 100}
	e.Health = HealthComponent{MaxHealth: 40, MinHealth: 40}
	e.Velocity = VelocityComponent{X: 0, Y: 20}
	e.Enemy.Active = true
	return e
}

// CreateExplosion
func (this *ShmupWarz) CreateExplosion(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_EXPIRES | C_SCALE | C_SCALETWEEN)

	e.Layer = LayerComponent{Ordinal: LAYER_EXPLOSION}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/explosion.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Expires = ExpiresComponent{Delay: 0.5}
	e.Scale = ScaleComponent{X: 1, Y: 1}
	e.ScaleTween = ScaleTweenComponent{Min: .001, Max: 1, Speed: -3, Repeat: false, Active: true}
	return e
}

// CreateBang
func (this *ShmupWarz) CreateBang(x float64, y float64) (e *Entity) {
	e = CreateEntity(C_EXPIRES | C_SCALE | C_SCALETWEEN)

	e.Layer = LayerComponent{Ordinal: LAYER_EXPLOSION}
	e.Position = PositionComponent{X: x, Y: y}
	e.Resource = ResourceComponent{Name: "assets/images/bang.png"}
	e.Sprite = SpriteComponent{Texture: this.LoadTexture(e.Resource.Name)}
	e.Expires = ExpiresComponent{Delay: 0.5}
	e.Scale = ScaleComponent{X: 1, Y: 1}
	e.ScaleTween = ScaleTweenComponent{Min: .001, Max: 1, Speed: -3, Repeat: false, Active: true}
	return e
}
