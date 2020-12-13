//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
// 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回 0。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。 
// Related Topics 广度优先搜索 
// 👍 661 👎 0

package main

import (
	"fmt"
)

// 预处理，构建图结构
//leetcode submit region begin(Prohibit modification and deletion)
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	// 预处理
//	existEndWord := false
//	dict := make(map[string][]string)
//	for _, word := range wordList {
//		if strings.EqualFold(endWord, word) {
//			existEndWord = true
//		}
//		for i := 0; i < len(word); i++ {
//			key := word[:i] + "*" + word[i+1:]
//			if _, ok := dict[key]; ok {
//				dict[key] = append(dict[key], word)
//			}else {
//				dict[key] = []string{word}
//			}
//		}
//	}
//	if !existEndWord {
//		return 0
//	}
//
//	// BFS搜索
//	type Node struct{ // 图节点，增加层数level信息
//		Word string
//		Level int
//	}
//	queue := []Node{{Word:  beginWord, Level: 1}}
//	visited := make(map[string]bool)
//	visited[beginWord] = true
//	for len(queue) > 0 {
//		curr := queue[0]
//		neighbors := getNeighbors(curr.Word, dict)
//		for _, nextWord := range neighbors {
//			if nextWord == endWord {
//				return curr.Level + 1
//			}
//			if visited[nextWord] {
//				continue
//			}
//			visited[nextWord] = true
//			queue = append(queue, Node{Word: nextWord, Level: curr.Level+1})
//		}
//		queue = queue[1:]
//	}
//	return 0
//}
//
//func getNeighbors(word string, dict map[string][]string) []string {
//	var neighbors []string
//	exist := make(map[string]bool)
//	for i := 0; i < len(word); i++ {
//		key := word[:i] + "*" + word[i+1:]
//		if s, ok := dict[key]; ok {
//			for _, nextWord := range s {
//				if !exist[nextWord] {
//					exist[nextWord] = true
//					neighbors = append(neighbors, nextWord)
//				}
//			}
//		}
//	}
//	return neighbors
//}

// 不做预处理
func ladderLength(beginWord string, endWord string, wordList []string) int {
	table := make(map[string]bool)
	for _, word := range wordList {
		table[word] = true
	}
	if !table[endWord] {
		return 0
	}
	visit := make(map[string]bool)
	level := 1
	queue := []string{beginWord}
	visit[beginWord] = true

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[i]
			neighbors := getNeighbors(word, table)
			for _, nextWord := range neighbors {
				if nextWord == endWord {
					return level + 1
				}
				if visit[nextWord] {
					continue
				}
				visit[nextWord] = true
				queue = append(queue, nextWord)
			}
		}
		level++
		queue = queue[size:]
	}
	return 0
}

// time:O(n*26)=O(n)，where n=len(word)
func getNeighbors(word string, table map[string]bool) []string {
	var neighbors []string
	for i := 0; i < len(word); i++ {
		for ch := 'a'; ch <= 'z'; ch++ {
			if word[i] == uint8(ch) {
				continue
			}
			nextWord := word[:i] + string(ch) + word[i+1:]
			if table[nextWord] {
				neighbors = append(neighbors, nextWord)
			}
		}
	}
	return neighbors
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Printf("%d\n", ladderLength(beginWord, endWord, wordList))
}
