package utils

import (
	"math"
)

func VectorLength(a *[]float64) float64 {
	var sum float64
	for _, v := range *a {
		sum += math.Pow(v, 2)
	}
	return math.Sqrt(sum)
}

func VectorScalarMul(a *float64, b *[]float64) (c *[]float64) {
	c = &[]float64{}
	for _, v := range *b {
		*c = append(*c, *a*v)
	}
	return
}

func VectorMul(a, b *[]float64) (c *[]float64) {
	c = &[]float64{}
	for i := 0; i < MinLength(a, b); i++ {
		*c = append(*c, (*a)[i]*(*b)[i])
	}
	return
}

func VectorSum(a, b *[]float64) (c *[]float64) {
	c = &[]float64{}
	for i := 0; i < MinLength(a, b); i++ {
		*c = append(*c, (*a)[i]+(*b)[i])
	}
	return
}

func VectorSub(a, b *[]float64) (c *[]float64) {
	c = &[]float64{}
	for i := 0; i < MinLength(a, b); i++ {
		*c = append(*c, (*a)[i]-(*b)[i])
	}
	return
}

func VectorLike(a *[]float64, b float64) (c *[]float64) {
	c = &[]float64{}
	for i := 0; i < len(*a); i++ {
		*c = append(*c, b)
	}
	return
}

func VectorMinValue(a, b *[]float64) (c *[]float64) {
	c = &[]float64{}
	for i := 0; i < MinLength(a, b); i++ {
		if math.Abs((*a)[i]) < math.Abs((*b)[i]) {
			*c = append(*c, (*a)[i])
		} else {
			*c = append(*c, (*b)[i])
		}
	}
	return
}

func MinLength(a, b *[]float64) int {
	return IntMin(len(*a), len(*b))
}

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
