package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

/*

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

*/

func main() {
	data, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	names := strings.Split(strings.Replace(string(data), `"`, "", -1), ",")
	sort.Strings(names)
	var totalScore int
	for i := 0; i < len(names); i++ {
		totalScore += getNameScore(names[i], i+1)
	}
	fmt.Printf("Total score = %d\n", totalScore) //outputs: 871198282
}

func getNameScore(name string, index int) int {
	var score int
	for i := 0; i < len(name); i++ {
		score += int(name[i] - 64)
	}
	return score * index
}
