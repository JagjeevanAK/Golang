package struc

import "testing"

func TestStruct (t *testing.T){
	t.Run("Circle", func (t *testing.T){
		circle := Circle{ 100 };
		got := circle.Area()
		want := 314.0

		if got != want {
			t.Errorf("Got %f instead of %f", got, want);
		}
	})

	t.Run("Rectangle", func(t *testing.T){
		r := Rectangle {10, 10};
		got := r.Area();
		want := 100;

		if got != want {
			t.Errorf("Got %d instead of %d", got, want);
		}
	})
}