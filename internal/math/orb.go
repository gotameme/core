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

package math

import (
	"github.com/paulmach/orb"
	"math"
)

const (
	DegToRad = math.Pi / 180
	RadToDeg = 180 / math.Pi
)

type Rect [2]orb.Point

func NewRect(x, y, w, h float64) Rect {
	return Rect{orb.Point{x - w/2, y - h/2}, orb.Point{x + w/2, y + h/2}}
}

func (r Rect) ToBox() (min, max [2]float64) {
	return [2]float64{r[0][0], r[0][1]}, [2]float64{r[1][0], r[1][1]}
}

func (r Rect) Size() (w, h float32) {
	return float32(r[1][0] - r[0][0]), float32(r[1][1] - r[0][1])
}

type Circle struct {
	orb.Point
	Radius float64
	Rect
}

func NewCircle(x, y, r float64) Circle {
	return Circle{
		orb.Point{x, y},
		r,
		Rect{
			orb.Point{x - r, y - r},
			orb.Point{x + r, y + r},
		},
	}
}

func CalcCenterOfGravity(min orb.Point, max orb.Point) orb.Point {
	return orb.Point{(min[0] + max[0]) / 2, (min[1] + max[1]) / 2}
}

func CalculateDirection(p1, p2 orb.Point) int {
	// Berechne den Winkel in Radiant
	angle := math.Atan2(p2.Y()-p1.Y(), p2.X()-p1.X())
	return int(math.Round(angle * RadToDeg))
}

// TouchesOrIntersects prüft, ob sich zwei Boxen berühren oder überschneiden.
func TouchesOrIntersects(b1, b2 Rect) bool {
	// Sie überschneiden oder berühren sich, wenn sie nicht getrennt sind.
	return !(b1[1][0] < b2[0][0] || b1[0][0] > b2[1][0] ||
		b1[1][1] < b2[0][1] || b1[0][1] > b2[1][1])
}
