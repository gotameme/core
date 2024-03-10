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
	"image/color"
	"sync"

	dll "github.com/emirpasic/gods/v2/lists/doublylinkedlist"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tidwall/rtree"
)

type Simulation struct {
	screenWidth, screenHeight int
	rtree                     *rtree.RTreeGN[float32, GameObject]
	addMarkingQueue           []*Marking
	removeMarkingQueue        map[*Marking]struct{}
	queueMutex                sync.Mutex

	RolesCount map[string]int
	SimulationConfig

	antHill *AntHill
	apple   *Apple
	ants    *dll.List[*AntOS]
	marks   *dll.List[*Marking]
	sugar   *dll.List[*Sugar]
}

func NewSimulation(screenWidth, screenHeight int, cnf SimulationConfig) *Simulation {
	var rTree = rtree.RTreeGN[float32, GameObject]{}
	var antHill = NewRandomAntHill(screenWidth, screenHeight, 30)
	var apple = NewRandomApple(screenWidth, screenHeight, 30)
	rTree.Insert(antHill.Bounds())
	simulation := &Simulation{
		ants:             dll.New[*AntOS](),
		marks:            dll.New[*Marking](),
		sugar:            dll.New[*Sugar](),
		screenWidth:      screenWidth,
		screenHeight:     screenHeight,
		antHill:          antHill,
		apple:            apple,
		rtree:            &rTree,
		SimulationConfig: cnf,
		RolesCount:       make(map[string]int),
	}

	return simulation
}

func (s *Simulation) Update() error {
	if s.ants.Size() < s.antDesiredValue {
		for i := 0; i < s.antDesiredValue-s.ants.Size(); i++ {
			s.AddNewAnt()
		}
	}
	if s.sugar.Size() < s.sugarDesiredValue {
		for i := 0; i < s.sugarDesiredValue-s.sugar.Size(); i++ {
			sugar := NewRandomSugar(s, 30)
			newMin, newMax, _ := sugar.Bounds()
			s.rtree.Insert(newMin, newMax, sugar)
			s.sugar.Append(sugar)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(s.ants.Size())
	mu := sync.Mutex{}
	s.ants.Each(func(i int, antOS *AntOS) {
		go func(antOS *AntOS) {
			defer wg.Done()
			// poc
			oldMin, oldMax := antOS.Bounds()
			antOS.Update()
			newMin, newMax := antOS.Bounds()
			mu.Lock()
			defer mu.Unlock()
			s.rtree.Replace(oldMin, oldMax, antOS, newMin, newMax, antOS)
		}(antOS)
	})

	s.sugar.Each(func(i int, sugar *Sugar) {
		wg.Add(1)
		go func(sugar *Sugar) {
			defer wg.Done()
			sugar.Update()
		}(sugar)
	})
	wg.Wait()

	s.ants.Each(func(i int, antOS *AntOS) {
		wg.Add(1)
		go func(antOS *AntOS) {
			defer wg.Done()
			s.rtree.Search(antOS.See())
		}(antOS)
		wg.Add(1)
		go func(antOS *AntOS) {
			defer wg.Done()
			s.rtree.Search(antOS.Smell())
		}(antOS)
		if antOS.Target != nil {
			// check if antOS is close to target
			// set antOS.Target to nil
			// calc box distance
			wg.Add(1)
			go func(antOS *AntOS) {
				defer wg.Done()
				s.rtree.Search(antOS.Collides())
			}(antOS)
		}
	})
	wg.Wait()
	s.marks.Each(func(_ int, mark *Marking) {
		mark.Update()
	})
	s.FlushMarkingChanges()
	return nil
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xc0, B: 0xa0, A: 0xff})

	s.marks.Each(func(_ int, mark *Marking) {
		mark.Draw(screen)
	})
	s.ants.Each(func(i int, antOS *AntOS) {
		antOS.Draw(screen)
	})
	s.sugar.Each(func(i int, sugar *Sugar) {
		sugar.Draw(screen)
	})
	s.antHill.Draw(screen)
	// s.apple.Draw(screen)
	// x, y := ebiten.CursorPosition()
	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nLen: %d\nSugar: %d", ebiten.ActualTPS(), ebiten.ActualFPS(), s.rtree.Len(), s.antHill.CurrentSugar)
	ebitenutil.DebugPrint(screen, msg)
}

func (s *Simulation) Layout(_, _ int) (int, int) {
	return s.screenWidth, s.screenHeight
}

func (s *Simulation) AddNewAnt() {
	roleName := s.chooseRole(s.RolesCount)
	antProperties := s.defaultRoleProperties
	if properties, ok := s.roles[roleName]; ok {
		antProperties = properties
		s.RolesCount[roleName]++
	}
	antOS := NewAntOS(s, WithAntHill(s.antHill), WithRole(roleName, antProperties))

	newAnt := s.antConstructor(antOS)
	antOS.Init(newAnt)

	s.ants.Append(antOS)
	newMin, newMax := antOS.Bounds()
	s.rtree.Insert(newMin, newMax, antOS)
}

func (s *Simulation) RemoveAnt(ant *AntOS) {
	s.ants.Each(func(i int, a *AntOS) {
		if a == ant {
			s.ants.Remove(i)
			vmin, vmax := ant.Bounds()
			s.rtree.Delete(vmin, vmax, ant)
			return
		}
	})
}

func (s *Simulation) RemoveSugar(sugar *Sugar) {
	s.rtree.Delete(sugar.Bounds())
	s.sugar.Each(func(i int, _sugar *Sugar) {
		if _sugar == sugar {
			s.sugar.Remove(i)
			return
		}
	})
}

func (s *Simulation) AddMarkingAtPosition(position [2]float32, radius int, information int) {
	marking := NewMarking(s, position, radius, information)
	s.AddMarking(marking)
}

func (s *Simulation) AddMarking(m *Marking) {
	s.queueMutex.Lock()
	s.addMarkingQueue = append(s.addMarkingQueue, m)
	s.queueMutex.Unlock()
}

func (s *Simulation) RemoveMarking(m *Marking) {
	s.queueMutex.Lock()
	s.removeMarkingQueue[m] = struct{}{}
	s.queueMutex.Unlock()
}

func (s *Simulation) FlushMarkingChanges() {
	s.queueMutex.Lock()
	defer s.queueMutex.Unlock()
	for _, m := range s.addMarkingQueue {
		s.rtree.Insert(m.Bounds())
		s.marks.Append(m)
	}
	s.addMarkingQueue = nil

	s.marks.Each(func(i int, mark *Marking) {
		if _, ok := s.removeMarkingQueue[mark]; ok {
			s.marks.Remove(i)
			MarkCache.RemoveMark(mark)
			s.rtree.Delete(mark.Bounds())
		}
	})
	s.removeMarkingQueue = make(map[*Marking]struct{})
}
