package queue

import "testing"

func BenchmarkSliced_InsertInt(b *testing.B) {
	object := Sliced[int]{}
	for i := 0; i < b.N; i++ {
		object.Insert(i)
	}
}

func BenchmarkSliced_InsertString(b *testing.B) {
	object := Sliced[string]{}
	for i := 0; i < b.N; i++ {
		object.Insert("hello world")
	}
}
