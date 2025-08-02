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

func TestFlatMapAllContinueBreak(t *testing.T) {
	fm := FlatMap[int, string]{
		data: []Pair[int, string]{
			{1, "4"},
			{2, "3"},
			{3, "2"},
			{4, "1"},
		},
	}
	count := 0
	for range fm.All() {
		count++
		break
	}
	if count != 1 {
		t.Error("break did not work")
	}

	count = 0
	for k := range fm.All() {
		if k%2 == 0 {
			continue
		}
		count++
	}
	if count != 2 {
		t.Error("continue did not work")
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

func TestFlatMapKeysContinueBreak(t *testing.T) {
	fm := FlatMap[int, string]{
		data: []Pair[int, string]{
			{1, "4"},
			{2, "3"},
			{3, "2"},
			{4, "1"},
		},
	}
	count := 0
	for range fm.Keys() {
		count++
		break
	}
	if count != 1 {
		t.Error("break did not work")
	}

	count = 0
	for k := range fm.Keys() {
		if k%2 == 0 {
			continue
		}
		count++
	}
	if count != 2 {
		t.Error("continue did not work")
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

func TestFlatMapValuesContinueBreak(t *testing.T) {
	fm := FlatMap[int, int]{
		data: []Pair[int, int]{
			{1, 4},
			{2, 3},
			{3, 2},
			{4, 1},
		},
	}
	count := 0
	for range fm.Values() {
		count++
		break
	}
	if count != 1 {
		t.Error("break did not work")
	}

	count = 0
	for v := range fm.Values() {
		if v%2 == 0 {
			continue
		}
		count++
	}
	if count != 2 {
		t.Error("continue did not work")
	}
}
