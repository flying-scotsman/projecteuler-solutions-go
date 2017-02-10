package main

import (
	"encoding/csv"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Matthew/projecteuler-solutions-go/22/sortNames"
)

func BenchmarkHello(b *testing.B) {
	for i := 1; i < b.N; i++ {
		// Connect to file
		file, errRead := ioutil.ReadFile("names.txt") // type []byte
		// Initialize a new csv reader that will read a string
		reader := csv.NewReader(strings.NewReader(string(file)))
		// Use the ReadAll method to read out all the names
		names, errNames := reader.ReadAll()
		check(errRead)
		check(errNames)

		// Sort the names
		sortNames.Quicksort(names[0], 0, len(names[0])-1) // Since test is a reference type, its actual value gets changed by quicksort - freakin' A!
		nameScore(names[0])
	}
}
