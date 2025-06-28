# 🚀 Space Invaders (Raylib-Go)

A simple 2D space shooter game written in Go using the [raylib-go](https://github.com/gen2brain/raylib-go) library. Destroy asteroids, protect the planet, and collect cargo to repair the planet!

---

## 🎮 Features

- Player spaceship with WASD movement and mouse-based aiming
- Asteroids that split into smaller pieces when shot
- Planet with health that takes damage from asteroid collisions
- Collect cargo from asteroid fragments by flying over them
- Deliver cargo to the planet to restore its health
- Basic sound effects and background music
- Sprite-based rendering

---

## 🕹️ Controls

| Action        | Key/Mouse       |
|---------------|-----------------|
| Rotate ship   | Move mouse      |
| Thrust forward| `W`             |
| Thrust backward | `S`           |
| Fire projectile | `SPACE`       |
| Restart on Game Over | `R`     |

---

## 📦 Installation

### Requirements

- Go 1.21+
- [raylib-go](https://github.com/gen2brain/raylib-go) and raylib C library installed

### Clone and Run

```bash
git clone https://github.com/yourusername/space-invaders-go.git
cd space-invaders-go
go run main.go
