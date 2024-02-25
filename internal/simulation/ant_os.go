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
	"github.com/gotameme/core/ant"
	"github.com/gotameme/core/internal/helper"
	"github.com/gotameme/core/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/paulmach/orb"
	"image/color"
)

var autoAntID int

type AntOptions func(*AntOS)

type AntOSConstructor func(AntOptions) *AntOS

func WithSightDistance(distance float64) AntOptions {
	return func(a *AntOS) {
		a.sightDistance = distance
	}
}

func WithPosition(x, y float64) AntOptions {
	return func(a *AntOS) {
		a.Position = orb.Point{x, y}
	}
}

func WithAntHill(antHill *AntHill) AntOptions {
	return func(a *AntOS) {
		a.Position = antHill.Position
		a.antHill = antHill
	}
}

type AntOS struct {
	id         int
	simulation *Simulation
	resources.AnimatedSprite
	ant ant.Ant

	antHill *AntHill
	Target  interface{}

	sightDistance      float64
	sightDistanceImage *ebiten.Image
	maxLoad            int

	CurrentDirection float64 // in degrees
	State            AntOSState
	CurrentSugarLoad int

	Lifespan int

	SetMarkThreshold int
	SetMarkResetTime int
}

func NewAntOS(simulation *Simulation, options ...AntOptions) *AntOS {
	autoAntID++
	antOS := &AntOS{
		id:             autoAntID,
		simulation:     simulation,
		AnimatedSprite: resources.NewAnimatedAnt(simulation.screenWidth, simulation.screenHeight),

		sightDistance:    60, // default value
		maxLoad:          5,  // default value
		SetMarkThreshold: 10, // in ticks
	}

	for _, option := range options {
		option(antOS)
	}

	antOS.sightDistanceImage = helper.NewCircle(int(antOS.sightDistance), color.RGBA{R: 128, G: 0, B: 0, A: 255})

	return antOS
}

func (a *AntOS) Init(ant ant.Ant) {
	a.ant = ant
}

func (a *AntOS) IsInitialized() bool {
	return a.ant != nil
}
