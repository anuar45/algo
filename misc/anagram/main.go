package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("a", "ab"))
	fmt.Println(isAnagram("ab", "a"))

}

func isAnagram(s string, t string) bool {
	smap, tmap := make(map[rune]int), make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, v := range []rune(s) {
		smap[v]++
	}

	for _, v := range []rune(t) {
		tmap[v]++
	}

	for k := range tmap {
		if smap[k] != tmap[k] {
			return false
		}
	}

	return true
}
