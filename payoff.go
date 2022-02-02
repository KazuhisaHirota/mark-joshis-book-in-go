package main

import (
	"fmt"
	"math"
)

type OptionType int

const (
	Call OptionType = iota
	Put
)

type Payoff struct {
	strike     float64
	optionType OptionType
}

func NewPayoff(strike float64, optionType OptionType) *Payoff {
	return &Payoff{strike, optionType}
}

func (p Payoff) Calc(spot float64) float64 {
	switch p.optionType {
	case Call:
		return math.Max(spot-p.strike, 0.)
	case Put:
		return math.Max(p.strike-spot, 0.)
	default:
		fmt.Printf("unknown OptionType")
		return 0.
	}
}

type IPayoff interface {
	Calc(spot float64) float64
}

type StrikedTypePayoff struct {
	strike float64
}

func NewStrikedTypePayoff(strike float64) *StrikedTypePayoff {
	return &StrikedTypePayoff{strike}
}

type PayoffCall struct {
	*StrikedTypePayoff // embedded
}

func NewPayoffCall(strike float64) *PayoffCall {
	return &PayoffCall{NewStrikedTypePayoff(strike)}
}

func (p PayoffCall) Calc(spot float64) float64 {
	return math.Max(spot-p.strike, 0.)
}

type PayoffPut struct {
	*StrikedTypePayoff // embedded
}

func NewPayoffPut(strike float64) *PayoffPut {
	return &PayoffPut{NewStrikedTypePayoff(strike)}
}

func (p PayoffPut) Calc(spot float64) float64 {
	return math.Max(p.strike-spot, 0.)
}

type PayoffDoubleDigital struct {
	lowerLevel float64
	upperLevel float64
}

func NewPayoffDoubleDigital(lowerLevel float64, upperLevel float64) *PayoffDoubleDigital {
	return &PayoffDoubleDigital{lowerLevel, upperLevel}
}

func (p PayoffDoubleDigital) Calc(spot float64) float64 {
	if spot <= p.lowerLevel {
		return 0.
	} else if spot >= p.upperLevel {
		return 0.
	} else {
		return 1.
	}
}
