// 5954. Watering Plants II
// https://leetcode.com/contest/weekly-contest-271/problems/watering-plants-ii/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{2, 2, 3, 3}, 5, 5))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{2, 2, 3, 3}, 3, 4))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{5}, 10, 8))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{1, 2, 4, 4, 5}, 6, 5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{2, 2, 5, 2, 2}, 5, 5))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minimumRefill([]int{2, 1, 1}, 2, 2))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minimumRefill(
		[]int{726, 739, 934, 116, 643, 648, 473, 984, 482, 85, 850, 806, 146, 764, 156, 66, 186, 339, 985, 237, 662, 552, 800, 78, 617, 933, 481, 652, 796, 594, 151, 82, 183, 241, 525, 221, 951, 732, 799, 483, 368, 354, 776, 175, 974, 187, 913, 842},
		1439,
		1207,
	))
	pp.Println(24)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - n == plants.length
 * - 1 <= n <= 10^5
 * - 1 <= plants[i] <= 10^6
 * - max(plants[i]) <= capacityA, capacityB <= 10^9
 */

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var count int
	capA, capB := capacityA, capacityB
	step := 0
	for a, b := 0, len(plants)-1; a <= b; a, b = a+1, b-1 {
		step++
		// fmt.Printf("--- Step %d: \n", step)

		if a == b {
			if plants[a] > capA && plants[b] > capB {
				count++
			}
		} else {
			// fmt.Printf("@Alice: %d\n", capA)
			// fmt.Printf("Target: %d\n", plants[a])
			if plants[a] <= capA {
				capA -= plants[a]
			} else {
				count++
				// fmt.Printf("Refilled! count: %d\n", count)
				// capA += capacityA - plants[a]
				capA = capacityA - plants[a]
			}
			// fmt.Printf("After: %d\n", capA)

			// fmt.Printf("@Bob: %d\n", capB)
			// fmt.Printf("Target: %d\n", plants[b])
			if plants[b] <= capB {
				capB -= plants[b]
			} else {
				count++
				// fmt.Printf("Refilled! count: %d\n", count)
				// capB += capacityB - plants[b]
				capB = capacityB - plants[b]
			}
			// fmt.Printf("After: %d\n", capB)
		}
	}
	return count
}
