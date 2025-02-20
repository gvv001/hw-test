package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {

	type MapItem struct {
		value int
		key   string
	}

	slice := strings.Fields(str)

	dict := make(map[string]int, len(slice))

	for _, value := range slice {
		dict[value]++
	}

	newSlice := make([]MapItem, 0, len(dict))

	for key, value := range dict {
		newSlice = append(newSlice, MapItem{key: key, value: value})
	}

	sort.Slice(newSlice, func(i, j int) bool {
		return newSlice[i].key > newSlice[j].key
	})

	sort.Slice(newSlice, func(i, j int) bool {
		return newSlice[i].value > newSlice[j].value
	})

	top10Slice := make([]string, 0)

	for i, v := range newSlice {
		if i >= 10 {
			break
		}
		top10Slice = append(top10Slice, v.key)
	}

	return top10Slice
}
