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
	const k = 12

	if len(bank) <= k {
		v, _ := strconv.Atoi(bank)
		return v
	}

	drop := len(bank) - k
	stack := make([]byte, 0, len(bank))

	for i := 0; i < len(bank); i++ {
		c := bank[i]

		for drop > 0 && len(stack) > 0 && c > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			drop--
		}

		stack = append(stack, c)
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	valueStr := string(stack)
	value, _ := strconv.Atoi(valueStr)
	return value
}
