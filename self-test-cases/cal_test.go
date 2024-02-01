package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type FixFunc func(x float64, min float64, max float64) float64

func DE(individualA float64, individualB float64, min float64, max float64, fixElement FixFunc) (float64, float64) {
	var (
		mu      float64 = 2
		randNum float64 = rand.Float64()
		alpha   float64 = 0
	)
	if randNum <= 0.5 {
		alpha = math.Pow(2.0*randNum, (1.0 / (mu + 1)))
	} else {
		alpha = math.Pow((1 / (2 * (1 - randNum))), (1.0 / (mu + 1)))
	}

	child1 := 0.5 * ((1+alpha)*individualA + (1-alpha)*individualB)
	child2 := 0.5 * ((1-alpha)*individualA + (1+alpha)*individualB)

	child1 = fixElement(child1, min, max)
	child2 = fixElement(child2, min, max)

	return child1, child2
}

func FixElement(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func f1(x float64) float64 {
	return -(x*math.Sin(10*math.Pi*x) + 2)
}

func realX(x float64, min float64, max float64) float64 {
	return x*(max-min) + min
}

func TestCal(t *testing.T) {
	var (
		spanMin float64 = -1
		spanMax float64 = 2
		x1      float64 = realX(rand.Float64(), spanMin, spanMax)
		x2      float64 = realX(rand.Float64(), spanMin, spanMax)
		iter    int     = 50
	)

	fv1, fv2 := f1(x1), f1(x2)
	if fv1 > fv2 {
		fv1, fv2 = fv2, fv1
		x1, x2 = x2, x1
	}

	globalX, globalFv := x1, fv1

	var count int = 0

	for i := 1; i <= iter; i++ {
		newX1, newX2 := DE(x1, x2, spanMin, spanMax, FixElement)
		newFv1, newFv2 := f1(newX1), f1(newX2)

		tempX, tempY := x1, fv1
		if newFv1 < fv1 {
			x1, fv1 = newX1, newFv1
			newX1, newFv1 = tempX, tempY
			x2, fv2 = newX1, newFv1
		} else if newFv1 < fv2 {
			x2, fv2 = newX1, newFv1
		}

		tempX, tempY = x1, fv1
		if newFv2 < fv1 {
			x1, fv1 = newX2, newFv2
			newX2, newFv2 = tempX, tempY
			x2, fv2 = newX2, newFv2
		} else if newFv2 < fv2 {
			x2, fv2 = newX2, newFv2
		}

		if fv1 < globalFv {
			globalX, globalFv = x1, fv1
		}

		if fv2-fv1 < 0.001 {
			count++
		}
		if count >= 5 {
			x1 = realX(rand.Float64(), spanMin, spanMax)
			x2 = realX(rand.Float64(), spanMin, spanMax)
			fv1, fv2 = f1(x1), f1(x2)
			if fv1 > fv2 {
				fv1, fv2 = fv2, fv1
				x1, x2 = x2, x1
			}
			count = 0
		}
	}

	fmt.Printf("f(%v)=%v\n", globalX, -globalFv)
}

func TestRandCal(t *testing.T) {
	var (
		spanMin float64 = -1
		spanMax float64 = 2
		iter    int     = 50
	)
	var x3, fv3 float64 = -10, math.MaxFloat64
	for i := 1; i <= iter; i++ {
		randX := realX(rand.Float64(), spanMin, spanMax)
		randFv3 := f1(randX)
		if randFv3 < fv3 {
			fv3 = randFv3
			x3 = randX
		}
	}

	fmt.Printf("f(%v)=%v\n", x3, -fv3)
}
