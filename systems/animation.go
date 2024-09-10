package systems

import (
	"image"
	"platformer/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Animation struct {
  tick int
}

func NewAnimation() Animation {
  return Animation{
    tick: 0,
  }
}

func (sys *Animation) UpdateAnimation(ecs *ecs.ECS) {
  components.Animation.Each(ecs.World, func(e *donburi.Entry) {
    anim := components.Animation.Get(e)
    ref := anim.StateMap[anim.CurrentState]
		now := float32(sys.tick) / 60 * 1000
    lastUpdate := anim.LastUpdateTime
    curFrame := anim.CurrentFrame
    animDuration := ref.Sheet[curFrame].Duration

    if now-lastUpdate >= float32(animDuration) {
      anim.CurrentFrame = (anim.CurrentFrame+1) % ref.TotalFrames
      anim.LastUpdateTime = now
    }
  })
  sys.tick += 1
}

func DrawAnimation(ecs *ecs.ECS, screen *ebiten.Image) {
  query := donburi.NewQuery(filter.Contains(components.Object, components.Animation))
  query.Each(ecs.World, func(e *donburi.Entry) {
    // extract animation and position
    anim := components.Animation.Get(e)
    obj := components.Object.Get(e)

		frame := anim.StateMap[anim.CurrentState].Sheet[anim.CurrentFrame]
    texture := anim.StateMap[anim.CurrentState].Texture
    offset := anim.Offset

    // the Animation type should handle this
		ops := &ebiten.DrawImageOptions{}
		// ops.GeoM.Translate(float64(obj.X)-float64(anim.Data.TileWidth/2), float64(obj.Y)-float64(anim.Data.TileHeight/2))

    // obj.

    // Will only work with the player entity
    player := components.MustFindPlayer(ecs.World)
    scaleX := 1

    if !player.FacingRight {
      scaleX = -1
    }


    ops.GeoM.Scale(float64(scaleX), 1)

    if scaleX < 0 {
      ops.GeoM.Translate(float64(24), 0)
    }

		ops.GeoM.Translate(float64(obj.X)+float64(offset.X), float64(obj.Y)+float64(offset.Y))

    tw := anim.StateMap[anim.CurrentState].TileWidth
    th := anim.StateMap[anim.CurrentState].TileWidth
		x := frame.Cell % (texture.Bounds().Dx() / tw)
		y := frame.Cell / (texture.Bounds().Dx() / tw)

		screen.DrawImage(
      texture.SubImage(
        image.Rect(
          x*tw, 
          y*th, 
          x*tw+tw, 
          y*th+th,
        ),
      ).(*ebiten.Image),
      ops,
    )
  })
}
