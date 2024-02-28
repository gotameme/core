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
	"github.com/paulmach/orb"
	"log"

	gmath "github.com/gotameme/core/internal/math"
)

type Sugar struct {
	simulation *Simulation
	resources.AnimatedSprite
	gmath.Rect
	CurrentSugar int
}

func NewSugar(simulation *Simulation, position orb.Point) *Sugar {
	sugar := resources.NewSugar(simulation.screenWidth, simulation.screenHeight)
	sugar.Position = position
	rect := gmath.NewRect(position[0], position[1], float64(sugar.FrameWidth), float64(sugar.FrameHeight))
	log.Println(position[0], position[1], float64(sugar.FrameWidth), float64(sugar.FrameHeight), rect)
	return &Sugar{
		simulation:     simulation,
		AnimatedSprite: sugar,
		Rect:           rect,
		CurrentSugar:   1000,
	}
}

func NewRandomSugar(simulation *Simulation, border int) *Sugar {
	var minValue = [2]float64{float64(border), float64(border)}
	var maxValue = [2]float64{float64(simulation.screenWidth - border), float64(simulation.screenHeight - border)}
	return NewSugar(simulation, rand.RandomPoint(minValue, maxValue))
}

func (s *Sugar) Update() {
	if s.CurrentSugar <= 0 {
		s.simulation.RemoveSugar(s)
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
	op.GeoM.Translate(s.Position[0], s.Position[1])
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
