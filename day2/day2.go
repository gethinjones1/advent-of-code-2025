package day2

import (
	"strconv"
	"strings"
)

func FindInvalidIds(idRange string) []int {
	var result []int
	splitIdRange := strings.Split(idRange, "-")
	from, _ := strconv.Atoi(splitIdRange[0])
	to, _ := strconv.Atoi(splitIdRange[1])

	for i := from; i <= to; i++ { // 10
		digits := convertIntToArray(i)
		var lastDigit int
		intString := strconv.Itoa(i)
		intLength := len(intString)

		if intLength >= 3 && CheckPatternRepeated(i) {
			result = append(result, i)
			continue
		}

		if intLength >= 4 && checkNumberSymmetry(intString, intLength) {
			result = append(result, i)
			continue
		}

		if intLength >= 3 && checkNumberRepetition(intString, intLength) {
			result = append(result, i)
			continue
		}

		for _, ch := range digits {
			if intLength%2 != 0 {
				continue
			}

			if ch == lastDigit && len(digits) == 2 {
				result = append(result, i)
			}

			lastDigit = ch
		}
	}
	return result
}

func CheckPatternRepeated(checkNumber int) bool {
	s := strconv.Itoa(checkNumber)

	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			p := s[:i]
			ok := true
			for j := 0; j < len(s); j += i {
				if s[j:j+i] != p {
					ok = false
					break
				}
			}
			if ok {
				return true
			}
		}
	}
	return false
}

func checkNumberRepetition(number string, length int) bool {
	for i := 0; i < length; i++ {
		if i == 0 {
			continue
		}

		if number[i] != number[i-1] {
			return false
		}
	}
	return true
}

func checkNumberSymmetry(number string, length int) bool {
	startRange := number[:length/2]
	endRange := number[length/2:]

	return startRange == endRange
}

func convertIntToArray(i int) []int {
	digits := []int{}
	for _, ch := range strconv.Itoa(i) {
		digits = append(digits, int(ch-'0'))
	}
	return digits
}
