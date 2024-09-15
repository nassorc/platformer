package systems

import (
	"image/color"

	"platformer/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func DrawDebug(ecs *ecs.ECS, screen *ebiten.Image) {
	const strokeWidth = 1
	settings := GetOrCreateSettings(ecs)
	if !settings.Debug {
		return
	}
	spaceEntry, ok := components.Space.First(ecs.World)
	if !ok {
		return
	}
	space := components.Space.Get(spaceEntry)

	for y := 0; y < space.Height(); y++ {

		for x := 0; x < space.Width(); x++ {

			cell := space.Cell(x, y)

			cw := float64(space.CellWidth)
			ch := float64(space.CellHeight)
			cx := float64(cell.X) * cw
			cy := float64(cell.Y) * ch

			drawColor := color.RGBA{20, 20, 20, 255}

			if cell.Occupied() {
				drawColor = color.RGBA{255, 255, 0, 255}
			}

			vector.StrokeLine(screen, float32(cx), float32(cy), float32(cx+cw), float32(cy), strokeWidth, drawColor, false)

			vector.StrokeLine(screen, float32(cx+cw), float32(cy), float32(cx+cw), float32(cy+ch), strokeWidth, drawColor, false)

			vector.StrokeLine(screen, float32(cx+cw), float32(cy+ch), float32(cx), float32(cy+ch), strokeWidth, drawColor, false)

			vector.StrokeLine(screen, float32(cx), float32(cy+ch), float32(cx), float32(cy), strokeWidth, drawColor, false)
		}

	}
}
