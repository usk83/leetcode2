// 5841. Find the Longest Valid Obstacle Course at Each Position
// https://leetcode.com/contest/weekly-contest-253/problems/find-the-longest-valid-obstacle-course-at-each-position/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))
	fmt.Println([]int{1, 2, 3, 3})
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))
	fmt.Println([]int{1, 2, 1})
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
	fmt.Println([]int{1, 1, 2, 3, 2, 2})
	pp.Println("=========================================")
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	prevs := make([]int, 0, n)
	length := make([]int, n)
	for i := 0; i < n; i++ {
		index := sort.SearchInts(prevs, obstacles[i]+1)
		if index >= len(prevs) {
			prevs = append(prevs, obstacles[i])
		} else {
			prevs[index] = obstacles[i]
		}
		length[i] = index + 1
	}
	return length
}

/**
 * �Է֤��ǰ��
 *   �Է֤�ͬ�����Է֤��С������
 *     �����ޤǤ��L����һ���L����Τ��x�k����
 *   �Է֤��󤭤��路�����ڤ��ʤ����Ϥ⤢��
 *     ���Έ��ϴ𤨤�1
 *
 * 2�Ĥ������ǥ��`�Ȥ��줿��Τ�ȡ�ä�����
 *   �Ϻ���θߤ�
 *   �����ޤǤΥ��`�����L��
 *
 *     n == obstacles.length
 *     1 <= n <= 10^5
 *     1 <= obstacles[i] <= 10^7
 *
 * nlogn ���餤�ޤǤ�Ӌ�㥪�`��
 * ��ǰ�I���OK
 *
 * ͬ���L���Υ��`���ʤ����ҤΤ�Τ��x�k���Ƥ��ޤäƤ褤
 *
 * λ�äǤϥ��`�Ȥ���Ƥ���
 * �ߤ��ǥ��`�Ȥ���
 * ���`�����L���ǥ��`�Ȥ���
 *
 * Binary Search
 *
 * λ�äȥ��`�����L��
 * ���`�����L����index
 */
// func longestObstacleCourseAtEachPosition(obstacles []int) []int {
// 	n := len(obstacles)
//
// 	// λ�äȂ��νM�ߺϤ碌�����ä�
// 	calcOrder := make([][2]int, n)
// 	for i, height := range obstacles {
// 		calcOrder[i] = [2]int{i, height}
// 	}
//
// 	// ���`��
// 	// ����С������Τ���
// 	// ͬ�����ΤȤ���λ�ä���Τ�Τ���
// 	sort.Slice(calcOrder, func(i, j int) bool {
// 		return calcOrder[i][1] < calcOrder[j][1] || (calcOrder[i][1] == calcOrder[j][1] && calcOrder[i][0] < calcOrder[j][0])
// 	})
//
// 	// λ�äȥ��`�����L��
// 	// ���`�����L����index
// 	optList := make([]int, 0, len(obstacles))
//
// 	length := make([]int, n)
// 	for i := o; i < n; i++ {
// 		// �Է֤��λ�ä���Τ�Τ�һ���Ҥˤ����Τ�̽��
// 		i := sort.Search(len(optList), func(i int) bool {
//
// 		})
// 		calcOrder[0]
//
// 	}
// 	return length
// }
