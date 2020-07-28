package main

import "testing"

type AddResult struct {
	x int
	y int
	expected int
}
var AddResults =[] AddResult{
	{10,20,30},
	{100,200,300},
	{1,2,3},
}
func TestAdd(t *testing.T) {
	for _, test := range AddResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatalf("Expected result Not found")
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Add(i,i)
	}
}
