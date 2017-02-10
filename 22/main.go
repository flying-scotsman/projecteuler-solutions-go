/*
Using names.txt (right click and 'Save Link/Target As...'),
a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out
the alphabetical value for each name, multiply this value by
 its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order,
 COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
 is the 938th name in the list.
 So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Matthew/projecteuler-solutions-go/22/sortNames"
)

func main() {
	// Connect to file
	file, errRead := ioutil.ReadFile("names.txt") // type []byte
	// Initialize a new csv reader that will read a string
	reader := csv.NewReader(strings.NewReader(string(file)))
	// Use the ReadAll method to read out all the names
	// The format is a slice of slice of strings with length one, i.e. there is only
	// one slice of strings in it
	names, errNames := reader.ReadAll()
	check(errRead)
	check(errNames)

	// Sort the names
	sortNames.Quicksort(names[0], 0, len(names[0])-1) // Since test is a reference type, its actual value gets changed by quicksort - freakin' A!
	namesScore := nameScore(names[0])
	fmt.Println(namesScore)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Each name's score is the sum of its alphabetical values multiplied by its place in the list
func nameScore(names []string) int {
	var sum int
	for i := 0; i < len(names); i++ {
		var nameSum int
		for j := 0; j < len(names[i]); j++ {
			nameSum += int(names[i][j]) - 64
		}
		sum += nameSum * (i + 1)
	}
	return sum
}
