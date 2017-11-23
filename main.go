package main

import (
	"fmt"
	"math"
)

func main() {
	// --------------------------------------------------------------------------------
	// Monthly Payment Formula
	// --------------------------------------------------------------------------------
	// Implementing the loan payments formula from here:
	// http://www.math.utah.edu/~pa/math/equations/equations.html
	// --------------------------------------------------------------------------------

	// We try the formula over a sequence of values from 1000 to 10,000 USD
	borrowedAmount := Seq(1000.0, 1000.0, 10000.0)

	// Some initializations
	rate := 2.0 // Percent
	months := 24.0

	// The formula
	monthlyPaymentUSD := Mul(
		Div(
			Mul(
				Div(Val(rate), Val(1200.0)),
				Exp(
					Add(
						Val(1.0),
						Div(
							Val(rate),
							Val(1200.0))),
					Val(months))),
			Sub(
				Exp(
					Add(
						Val(1.0),
						Div(
							Val(rate),
							Val(1200.0))),
					Val(months)),
				Val(1.0))),
		borrowedAmount)

	// Print out all the resulting monthly payments:
	borrowedAmountForPrint := Seq(1000.0, 100.0, 10000.0)
	for monthPay := range monthlyPaymentUSD {
		borrowed := <-borrowedAmountForPrint
		fmt.Printf("Monthly payment for 24 months, when borrowing %.2f USD: %.2f USD\n", borrowed, monthPay)
	}
}

// --------------------------------------------------------------------------------
// Components
// --------------------------------------------------------------------------------

type valstream chan float64

func Add(x valstream, y valstream) valstream {
	return Apply2(func(x float64, y float64) float64 { return x + y }, x, y)
}

func Sub(x valstream, y valstream) valstream {
	return Apply2(func(x float64, y float64) float64 { return x - y }, x, y)
}

func Mul(x valstream, y valstream) valstream {
	return Apply2(func(x float64, y float64) float64 { return x * y }, x, y)
}

func Div(x valstream, y valstream) valstream {
	return Apply2(func(x float64, y float64) float64 { return x / y }, x, y)
}

func Exp(x valstream, y valstream) valstream {
	return Apply2(func(x float64, y float64) float64 { return math.Pow(x, y) }, x, y)
}

func Apply2(fn func(x float64, y float64) float64, xs valstream, ys valstream) valstream {
	zs := make(valstream)
	go func() {
		defer close(zs)
		for x := range xs {
			y := <-ys
			zs <- fn(x, y)
		}
	}()
	return zs
}

func Val(x float64) valstream {
	xs := make(valstream)
	go func() {
		defer close(xs)
		for i := 0; i < 10; i++ {
			xs <- x
		}
	}()
	return xs
}

func Seq(start float64, step float64, end float64) valstream {
	res := make(valstream)
	go func() {
		defer close(res)
		val := start
		for (end + val - val) > 0.001 { // Same as val <= end, but take care of propagating float errors
			res <- val
			val = val + step
		}
	}()
	return res
}
