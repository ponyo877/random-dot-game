package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 1000
	screenHeight = 1000
	numPoints    = 1000
	radius       = 10
)

type Point struct {
	x, y int
}

type Game struct {
	points         []Point
	showPoints     bool
	clickedX       int
	clickedY       int
	clicked        bool
	pointsInCircle int
}

func (g *Game) Update() error {
	if !g.showPoints {
		time.Sleep(1 * time.Second)
		g.showPoints = true
	}

	if g.showPoints && !g.clicked && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.clickedX, g.clickedY = ebiten.CursorPosition()
		g.clicked = true
		g.pointsInCircle = g.countPointsInCircle(g.clickedX, g.clickedY)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.showPoints {
		for _, p := range g.points {
			vector.DrawFilledRect(screen, float32(p.x), float32(p.y), 2, 2, color.White, true)
		}
	}

	if g.clicked {
		vector.DrawFilledRect(screen, float32(g.clickedX), float32(g.clickedY), 2, 2, color.RGBA{255, 0, 0, 255}, true)
		msg := fmt.Sprintf("Points in circle: %d", g.pointsInCircle)
		ebitenutil.DebugPrint(screen, msg)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) countPointsInCircle(x, y int) int {
	count := 0
	for _, p := range g.points {
		if math.Sqrt(math.Pow(float64(p.x-x), 2)+math.Pow(float64(p.y-y), 2)) <= radius {
			count++
		}
	}
	return count
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	points := make([]Point, numPoints)
	// for i := 0; i < numPoints; i++ {
	// 	points[i] = Point{r.Intn(screenWidth), r.Intn(screenHeight)}
	// }
	// meanX, meanY := screenWidth/2, screenHeight/2
	// stdDev := 50.0

	for i := 0; i < numPoints; i++ {
		meanX, meanY := r.Intn(screenWidth), r.Intn(screenHeight)
		// x := int(r.NormFloat64()*stdDev + float64(meanX))
		// y := int(r.NormFloat64()*stdDev + float64(meanY))
		// points[i] = Point{x, y}
		points[i] = Point{meanX, meanY}
	}

	game := &Game{
		points: points,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Circle Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
