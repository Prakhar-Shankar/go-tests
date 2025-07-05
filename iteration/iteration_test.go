package iteration

import "testing"

func TestIteration(t *testing.T){
	repeated := Iteration("a")
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkIteration(b *testing.B){
	for b.Loop(){
		Iteration("a")
	}
}