package day3

import "strconv"

func MaxBankJoltage(bank string) int {
	best := 0
	bestIndex := 0
	secondBest := 0

	for i := 0; i < len(bank)-1; i++ {
		check := int(bank[i] - '0')
		if check > best {
			best = check
			bestIndex = i
			secondBest = 0
		}

		for x := bestIndex + 1; x < len(bank); x++ {
			secondCheck := int(bank[x] - '0')
			if secondCheck > secondBest {
				secondBest = secondCheck
			}
		}
	}

	return (best * 10) + secondBest
}

func MaxBankJoltagePt2(bank string) int {
	best := 0
	bestIndex := 0
	var result []int

	for i := 0; i < len(bank)-12; i++ {
		check := int(bank[i] - '0')
		if check > best {
			best = check
			bestIndex = i
		}
	}
	valueStr := bank[bestIndex : bestIndex+12]
	value, _ := strconv.Atoi(valueStr)
	return value

}
