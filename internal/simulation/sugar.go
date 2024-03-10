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
	"github.com/gotameme/core/internal/resources"
	"github.com/gotameme/core/rand"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sugar struct {
	resources.AnimatedSprite
	gmath.Rect
	CurrentSugar int
}

func NewSugar(position [2]float32) *Sugar {
	sugar := resources.NewSugar()
	sugar.Position = position
	rect := gmath.NewRect(position[0], position[1], float32(sugar.FrameWidth), float32(sugar.FrameHeight))
	return &Sugar{
		AnimatedSprite: sugar,
		Rect:           rect,
		CurrentSugar:   1000,
	}
}

func NewRandomSugar(simulation *Simulation, border int) *Sugar {
	var minValue = [2]float32{float32(border), float32(border)}
	var maxValue = [2]float32{float32(simulation.screenWidth - border), float32(simulation.screenHeight - border)}
	return NewSugar(rand.RandomPoint(minValue, maxValue))
}

func (s *Sugar) Update(simulation *Simulation) {
	if s.CurrentSugar <= 0 {
		simulation.RemoveSugar(s)
	}
}

func (s *Sugar) Draw(screen *ebiten.Image) {
	if s.CurrentSugar <= 0 {
		return
	}

	s.CurrentAnimation = s.CurrentSugar/100 + 1
	// Stelle sicher, dass CurrentAnimation nicht größer als 10 wird.
	if s.CurrentAnimation > 10 {
		s.CurrentAnimation = 10
	}
	if s.CurrentSugar < 10 {
		s.CurrentAnimation = 1
	}

	img := s.AnimatedSprite.Draw()
	// rect := helper.DrawRect(s.Rect, colornames.Beige) // debug
	// rect.DrawImage(img, nil) // Debug
	// draw ant at position 100, 100
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.GetCenteredRotationOffset())
	op.GeoM.Translate(float64(s.Position[0]), float64(s.Position[1]))
	screen.DrawImage(img, op)
}

func (s *Sugar) GetLoad(i int) int {
	if i > s.CurrentSugar {
		s.CurrentSugar = 0
		return s.CurrentSugar
	}
	s.CurrentSugar -= i
	return i
}

func (s *Sugar) GetCurrentSugar() int {
	return s.CurrentSugar
}
