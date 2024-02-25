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
	"errors"
	"github.com/gotameme/core/internal"
	"github.com/gotameme/core/internal/simulation"
	"github.com/hajimehoshi/ebiten/v2"
)

func Run(options ...Option) {
	cnf := NewConfiguration(options...)
	Options(options).Apply(cnf)
	if cnf.headless != true {
		ebiten.SetWindowSize(cnf.ScreenWidth, cnf.ScreenHeight)
		ebiten.SetWindowTitle("Go! Tame! Me!")
		ebiten.SetTPS(60)
		if err := ebiten.RunGame(internal.NewGame(cnf.ScreenWidth, cnf.ScreenHeight, cnf.GameConfiguration)); err != nil {
			if !errors.Is(err, internal.Quit) {
				panic(err)
			}
		}
	} else {
		s := simulation.NewSimulation(cnf.ScreenWidth, cnf.ScreenHeight, cnf.SimulationConfig)
		for {
			if err := s.Update(); err != nil {
				if !errors.Is(err, internal.Quit) {
					panic(err)
				}
				break
			}
		}
	}

}
