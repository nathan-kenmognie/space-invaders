package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand/v2"
	"strconv"
)

type collisionObject struct {
	pos    rl.Vector2
	Width  float32
	Height float32
	health int
}

type Asteroids struct {
	collisionObject
	color    rl.Color
	velocity rl.Vector2
}

type Enemies struct {
	list []Asteroids
}

type Player struct {
	collisionObject
	speed  float32
	points int
	angle  float32
	cargo  int
}

type Projectile struct {
	collisionObject
	angle float32
	speed float32
	color rl.Color
}

func (p *Player) fireProjectile() Projectile {

	projectileSprite := rl.LoadTexture("textures/projectile.png")
	projectileColor := rl.Purple

	projectile := Projectile{collisionObject: collisionObject{pos: rl.Vector2{X: p.pos.X, Y: p.pos.Y}, Width: float32(projectileSprite.Width) / 4, Height: float32(projectileSprite.Height) / 4, health: 0}, angle: p.angle, speed: 20, color: projectileColor}

	return projectile

}

func (proj *Projectile) projectileUpdate() {
	rad := float64(proj.angle * (math.Pi / 180))
	dx := math.Cos(rad) * float64(proj.speed)
	dy := math.Sin(rad) * float64(proj.speed)

	proj.pos.X += float32(dx)
	proj.pos.Y += float32(dy)

}

func (p *Player) rotate() {
	if rl.IsKeyDown(rl.KeyQ) || rl.IsKeyDown(rl.KeyA) {
		p.angle -= p.speed / 4
	}
	if rl.IsKeyDown(rl.KeyR) || rl.IsKeyDown(rl.KeyD) {
		p.angle += p.speed / 4
	}

}

func (p *Player) collisionDetection(obj *collisionObject) bool {
	scaledWidth := p.Width * 0.125
	scaledHeight := p.Height * 0.125

	playerLeft := p.pos.X - scaledWidth/2
	playerRight := p.pos.X + scaledWidth/2
	playerTop := p.pos.Y - scaledHeight/2
	playerBottom := p.pos.Y + scaledHeight/2

	objectLeft := obj.pos.X
	objectRight := obj.pos.X + obj.Width
	objectTop := obj.pos.Y
	objectBottom := obj.pos.Y + obj.Height

	if playerRight > objectLeft && playerLeft < objectRight &&
		playerBottom > objectTop && playerTop < objectBottom {
		return true
	}
	return false
}

func (p *Player) borderDetection() {
	window := rl.Vector2{X: 1500, Y: 1000}
	/*
		topBorder := collisionObject{pos: rl.Vector2{X: 0, Y: 0}, Width: window.X, Height: 1}
		bottomBorder := collisionObject{pos: rl.Vector2{X: 0, Y: window.Y}, Width: window.X, Height: window.Y}
		leftBorder := collisionObject{pos: rl.Vector2{X: -1, Y: 0}, Width: 1, Height: window.Y}
		rightBorder := collisionObject{pos: rl.Vector2{X: window.X, Y: 0}, Width: 1, Height: window.Y}
	*/
	minX := float32(0)
	maxX := window.X
	minY := float32(0)
	maxY := window.Y

	if p.pos.X < minX {
		p.pos.X = minX
	}
	if p.pos.X > maxX {
		p.pos.X = maxX
	}
	if p.pos.Y < minY {
		p.pos.Y = minY
	}
	if p.pos.Y > maxY {
		p.pos.Y = maxY
	}

}

func (p *Player) move() {
	rad := float64(p.angle * (math.Pi / 180))

	dx := math.Cos(rad) * float64(p.speed)

	dy := math.Sin(rad) * float64(p.speed)

	if rl.IsKeyDown(rl.KeyW) {
		p.pos.X += float32(dx)
		p.pos.Y += float32(dy)
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.pos.X -= float32(dx)
		p.pos.Y -= float32(dy)
	}
}

func randColor() rl.Color {
	num := rand.IntN(6)

	switch num {
	case 0:
		return rl.Lime
	case 1:
		return rl.Beige
	case 2:
		return rl.Blue
	case 3:
		return rl.Green
	case 4:
		return rl.Red
	case 5:
		return rl.Purple
	}
	return rl.Purple
}
func newAsteroids() Enemies {
	numAsteroids := 5
	var asteroids Enemies
	window := rl.Vector2{X: 1500, Y: 1000}

	planetSprite := rl.LoadTexture("textures/planet.png")
	planet := collisionObject{
		pos:    rl.NewVector2(window.X/2-float32(planetSprite.Width)/2, float32(planetSprite.Height)/2),
		Width:  float32(planetSprite.Width),
		Height: float32(planetSprite.Height),
	}

	for i := 0; i < numAsteroids; i++ {
		var newAsteroid Asteroids

		asteroidWidth := float32(rl.GetRandomValue(20, 30))
		asteroidHeight := float32(rl.GetRandomValue(20, 30))
		asteroidPos := rl.Vector2{X: float32(rl.GetRandomValue(0, int32(window.X))), Y: float32(rl.GetRandomValue(0, int32(window.Y)))}

		overlap := true
		for overlap {
			overlap = false
			planetLeft := planet.pos.X
			planetRight := planet.pos.X + planet.Width
			planetTop := planet.pos.Y
			planetBottom := planet.pos.Y + planet.Height

			asteroidLeft := asteroidPos.X
			asteroidRight := asteroidPos.X + asteroidWidth
			asteroidTop := asteroidPos.Y
			asteroidBottom := asteroidPos.Y + asteroidHeight

			if asteroidRight > planetLeft && asteroidLeft < planetRight &&
				asteroidBottom > planetTop && asteroidTop < planetBottom {
				asteroidPos = rl.Vector2{X: float32(rl.GetRandomValue(0, int32(window.X))), Y: float32(rl.GetRandomValue(0, int32(window.Y)))}
				overlap = true
			}
		}

		color := randColor()

		newAsteroid = Asteroids{collisionObject: collisionObject{pos: asteroidPos, Width: asteroidWidth, Height: asteroidHeight, health: 1}, color: color}
		asteroids.list = append(asteroids.list, newAsteroid)
	}

	return asteroids
}
func updateAsteroids(asteroids *Enemies, planet collisionObject) {

	window := rl.Vector2{X: 1500, Y: 1000}

	planetCenter := rl.Vector2{
		X: planet.pos.X + planet.Width/2,
		Y: planet.pos.Y + planet.Height/2,
	}
	for i := 0; i < len(asteroids.list); i++ {
		asteroids.list[i].pos.X += asteroids.list[i].velocity.X
		asteroids.list[i].pos.Y += asteroids.list[i].velocity.Y
	}

	for i := range asteroids.list {
		asteroid := &asteroids.list[i]

		asteroidCenter := rl.Vector2{
			X: asteroid.pos.X + asteroid.Width/2,
			Y: asteroid.pos.Y + asteroid.Height/2,
		}

		if asteroid.pos.X <= 0 || asteroid.pos.X >= window.X {
			asteroid.velocity.X *= -1
		}
		if asteroid.pos.Y <= 0 || asteroid.pos.Y >= window.Y {
			asteroid.velocity.Y *= -1
		}

		direction := rl.Vector2Subtract(planetCenter, asteroidCenter)
		distance := rl.Vector2Length(direction)

		if distance > 0 {
			direction = rl.Vector2Scale(rl.Vector2Normalize(direction), 0.25)
			asteroid.pos = rl.Vector2Add(asteroid.pos, direction)
		}
	}
}

func AsteroidPlanet(asteroid Asteroids, planet collisionObject) bool {
	asteroidLeft := asteroid.pos.X
	asteroidRight := asteroid.pos.X + asteroid.Width
	asteroidTop := asteroid.pos.Y
	asteroidBottom := asteroid.pos.Y + asteroid.Height

	planetLeft := planet.pos.X
	planetRight := planet.pos.X + planet.Width
	planetTop := planet.pos.Y
	planetBottom := planet.pos.Y + planet.Height

	if asteroidRight > planetLeft && asteroidLeft < planetRight &&
		asteroidBottom > planetTop && asteroidTop < planetBottom {
		return true
	}
	return false
}

func AsteroidProjectile(asteroid Asteroids, projectile Projectile) bool {
	asteroidLeft := asteroid.pos.X
	asteroidRight := asteroid.pos.X + asteroid.Width
	asteroidTop := asteroid.pos.Y
	asteroidBottom := asteroid.pos.Y + asteroid.Height

	projectileLeft := projectile.pos.X
	projectileRight := projectile.pos.X + projectile.Width
	projectileTop := projectile.pos.Y
	projectileBottom := projectile.pos.Y + projectile.Height

	if projectileRight > asteroidLeft && projectileLeft < asteroidRight &&
		projectileBottom > asteroidTop && projectileTop < asteroidBottom {
		return true
	}
	return false
}

func (p *collisionObject) gameOver() {
	rl.EndDrawing()
	paused := true
out:

	for paused {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Red)
		rl.DrawText("Game Over!\nPress R to Restart", 400, 500, 100, rl.Black)

		if rl.IsKeyDown(rl.KeyR) {
			p.health = 100

			paused = false

			break out

		}
		rl.EndDrawing()
	}

}
func spawnSmallerAsteroids(parent Asteroids, asteroids *Enemies) {
	if parent.Width/2 < 10 || parent.Height/2 < 10 {
		return
	}

	for i := 0; i < 2; i++ {
		angle := float32(rl.GetRandomValue(0, 360)) * (math.Pi / 180)
		speed := float32(rl.GetRandomValue(2, 5))

		newAsteroid := Asteroids{
			collisionObject: collisionObject{
				pos:    rl.Vector2{X: parent.pos.X, Y: parent.pos.Y},
				Width:  parent.Width / 2,
				Height: parent.Height / 2,
				health: 1,
			},
			color:    rl.White,
			velocity: rl.Vector2{X: float32(math.Cos(float64(angle))) * speed, Y: float32(math.Sin(float64(angle))) * speed},
		}

		newAsteroid.pos.X += newAsteroid.velocity.X * 5
		newAsteroid.pos.Y += newAsteroid.velocity.Y * 5

		asteroids.list = append(asteroids.list, newAsteroid)
	}
}

func removeOffScreenProjectiles(projectiles *[]Projectile, window rl.Vector2) {
	updatedProjectiles := (*projectiles)[:0]
	for _, proj := range *projectiles {
		if proj.pos.X >= 0 && proj.pos.X <= window.X && proj.pos.Y >= 0 && proj.pos.Y <= window.Y {
			updatedProjectiles = append(updatedProjectiles, proj)
		}
	}

	*projectiles = updatedProjectiles
}

func (p *Player) updateRotation() {
	mousePos := rl.GetMousePosition()

	dx := mousePos.X - p.pos.X
	dy := mousePos.Y - p.pos.Y

	p.angle = float32(math.Atan2(float64(dy), float64(dx)) * (180 / math.Pi))
}

func main() {
	window := rl.Vector2{X: 1500, Y: 1000}
	rl.SetTargetFPS(60)
	rl.InitWindow(int32(window.X), int32(window.Y), "Space Invaders")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	ambience := rl.LoadMusicStream("music/ambience.ogg")
	laser := rl.LoadSound("music/laser.mp3")
	explosion := rl.LoadSound("music/explosion.mp3")
	pickup := rl.LoadSound("music/pickup.mp3")
	heal := rl.LoadSound("music/heal.mp3")

	rl.PlayMusicStream(ambience)

	playerSprite := rl.LoadTexture("textures/spaceship.png")
	playerWidth := float32(playerSprite.Width)
	playerHeight := float32(playerSprite.Height)
	playerColor := rl.Purple
	player := Player{collisionObject{rl.NewVector2(200, 200), playerWidth, playerHeight, 1}, 10, 0, 0, 0}

	planetSprite := rl.LoadTexture("textures/planet.png")
	planetColor := rl.Purple
	planet := collisionObject{pos: rl.NewVector2(window.X/2-float32(planetSprite.Width)/2, float32(planetSprite.Height/2)), Width: float32(planetSprite.Width), Height: float32(planetSprite.Height), health: 100}
	// rl.DrawTextureEx(playerSprite, rl.NewVector2(200, 200), 0, 2, playerColor)
	projectileSprite := rl.LoadTexture("textures/projectile.jpg")
	asteroidSprite := rl.LoadTexture("textures/asteroid.png")
	var projectiles []Projectile
	obstacles := newAsteroids()
	mouseLook := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.UpdateMusicStream(ambience)

		player.rotate()
		player.move()
		player.collisionDetection(&planet)
		player.borderDetection()
		updateAsteroids(&obstacles, planet)
		removeOffScreenProjectiles(&projectiles, window)
		rl.DrawText("Health: "+strconv.Itoa(planet.health), 50, 25, 50, playerColor)
		rl.DrawText("Cargo: "+strconv.Itoa(player.cargo), 50, 55, 50, playerColor)

		if rl.IsKeyPressed(rl.KeyE) {
			mouseLook = true
		} else if rl.IsKeyPressed(rl.KeyT) {
			mouseLook = false
		}

		if mouseLook {
			player.updateRotation()
		}

	label1:
		for i := range obstacles.list {
			rl.DrawTextureEx(asteroidSprite, obstacles.list[i].pos, 0, 0.167, obstacles.list[i].color)
			if AsteroidPlanet(obstacles.list[i], planet) {
				planet.health -= 4
				obstacles.list = append(obstacles.list[:i], obstacles.list[i+1:]...)
				print(planet.health)
				goto label1
			}

		}
		for i := range projectiles {
			projectiles[i].projectileUpdate()
			src := rl.NewRectangle(0, 0, float32(projectileSprite.Width), float32(projectileSprite.Height))
			dest := rl.NewRectangle(projectiles[i].pos.X, projectiles[i].pos.Y, float32(projectileSprite.Width)*0.125, float32(projectileSprite.Height)*0.125)
			origin := rl.NewVector2(dest.Width/2, dest.Height/2)
			rl.DrawTexturePro(projectileSprite, src, dest, origin, projectiles[i].angle, projectiles[i].color)
			for j := len(obstacles.list) - 1; j >= 0; j-- {

				if AsteroidProjectile(obstacles.list[j], projectiles[i]) {
					spawnSmallerAsteroids(obstacles.list[j], &obstacles)
					obstacles.list = append(obstacles.list[:j], obstacles.list[j+1:]...)
					projectiles = append(projectiles[:i], projectiles[i+1:]...)
					rl.PlaySound(explosion)
					i--
					break
				}
			}
		}

		for i := len(obstacles.list) - 1; i >= 0; i-- {
			if player.collisionDetection(&obstacles.list[i].collisionObject) && obstacles.list[i].color == rl.White {
				player.cargo += 4
				obstacles.list = append(obstacles.list[:i], obstacles.list[i+1:]...)
				rl.PlaySound(pickup)
			}
		}
		if planet.health == 0 {
			planet.gameOver()
		}
		if player.collisionDetection(&planet) {
			planet.health += player.cargo
			if player.cargo != 0 {
				rl.PlaySound(heal)
			}

			player.cargo = 0
		}
		if len(obstacles.list) == 0 {
			obstacles = newAsteroids()
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			proj := player.fireProjectile()
			projectiles = append(projectiles, proj)
			fmt.Println("Projectile fired:", proj.pos)
			rl.PlaySound(laser)
		}

		src := rl.NewRectangle(0, 0, playerWidth, playerHeight)
		dest := rl.NewRectangle(player.pos.X, player.pos.Y, playerWidth*0.125, playerHeight*0.125)
		origin := rl.NewVector2(dest.Width/2, dest.Height/2)

		rl.DrawTexturePro(playerSprite, src, dest, origin, player.angle, playerColor)

		rl.DrawTexture(planetSprite, int32(planet.pos.X), int32(planet.pos.Y), planetColor)
		rl.EndDrawing()
	}

}
