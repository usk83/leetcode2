// 5826. Delete Duplicate Folders in System
// https://leetcode.com/contest/weekly-contest-251/problems/delete-duplicate-folders-in-system/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

/**
 * まずはTreeの形に変換する
 * TreeNode {
 * 	name string
 * 	childrenNodes []TreeNode // ソート済みである必要あり
 * 	structure string // 自分以下の構造を特定の文字列で表現
 * 	deleteFlag bool
 * }
 * その際、子ノードを名前でソートしておく

 * 次にルートからDFSして、structure文字列を当てはめていく
 * 	structure について
 * 	子ノードの表現をフォルダに含まれない文字で区切り、カッコで囲ったものを親の表現とする
 * 	例えば以下のような感じにする
 * 		f(z(a|b|c)|e(d))|g(h(i|j)|k)
 * その際、完成した表現を別に用意したmapに格納し、重複があった場合は、そのノードとすでに出現しているノードについて
 * 	map[structure string]*TreeNode
 * 削除のフラグを立てる

 * 最後にフラグを意識しながらもう一度ルートからDFSして、元の構造表現に戻す
 */
func deleteDuplicateFolder(paths [][]string) [][]string {
	return nil
}
