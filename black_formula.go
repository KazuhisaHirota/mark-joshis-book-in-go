package main

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func BlackFormula(
	forward float64, vol float64, expiry float64, strike float64,
	callPut OptionType, annuity float64) float64 {

	var logMoneyness = math.Log(forward / strike)
	var stdDeviation = vol * math.Sqrt(expiry)

	var d1 = logMoneyness/stdDeviation + stdDeviation
	var d2 = d1 - stdDeviation

	var normal = distuv.Normal{Mu: 0., Sigma: 1.}
	var flag = 1.
	if callPut == Put {
		flag = -1.
	}
	return annuity * flag * (forward*normal.CDF(flag*d1) -
		strike*normal.CDF(flag*d2))
}

func BSFormula(
	spot float64, r float64, vol float64,
	expiry float64, strike float64, callPut OptionType) float64 {

	var df = math.Exp(-r * expiry)
	var forward = spot / df
	return BlackFormula(forward, vol, expiry, strike, callPut, df)
}
