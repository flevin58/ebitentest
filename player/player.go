package player

import (
	"image"

	"github.com/flevin58/ebitentest/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image         *ebiten.Image
	frameOX       int
	frameOY       int
	frameWidth    int
	frameHeight   int
	frameCount    int
	ticksPerFrame int
	ticksCount    int
	currentFrame  int
}

func New() *Player {
	return &Player{
		image:         resources.GetImage("images/runner/runner.png"),
		frameOX:       0,  // x offset of the sprite sheet
		frameOY:       32, // y offset of the sprite sheet
		frameWidth:    32, // width of a single frame
		frameHeight:   32, // height of a single frame
		frameCount:    8,  // number of frames in the sprite sheet
		ticksPerFrame: 5,  // duration of each frame in ticks
		ticksCount:    0,  // current tick count
		currentFrame:  0,  // current frame index to display
	}
}

func (p *Player) Update() {
	// Update player state at every tick of the game (60 times per second)
	p.ticksCount++
	if p.ticksCount >= p.ticksPerFrame {
		p.ticksCount = 0
		p.currentFrame = (p.currentFrame + 1) % p.frameCount
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw the player on the screen
	b := screen.Bounds()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(p.frameWidth)/2, -float64(p.frameHeight)/2)
	op.GeoM.Translate(float64(b.Dx())/2, float64(b.Dy())/2)
	sx, sy := p.frameOX+p.currentFrame*p.frameWidth, p.frameOY
	frameRect := image.Rect(sx, sy, sx+p.frameWidth, sy+p.frameHeight)
	frameImage := p.image.SubImage(frameRect).(*ebiten.Image)
	screen.DrawImage(frameImage, op)
}
