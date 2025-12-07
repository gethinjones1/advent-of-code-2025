package day4

import (
	"strings"
)

func CountAccessPaths(path string) int {
	count := 0
	rowsFlat := strings.Split(path, "\n")

	rows := make([][]rune, len(rowsFlat))
	for i, row := range rowsFlat {
		rows[i] = []rune(row)
	}
	var rollCounter int

	for i := range rows {
		for x := range rows[i] {

			rollCounter = 0
			if rows[i][x] != '@' {
				continue
			}

			if x > 0 && rows[i][x-1] == '@' {
				rows[i][x-1] = '.'
				rollCounter++
			}

			if x+1 < len(rows[i]) && rows[i][x+1] == '@' {
				rows[i][x+1] = '.'
				rollCounter++
			}

			if i > 0 && x < len(rows[i-1]) && rows[i-1][x] == '@' {
				rows[i-1][x] = '.'
				rollCounter++
			}

			if i+1 < len(rows) && x < len(rows[i+1]) && rows[i+1][x] == '@' {
				rows[i+1][x] = '.'
				rollCounter++
			}

			if i > 0 && x+1 < len(rows[i-1]) && rows[i-1][x+1] == '@' {
				rows[i-1][x+1] = '.'
				rollCounter++
			}

			if i > 0 && x > 0 && x-1 < len(rows[i-1]) && rows[i-1][x-1] == '@' {
				rows[i-1][x-1] = '.'
				rollCounter++
			}

			if i+1 < len(rows) && x+1 < len(rows[i+1]) && rows[i+1][x+1] == '@' {
				rows[i+1][x+1] = '.'
				rollCounter++
			}

			if i+1 < len(rows) && x > 0 && x-1 < len(rows[i+1]) && rows[i+1][x-1] == '@' {
				rows[i+1][x-1] = '.'
				rollCounter++
			}

			if rollCounter < 4 {
				rows[i][x] = '.'
				count++
			}
		}
	}
	return count
}

func CountTotalRemoved(path string) int {
	removedTotal := 0
	rowsFlat := strings.Split(path, "\n")

	rows := make([][]rune, len(rowsFlat))
	for i, row := range rowsFlat {
		rows[i] = []rune(row)
	}

	for {
		toRemove := [][2]int{}

		for i := range rows {
			for x := range rows[i] {

				if rows[i][x] != '@' {
					continue
				}

				rollCounter := 0

				if x > 0 && rows[i][x-1] == '@' {
					rollCounter++
				}

				if x+1 < len(rows[i]) && rows[i][x+1] == '@' {
					rollCounter++
				}

				if i > 0 && x < len(rows[i-1]) && rows[i-1][x] == '@' {
					rollCounter++
				}

				if i+1 < len(rows) && x < len(rows[i+1]) && rows[i+1][x] == '@' {
					rollCounter++
				}

				if i > 0 && x+1 < len(rows[i-1]) && rows[i-1][x+1] == '@' {
					rollCounter++
				}

				if i > 0 && x > 0 && x-1 < len(rows[i-1]) && rows[i-1][x-1] == '@' {
					rollCounter++
				}

				if i+1 < len(rows) && x+1 < len(rows[i+1]) && rows[i+1][x+1] == '@' {
					rollCounter++
				}

				if i+1 < len(rows) && x > 0 && x-1 < len(rows[i+1]) && rows[i+1][x-1] == '@' {
					rollCounter++
				}

				if rollCounter < 4 {
					toRemove = append(toRemove, [2]int{i, x})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			i, x := pos[0], pos[1]
			rows[i][x] = '.'
		}

		removedTotal += len(toRemove)
	}

	return removedTotal
}
