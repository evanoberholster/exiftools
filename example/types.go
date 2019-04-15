package main

type Rational struct {
	num, denom		int64 // Numerator and Denominator
}

func (r Rational) Set(n, d int64) Rational{
	r.num = n
	r.denom = d
	return r
}

/////



