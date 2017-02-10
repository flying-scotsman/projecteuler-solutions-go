package sortNames

// Returns true if names[j] is alphabetically before or equal to pivot
func compareNames(names []string, j int, pivot string) bool {
	var maxCompLength int

	// Check if names are equal
	if names[j] == pivot {
		return true
	}

	// Don't want to compare bytes past the length of the shortest string
	if len(names[j]) < len(pivot) || len(names[j]) == len(pivot) {
		maxCompLength = len(names[j])
	} else {
		maxCompLength = len(pivot)
	}

	// As soon as a letter in names[j] is alphabetically lower than the corresponding one in pivot, condition is true
	for i := 0; i < maxCompLength; i++ {
		if names[j][i] < pivot[i] {
			return true
		} else if names[j][i] == pivot[i] {
			continue
		} else {
			return false
		}
	}

	// If e.g. matt and matthew come up, matt comes before matthew
	if len(pivot) > maxCompLength {
		return true
	}
	return false
}
