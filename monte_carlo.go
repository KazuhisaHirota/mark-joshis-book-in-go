package main

import (
	"math"
	"math/rand"
)

func SimpleMonteCarlo(expiry float64, strike float64, callPut OptionType, spot float64, vol float64, r float64, numPath int) float64 {
	var variance = vol * vol * expiry
	var rootVariance = math.Sqrt(variance)
	var itoCorrection = -0.5 * variance

	var movedSpot = spot * math.Exp(r*expiry+itoCorrection)
	var runningSum = 0.
	rand.Seed(1234)

	for i := 0; i < numPath; i++ {
		var thisGaussian = rand.NormFloat64()
		var thisSpot = movedSpot * math.Exp(rootVariance*thisGaussian)

		var callPutflag = 1.
		if callPut == Put {
			callPutflag = -1.
		}
		var thisPayoff = math.Max(callPutflag*(thisSpot-strike), 0.) // NOTE

		runningSum += thisPayoff
	}

	var mean = runningSum / float64(numPath)
	mean *= math.Exp(-r * expiry)
	return mean
}

func SimpleMonteCarlo2(payoff IPayoff, expiry float64, spot float64, vol float64, r float64, numPath int) float64 {
	var variance = vol * vol * expiry
	var rootVariance = math.Sqrt(variance)
	var itoCorrection = -0.5 * variance

	var movedSpot = spot * math.Exp(r*expiry+itoCorrection)
	var runningSum float64 = 0
	rand.Seed(1234)

	for i := 0; i < numPath; i++ {
		var thisGaussian float64 = rand.NormFloat64()
		var thisSpot float64 = movedSpot * math.Exp(rootVariance*thisGaussian)
		var thisPayoff float64 = payoff.Calc(thisSpot) // NOTE
		runningSum += thisPayoff
	}

	var mean float64 = runningSum / float64(numPath)
	mean *= math.Exp(-r * expiry)
	return mean
}

func SimpleMonteCarlo3(theOption VanillaOption, spot float64, vol float64, r float64, numPath int) float64 {
	var expiry float64 = theOption.expiry // NOTE
	var variance float64 = vol * vol * expiry
	var rootVariance float64 = math.Sqrt(variance)
	var itoCorrection = -0.5 * variance

	var movedSpot = spot * math.Exp(r*expiry+itoCorrection)
	var runningSum float64 = 0
	rand.Seed(1234)

	for i := 0; i < numPath; i++ {
		var thisGaussian float64 = rand.NormFloat64()
		var thisSpot float64 = movedSpot * math.Exp(rootVariance*thisGaussian)
		var thisPayoff float64 = theOption.Payoff(thisSpot) // NOTE
		runningSum += thisPayoff
	}

	var mean float64 = runningSum / float64(numPath)
	mean *= math.Exp(-r * expiry)
	return mean
}
