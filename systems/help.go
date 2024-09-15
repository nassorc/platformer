package systems

import (
	"fmt"
	"image/color"
	"platformer/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func DrawHelp(ecs *ecs.ECS, screen *ebiten.Image) {
	settings := GetOrCreateSettings(ecs)
	if settings.ShowHelpText {
		drawText(screen, 16, 16,
			"~ Platformer Demo ~",
			"Move Player: Left, Right Arrow",
			"Jump: X Key",
			"Wallslide: Move into wall in air",
			"Walljump: Jump while wallsliding",
			"Fall through platforms: Down + X",
			"",
			"F1: Toggle Debug View",
			"F2: Show / Hide help text",
			fmt.Sprintf("%d FPS (frames per second)", int(ebiten.ActualFPS())),
			fmt.Sprintf("%d TPS (ticks per second)", int(ebiten.ActualFPS())),
		)
	}
}

func drawText(screen *ebiten.Image, x, y int, textLines ...string) {
	rectHeight := 10
	face := assets.Excel.GetFontFace(10)

	for _, txt := range textLines {
		w, _ := text.Measure(txt, face, 1)
		vector.DrawFilledRect(screen, float32(x), float32(y-8), float32(w), float32(rectHeight), color.RGBA{0, 0, 0, 192}, false)

		opt := &text.DrawOptions{}
		opt.GeoM.Translate(float64(x+1), float64(y+1))
		text.Draw(screen, txt, face, opt)
		y += rectHeight
	}
}
