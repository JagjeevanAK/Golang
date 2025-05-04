package function

import "testing"

func TestDivide (t *testing.T){
	q_got, r_got, _ := Divide(2,1);
	q_want,r_want := 2.0, 0;

	if q_got != q_want {
		t.Errorf("Qutiont got %f instead of %f ",q_got, q_want);
		t.Errorf("Qutiont got %d instead of %d ",r_got, r_want);
	}
}