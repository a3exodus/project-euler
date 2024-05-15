package helpers

type Fraction struct {
	Numerator   int
	Denominator int
}

func FractionValue(f Fraction) float64 {
	return float64(f.Numerator) / float64(f.Denominator)
}
