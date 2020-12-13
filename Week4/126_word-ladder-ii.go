//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换
//需遵循如下规则： 
//
// 
// 每次转换只能改变一个字母。 
// 转换后得到的单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回一个空列表。 
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
//输出:
//[
//  ["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。 
// Related Topics 广度优先搜索 数组 字符串 回溯算法 
// 👍 368 👎 0

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string
	visited := make(map[string]bool)
	// 预处理
	existEndWord := false
	dict := make(map[string][]string)
	for _, word := range wordList {
		if strings.EqualFold(endWord, word) {
			existEndWord = true
		}
		for i := 0; i < len(word); i++ {
			key := word[:i] + "*" + word[i+1:]
			if _, ok := dict[key]; ok {
				dict[key] = append(dict[key], word)
			} else {
				dict[key] = []string{word}
			}
		}
	}
	if !existEndWord {
		return res
	}

	// 找到该元素的所有邻接节点
	var getNeighbors func(word string) []string
	getNeighbors = func(word string) []string {
		var res []string
		for i := 0; i < len(word); i++ {
			key := word[:i] + "*" + word[i+1:]
			s, ok := dict[key]
			if !ok {
				continue
			}
			res = append(res, s...)
		}
		return res
	}

	// debug
	//for key, value := range dict {
	//	fmt.Printf("%v:%v\n", key, value)
	//}

	// BFS搜索
	queue := [][]string{{beginWord}} // 队列里存放的是路径数组[]string
	visited[beginWord] = true
	isFound := false
	for len(queue) > 0 {
		size := len(queue)
		var subVisit []string
		for i := 0; i < size; i++ {
			currPath := queue[i]
			//fmt.Printf("currPath: %v\n", currPath)
			lastWord := currPath[len(currPath)-1] // 取出当前路径的末尾元素
			neighbors := getNeighbors(lastWord)   // 找到该元素的所有邻接节点
			for _, nextWord := range neighbors {
				if strings.EqualFold(nextWord, endWord) {
					isFound = true
					// 拼接上最后一个单词，就是合法的路径
					//fmt.Printf("currPath:%v\n", currPath)
					currPath = append(currPath, nextWord)
					res = append(res, currPath)
				}
				if visited[nextWord] {
					continue
				}
				//visited[nextWord] = true
				subVisit = append(subVisit, nextWord)
				currPath = append(currPath, nextWord)
				queue = append(queue, append([]string{}, currPath...))
				currPath = currPath[:len(currPath)-1]
			}
		}
		for _, word := range subVisit {
			visited[word] = true
		}
		queue = queue[size:]
		if isFound {
			break
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Printf("%v\n", findLadders(beginWord, endWord, wordList))

	beginWord = "a"
	endWord = "c"
	wordList = []string{"a", "b", "c"}
	fmt.Printf("%v\n", findLadders(beginWord, endWord, wordList))

	beginWord = "red"
	endWord = "tax"
	wordList = []string{"ted","tex","red","tax","tad","den","rex","pee"}
	fmt.Printf("%v\n", findLadders(beginWord, endWord, wordList))
}
