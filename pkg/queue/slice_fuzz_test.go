package queue

import "testing"

func FuzzSliced_InsertInt(f *testing.F) {
	object := Sliced[int]{}

	testCases := []int{13, 843, 0, -99, 1e8}
	for _, tc := range testCases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, value int) {
		object.Insert(value)
	})
}

func FuzzSliced_InsertString(f *testing.F) {
	object := Sliced[string]{}

	testCases := []string{"", "hello", "!123456", "Veronika"}
	for _, tc := range testCases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, value string) {
		object.Insert(value)
	})
}
