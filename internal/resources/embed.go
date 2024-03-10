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
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

var (
	//go:embed ant_simple.png
	AntSimpleSprites_png []byte

	//go:embed arrow.png
	Arrow_png []byte

	//go:embed sugar_sprite.png
	Sugar2_png []byte

	//go:embed ant_hill.png
	AntHill_png []byte

	//go:embed gotameme_single.png
	GoTameMe_png []byte

	//go:embed green_apple.png
	GreenApple_png []byte
)

func NewAnimatedAnt(screenWidth, screenHeight int) AnimatedSprite {
	return newAnimatedSprite(
		DecodeEbitenImage(AntSimpleSprites_png),
		screenWidth,
		screenHeight,
		8,
		4,
		[]Animation{
			{0, 0, 2},
			{0, 4, 2},
		},
		1,
	)
}

func NewAntHill(screenWidth, screenHeight int) AnimatedSprite {
	img, _, err := image.Decode(bytes.NewReader(AntHill_png))
	if err != nil {
		panic(err)
	}
	w, h := 32, 32
	return newAnimatedSprite(
		ebiten.NewImageFromImage(img),
		screenWidth,
		screenHeight,
		w,
		h,
		[]Animation{
			{0, 0, 1},
		},
		1,
	)
}

func NewGreenApple(screenWidth, screenHeight int) AnimatedSprite {
	img, _, err := image.Decode(bytes.NewReader(GreenApple_png))
	if err != nil {
		panic(err)
	}
	w, h := 19, 24
	return newAnimatedSprite(
		ebiten.NewImageFromImage(img),
		screenWidth,
		screenHeight,
		w,
		h,
		[]Animation{
			{0, 0, 1},
		},
		1,
	)
}

func NewSugar(screenWidth, screenHeight int) AnimatedSprite {
	return newAnimatedSprite(
		DecodeEbitenImage(Sugar2_png),
		screenWidth,
		screenHeight,
		32,
		24,
		[]Animation{
			{320, 0, 1},
			{288, 0, 1},
			{256, 0, 1},
			{224, 0, 1},
			{192, 0, 1},
			{160, 0, 1},
			{128, 0, 1},
			{96, 0, 1},
			{64, 0, 1},
			{32, 0, 1},
			{0, 0, 1},
		},
		1,
	)
}

func DecodeEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
