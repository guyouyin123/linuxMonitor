package src

import "github.com/guyouyin123/tools/qmixCompute"

func MixCompute(formula string, nums map[rune]float64) float64 {
	cal := qmixCompute.NewStruct(len(formula) + 1)
	cal.GiveRule(formula)
	for s, f := range nums {
		cal.Set(s, f)
	}
	return cal.Compute()
}
