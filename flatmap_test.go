package flatmap

import (
	"slices"
	"testing"
)

func TestUpperBound(t *testing.T) {
	fm := FlatMap[int, int]{
		data: []Pair[int, int]{
			{1, 1},
			{2, 2},
		},
	}
	bound := fm.upperBound(1)
	if bound != 1 {
		t.Errorf("excepted 1, got %v", bound)
	}

	bound = fm.upperBound(2)
	if bound != 2 {
		t.Errorf("excepted 2, got %v", bound)
	}

	bound = fm.lowerBound(3)
	if bound != 2 {
		t.Errorf("excepted 2, got %v", bound)
	}

	bound = fm.upperBound(-1)
	if bound != 0 {
		t.Errorf("excepted 0, got %v", bound)
	}
}

func TestLowerBound(t *testing.T) {
	fm := FlatMap[int, int]{
		data: []Pair[int, int]{
			{1, 1},
			{2, 2},
		},
	}
	bound := fm.lowerBound(1)
	if bound != 0 {
		t.Errorf("excepted 0, got %v", bound)
	}

	bound = fm.lowerBound(2)
	if bound != 1 {
		t.Errorf("excepted 1, got %v", bound)
	}

	bound = fm.lowerBound(3)
	if bound != 2 {
		t.Errorf("excepted 2, got %v", bound)
	}

	bound = fm.lowerBound(-1)
	if bound != 0 {
		t.Errorf("excepted 0, got %v", bound)
	}
}

func TestFlatMapGetSet(t *testing.T) {
	fm := FlatMap[string, any]{}
	testCases := []struct {
		testData Pair[string, any]
		want     any
	}{
		{
			Pair[string, any]{"hello", 1},
			1,
		},
		{
			Pair[string, any]{"hello", "2"},
			"2",
		},
		{
			Pair[string, any]{"hello world", "testing"},
			"testing",
		},
		{
			Pair[string, any]{"hello1", 4},
			4,
		},
		{
			Pair[string, any]{"Hello", 5},
			5,
		},
		{
			Pair[string, any]{"zzzzzzzzzz", nil},
			nil,
		},
		{
			Pair[string, any]{"", 9999},
			9999,
		},
	}
	for _, tc := range testCases {
		fm.Set(tc.testData.First, tc.testData.Second)
		v, ok := fm.Get(tc.testData.First)
		if !ok {
			t.Errorf("%#v not found", tc.testData)
		}
		if v != tc.want {
			t.Errorf("want: %#v, got %#v", tc.want, v)
		}
	}
	_, ok := fm.Get("Not Found")
	if ok {
		t.Error("want false, got true")
	}
	if fm.Length() != len(testCases)-1 {
		t.Errorf("want %v, got %v", len(testCases), fm.Length())
	}
}

func TestFlatMapDelete(t *testing.T) {
	fm := FlatMap[string, int]{}
	data := []string{"hello", "Hello", "world", "hello world", "", "123", "abc"}
	for i, v := range data {
		fm.Set(v, i)
	}
	if fm.Length() != len(data) {
		t.Errorf("want %v, got %v", len(data), fm.Length())
	}

	fm.Delete("Not Found")
	if fm.Length() != len(data) {
		t.Errorf("want %v, got %v", len(data), fm.Length())
	}

	for i, v := range data {
		fm.Delete(v)
		if fm.Length() != len(data)-i-1 {
			t.Errorf("want %v, got %v", len(data)-i-1, fm.Length())
		}
		if _, ok := fm.Get(v); ok {
			t.Errorf("%v delete failed", v)
		}
	}

	fm.Delete("Not Found")
	if fm.Length() != 0 {
		t.Errorf("want %v, got %v", 0, fm.Length())
	}
}

func TestFlatMapReset(t *testing.T) {
	fm := FlatMap[string, int]{}
	data := []string{"hello", "Hello", "world", "hello world", "", "123", "abc"}
	for i, v := range data {
		fm.Set(v, i)
	}
	if fm.Length() != len(data) {
		t.Errorf("want %v, got %v", len(data), fm.Length())
	}

	fm.Reset()
	if fm.Length() != 0 {
		t.Errorf("error length after reset, want 0, got: %v", fm.Length())
	}
}

func TestFlatMapRange(t *testing.T) {
	fm := FlatMap[int, int]{}
	data := []int{100, 4, 5, 3, 2, 1}
	for i, v := range data {
		fm.Set(v, i)
	}
	slices.Sort(data)
	iteration := make([]int, 0)
	for iter := fm.Range(); !iter.HasNext(); iter.Next() {
		k, v := iter.Val()
		if value, ok := fm.Get(k); !ok || v != value {
			t.Error("range got wrong key/value")
		}
		iteration = append(iteration, k)
	}
	if !slices.Equal(data, iteration) {
		t.Errorf("keys not sort, want %v, got %v", iteration, data)
	}
}

func TestNewFlatMap(t *testing.T) {
	allocs := uint64(testing.AllocsPerRun(10, func() {
		const length = 10
		fm := NewFlatMap[int, int](length)
		if fm.Length() != 0 {
			t.Fatal("NewFlatMap can only set caps, but got length")
		}
		for i := 0; i < length; i++ {
			fm.Set(i, i+1)
		}
	}))

	if allocs > 1 {
		t.Fatalf("want at most 1 allocation, but got: %v", allocs)
	}
}
