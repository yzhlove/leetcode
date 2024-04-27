package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("1234567829"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("12345"))
	fmt.Println(lengthOfLongestSubstring("123452"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("abb"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var left, right, length, maxLength int
	var validate = make(map[rune]struct{}, 128)

	for left < len(s) {
		if _, ok := validate[rune(s[left])]; !ok {
			validate[rune(s[left])] = struct{}{}
			left++
			length++
			if length > maxLength {
				maxLength = length
			}
		} else {
			for {
				delete(validate, rune(s[right]))
				right++
				if _, ok := validate[rune(s[left])]; !ok {
					validate[rune(s[left])] = struct{}{}
					left++
					break
				}
				length--
			}
		}
	}
	return maxLength
}
