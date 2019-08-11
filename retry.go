package main

import (
	"log"
	"strconv"
	"time"
)

// RunRetry retries the given function the given number of tries.
// funcToRetry Should return nil if successful and an error message if it should be retried
// numRetries is -1, then it'll retry forever.
// retryBackoffMillis
func RunRetry(funcToRetry func() *string, numRetries int, retryBackoffMillis int) {
	runForever := false
	if numRetries < 1 {
		runForever = true
		numRetries = 1
	}

	retryBackoff := retryBackoffMillis
	i := 0
	for i < numRetries {
		err := funcToRetry()

		if err != nil {
			log.Printf("Retry: %v", *err)
			log.Printf("Sleep for %v ms", retryBackoff)

			backoffDuration, _ := time.ParseDuration(strconv.Itoa(retryBackoff) + "ms")
			time.Sleep(backoffDuration)

			// Double the time before the next retry
			retryBackoff *= 2
			// Only increase time between retries until it hits 5 minutes
			if retryBackoff > 300000 {
				retryBackoff = 300000
			}

			if !runForever {
				i++
			}
		} else {
			break
		}
	}

}
