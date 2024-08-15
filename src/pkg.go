package src

import calSymbol2 "github.com/FITLOSS/GoCalSymbol"

func MixCompute(formula string, nums map[rune]float64) float64 {
	cal := calSymbol2.NewStruct(len(formula) + 1)
	cal.GiveRule(formula)
	for s, f := range nums {
		cal.Set(s, f)
	}
	return cal.Compute()
}
