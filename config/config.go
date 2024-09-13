package config

type Config struct {
	Width  int
	Height int
  SpawnX float64
  SpawnY float64
}

var C *Config

func init() {
	C = &Config{
		Width:  640,
		Height: 360,
    SpawnX: 32,
    SpawnY: 128,
	}
}
