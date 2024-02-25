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
	"github.com/gotameme/core/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/paulmach/orb"
	"image/color"
)

type Marking struct {
	simulation  *Simulation
	img         *ebiten.Image
	op          *ebiten.DrawImageOptions
	Position    orb.Point
	Radius      int
	Information int
	Lifespan    int
}

func (m *Marking) GetPosition() orb.Point {
	return m.Position
}

func NewMarking(simulation *Simulation, position orb.Point, radius int, information int) *Marking {
	// return &Marking{
	// 	simulation:  simulation,
	// 	Position:    position,
	// 	Radius:      radius,
	// 	Information: information,
	// 	Lifespan:    150, // ticks
	// }
	m := GetMarking()
	m.simulation = simulation
	m.Position = position
	m.Radius = radius
	m.Information = information
	m.Lifespan = 200 // ticks
	return m
}

func (m *Marking) Draw(screen *ebiten.Image) {
	if m.img == nil {
		circle := helper.NewCircle(m.Radius, color.RGBA{R: 128, G: 128, B: 0, A: 255})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(m.Position[0]-float64(m.Radius), m.Position[1]-float64(m.Radius))
		m.img = circle
		m.op = op
	}
	screen.DrawImage(m.img, m.op)
}

func (m *Marking) Update() {
	m.Lifespan--
	if m.Lifespan <= 0 {
		m.simulation.RemoveMarking(m)
	}
}

func (m *Marking) Bounds() ([2]float64, [2]float64, *Marking) {
	vmin := [2]float64{m.Position[0] - float64(m.Radius), m.Position[1] - float64(m.Radius)}
	vmax := [2]float64{m.Position[0] + float64(m.Radius), m.Position[1] + float64(m.Radius)}
	return vmin, vmax, m
}

func (m *Marking) GetInformation() int {
	return m.Information
}
