// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println([]int{1})
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= nums.length <= 10^5
 * - k is in the range [1, the number of unique elements in the array].
 * - It is guaranteed that the answer is unique.
 *
 * Follow up:
 * Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
 */

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than 100.00%)
 * > Memory Usage: X MB (less than 100.00%)
 */
func topKFrequent(nums []int, k int) []int {

}
