package flatmap

import (
	"fmt"
	"strings"
	"testing"
)

func genIntFlatMap(length int) *FlatMap[int, string] {
	fm := NewFlatMap[int, string](length)
	for i := 0; i < length; i++ {
		fm.Set(i, fmt.Sprintf("result: %d", i))
	}
	return fm
}

func genIntMap(length int) map[int]string {
	m := make(map[int]string, length)
	for i := 0; i < length; i++ {
		m[i] = fmt.Sprintf("result: %d", i)
	}
	return m
}

func genStringFlatMap(length int) *FlatMap[string, int] {
	fm := NewFlatMap[string, int](length)
	for i := 0; i < length; i++ {
		fm.Set(fmt.Sprintf("test:%d", i), i)
	}
	return fm
}

func genStringMap(length int) map[string]int {
	m := make(map[string]int, length)
	for i := 0; i < length; i++ {
		m[fmt.Sprintf("test:%d", i)] = i
	}
	return m
}

// length greater than 64
func genLongStringFlatMap(length int) *FlatMap[string, int] {
	fm := NewFlatMap[string, int](length)
	for i := 0; i < length; i++ {
		fm.Set(strings.Repeat("a", 64)+fmt.Sprintf("test:%d", i), i)
	}
	return fm
}

func genLongStringMap(length int) map[string]int {
	m := make(map[string]int, length)
	for i := 0; i < length; i++ {
		m[strings.Repeat("a", 64)+fmt.Sprintf("test:%d", i)] = i
	}
	return m
}

const lengthLimit = 64

func BenchmarkFlatMapSetInt(b *testing.B) {
	fm := NewFlatMap[int, int](lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			fm.Set(j, j+1)
		}
	}
}

func BenchmarkMapSetInt(b *testing.B) {
	m := make(map[int]int, lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			m[j] = j + 1
		}
	}
}

func BenchmarkFlatMapSetString(b *testing.B) {
	fm := NewFlatMap[string, int](lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			fm.Set(fmt.Sprintf("testing: %d", j), j+1)
		}
	}
}

func BenchmarkMapSetString(b *testing.B) {
	m := make(map[string]int, lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			m[fmt.Sprintf("testing: %d", j)] = j + 1
		}
	}
}

func BenchmarkFlatMapSetLongString(b *testing.B) {
	fm := NewFlatMap[string, int](lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			fm.Set(strings.Repeat("a", 64)+fmt.Sprintf("testing: %d", j), j+1)
		}
	}
}

func BenchmarkMapSetLongString(b *testing.B) {
	m := make(map[string]int, lengthLimit)
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			m[strings.Repeat("a", 64)+fmt.Sprintf("testing: %d", j)] = j + 1
		}
	}
}

func BenchmarkFlatMapGetInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_, _ = fm.Get(j)
		}
	}
}

func BenchmarkMapGetInt(b *testing.B) {
	m := genIntMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_ = m[j]
		}
	}
}

func BenchmarkFlatMapGetString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_, _ = fm.Get("test:50")
		}
	}
}

func BenchmarkMapGetString(b *testing.B) {
	m := genStringMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_ = m["test:50"]
		}
	}
}

func BenchmarkFlatMapGetLongString(b *testing.B) {
	fm := genLongStringFlatMap(lengthLimit)
	key := strings.Repeat("a", 64) + "test:50"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_, _ = fm.Get(key)
		}
	}
}

func BenchmarkMapGetLongString(b *testing.B) {
	m := genLongStringMap(lengthLimit)
	key := strings.Repeat("a", 64) + "test:50"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < lengthLimit; j++ {
			_ = m[key]
		}
	}
}

func BenchmarkFlatMapRangeInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for iter := fm.Range(); !iter.HasNext(); iter.Next() {
			_, _ = iter.Val()
		}
	}
}

func BenchmarkFlatMapIterateInt(b *testing.B) {
	fm := genIntFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range fm.All() {
		}
	}
}

func BenchmarkMapRangeInt(b *testing.B) {
	m := genIntMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkFlatMapRangeString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for iter := fm.Range(); !iter.HasNext(); iter.Next() {
			_, _ = iter.Val()
		}
	}
}

func BenchmarkFlatMapIterateString(b *testing.B) {
	fm := genStringFlatMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range fm.All() {
		}
	}
}

func BenchmarkMapRangeString(b *testing.B) {
	m := genStringMap(lengthLimit)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkFlatMapDeleteInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fm := genIntFlatMap(lengthLimit)
		b.StartTimer()
		for j := 0; j < lengthLimit; j++ {
			fm.Delete(j)
		}
	}
}

func BenchmarkMapDeleteInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := genIntMap(lengthLimit)
		b.StartTimer()
		for j := 0; j < lengthLimit; j++ {
			delete(m, j)
		}
	}
}

func BenchmarkFlatMapDeleteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fm := genStringFlatMap(lengthLimit)
		b.StartTimer()
		for j := 0; j < lengthLimit; j++ {
			fm.Delete(fmt.Sprintf("test:%d", j))
		}
	}
}

func BenchmarkMapDeleteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := genStringMap(lengthLimit)
		b.StartTimer()
		for j := 0; j < lengthLimit; j++ {
			delete(m, fmt.Sprintf("test:%d", j))
		}
	}
}
