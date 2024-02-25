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

package resources

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"image"
	"time"
)

type AnimationFunc func(screen *ebiten.Image, image *ebiten.Image, elapsed float64)

type TimedAnimation struct {
	Images     []*ebiten.Image
	ImageIndex int
	Timer      *time.Timer
	StartTime  time.Time
	AnimationFunc
}

func NewTimedAnimation(images []*ebiten.Image, animationFunc AnimationFunc) *TimedAnimation {
	return &TimedAnimation{
		Images:        images,
		AnimationFunc: animationFunc,
	}
}

func (a *TimedAnimation) Start() {
	a.Timer = time.NewTimer(time.Second)
	a.StartTime = time.Now()
}

func (a *TimedAnimation) Stop() {
	a.Timer.Stop()
	a.Timer = nil
}

func (a *TimedAnimation) Reset() {
	a.Stop()
	a.ImageIndex = 0
	// a.Start()
}

func (a *TimedAnimation) Update() bool {
	if a.Timer == nil {
		a.Start()
	} else {
		select {
		case <-a.Timer.C:
			a.ImageIndex++
			a.Timer.Reset(time.Second)
			a.StartTime = time.Now()
			if a.ImageIndex >= len(a.Images) {
				a.ImageIndex = 0
				return false
			}
		default:
			// Do nothing
		}
	}
	return true
}

func (a *TimedAnimation) GetCurrentImage() *ebiten.Image {
	return a.Images[a.ImageIndex]
}

func (a *TimedAnimation) Draw(screen *ebiten.Image) {
	elapsed := float64(time.Since(a.StartTime).Milliseconds())
	a.AnimationFunc(screen, a.GetCurrentImage(), elapsed)
}

func NewGoTameMeAnimation(screenWidth, screenHeight float64) *TimedAnimation {
	decodedImage, _, decodeErr := image.Decode(bytes.NewReader(GoTameMe_png))
	if decodeErr != nil {
		panic(decodeErr)
	}
	ebitenImage := ebiten.NewImageFromImage(decodedImage)
	imageWidth, imageHeight := 69, 22
	subImages := make([]*ebiten.Image, 3)
	for i := 0; i < 3; i++ {
		subImages[i] = ebitenImage.SubImage(image.Rect(i*imageWidth, 0, (i+1)*imageWidth, imageHeight)).(*ebiten.Image)
	}
	return NewTimedAnimation(subImages, func(screen *ebiten.Image, image *ebiten.Image, elapsed float64) {
		alpha := 1.0 - elapsed/1000.0
		scaleFactor := 1.0 + elapsed/100.0

		op := &colorm.DrawImageOptions{}
		op.GeoM.Scale(scaleFactor, scaleFactor)
		op.GeoM.Translate(screenWidth/2-float64(image.Bounds().Dx())*scaleFactor/2, screenHeight/2-float64(image.Bounds().Dy())*scaleFactor/2)

		colorM := colorm.ColorM{}
		colorM.Scale(1, 1, 1, alpha)

		colorm.DrawImage(screen, image, colorM, op)
	})
}
