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
