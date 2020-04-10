package iq

import (
    "math"
)

// ComputeGaussianDistribution - get normal distribution
func ComputeGaussianDistribution(mean float64, standardDeviation float64, x float64) float64 {
    exp := math.Exp(-((x - mean) * (x - mean) / (2 * standardDeviation * standardDeviation)))
    res := 1 / (standardDeviation * math.Sqrt(2 * math.Pi)) * exp

    return res
}
