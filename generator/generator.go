package generator

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
