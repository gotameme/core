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
	"math"
)

func (a *AntOS) sightBox() ([2]float64, [2]float64) {
	return gmath.NewCircle(a.Position[0], a.Position[1], a.sightDistance).ToBox()
}

func (a *AntOS) bounds() ([2]float64, [2]float64) {
	return gmath.NewRect(a.Position[0], a.Position[1], float64(a.FrameWidth), float64(a.FrameHeight)).ToBox()
}

func distance(a, b [2]float64) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(dx*dx + dy*dy)
}

func (a *AntOS) See() ([2]float64, [2]float64, SearchIter) {
	start, end := a.sightBox()
	return start, end, func(min, max [2]float64, data GameObject) bool {
		if data == a {
			return true
		}
		switch data.(type) {
		// case *AntOS:
		// 	// println("See: ", data.(*AntOS).Name)
		// 	// log.Printf("#%d see AntOS #%d", a.GetId(), data.(*AntOS).GetId())
		// case *AntHill:
		// 	// Do nothing
		case *Sugar:
			sugar := data.(*Sugar)
			// do not see sugar if empty
			if sugar.CurrentSugar <= 0 {
				return true
			}
			if distance(a.Position, sugar.Position) <= a.sightDistance && a.Target != sugar {
				a.ant.SeeSugar(sugar)
			}
			// case *Marking:
			// 	if int(distance(a.Position, data.(*Marking).Position)) > data.(*Marking).Radius {
			// 		return true
			// 	}
			// 	if MarkCache.HasMark(a, data.(*Marking)) {
			// 		return true
			// 	}
			// 	MarkCache.AddMark(a, data.(*Marking))
			// 	a.ant.SeeMark(data.(*Marking))
			// default:
			// 	// log.Printf("See something unknown %T", data)
		}
		return true
	}
}

func (a *AntOS) Smell() ([2]float64, [2]float64, SearchIter) {
	start, end := a.bounds()
	return start, end, func(min, max [2]float64, data GameObject) bool {
		if data == a {
			return true
		}
		switch data.(type) {
		// case *AntOS:
		// 	// println("See: ", data.(*AntOS).Name)
		// case *AntHill:
		// 	// Do nothing
		// case *Sugar:
		// 	// Do nothing
		case *Marking:
			if MarkCache.HasMark(a, data.(*Marking)) {
				return true
			}
			MarkCache.AddMark(a, data.(*Marking))
			a.ant.SeeMark(data.(*Marking))
			// default:
			// 	// log.Printf("See something unknown %T", data)
		}
		return true
	}
}

func (a *AntOS) Collides() ([2]float64, [2]float64, SearchIter) {
	start, end := a.bounds()
	return start, end, func(min, max [2]float64, data GameObject) bool {
		if data == a {
			return true
		}
		if a.Target == data {
			a.Target = nil
			a.State = nil
			switch data.(type) {
			case *AntHill:
				// log.Printf("AntOS #%d reached the ant hill", a.GetId())
				data.(*AntHill).CurrentSugar += a.CurrentSugarLoad
				a.CurrentSugarLoad = 0
			case *Sugar:
				a.ant.ReachedSugar(data.(*Sugar))
				// default:
				// 	log.Println("See something unknown")
			}
		}
		return true
	}
}
