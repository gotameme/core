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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAndRemoveMarking(t *testing.T) {
	// 1. Create a new instance of the `Simulation` structure.
	sim := NewSimulation(800, 600, SimulationConfig{})

	// 2. Add a new `Marking` to the simulation.
	position := [2]float32{100, 100}
	radius := 10
	information := 1
	sim.AddMarkingAtPosition(position, radius, information)
	sim.FlushMarkingChanges()

	// 3. Check if the `Marking` is in the tree.
	markingExists := false
	sim.rtree.Search([2]float32{0, 0}, [2]float32{800, 600}, func(min, max [2]float32, value GameObject) bool {
		if marking, ok := value.(*Marking); ok {
			if marking.Position == position && marking.Radius == radius && marking.Information == information {
				markingExists = true
			}
		}
		return true
	})
	assert.True(t, markingExists, "Marking should exist in the tree after being added")

	// 4. Remove the `Marking` from the simulation.
	sim.marks.Each(func(i int, m *Marking) {
		if m.Position == position && m.Radius == radius && m.Information == information {
			sim.RemoveMarking(m)
		}
	})
	sim.FlushMarkingChanges()

	// 5. Check if the `Marking` has been removed from the tree.
	markingExists = false
	sim.rtree.Search([2]float32{0, 0}, [2]float32{800, 600}, func(min, max [2]float32, value GameObject) bool {
		if marking, ok := value.(*Marking); ok {
			if marking.Position == position && marking.Radius == radius && marking.Information == information {
				markingExists = true
			}
		}
		return true
	})
	assert.False(t, markingExists, "Marking should not exist in the tree after being removed")
}

func TestMarkingRemovalAfterLifespanExceeded(t *testing.T) {
	// 1. Create a new instance of the `Simulation` structure.
	sim := NewSimulation(800, 600, SimulationConfig{})

	// 2. Add a new `Marking` to the simulation with a very small lifespan.
	position := [2]float32{100, 100}
	radius := 10
	information := 1
	marking := NewMarking(position, radius, information)
	marking.Lifespan = 0
	sim.AddMarking(marking)
	sim.FlushMarkingChanges()

	// 3. Execute the `Tick()` method of the simulation.
	sim.Tick()

	// 4. Check if the `Marking` has been removed from the tree.
	markingExists := false
	sim.rtree.Search([2]float32{0, 0}, [2]float32{800, 600}, func(min, max [2]float32, value GameObject) bool {
		if _, ok := value.(*Marking); ok {
			markingExists = true
		}
		return true
	})
	assert.False(t, markingExists, "Marking should not exist in the tree after its lifespan has been exceeded")
	assert.Equal(t, 0, sim.rtree.Len(), "Tree should be empty after the marking has been removed")
}
