package mytest

func SumElement(a []float64) float64 {
	sum := float64(0)

	if len(a) == 0 {
		return sum
	}
	for _, x := range a {
		sum += x
	}
	return sum

}
