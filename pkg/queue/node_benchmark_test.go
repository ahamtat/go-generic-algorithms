package queue

import "testing"

func BenchmarkNodes_InsertInt(b *testing.B) {
	object := Nodes[int]{}
	for i := 0; i < b.N; i++ {
		object.Insert(i)
	}
}

func BenchmarkNodes_InsertString(b *testing.B) {
	object := Nodes[string]{}
	for i := 0; i < b.N; i++ {
		object.Insert("hello world")
	}
}
