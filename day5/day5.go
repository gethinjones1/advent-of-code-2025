package day5

import (
	"slices"
	"strconv"
	"strings"
)

func CheckFoodIsFresh(idRange []string, idToCheck int) int {
	var foundIds []int
	for _, v := range idRange {
		splitIdRange := strings.Split(v, "-")
		from, _ := strconv.Atoi(splitIdRange[0])
		to, _ := strconv.Atoi(splitIdRange[1])

		if from == idToCheck || to == idToCheck || (idToCheck > from && idToCheck < to) {
			foundIds = append(foundIds, idToCheck)
		}
	}

	uniqueIds := CreateUnique(foundIds)

	return len(uniqueIds)
}

func CreateUnique(ids []int) []int {
	var unique []int
	for _, v := range ids {
		if !slices.Contains(unique, v) {
			unique = append(unique, v)
		}
	}
	return unique
}
