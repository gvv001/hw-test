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
		return newSlice[i].key < newSlice[j].key
	})

	sort.SliceStable(newSlice, func(i, j int) bool {
		return newSlice[i].value > newSlice[j].value
	})

	top := make([]string, 0)

	for i, v := range newSlice {
		if i == 10 {
			break
		}
		top = append(top, v.key)
	}

	return top
}
