package main

import "testing"

// Test when no retry should be necessary
func TestNoRetryNecessary(t *testing.T) {
	count := 0
	RunRetry(
		func() *string {
			count++
			return nil
		},
		5,
		1)

	if count != 1 {
		t.Error("Expected count == 1, got", count)
	}
}

// Test when retry is necessary, but eventually succeeds
func TestLessThanMaxRetryNecessary(t *testing.T) {
	count := 0
	RunRetry(
		func() *string {
			count++
			if count < 3 {
				ret := "value"
				return &ret
			}
			return nil
		},
		5,
		1)

	if count != 3 {
		t.Error("Expected count == 3, got", count)
	}
}

// Test that it will only retry max times and then stop if the function is never successful
func TestStopAtMax(t *testing.T) {
	count := 0
	RunRetry(
		func() *string {
			count++
			ret := "value"
			return &ret
		},
		4,
		1)

	if count != 4 {
		t.Error("Expected count == 4, got", count)
	}
}
