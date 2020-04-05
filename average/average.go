package average

// Average should be return everage
func Average(grade []float64) float64 {
	return sum(grade) / float64(len(grade))
}

func sum(grade []float64) float64 {
	result := 0.0
	for _, value := range grade {
		result += value
	}

	return result
}
