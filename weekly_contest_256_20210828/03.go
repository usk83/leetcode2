// 5856. Minimum Number of Work Sessions to Finish the Tasks
// https://leetcode.com/contest/weekly-contest-256/problems/minimum-number-of-work-sessions-to-finish-the-tasks/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minSessions([]int{1, 2, 3}, 3))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minSessions([]int{3, 1, 3, 1, 1}, 8))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minSessions([]int{1, 2, 3, 4, 5}, 15))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minSessions([]int{2, 3, 3, 4, 4, 4, 5, 6, 7, 10}, 12))
	pp.Println(4)
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */

// n == tasks.length
// 1 <= n <= 14
// 1 <= tasks[i] <= 10
// max(tasks[i]) <= sessionTime <= 15
func minSessions(tasks []int, sessionTime int) int {
	sort.Ints(tasks)
	taskTimes, taskCounts := uniqueTimesAndCounts(tasks)

	// pp.Println("=== DEBUG START ======================================")
	// fmt.Println(taskTimes)
	// fmt.Println(taskCounts)
	// pp.Println("=== DEBUG END ========================================")

	memo := map[string]int{}
	var recursive func(int, int)
	recursive = func(sessionCount, remainTime int) {
		mKey, mVal := getMapKey(taskCounts), getMapValue(sessionCount, remainTime)

		if best, ok := memo[mKey]; ok && compare(best, mVal) {
			return
		} else {
			memo[mKey] = mVal
		}
		for i := range taskCounts {
			if taskCounts[i] == 0 {
				continue
			}

			taskCounts[i]--
			if remainTime >= taskTimes[i] {
				recursive(sessionCount, remainTime-taskTimes[i])
			} else {
				recursive(sessionCount+1, sessionTime-taskTimes[i])
			}
			taskCounts[i]++
		}
	}
	recursive(1, sessionTime)

	minKey := getMapKey(make([]int, len(taskCounts)))
	return memo[minKey] / 100
}

func compare(left int, right int) bool {
	if left/100 > right/100 {
		return false
	} else if left/100 < right/100 {
		return true
	}
	return left%100 >= right%100
}

func getMapKey(taskCounts []int) string {
	key := ""
	for _, count := range taskCounts {
		key += fmt.Sprintf("%02d", count)
	}
	return key
}

func getMapValue(sessionCount int, remainTime int) int {
	return sessionCount*100 + remainTime
}

func uniqueTimesAndCounts(ts []int) ([]int, []int) {
	m := map[int]int{}
	for _, t := range ts {
		m[t] = m[t] + 1
	}
	times := make([]int, 0, len(m))
	for key := range m {
		times = append(times, key)
	}
	sort.Ints(times)

	counts := make([]int, 0, len(m))
	for _, key := range times {
		counts = append(counts, m[key])
	}

	return times, counts
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
