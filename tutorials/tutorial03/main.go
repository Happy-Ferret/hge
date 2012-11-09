package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/particle"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	"github.com/losinggeneration/hge-go/hge"
	"github.com/losinggeneration/hge-go/hge/gfx"
	"github.com/losinggeneration/hge-go/hge/input"
	"github.com/losinggeneration/hge-go/hge/sound"
	"github.com/losinggeneration/hge-go/hge/timer"
)

// An example of using closures
func main() {
	func() {
		const (
			speed    = 90.0
			friction = 0.98
		)

		var (
			spr sprite.Sprite
			fnt *font.Font
			par *particle.ParticleSystem

			snd *sound.Effect

			x  = 100.0
			y  = 100.0
			dx = 0.0
			dy = 0.0
		)

		h := hge.New()

		h.SetState(hge.LOGFILE, "tutorial03.log")
		h.SetState(hge.TITLE, "HGE Tutorial 03 - Using helper classes")
		h.SetState(hge.FPS, 100)
		h.SetState(hge.WINDOWED, true)
		h.SetState(hge.SCREENWIDTH, 800)
		h.SetState(hge.SCREENHEIGHT, 600)
		h.SetState(hge.SCREENBPP, 32)
		h.SetState(hge.FRAMEFUNC, func() int {
			dt := float64(timer.Delta())

			boom := func() {
				pan := int((x - 400) / 4)
				pitch := (dx*dx+dy*dy)*0.0005 + 0.2
				snd.PlayEx(100, pan, pitch)
			}

			// Process keys
			if input.K_ESCAPE.State() {
				return 1
			}
			if input.K_LEFT.State() {
				dx -= speed * dt
			}
			if input.K_RIGHT.State() {
				dx += speed * dt
			}
			if input.K_UP.State() {
				dy -= speed * dt
			}
			if input.K_DOWN.State() {
				dy += speed * dt
			}

			// Do some movement calculations and collision detection
			dx *= friction
			dy *= friction
			x += dx
			y += dy
			if x > 784 {
				x = 784 - (x - 784)
				dx = -dx
				boom()
			}
			if x < 16 {
				x = 16 + 16 - x
				dx = -dx
				boom()
			}
			if y > 584 {
				y = 584 - (y - 584)
				dy = -dy
				boom()
			}
			if y < 16 {
				y = 16 + 16 - y
				dy = -dy
				boom()
			}

			// Update particle system
			par.Info.Emission = (int)(dx*dx+dy*dy) * 2
			par.MoveTo(x, y)
			par.Update(dt)

			return 0
		})

		h.SetState(hge.RENDERFUNC, func() int {
			gfx.BeginScene()
			gfx.Clear(0)
			par.Render()
			spr.Render(x, y)
			fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", timer.Delta(), timer.FPS())
			gfx.EndScene()

			return 0
		})

		if err := h.Initiate(); err == nil {
			defer h.Shutdown()

			snd = sound.NewEffect("menu.ogg")
			tex := gfx.LoadTexture("particles.png")
			if snd == nil || tex == nil {
				fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
				return
			}

			spr = sprite.New(tex, 96, 64, 32, 32)
			spr.SetColor(0xFFFFA000)
			spr.SetHotSpot(16, 16)

			if fnt = font.New("font1.fnt"); fnt == nil {
				fmt.Println("Error loading font1.fnt")
				return
			}

			spt := sprite.New(tex, 32, 32, 32, 32)
			spt.SetBlendMode(gfx.BLEND_COLORMUL | gfx.BLEND_ALPHAADD | gfx.BLEND_NOZWRITE)
			spt.SetHotSpot(16, 16)

			par = particle.New("trail.psi", spt)
			if par == nil {
				fmt.Println("Error loading trail.psi")
				return
			}
			par.Fire()

			h.Start()
		}
	}()
}
