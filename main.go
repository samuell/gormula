package main

import "fmt"

func main() {

	// --------------------------------------------------------------------------------
	// Monthly Payment Formula
	// --------------------------------------------------------------------------------
	// Implementing the loan payments formula from here:
	// http://www.math.utah.edu/~pa/math/equations/equations.html
	// --------------------------------------------------------------------------------

	// We try the formula over a sequence of values from 1000 to 10,000 USD
	borrowedAmount := Seq(1000.0, 100.0, 10000.0)

	// The formula
	monthlyPaymentUSD := Mult(
		Div(
			Mult(
				Div(rate, X),
				Exp(X, Val(months))),
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
	for monthPay := range monthlyPaymentUSD {
		fmt.Printf("Monthly payment: %f\n", monthPay)
	}
}
