package mathematics

import (
	"math"
	"math/big"
)

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

// Calculates the mean (average) of variadic *big.Float values
func MeanBigFloat(precisionBits uint, elements ...*big.Float) *big.Float {
	sum := new(big.Float).SetPrec(precisionBits).SetFloat64(0)

	for _, e := range elements {
		sum.Add(sum, e)
	}

	count := new(big.Float).SetPrec(precisionBits).SetInt64(int64(len(elements)))
	mean := new(big.Float).SetPrec(precisionBits).Quo(sum, count)

	return mean
}

// Calculates the standard deviation of variadic *big.Float values
func StandardDeviationBigFloat(precisionBits uint, elements ...*big.Float) *big.Float {
	if len(elements) <= 1 {
		// Return 0 if there's only one or no element to avoid division by zero
		return big.NewFloat(0).SetPrec(precisionBits)
	}

	mean := MeanBigFloat(precisionBits, elements...)
	varianceSum := new(big.Float).SetPrec(precisionBits).SetFloat64(0)

	for _, e := range elements {
		diff := new(big.Float).SetPrec(precisionBits).Sub(e, mean)
		squared := new(big.Float).SetPrec(precisionBits).Mul(diff, diff)
		varianceSum.Add(varianceSum, squared)
	}

	// Divide by (n - 1) for sample standard deviation
	nMinusOne := new(big.Float).SetPrec(precisionBits).SetInt64(int64(len(elements) - 1))
	variance := new(big.Float).SetPrec(precisionBits).Quo(varianceSum, nMinusOne)

	// Return the square root of the variance
	stdDev := new(big.Float).SetPrec(precisionBits).Sqrt(variance)
	return stdDev
}
