package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T){
	repeated := Iteration("a", 5)
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkIteration(b *testing.B){
	for b.Loop(){
		Iteration("a", 5)
	}
}

func ExampleIteration(){
	fmt.Println(Iteration("a", 5))
	//Output: aaaaa
}