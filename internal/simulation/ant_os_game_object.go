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
	gmath "github.com/gotameme/core/internal/math"
	"github.com/hajimehoshi/ebiten/v2"
)

func (a *AntOS) Update() {
	a.Range++
	if a.State != nil {
		a.State.Update(a)
	} else {
		if waitAnt, ok := a.ant.(interface{ Waits() }); ok {
			waitAnt.Waits()
		}
	}
	a.AnimatedSprite.Update()
	if tickAnt, ok := a.ant.(interface{ Tick() }); ok {
		tickAnt.Tick()
	}
}

func (a *AntOS) Draw(screen *ebiten.Image) {
	if a.CurrentSugarLoad > 0 {
		// Draw the sugar load
		a.AnimatedSprite.CurrentAnimation = 1
	} else {
		a.AnimatedSprite.CurrentAnimation = 0
	}
	img := a.AnimatedSprite.Draw()
	// if a.visionImage != nil {
	// 	op := &ebiten.DrawImageOptions{}
	// 	// halfSigthDistance := a.sightDistance / 2
	// 	op.GeoM.Translate(a.Position.X()-float64(a.Vision), a.Position.Y()-float64(a.Vision))
	// 	screen.DrawImage(a.visionImage, op)
	// }
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.AnimatedSprite.GetCenteredRotationOffset())
	op.GeoM.Rotate(a.CurrentDirection * gmath.DegToRad)
	op.GeoM.Translate(a.Position[0], a.Position[1])
	screen.DrawImage(img, op)
}
