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
	"fmt"
	"github.com/gotameme/core/ant"
	"math"

	gmath "github.com/gotameme/core/internal/math"
)

func (a *AntOS) GetId() int {
	return a.id
}

func (a *AntOS) GetRole() string {
	return a.role
}

func (a *AntOS) GetCurrentLoad() int {
	return a.CurrentSugarLoad
}

func (a *AntOS) GoForwards() {
	a.GoForward(math.MaxInt32)
}

func (a *AntOS) GoForward(steps int) {
	if _, ok := a.State.(*AntOSMoving); !ok {
		a.State = &AntOSMoving{
			Steps: &steps, // max int
		}
	} else {
		a.State.(*AntOSMoving).Steps = &steps
	}
}

func (a *AntOS) Turn(direction int) {
	fDirection := float32(direction)
	if _, ok := a.State.(*AntOSMoving); !ok {
		a.State = &AntOSMoving{
			TargetDirection: &fDirection,
		}
	} else {
		a.State.(*AntOSMoving).TargetDirection = &fDirection
	}

}

func (a *AntOS) GetDirectionToSugar(sugar ant.Sugar) int {
	// Calculate the direction to the object
	return gmath.CalculateDirection(a.GetPosition(), sugar.(*Sugar).GetPosition())
}

func (a *AntOS) GoToSugar(sugar ant.Sugar) {
	a.Target = sugar
	// log.Printf("Ant #%d turns to %v\n", a.GetId(), direction)
	a.Turn(a.GetDirectionToSugar(sugar))
	a.GoForwards()
}

func (a *AntOS) TakeSugar(sugar ant.Sugar) error {
	b1Min, b1Max := a.bounds()
	b2Min, b2Max, _ := sugar.(*Sugar).Bounds()
	if b1Max[0] < b2Min[0] || b1Min[0] > b2Max[0] || b1Max[1] < b2Min[1] || b1Min[1] > b2Max[1] {
		return fmt.Errorf("ant is not near the sugar")
	}
	// log.Printf("Ant #%d takes sugar\n", a.GetId())
	a.CurrentSugarLoad += sugar.(*Sugar).GetLoad(int(a.Load) - a.CurrentSugarLoad)
	return nil
}

func (a *AntOS) GotToAntHill() {
	a.Target = a.antHill
	// Calculate the direction to the object
	direction := gmath.CalculateDirection(a.GetPosition(), a.antHill.Position)
	// log.Printf("Ant #%d turns to %v\n", a.GetId(), direction)
	a.Turn(direction)
	a.GoForwards()
}

func (a *AntOS) SetMark(radius, information int) {
	if a.Range >= a.SetMarkResetTime {
		a.SetMarkResetTime = a.Range + a.SetMarkThreshold
		a.simulation.AddMarkingAtPosition(a.GetPosition(), radius, information)
	}
}
