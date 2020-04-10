package iq

import (
	"os"
	"fmt"
)

const (
	min = 0
	max = 200
)

func computeDensityFunction() {
	for i := min; i < max + 1; i++ {
		res := ComputeGaussianDistribution(float64(U), float64(S), float64(i))
		fmt.Printf("%d %.5f\n", i, res)
	}
}

func computePercentagePeopleIQInferior(IQ int) float64 {
	var res float64 = 0.0
	var incrementor = 0.001

	for i := 0.0; i < float64(IQ); i += incrementor {
		lhs := ComputeGaussianDistribution(float64(U), float64(S), float64(i))
		rhs := ComputeGaussianDistribution(float64(U), float64(S), float64(i + incrementor))
		res += incrementor * (lhs + rhs) / 2
	}
	percent := res * 100
	return percent
}

func computePercentagePeopleIQBetweenThose() {
	IQ1Inferior := computePercentagePeopleIQInferior(IQ1)
	IQ2Inferior := computePercentagePeopleIQInferior(IQ2)
	percent := IQ2Inferior - IQ1Inferior
	fmt.Printf("%.1f%% of people have an IQ between %d and %d\n", percent, IQ1, IQ2)
}

// IQ - main
func IQ() {
	argsWithoutProg := os.Args[1:]
	nbArgs := len(argsWithoutProg)

	if (nbArgs == 2) {
		computeDensityFunction();
	}
	if (nbArgs == 3) {
		res := computePercentagePeopleIQInferior(IQ1);
		fmt.Printf("%.1f%% of people have an IQ inferior to %d\n", res, IQ1)
	}
	if (nbArgs == 4) {
		computePercentagePeopleIQBetweenThose();
	}
}