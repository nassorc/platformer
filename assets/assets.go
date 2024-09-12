package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
  Excel FontName = "excel"

  // Animation State
  Idle = iota
  Walking
  Running
)

var (
  PlatformLevel Level
  Background *ebiten.Image
  //go:embed tiles/*
  LevelFS embed.FS
  //go:embed tiles/sprites/world_tileset.png
  TileSheetData []byte
  //go:embed dino/sheets/DinoSprites-vita.png
  DinoGreenData []byte
  //go:embed fonts/excel.ttf
  excelFont []byte

	fonts = map[FontName]font.Face{}
  TileSheet *ebiten.Image
  DinoGreenSheet *ebiten.Image

  DinoGreenIdle = Animation {
    Texture: nil,
    TileWidth: 24,
    TileHeight: 24,
    TotalFrames: 4,
    Sheet: Sheet{
      Frame{ 0, 120 },
      Frame{ 1, 120 },
      Frame{ 2, 120 },
      Frame{ 3, 120 },
    },
  }

  DinoGreenRun = Animation {
    Texture: nil,
    TileWidth: 24,
    TileHeight: 24,
    TotalFrames: 7,
    Sheet: Sheet{
      Frame{ 17, 120 },
      Frame{ 18, 120 },
      Frame{ 19, 120 },
      Frame{ 20, 120 },
      Frame{ 21, 120 },
      Frame{ 22, 120 },
      Frame{ 23, 120 },
    },
  }

  DinoGreenWallClimb = Animation {
    Texture: nil,
    TileWidth: 24,
    TileHeight: 24,
    TotalFrames: 1,
    Sheet: Sheet{
      Frame{ 12, 1 },
    },
  }

  DinoGreenJumpAsc = Animation {
    Texture: nil,
    TileWidth: 24,
    TileHeight: 24,
    TotalFrames: 1,
    Sheet: Sheet{
      Frame{ 12, 1 },
    },
  }

  DinoGreenJumpDesc = Animation {
    Texture: nil,
    TileWidth: 24,
    TileHeight: 24,
    TotalFrames: 1,
    Sheet: Sheet{
      Frame{ 17, 1 },
    },
  }
)

type Level struct {
  Background *ebiten.Image
  Objects []Object
}

type Object struct {
  ID uint32
  Class string
  X, Y, Width, Height float64
}

type FontName string

func (f FontName) Get() font.Face {
	return getFont(f)
}

func NewFrame(Cell, Duration int) Frame {
	return Frame{Cell, Duration}
}

type Frame struct{ 
  Cell int 
  Duration int
}

type Sheet []Frame

type Animation struct {
	Sheet       Sheet
	Texture     *ebiten.Image
	TotalFrames int
	TileWidth   int
	TileHeight  int
}

// Type
type Class int

// State
type State int

func MustLoadAssets() {
  MustLoadLevel()
  img, _, err := image.Decode(bytes.NewReader(TileSheetData))

  if err != nil {
    panic(err)
  }

  TileSheet = ebiten.NewImageFromImage(img)

  img, _, err = image.Decode(bytes.NewReader(DinoGreenData))

  if err != nil {
    panic(err)
  }

  DinoGreenSheet = ebiten.NewImageFromImage(img)
  DinoGreenIdle.Texture = DinoGreenSheet
  DinoGreenRun.Texture = DinoGreenSheet
  DinoGreenWallClimb.Texture = DinoGreenSheet
  DinoGreenJumpAsc.Texture = DinoGreenSheet
  DinoGreenJumpDesc.Texture = DinoGreenSheet

	LoadFont(Excel, excelFont)
} 

func MustLoadLevel() {
  gameMap, err := tiled.LoadFile("./assets/levels/platformer.tmx")
  if err != nil {
    fmt.Printf("error parsing map: %s", err.Error())
    os.Exit(2)
  }
  for _, obj := range gameMap.ObjectGroups[0].Objects {
    fmt.Println(obj)
  }

  renderer, err := render.NewRenderer(gameMap)
  if err != nil {
    fmt.Printf("map unsupported for rendering: %s", err.Error())
    os.Exit(2)
  }

  // level.Background
  if err = renderer.RenderVisibleLayers(); err != nil {
    fmt.Printf("layer unsupported for rendering: %s", err.Error())
    os.Exit(2)
  }
  PlatformLevel.Background = ebiten.NewImageFromImage(renderer.Result)

  for _, obj := range gameMap.ObjectGroups[0].Objects {
    PlatformLevel.Objects = append(PlatformLevel.Objects, Object{
      ID: obj.ID,
      Class: obj.Class,
      X: obj.X,
      Y: obj.Y,
      Width: obj.Width,
      Height: obj.Height,
    })
  }
}

func LoadFont(name FontName, ttf []byte) {
	fontData, _ := truetype.Parse(ttf)
	fonts[name] = truetype.NewFace(fontData, &truetype.Options{Size: 10})
}

func getFont(name FontName) font.Face {
	f, ok := fonts[name]
	if !ok {
		panic(fmt.Sprintf("Font %s not found", name))
	}
	return f
}
