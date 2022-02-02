package main

import (
	"fmt"
	"strconv"
)

func main() {
	var spot float64 = 100.
	var r float64 = 0.01
	var vol float64 = 0.3
	var expiry float64 = 1.
	var strike float64 = 110.
	var lowerLevel float64 = 90.
	var upperLevel float64 = 130.
	var numPath int = 10000

	var bsCall = BSFormula(spot, r, vol, expiry, strike, Call)
	fmt.Println("BS call price is " + strconv.FormatFloat(bsCall, 'f', -1, 64))
	var bsPut = BSFormula(spot, r, vol, expiry, strike, Put)
	fmt.Println("BS put price is " + strconv.FormatFloat(bsPut, 'f', -1, 64))

	TestMonteCarlo(spot, r, vol, expiry, strike, numPath)
	TestMonteCarlo2(spot, r, vol, expiry, strike, lowerLevel, upperLevel, numPath)
	TestMonteCarlo3(spot, r, vol, expiry, strike, lowerLevel, upperLevel, numPath)
}
