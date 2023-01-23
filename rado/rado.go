package rado

import (
	"math"
)

func RadoTM(n float64) float64 {
	return math.Pow((4 * (n + 1)), 2*n)
}
