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
	"github.com/paulmach/orb"
	"math"
)

type AntOSState interface {
	Update(os *AntOS)
}

type AntOSMoving struct {
	TargetDirection *float64 // The direction the Ant should move to in degrees
	Steps           *int     // The number of steps the Ant should take
}

func (a *AntOSMoving) Update(os *AntOS) {
	if a.TargetDirection != nil {
		// Check if the Ant is already facing the target direction
		if os.CurrentDirection != *a.TargetDirection {
			a.rotate(os)
			return
		}
		a.TargetDirection = nil // Stop rotating
	}
	if a.Steps != nil {
		a.move(os)
		return
	}

	os.State = nil // Stop moving
}

func (a *AntOSMoving) rotate(os *AntOS) {
	if a.TargetDirection == nil {
		return
	}

	// Rotate the Ant to the target direction x degree per step
	// TODO: Add a rotation speed to the Ant and rotate by that speed
	directionDifference := *a.TargetDirection - os.CurrentDirection
	if directionDifference < -180 {
		directionDifference += 360
	} else if directionDifference > 180 {
		directionDifference -= 360
	}

	rotationSpeed := float64(os.Rotation)
	if math.Abs(directionDifference) < rotationSpeed {
		os.CurrentDirection = *a.TargetDirection
		a.TargetDirection = nil // Stop rotating
	} else {
		if directionDifference > 0 {
			os.CurrentDirection += rotationSpeed
		} else {
			os.CurrentDirection -= rotationSpeed
		}
	}

	// Keep the current direction within the range of [0, 360)
	if os.CurrentDirection >= 360 {
		os.CurrentDirection -= 360
	} else if os.CurrentDirection < 0 {
		os.CurrentDirection += 360
	}
}

func (a *AntOSMoving) move(os *AntOS) {
	if a.Steps == nil || *a.Steps <= 0 {
		a.Steps = nil // Stop moving
		return
	}

	speed := float64(os.Speed)
	if os.CurrentSugarLoad > 0 {
		// Ant is carrying sugar, reduce speed
		speed /= 2.0
	}
	// calculate the new position
	*a.Steps -= int(speed)
	x := os.Position.X() + speed*math.Cos(os.CurrentDirection*gmath.DegToRad)
	y := os.Position.Y() + speed*math.Sin(os.CurrentDirection*gmath.DegToRad)
	os.Position = orb.Point{x, y}
	if x < 0 {
		// bounce left edge of screen reached
		os.Position[0] = -x
		os.CurrentDirection = 180 - os.CurrentDirection
		// log.Printf("bounce left edge of screen reached: %f, %v", s.angle, s.v)
	} else if mx := float64(os.ScreenWidth) - float64(os.FrameWidth)/2; mx <= x {
		// bounce right edge of screen reached
		os.Position[0] = 2*mx - x
		os.CurrentDirection = 180 - os.CurrentDirection
		// log.Printf("bounce right edge of screen reached: %f, %v", s.angle, s.v)
	} else if y < 0 {
		// bounce top edge of screen reached
		os.Position[1] = -y
		os.CurrentDirection = -os.CurrentDirection
		// log.Printf("bounce top edge of screen reached: %f, %v", s.angle, s.v)
	} else if my := float64(os.ScreenHeight - os.FrameHeight/2); my <= y {
		// bounce bottom edge of screen reached
		os.Position[1] = 2*my - y
		os.CurrentDirection = -os.CurrentDirection
		// log.Printf("bounce bottom edge of screen reached: %f, %v", s.angle, s.v)
	}

}
