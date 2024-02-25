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

import "sync"

type markCache map[*AntOS]map[*Marking]struct{}

var markCacheSyncMutex = sync.RWMutex{}

func (m *markCache) AddMark(pa *AntOS, pm *Marking) {
	// Lock the mutex to prevent concurrent map writes
	markCacheSyncMutex.Lock()
	defer markCacheSyncMutex.Unlock()
	if _, ok := (*m)[pa]; !ok {
		(*m)[pa] = make(map[*Marking]struct{})
	}
	(*m)[pa][pm] = struct{}{}
}

func (m *markCache) RemoveMark(pm *Marking) {
	// Lock the mutex to prevent concurrent map writes
	markCacheSyncMutex.Lock()
	defer markCacheSyncMutex.Unlock()
	for _, marks := range *m {
		delete(marks, pm)
	}
}

func (m *markCache) RemoveAnt(pa *AntOS) {
	// Lock the mutex to prevent concurrent map writes
	markCacheSyncMutex.Lock()
	defer markCacheSyncMutex.Unlock()
	delete(*m, pa)
}

func (m *markCache) HasMark(pa *AntOS, pm *Marking) (ok bool) {
	// Lock the mutex to prevent concurrent map reads
	markCacheSyncMutex.RLock()
	defer markCacheSyncMutex.RUnlock()
	if _, ok = (*m)[pa]; ok {
		_, ok = (*m)[pa][pm]
	}
	return
}

var MarkCache = markCache{}
