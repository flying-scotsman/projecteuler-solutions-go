package sortNames

// This is the Lomuto implementation of the quicksort algorithm! Awesome

func Quicksort(names []string, low, high int) {
	if low < high {
		p := partition(names, low, high)
		Quicksort(names, low, p-1)
		Quicksort(names, p+1, high)
	}
}

func partition(names []string, low, high int) int {
	pivot := names[high]
	i := low
	for j := low; j <= high-1; j++ {
		// compareNames returns true if names[j] <= pivot
		if compareNames(names, j, pivot) {
			names[i], names[j] = names[j], names[i]
			i = i + 1
		}
	}
	// All elements lower than the pivot are now below it --
	// we swap the remaining high term
	names[i], names[high] = names[high], names[i]
	return i
}
