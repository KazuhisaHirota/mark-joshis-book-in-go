package main

type VanillaOption struct {
	expiry float64
	payoff IPayoff
}

func NewVanillaOption(expiry float64, payoff IPayoff) *VanillaOption {
	return &VanillaOption{expiry: expiry, payoff: payoff}
}

func (o VanillaOption) Payoff(spot float64) float64 {
	return o.payoff.Calc(spot)
}
