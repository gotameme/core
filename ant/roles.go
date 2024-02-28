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

package ant

type Level int

const (
	Decreased Level = iota - 1
	Default
	Increased
	Boosted
)

type Adjustments struct {
	Speed    Level
	Rotation Level
	Load     Level
	Vision   Level
	Range    Level
	Energy   Level
	Attack   Level
}

func (a Adjustments) IsValid() bool {
	// sum must be 0
	return a.Speed+a.Rotation+a.Load+a.Vision+a.Range+a.Energy+a.Attack == 0
}

type Roles map[string]Adjustments

type RolesCount map[string]int

type ChooseRole func(RolesCount) string
