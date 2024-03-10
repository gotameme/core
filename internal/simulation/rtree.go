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
	"github.com/tidwall/rtree"
	"sync"
)

type RTreeG[T any] struct {
	base rtree.RTreeGN[float32, T]
	mu   sync.RWMutex
}

func (r *RTreeG[T]) Insert(min, max [2]float32, value T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Insert(min, max, value)
}

func (r *RTreeG[T]) Delete(min, max [2]float32, value T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Delete(min, max, value)
}

func (r *RTreeG[T]) Replace(
	oldMin, oldMax [2]float32, oldData T,
	newMin, newMax [2]float32, newData T,
) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Replace(
		oldMin, oldMax, oldData,
		newMin, newMax, newData,
	)
}

func (r *RTreeG[T]) Search(
	min, max [2]float32,
	iter func(min, max [2]float32, data T) bool,
) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Search(min, max, iter)
}

func (r *RTreeG[T]) Scan(iter func(min, max [2]float32, data T) bool) {
	r.mu.Lock()
	defer r.mu.RUnlock()
	r.base.Scan(iter)
}

func (r *RTreeG[T]) Nearby(
	algo func(min, max [2]float32, data T, item bool) (dist float64),
	iter func(min, max [2]float32, data T, dist float64) bool,
) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Nearby(algo, iter)
}

func (r *RTreeG[T]) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.base.Len()
}

func (r *RTreeG[T]) Copy() *RTreeG[T] {
	r.mu.Lock()
	defer r.mu.Unlock()
	return &RTreeG[T]{
		base: *r.base.Copy(),
	}
}

func (r *RTreeG[T]) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.base.Clear()
}
