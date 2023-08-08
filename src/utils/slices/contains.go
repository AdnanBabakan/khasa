package slices

func Contains(slice interface{}, element interface{}) bool {
	switch slice.(type) {
	case []string:
		for _, s := range slice.([]string) {
			if s == element {
				return true
			}
		}
	case []int:
		for _, s := range slice.([]int) {
			if s == element {
				return true
			}
		}
	case []float64:
		for _, s := range slice.([]float64) {
			if s == element {
				return true
			}
		}
	case []bool:
		for _, s := range slice.([]bool) {
			if s == element {
				return true
			}
		}
	case []interface{}:
		for _, s := range slice.([]interface{}) {
			if s == element {
				return true
			}
		}
	}
	return false
}
