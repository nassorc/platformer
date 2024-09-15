package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type PlayerData struct {
	SpeedX         float64
	SpeedY         float64
	JumpBuffer     int
	CoyoteTimer    int
	OnGround       *resolv.Object
	WallSliding    *resolv.Object
	IgnorePlatform *resolv.Object
	FacingRight    bool
	Jumping        bool
	CanCoyote      bool
}

var Player = donburi.NewComponentType[PlayerData]()

func MustFindPlayer(world donburi.World) *PlayerData {
	entry, ok := Player.First(world)
	if !ok {
		panic("no player found")
	}
	return Player.Get(entry)
}
