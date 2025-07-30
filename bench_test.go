package flatmap

import (
	"fmt"
	"strings"
	"testing"
)

func genIntFlatMap(length int) *FlatMap[int, string] {
	fm := NewFlatMap[int, string](length)
	for i := range length {
		fm.Set(i, fmt.Sprintf("result: %d", i))
	}
	return fm
}

func genIntMap(length int) map[int]string {
	m := make(map[int]string, length)
	for i := range length {
		m[i] = fmt.Sprintf("result: %d", i)
	}
	return m
}

func genStringFlatMap(length int) *FlatMap[string, int] {
	fm := NewFlatMap[string, int](length)
	for i := range length {
		fm.Set(fmt.Sprintf("test:%d", i), i)
	}
	return fm
}

func genStringMap(length int) map[string]int {
	m := make(map[string]int, length)
	for i := range length {
		m[fmt.Sprintf("test:%d", i)] = i
	}
	return m
}

// length greater than 64
func genLongStringFlatMap(length int) *FlatMap[string, int] {
	fm := NewFlatMap[string, int](length)
	for i := range length {
		fm.Set(strings.Repeat("a", 64)+fmt.Sprintf("test:%d", i), i)
	}
	return fm
}

func genLongStringMap(length int) map[string]int {
	m := make(map[string]int, length)
	for i := range length {
		m[strings.Repeat("a", 64)+fmt.Sprintf("test:%d", i)] = i
	}
	return m
}

const lengthLimit = 64

func BenchmarkFlatMapSetInt(b *testing.B) {
	fm := NewFlatMap[int, int](lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			fm.Set(i, i+1)
		}
	}
}

func BenchmarkMapSetInt(b *testing.B) {
	m := make(map[int]int, lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			m[i] = i + 1
		}
	}
}

func BenchmarkFlatMapSetString(b *testing.B) {
	fm := NewFlatMap[string, int](lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			fm.Set(fmt.Sprintf("testing: %d", i), i+1)
		}
	}
}

func BenchmarkMapSetString(b *testing.B) {
	m := make(map[string]int, lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			m[fmt.Sprintf("testing: %d", i)] = i + 1
		}
	}
}

func BenchmarkFlatMapSetLongString(b *testing.B) {
	fm := NewFlatMap[string, int](lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			fm.Set(strings.Repeat("a", 64)+fmt.Sprintf("testing: %d", i), i+1)
		}
	}
}

func BenchmarkMapSetLongString(b *testing.B) {
	m := make(map[string]int, lengthLimit)
	for b.Loop() {
		for i := range lengthLimit {
			m[strings.Repeat("a", 64)+fmt.Sprintf("testing: %d", i)] = i + 1
		}
	}
}

func BenchmarkFlatMapGetInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)

	for b.Loop() {
		for i := range lengthLimit {
			_, _ = fm.Get(i)
		}
	}
}

func BenchmarkMapGetInt(b *testing.B) {
	m := genIntMap(lengthLimit)

	for b.Loop() {
		for i := range lengthLimit {
			_ = m[i]
		}
	}
}

func BenchmarkFlatMapGetString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)

	for b.Loop() {
		for range lengthLimit {
			_, _ = fm.Get("test:50")
		}
	}
}

func BenchmarkMapGetString(b *testing.B) {
	m := genStringMap(lengthLimit)

	for b.Loop() {
		for range lengthLimit {
			_ = m["test:50"]
		}
	}
}

func BenchmarkFlatMapGetLongString(b *testing.B) {
	fm := genLongStringFlatMap(lengthLimit)
	key := strings.Repeat("a", 64) + "test:50"

	for b.Loop() {
		for range lengthLimit {
			_, _ = fm.Get(key)
		}
	}
}

func BenchmarkMapGetLongString(b *testing.B) {
	m := genLongStringMap(lengthLimit)
	key := strings.Repeat("a", 64) + "test:50"

	for b.Loop() {
		for range lengthLimit {
			_ = m[key]
		}
	}
}

func BenchmarkFlatMapRangeInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)

	for b.Loop() {
		for iter := fm.Range(); !iter.HasNext(); iter.Next() {
			_, _ = iter.Val()
		}
	}
}

func BenchmarkFlatMapIterateInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)

	for b.Loop() {
		for range fm.All() {
		}
	}
}

func BenchmarkMapRangeInt(b *testing.B) {
	m := genIntMap(lengthLimit)

	for b.Loop() {
		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkFlatMapRangeString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)

	for b.Loop() {
		for iter := fm.Range(); !iter.HasNext(); iter.Next() {
			_, _ = iter.Val()
		}
	}
}

func BenchmarkFlatMapIterateString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)

	for b.Loop() {
		for range fm.All() {
		}
	}
}

func BenchmarkMapRangeString(b *testing.B) {
	m := genStringMap(lengthLimit)

	for b.Loop() {
		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkFlatMapDeleteInt(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		fm := genIntFlatMap(lengthLimit)
		b.StartTimer()
		for i := range lengthLimit {
			fm.Delete(i)
		}
	}
}

func BenchmarkMapDeleteInt(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		m := genIntMap(lengthLimit)
		b.StartTimer()
		for i := range lengthLimit {
			delete(m, i)
		}
	}
}

func BenchmarkFlatMapDeleteString(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		fm := genStringFlatMap(lengthLimit)
		b.StartTimer()
		for i := range lengthLimit {
			fm.Delete(fmt.Sprintf("test:%d", i))
		}
	}
}

func BenchmarkMapDeleteString(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		m := genStringMap(lengthLimit)
		b.StartTimer()
		for i := range lengthLimit {
			delete(m, fmt.Sprintf("test:%d", i))
		}
	}
}
