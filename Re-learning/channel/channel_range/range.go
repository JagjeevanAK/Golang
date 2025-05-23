package channelrange

func concurrentFib(n int) []int {
	s := make(chan int)

	go fibonacci(n, s)

	var sl []int

	for i:= range s {
		sl = append(sl, i)
	}

	return sl
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
