package grade

// Calculate Grade must Return Grade
func Calculate(score float64) string {
	if score >= 80 {
		return "A"
	} else if score >= 75.0 {
		return "B+"
	} else if score >= 70.0 {
		return "B"
	} else if score >= 65.0 {
		return "C+"
	} else if score >= 60.0 {
		return "C"
	} else if score >= 55.0 {
		return "D+"
	} else if score >= 50.0 {
		return "D"
	} else {
		return "F"
	}
}
