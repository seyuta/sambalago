package generator

import (
	"math/rand"
	"time"
)

// GenerateCircularSequenceInRange function generates a circular sequence of integers within a specified range.
// Parameters:
// minVal: The minimum value allowed in the sequence (inclusive).
// maxVal: The maximum value allowed in the sequence (inclusive).
// startVal: The starting value for the sequence.
// amount: The number of elements to generate in the sequence.
// Returns: A slice containing the generated circular sequence.
func GenerateCircularSequenceInRange(minVal, maxVal, startVal, amount int) []int {
	sequence := make([]int, amount)
	for i := 0; i < amount; i++ {
		if startVal > maxVal {
			startVal = minVal
		}
		sequence[i] = startVal
		startVal++
	}

	return sequence
}

// GenerateRandomNumber will generate random numbers with the desired number of digits
func GenerateRandomNumber(digits int) int {
	if digits <= 0 {
		return 0
	}

	// Calculates the lower and upper limits for the desired number of digits.
	min := 1
	for i := 1; i < digits; i++ {
		min *= 10
	}
	max := min*10 - 1

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
