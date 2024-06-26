package flatmap

import (
	"cmp"
	"slices"
)

type FlatMap[K cmp.Ordered, V any] struct {
	data []Pair[K, V]
}

func NewFlatMap[K cmp.Ordered, V any](capacity int) *FlatMap[K, V] {
	return &FlatMap[K, V]{
		data: make([]Pair[K, V], 0, capacity),
	}
}

func (f *FlatMap[K, V]) upperBound(key K) int {
	count := len(f.data)
	first := 0
	for count > 0 {
		middle := first
		step := int(uint(count) >> 1)
		middle += step
		if !(key < f.data[middle].First) {
			middle++
			first = middle
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func (f *FlatMap[K, V]) lowerBound(key K) int {
	count := len(f.data)
	first := 0
	for count > 0 {
		middle := first
		step := int(uint(count) >> 1)
		middle += step
		if key > f.data[middle].First {
			middle++
			first = middle
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func (f *FlatMap[K, V]) Get(key K) (V, bool) {
	index := f.lowerBound(key)
	if index >= len(f.data) || f.data[index].First != key {
		var emptyValue V
		return emptyValue, false
	}

	return f.data[index].Second, true
}

func (f *FlatMap[K, V]) Set(key K, value V) {
	index := f.upperBound(key)
	switch {
	case index > 0 && index <= len(f.data) && f.data[index-1].First == key:
		f.data[index-1].Second = value
	default:
		f.data = slices.Insert(f.data, index, Pair[K, V]{key, value})
	}
}

func (f *FlatMap[K, V]) Delete(key K) {
	index := f.lowerBound(key)
	if index >= len(f.data) {
		return
	}
	if f.data[index].First != key {
		return
	}
	f.data = slices.Delete(f.data, index, index+1)
}

func (f *FlatMap[K, V]) Length() int {
	return len(f.data)
}

func (f *FlatMap[K, V]) Reset() {
	f.data = nil
}

type FlatMapIter[K cmp.Ordered, V any] struct {
	pos  int
	data []Pair[K, V]
}

func (fi *FlatMapIter[K, V]) Next() {
	fi.pos++
}

func (fi *FlatMapIter[K, V]) Val() (K, V) {
	var key K
	var value V
	if fi.pos >= 0 && fi.pos < len(fi.data) {
		key = fi.data[fi.pos].First
		value = fi.data[fi.pos].Second
	}
	return key, value
}

func (fi *FlatMapIter[K, V]) HasNext() bool {
	return fi.pos >= len(fi.data)
}

func newFlatMapIter[K cmp.Ordered, V any](pos int, data []Pair[K, V]) *FlatMapIter[K, V] {
	return &FlatMapIter[K, V]{
		pos:  pos,
		data: data,
	}
}

func (f *FlatMap[K, V]) Range() *FlatMapIter[K, V] {
	return newFlatMapIter(0, f.data)
}
