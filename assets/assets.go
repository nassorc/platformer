package assets

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type FontName string

const (
  Excel FontName = "excel"
)

func (f FontName) Get() font.Face {
	return getFont(f)
}

func NewFrame(Cell, Duration int) Frame {
	return Frame{Cell, Duration}
}

type Frame struct{ Cell, Duration int }

type Sheet []Frame

type Animation struct {
	Sheet       Sheet
	Texture     *ebiten.Image
	TotalFrames int
	TileWidth   int
	TileHeight  int
}

var (
  //go:embed dino/sheets/DinoSprites-vita.png
  DinoGreenData []byte

  //go:embed fonts/excel.ttf
  excelFont []byte
)

var (
	fonts = map[FontName]font.Face{}

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
)

func MustLoadAssets() {
  img, _, err := image.Decode(bytes.NewReader(DinoGreenData))
  if err != nil {
    panic(err)
  }

  DinoGreenSheet = ebiten.NewImageFromImage(img)
  DinoGreenIdle.Texture = DinoGreenSheet
  DinoGreenRun.Texture = DinoGreenSheet

	LoadFont(Excel, excelFont)
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
