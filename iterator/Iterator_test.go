package iterator

import "testing"

func TestInterator(t *testing.T){
	got := loop("a");
	want := "aaaaaa";

	if got != want {
		t.Errorf("Got %q instead of %q", got, want)
	}
}