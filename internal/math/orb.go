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
	"github.com/chewxy/math32"
	"math"
)

const (
	DegToRad = math.Pi / 180
	RadToDeg = 180 / math.Pi
)

type Rect [2][2]float32

func NewRect(x, y, w, h float32) Rect {
	return Rect{[2]float32{x - w/2, y - h/2}, [2]float32{x + w/2, y + h/2}}
}

func (r Rect) ToBox() (min, max [2]float32) {
	return [2]float32{r[0][0], r[0][1]}, [2]float32{r[1][0], r[1][1]}
}

func (r Rect) Size() (w, h float32) {
	return r[1][0] - r[0][0], r[1][1] - r[0][1]
}

type Circle struct {
	Position [2]float32
	Radius   float32
	Rect
}

func NewCircle(x, y, r float32) Circle {
	return Circle{
		[2]float32{x, y},
		r,
		NewRect(x, y, r*2, r*2),
	}
}

func CalcCenterOfGravity(min, max [2]float32) [2]float32 {
	return [2]float32{(min[0] + max[0]) / 2, (min[1] + max[1]) / 2}
}

func CalculateDirection(p1, p2 [2]float32) int {
	// Berechne den Winkel in Radiant
	angle := math32.Atan2(p2[1]-p1[1], p2[0]-p1[0])
	return int(math32.Round(angle * RadToDeg))
}

// TouchesOrIntersects prüft, ob sich zwei Boxen berühren oder überschneiden.
func TouchesOrIntersects(b1, b2 Rect) bool {
	// Sie überschneiden oder berühren sich, wenn sie nicht getrennt sind.
	return !(b1[1][0] < b2[0][0] || b1[0][0] > b2[1][0] ||
		b1[1][1] < b2[0][1] || b1[0][1] > b2[1][1])
}
