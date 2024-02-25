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

package helper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/paulmach/orb"
	"image/color"
)

var circleCache = make(map[int]map[color.RGBA]*ebiten.Image)

// NewCircle returns a new circle image with the given radius and color.
// The returned image is cached, so the same radius and color won't create a new image.
func NewCircle(r int, clr color.RGBA) *ebiten.Image {
	if img, ok := circleCache[r][clr]; ok {
		return img
	}
	size := r*2 + 1
	img := ebiten.NewImage(size, size)
	// make the fill color transparent
	transparentColor := color.RGBA{R: clr.R, G: clr.G, B: clr.B, A: 1}
	transparentColor2 := color.RGBA{R: clr.R, G: clr.G, B: clr.B, A: 8}
	// vector.DrawFilledRect(img, 0, 0, float32(2*r), float32(2*r), transparentColor, true)
	// vector.StrokeRect(img, 0, 0, float32(2*r), float32(2*r), 3, clr, true)
	vector.DrawFilledCircle(img, float32(r), float32(r), float32(r), transparentColor, true)
	vector.StrokeCircle(img, float32(r), float32(r), float32(r)-1.5, 3, transparentColor2, true)
	if circleCache[r] == nil {
		circleCache[r] = make(map[color.RGBA]*ebiten.Image)
	}
	circleCache[r][clr] = img
	return img
}

// DrawRect returns a new rectangle image with the given rectangle and color.
func DrawRect(r [2]orb.Point, clr color.RGBA) *ebiten.Image {
	w, h := float32(r[1][0]-r[0][0]), float32(r[1][1]-r[0][1])
	img := ebiten.NewImage(int(w), int(h))
	transparentColor := color.RGBA{R: clr.R, G: clr.G, B: clr.B, A: 128}
	vector.DrawFilledRect(img, 0, 0, w, h, transparentColor, true)
	vector.StrokeRect(img, 0, 0, w-3, h-3, 3, clr, true)
	return img
}
