package mathematics

import "math"

// Function to calculate the mean / average with variadic parameter
func Mean[T float32 | float64](elements ...T) float64 {
	var sum float64
	for _, e := range elements {
		sum += float64(e)
	}
	return sum / float64(len(elements))
}

// Function to calculate standard deviation with variadic parameter
func StandardDeviation[T float32 | float64](elements ...T) float64 {
	m := Mean(elements...)

	var varianceSum float64
	for _, e := range elements {
		diff := float64(e) - m
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(len(elements)-1)

	return math.Sqrt(variance)
}
