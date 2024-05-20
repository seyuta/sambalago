package generator

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
