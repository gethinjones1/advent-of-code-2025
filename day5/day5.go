package day5

import (
	"slices"
	"sort"
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

func CheckFoodIsFreshPt2(idRange []string) []string {
	// var foundIds []int
	sort.Slice(idRange, func(i, j int) bool {
		// Split left item
		partsI := strings.SplitN(idRange[i], "-", 2)
		a1, _ := strconv.ParseInt(partsI[0], 10, 64)
		a2, _ := strconv.ParseInt(partsI[1], 10, 64)

		// Split right item
		partsJ := strings.SplitN(idRange[j], "-", 2)
		b1, _ := strconv.ParseInt(partsJ[0], 10, 64)
		b2, _ := strconv.ParseInt(partsJ[1], 10, 64)

		// Primary sort: first number
		if a1 != b1 {
			return a1 < b1
		}

		// Secondary sort: second number
		return a2 < b2
	})

	// return idRange

	for i, v := range idRange {
		splitIdRange := strings.Split(v, "-")
		from, _ := strconv.Atoi(splitIdRange[0])
		to, _ := strconv.Atoi(splitIdRange[1])

		for x := i + 1; x < len(idRange); x++ {
			splitIdRangeNext := strings.Split(idRange[x], "-")
			fromNext, _ := strconv.Atoi(splitIdRange[0])
			toNext, _ := strconv.Atoi(splitIdRange[1])

			if from == fromNext {
				idRange = append(idRange[:i], idRange[i+1:]...)
			}

			if 

			// 1-5
			// 1-10

		}

	}
	//
	// uniqueIds := CreateUnique(foundIds)
	//
	// return len(uniqueIds)
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
