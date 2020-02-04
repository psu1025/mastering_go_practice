package cmdcalc

// AvgAllVariable is avg all parameters
func AvgAllVariable(vars []string) (avg float64) {
	avg = 0
	total := AddAllVariable(vars)
	if total != 0 {
		return float64(total) / float64(len(vars))
	}

	return 0.0
}
