package generic

func getLast[T any](s []T) T {
	if len(s) ==0 {
		var Var T
		return Var
	}
	return s[len(s)-1]
}