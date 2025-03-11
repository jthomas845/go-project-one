package main

import "testing"

func TestStrain(t *testing.T) {
	n := 50
	threshold := 20 //keep threshold LESS than n.

	signal := "Welcome to My App"
	altSignal := "We are testing in CI/CD"
	for x := range n {
		main()
		if x == 2 {
			writeToFile(signal, "signal.txt")
		}
		if x == threshold {
			writeToFile(altSignal, "signal.txt")
		}
	}
}
