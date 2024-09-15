package components

import "github.com/yohamta/donburi"

type SettingsData struct {
	Debug        bool
	ShowHelpText bool
}

var Settings = donburi.NewComponentType[SettingsData]()

func MustFindSettings(world donburi.World) *SettingsData {
	entry := Settings.MustFirst(world)
	return Settings.Get(entry)
}
