package canvas

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

var square = ebiten.NewImage(28, 28)

func (g *Game) Update() error {
	// fmt.Println(ebiten.CursorPosition())

	x, y := ebiten.CursorPosition()

	if x > 0 && y > 0 && x < 320 && y < 240 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			g.Draw(square)
		}
	}

	return nil
}

func d() {
	x, y := ebiten.CursorPosition()
	vector.DrawFilledRect(square, float32(x), float32(y), 28, 28, color.White, false)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "What is happening")
	d()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Test() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
