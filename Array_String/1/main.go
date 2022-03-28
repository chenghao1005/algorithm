package main

import (
	"fmt"
	"sort"
)

// LeetCode 第 242 题：给定两个字符串 s 和 t，编写一个函数来判断 t 是否是 s 的字母异位词。

//示例 1
//
//输入: s = "anagram", t = "nagaram"
//
//输出: true
//
//
//
//示例 2
//
//输入: s = "rat", t = "car"
//
//输出: false

func isAnagram(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

func isAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

func main() {
	s1 := "array"
	s2 := "rraay"
	res1 := isAnagram(s1, s2)
	res2 := isAnagram2(s1, s2)
	fmt.Println(res1)
	fmt.Println(res2)
}
