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

package core

import (
	"fmt"
	"github.com/gotameme/core/ant"
	"github.com/gotameme/core/internal"
	"github.com/gotameme/core/internal/simulation"
)

const (
	defaultScreenWidth  = 800
	defaultScreenHeight = 600
	defaultTPS          = 60
	maxSugar            = 40
)

type Option func(*Configuration)

type Options []Option

func (o Options) Apply(globalOptions *Configuration) {
	for _, option := range o {
		option(globalOptions)
	}
}

// region Global Options

type Configuration struct {
	ScreenWidth, ScreenHeight, TPS int
	headless                       bool
	internal.GameConfiguration
}

func NewConfiguration(opts ...Option) *Configuration {
	globalOptions := &Configuration{
		ScreenWidth:       defaultScreenWidth,
		ScreenHeight:      defaultScreenHeight,
		TPS:               defaultTPS,
		GameConfiguration: *internal.NewGameConfiguration(),
	}
	Options(opts).Apply(globalOptions)
	return globalOptions
}

func WithLayout(screenWidth, screenHeight int) Option {
	return func(cnf *Configuration) {
		cnf.ScreenWidth = screenWidth
		cnf.ScreenHeight = screenHeight
	}
}

func WithTPS(tps int) Option {
	if tps <= 0 {
		panic("TPS must be greater than 0")
	}
	return func(cnf *Configuration) {
		cnf.TPS = tps
	}

}

func Headless() Option {
	return func(cnf *Configuration) {
		cnf.headless = true
	}
}

// endregion

// region Game Options

func StartImmediately() Option {
	return func(cnf *Configuration) {
		internal.StartImmediately()(&cnf.GameConfiguration)
	}
}

// endregion

// region Simulation Options

func WithDesiredAnts(ants int) Option {
	return func(cnf *Configuration) {
		simulation.WithAntDesiredValue(ants)(&cnf.GameConfiguration.SimulationConfig)
	}
}

func WithDesiredSugar(sugar int) Option {
	if sugar < 0 {
		panic("Sugar must be greater than 0")
	}
	if sugar > maxSugar {
		sugar = maxSugar
		fmt.Printf("I'd say %d sugar is more than enough.\n", maxSugar)
	}
	return func(cnf *Configuration) {
		simulation.WithSugarDesiredValue(sugar)(&cnf.GameConfiguration.SimulationConfig)
	}
}

func WithAntConstructor(antConstructor ant.AntConstructor) Option {
	return func(cnf *Configuration) {
		simulation.WithAntConstructor(antConstructor)(&cnf.GameConfiguration.SimulationConfig)
	}
}

// endregion
