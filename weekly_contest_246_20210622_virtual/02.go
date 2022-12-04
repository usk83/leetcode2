// 1904. The Number of Full Rounds You Have Played
// https://leetcode.com/contest/weekly-contest-246/problems/the-number-of-full-rounds-you-have-played/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numberOfRounds("12:01", "12:44"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numberOfRounds("20:00", "06:00"))
	pp.Println(40)
	pp.Println("=========================================")
	pp.Println(numberOfRounds("00:00", "23:59"))
	pp.Println(95)
	pp.Println("=========================================")
	pp.Println(numberOfRounds("23:59", "23:58"))
	pp.Println(95)
	pp.Println("=========================================")
	pp.Println(numberOfRounds("23:15", "23:30"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numberOfRounds("01:45", "15:19"))
	pp.Println(54)
	pp.Println("=========================================")
}

/**
 * 終わった時間のほうが早い場合は計算
 *
 * 最初と最後の時間以外は4回
 * 最初と最後の時間を計算
 */
func numberOfRounds(startTime string, finishTime string) int {
	startHour, _ := strconv.Atoi(string(startTime[0:2]))
	startMinutes, _ := strconv.Atoi(string(startTime[3:5]))
	finishHour, _ := strconv.Atoi(string(finishTime[0:2]))
	finishMinutes, _ := strconv.Atoi(string(finishTime[3:5]))

	if startHour > finishHour || (startHour == finishHour && startMinutes > finishMinutes) {
		finishHour += 24
	}

	count := 0

	if startHour+2 <= finishHour {
		count += (finishHour - startHour - 1) * 4
	}

	if startHour == finishHour {
		if startMinutes == 0 {
			if finishMinutes >= 45 {
				count += 3
			} else if finishMinutes >= 30 {
				count += 2
			} else if finishMinutes >= 15 {
				count += 1
			}
		} else if startMinutes <= 15 {
			if finishMinutes >= 45 {
				count += 2
			} else if finishMinutes >= 30 {
				count += 1
			}
		} else if startMinutes <= 30 {
			if finishMinutes >= 45 {
				count += 1
			}
		}
	} else {
		if startMinutes == 0 {
			count += 4
		} else if startMinutes <= 15 {
			count += 3
		} else if startMinutes <= 30 {
			count += 2
		} else if startMinutes <= 45 {
			count += 1
		}

		if finishMinutes >= 45 {
			count += 3
		} else if finishMinutes >= 30 {
			count += 2
		} else if finishMinutes >= 15 {
			count += 1
		}
	}

	return count
}
