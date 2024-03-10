/*
Copyright (c) 2024 Sebastian Kroczek <me@xbug.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package simulation

import (
	"github.com/gotameme/core/internal/resources"
	"github.com/gotameme/core/rand"
	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	resources.AnimatedSprite
	img *ebiten.Image
}

func NewApple(screenWidth, screenHeight int, position [2]float32) *Apple {
	antHillAnimatedSprite := resources.NewGreenApple()
	antHillAnimatedSprite.Position = position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position[0]-float32(antHillAnimatedSprite.FrameWidth)/2), float64(position[1]-float32(antHillAnimatedSprite.FrameHeight)/2))
	// Anthill is currently a static image, so we can draw it once and forget about it
	anthillImage := ebiten.NewImage(screenWidth, screenHeight)
	anthillImage.DrawImage(antHillAnimatedSprite.Draw(), op)
	return &Apple{
		img:            anthillImage,
		AnimatedSprite: antHillAnimatedSprite,
	}
}

func NewRandomApple(screenWidth, screenHeight, border int) *Apple {
	var minValue = [2]float32{float32(border), float32(border)}
	var maxValue = [2]float32{float32(screenWidth - border), float32(screenHeight - border)}
	return NewApple(screenWidth, screenHeight, rand.RandomPoint(minValue, maxValue))
}

func (a *Apple) Draw(screen *ebiten.Image) {
	screen.DrawImage(a.img, nil)
}

func (a *Apple) Bounds() ([2]float32, [2]float32, *Apple) {
	var aMin, aMax = a.AnimatedSprite.Bounds()
	return aMin, aMax, a
}

func (a *Apple) Update(*Simulation) {
	// Overwrite the Update method to do nothing
}
