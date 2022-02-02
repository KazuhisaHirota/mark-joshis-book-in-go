package main

import (
	"fmt"
	"strconv"
)

func TestMonteCarlo(spot float64, r float64, vol float64, expiry float64, strike float64, numPath int) {
	var callPrice float64 = SimpleMonteCarlo(expiry, strike, Call, spot, vol, r, numPath)
	fmt.Println("MC call price is " + strconv.FormatFloat(callPrice, 'f', -1, 64))

	var putPrice float64 = SimpleMonteCarlo(expiry, strike, Put, spot, vol, r, numPath)
	fmt.Println("MC put price is " + strconv.FormatFloat(putPrice, 'f', -1, 64))
}

func TestMonteCarlo2(spot float64, r float64, vol float64, expiry float64, strike float64,
	lowerLevel float64, upperLevel float64, numPath int) {

	var callPayoff IPayoff = NewPayoffCall(strike)
	var callPrice float64 = SimpleMonteCarlo2(callPayoff, expiry, spot, vol, r, numPath)
	fmt.Println("MC call price is " + strconv.FormatFloat(callPrice, 'f', -1, 64))

	var putPayoff IPayoff = NewPayoffPut(strike)
	var putPrice float64 = SimpleMonteCarlo2(putPayoff, expiry, spot, vol, r, numPath)
	fmt.Println("MC put price is " + strconv.FormatFloat(putPrice, 'f', -1, 64))

	var digitalPayoff IPayoff = NewPayoffDoubleDigital(lowerLevel, upperLevel)
	var digitalPrice float64 = SimpleMonteCarlo2(digitalPayoff, expiry, spot, vol, r, numPath)
	fmt.Println("MC digital price is " + strconv.FormatFloat(digitalPrice, 'f', -1, 64))
}

func TestMonteCarlo3(spot float64, r float64, vol float64, expiry float64, strike float64,
	lowerLevel float64, upperLevel float64, numPath int) {

	var callPayoff IPayoff = NewPayoffCall(strike)
	var callOption *VanillaOption = NewVanillaOption(expiry, callPayoff)
	var callPrice float64 = SimpleMonteCarlo3(*callOption, spot, vol, r, numPath)
	fmt.Println("MC call price is " + strconv.FormatFloat(callPrice, 'f', -1, 64))

	var putPayoff IPayoff = NewPayoffPut(strike)
	var putOption *VanillaOption = NewVanillaOption(expiry, putPayoff)
	var putPrice float64 = SimpleMonteCarlo3(*putOption, spot, vol, r, numPath)
	fmt.Println("MC put price is " + strconv.FormatFloat(putPrice, 'f', -1, 64))

	var ditialPayoff IPayoff = NewPayoffDoubleDigital(lowerLevel, upperLevel)
	var digitalOption *VanillaOption = NewVanillaOption(expiry, ditialPayoff)
	var digitalPrice float64 = SimpleMonteCarlo3(*digitalOption, spot, vol, r, numPath)
	fmt.Println("MC digital price is " + strconv.FormatFloat(digitalPrice, 'f', -1, 64))
}
