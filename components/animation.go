package components

import (
	"image"
	"platformer/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nassorc/go-codebase/lib/math"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
  CurrentFrame int
  LastUpdateTime float32
  Offset math.Vec2
  CurrentState string
  StateMap map[string]*assets.Animation
}

func (anim *AnimationData) NextFrame(dt float32) {
  ref := anim.StateMap[anim.CurrentState]
  now := dt / 60 * 1000 // ms
  lastUpdate := anim.LastUpdateTime
  curFrame := anim.CurrentFrame
  animDuration := ref.Sheet[curFrame].Duration

  if now-lastUpdate >= float32(animDuration) {
    anim.CurrentFrame = (anim.CurrentFrame+1) % ref.TotalFrames
    anim.LastUpdateTime = now
  }
}

func (anim *AnimationData) CurrentSprite() *ebiten.Image {
  frame := anim.StateMap[anim.CurrentState].Sheet[anim.CurrentFrame]
  texture := anim.StateMap[anim.CurrentState].Texture

  tw := anim.StateMap[anim.CurrentState].TileWidth
  th := anim.StateMap[anim.CurrentState].TileWidth
  x := frame.Cell % (texture.Bounds().Dx() / tw)
  y := frame.Cell / (texture.Bounds().Dx() / tw)

  return texture.SubImage(
    image.Rect(
      x*tw, 
      y*th, 
      x*tw+tw, 
      y*th+th,
    ),
  ).(*ebiten.Image)
}

func (anim *AnimationData) MustChangeState(state string) {
  if _, ok := anim.StateMap[state]; !ok {
    panic("Animation state not found")
  }
  if state == anim.CurrentState {
    return
  }
  anim.CurrentState = state
  anim.CurrentFrame = 0
  anim.LastUpdateTime = 0
}

var Animation = donburi.NewComponentType[AnimationData]() 
