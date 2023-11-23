package util

import (
	"math"
)

func MeanSquaredError(yTrue []float64, yPred []float64) float64 {
	if len(yTrue) != len(yPred) {
		panic("Input arrays must have the same length")
	}

	n := float64(len(yTrue))
	var sumSquaredErrors float64

	for i := 0; i < len(yTrue); i++ {
		squaredError := math.Pow(yTrue[i]-yPred[i], 2)
		sumSquaredErrors += squaredError
	}

	mse := sumSquaredErrors / n
	return mse
}
