package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	poolSize, totalJobs := 5, 50
	pool, jobDone, allDone := MakePool(poolSize, totalJobs)

	fmt.Printf("Start jobs \n \n")

	for i := 0; i < totalJobs; i++ {

		go func(job int) {
			<-pool
			defer func() { jobDone <- true }()

			fmt.Printf("Starting job %d \n", job)

			time.Sleep(500 * time.Millisecond)

			fmt.Printf("Ending job %d \n", job)
		}(i + 1)
	}
	<-allDone

	fmt.Printf(" \nEnd jobs")
}

// END OMIT
func MakePool(poolSize, jobs int) (pool, jobDone, allDone chan bool) {
	// Channel to coordinate the number of concurrent goroutines.
	pool = make(chan bool, poolSize)
	// Fill the channel with tokens.
	for i := 0; i < poolSize; i++ {
		pool <- true
	}
	// The jobDone channel indicates when a single goroutine has finished its job.
	jobDone = make(chan bool)
	// The allDone channel allows the main program to wait until all the jobs are done.
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
