package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var imgX int
var imgY int

func init() {
	var err error
	imgX = 480
	imgY = 720
	img, _, err = ebitenutil.NewImageFromFile("images/1p.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	imgX, imgY = img.Size()
	return imgX, imgY
}

func main() {
	imgX, imgY = img.Size()
	ebiten.SetWindowSize(imgX, imgY)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
