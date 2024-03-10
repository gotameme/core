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
	"github.com/gotameme/core/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var autoAntID int

type AntOptions func(*AntOS)

type AntOSConstructor func(AntOptions) *AntOS

func WithPosition(x, y float32) AntOptions {
	return func(a *AntOS) {
		a.Position = [2]float32{x, y}
	}
}

func WithAntHill(antHill *AntHill) AntOptions {
	return func(a *AntOS) {
		a.Position = antHill.Position
		a.antHill = antHill
	}
}

func WithRole(role string, properties Properties) AntOptions {
	return func(a *AntOS) {
		a.role = role
		a.Properties = properties
	}
}

type AntOS struct {
	id         int
	role       string
	simulation *Simulation
	resources.AnimatedSprite
	ant interface{}

	antHill *AntHill
	Target  interface{}

	visionImage *ebiten.Image

	CurrentDirection float32 // in degrees
	State            AntOSState
	CurrentSugarLoad int

	Properties

	SetMarkThreshold int
	SetMarkResetTime int
}

func NewAntOS(simulation *Simulation, options ...AntOptions) *AntOS {
	autoAntID++
	antOS := &AntOS{
		id:               autoAntID,
		simulation:       simulation,
		AnimatedSprite:   resources.NewAnimatedAnt(simulation.screenWidth, simulation.screenHeight),
		Properties:       simulation.defaultRoleProperties,
		SetMarkThreshold: 10, // in ticks
	}

	for _, option := range options {
		option(antOS)
	}

	antOS.visionImage = helper.NewCircle(int(antOS.Vision), color.RGBA{R: 128, G: 0, B: 0, A: 255})

	return antOS
}

func (a *AntOS) Init(ant interface{}) {
	a.ant = ant
}

func (a *AntOS) IsInitialized() bool {
	return a.ant != nil
}
