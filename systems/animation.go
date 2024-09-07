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
    ref := anim.Data
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

    
		frame := anim.Data.Sheet[anim.CurrentFrame]
    texture := anim.Data.Texture

    // the Animation type should handle this
		ops := &ebiten.DrawImageOptions{}
		// ops.GeoM.Translate(float64(obj.X)-float64(anim.Data.TileWidth/2), float64(obj.Y)-float64(anim.Data.TileHeight/2))
		ops.GeoM.Translate(float64(obj.X), float64(obj.Y))

    tw := anim.Data.TileWidth
    th := anim.Data.TileWidth
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
