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
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Animation struct {
	FrameOX    int
	FrameOY    int
	FrameCount int
}

type AnimatedSprite struct {
	Image                     *ebiten.Image
	Position                  [2]float32
	ScreenWidth, ScreenHeight int
	FrameWidth, FrameHeight   int
	centerX, centerY          float32
	Animations                []Animation
	CurrentAnimation          int
	Count                     int
	//
	AnimationSpeed int
}

func newAnimatedSprite(image *ebiten.Image, screenWidth, screenHeight int, frameWidth, frameHeight int, animations []Animation, animationSpeed int) AnimatedSprite {
	return AnimatedSprite{
		Image:            image,
		ScreenWidth:      screenWidth,
		ScreenHeight:     screenHeight,
		FrameWidth:       frameWidth,
		FrameHeight:      frameHeight,
		centerX:          float32(frameWidth) / 2,
		centerY:          float32(frameHeight) / 2,
		Animations:       animations,
		CurrentAnimation: 0,
		Count:            0,
		AnimationSpeed:   animationSpeed,
	}
}

// GetCenteredRotationOffset calculates the offsets necessary for centering the rotation point of the sprite.
// This function is crucial for enabling the sprite to be rotated around its center, ensuring that rotations are
// visually centered and consistent.
// It returns the X and Y offsets based on the sprite's frame dimensions. These offsets are applied to adjust
// the rotation anchor point to the sprite's center, facilitating centered rotations within Ebiten's graphics
// framework.
func (a *AnimatedSprite) GetCenteredRotationOffset() (float64, float64) {
	// Calculate and return the offsets for centered rotation based on frame dimensions
	return float64(-a.centerX), float64(-a.centerY)
}

func (a *AnimatedSprite) Update() {
	a.Count++
	currentAnimation := a.Animations[a.CurrentAnimation]
	if a.Count >= currentAnimation.FrameCount*a.AnimationSpeed {
		a.Count = 0
	}
}

func (a *AnimatedSprite) Draw() *ebiten.Image {
	currentAnimation := a.Animations[a.CurrentAnimation]
	i := (a.Count / a.AnimationSpeed) % currentAnimation.FrameCount
	sx, sy := currentAnimation.FrameOX+i*a.FrameWidth, currentAnimation.FrameOY
	return a.Image.SubImage(image.Rect(sx, sy, sx+a.FrameWidth, sy+a.FrameHeight)).(*ebiten.Image)
}

func (a *AnimatedSprite) GetPosition() [2]float32 {
	return a.Position
}

func (a *AnimatedSprite) Bounds() (min, max [2]float32) {
	min[0] = float32(a.Position[0]) - a.centerX
	min[1] = float32(a.Position[1]) - a.centerY
	max[0] = float32(a.Position[0]) + a.centerX
	max[1] = float32(a.Position[1]) + a.centerY
	return
}
