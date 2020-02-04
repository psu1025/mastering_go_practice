package cmdcalc

// AvgAllVariable is avg all parameters
func AvgAllVariable(vars []string) float64 {
	var total, length = AddAllVariable(vars)
	if total != 0 {
		return float64(total) / float64(length)
	}

	return 0.0
}
