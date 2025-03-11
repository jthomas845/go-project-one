package main

import "testing"

func TestStrain(t *testing.T) {
	n := 50
	threshold := 20 //keep threshold LESS than n.

	signal := "Welcome to My Go Application. This is the initial signal, but it will soon change."
	altSignal := "This is the NEW signal. You're likely reading this inside my CI/CD testing."
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
