package flatmap

import "iter"

func (f *FlatMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := range f.data {
			pair := f.data[i]
			yield(pair.First, pair.Second)
		}
	}
}

func (f *FlatMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for i := range f.data {
			yield(f.data[i].First)
		}
	}
}

func (f *FlatMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range f.data {
			yield(f.data[i].Second)
		}
	}
}
