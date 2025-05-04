package array

import "testing"

func TestArray(t *testing.T){
	num := []int {1, 2, 3, 4};

	got := Sum(num);
	want := 10;

	if got != want {
		t.Errorf("Got %d instead of %d", got, want);
	}
}