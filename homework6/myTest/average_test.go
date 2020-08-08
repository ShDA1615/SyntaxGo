package mytest

import "testing"

type TestSums struct {
	Num      int
	Value    []float64
	Expected float64
}

var testSum = []TestSums{
	{
		Num:      0,
		Value:    []float64{0, 1, 2, 4},
		Expected: 7,
	},
	{
		Num:      1,
		Value:    []float64{0, 2.5, 7, 0.5},
		Expected: 10,
	},
	{
		Num:      2,
		Value:    []float64{},
		Expected: 0,
	},
	{
		Num:      3,
		Value:    []float64{},
		Expected: 5, //эмуляция ошибки
	},
}

func TestSumElement(t *testing.T) {
	for _, testSum := range testSum {
		result := SumElement(testSum.Value)
		if result != testSum.Expected {
			t.Errorf("Номер теста: %d ожидаемая величина: %f.00  результат: %f.00",
				testSum.Num,
				testSum.Expected,
				result,
			)
		}
	}
}
