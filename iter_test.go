package flatmap

import (
	"slices"
	"testing"
)

func TestFlatMapAll(t *testing.T) {
	fm := FlatMap[int, string]{
		data: []Pair[int, string]{
			{1, "4"},
			{2, "3"},
			{3, "2"},
			{4, "1"},
		},
	}
	count := 0
	for k, v := range fm.All() {
		count++
		if v2, ok := fm.Get(k); !ok || v2 != v {
			t.Errorf("got: %v, want: %v", v, v2)
		}
	}
	if count != fm.Length() {
		t.Errorf("length got: %v, want: %v", count, fm.Length())
	}
}

func TestFlatMapKeys(t *testing.T) {
	fm := FlatMap[int, string]{
		data: []Pair[int, string]{
			{1, "4"},
			{2, "3"},
			{3, "2"},
			{4, "1"},
		},
	}
	values := []int{1, 2, 3, 4}
	got := slices.Collect(fm.Keys())
	if !slices.Equal(got, values) {
		t.Errorf("keys got: %v, want: %v", got, values)
	}
}

func TestFlatMapValues(t *testing.T) {
	fm := FlatMap[int, string]{
		data: []Pair[int, string]{
			{1, "4"},
			{2, "3"},
			{3, "2"},
			{4, "1"},
		},
	}
	values := []string{"4", "3", "2", "1"}
	got := slices.Collect(fm.Values())
	if !slices.Equal(got, values) {
		t.Errorf("values got: %v, want: %v", got, values)
	}
}
