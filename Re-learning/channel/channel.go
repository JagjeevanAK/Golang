package channel

func countReports(numSentCh chan int) int {
	var totalReport int ;

	for{
		msg, ok := <- numSentCh;
		if !ok {
			break;
		}
		totalReport += msg ;
	}

	return totalReport;
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
