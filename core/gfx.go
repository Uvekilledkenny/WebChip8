package core

import (
	"image"
	"image/color"
	"image/draw"
)

var (
	screenHeight = 32
	screenWidth  = 64
	background   = color.NRGBA{0, 0, 0, 255}
	foreground   = color.NRGBA{65, 255, 0, 255} // P1 Phosphore
)

// GFX :
type gfx struct {
	Screen   *image.NRGBA
	DrawFlag bool
}

func blackScreen() *image.NRGBA {
	var Screen *image.NRGBA
	Screen = image.NewNRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	draw.Draw(Screen, Screen.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	return Screen
}

func newGFX() gfx {
	return gfx{
		Screen: blackScreen(),
	}
}

func (g *gfx) clearScreen() {
	g.Screen = blackScreen()
	g.DrawFlag = true
}

func (g *gfx) drawSprite(s []byte, x, y, n uint8) uint8 {
	var vf uint8
	for ay := 0; ay < int(n); ay++ {
		for ax := 0; ax < 8; ax++ {
			var (
				dx = int(x) + ax
				dy = int(y) + ay
				on = (s[ay]&(0x80>>uint(ax)) != 0)
			)

			if on {
				if g.Screen.At(dx, dy) == foreground {
					vf = 1
					g.Screen.SetNRGBA(dx, dy, background)
				} else {
					g.Screen.SetNRGBA(dx, dy, foreground)
					vf = 0
				}
			}
		}
	}
	g.DrawFlag = true
	return vf
}
