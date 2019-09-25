package main

// START OMIT
func MakePool(poolSize, jobs int) (pool, jobDone, allDone chan bool) {
	// Channel to coordinate the number of concurrent goroutines.
	pool = make(chan bool, poolSize)

	// Fill the channel with tokens.
	for i := 0; i < poolSize; i++ {
		pool <- true
	}

	jobDone = make(chan bool)
	allDone = make(chan bool)

	go func() {
		for i := 0; i < jobs; i++ {
			// Add a token each time a job is done
			<-jobDone
			pool <- true
		}
		allDone <- true
	}()

	return
}

// END OMIT
