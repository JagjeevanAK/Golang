package function

import "errors"

func Divide (x, y int) (float64, int, error){
	if (y == 0){
		return 0, 0, errors.New("Denominator is zero");
	} else {
		return float64(x/y), x%y, nil
	}
}