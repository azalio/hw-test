package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var maxWordsInResultSlice = 10

var re = regexp.MustCompile(`\p{L}+-?(\p{L}+)?`) // find in unicode too

func clearWords(s []string) []string {
	result := []string{}
	for _, w := range s {
		tString := re.FindString(w) // get only letters and hyphen
		if tString == "" {
			continue
		}
		result = append(result, strings.ToLower(tString))
	}

	return result
}

func Top10(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	sSlice := strings.Fields(s) // get words
	sSlice = clearWords(sSlice)

	sMap := make(map[string]int)

	for _, w := range sSlice { // get freq
		sMap[w]++
	}

	iMap := make(map[int][]string)

	for k, v := range sMap { // freq to key and word to slice
		iMap[v] = append(iMap[v], k)
	}

	iSort := []int{}
	for k := range iMap { // keys to int slice for sorting
		iSort = append(iSort, k)
	}

	sort.Ints(iSort)

	counter := maxWordsInResultSlice
	var k int
	resultStrings := []string{}
	for len(iSort) > 0 && counter > 0 {
		k, iSort = iSort[len(iSort)-1], iSort[:len(iSort)-1] // max word count and rest of slice
		v := iMap[k]                                         // get words slice by key
		sort.Strings(v)
		counter -= len(v)
		sLen := maxWordsInResultSlice - len(resultStrings)
		if sLen > len(v) { // boundaries
			sLen = len(v)
		}
		resultStrings = append(resultStrings, v[:sLen]...)
	}

	return resultStrings
}
