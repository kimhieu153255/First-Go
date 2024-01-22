package practiceleetcode

import (
	"sort"
)

// 4. Median of Two Sorted Arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortArr := append(nums1, nums2...)
	sort.Ints(sortArr)

	mid := len(sortArr) / 2
	if len(sortArr)%2 == 0 {
		return float64(sortArr[mid-1]+sortArr[mid]) / 2.0
	}
	return float64(sortArr[mid])
}

// 5. Longest Palindromic Substring
func (s *string) reverse() string {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func longestPalindrome(s string) string {
	var result, temp string
	for i, v := range s {
		temp += string(v)
		if(len(temp) > len(result) && temp == temp.reverse()) {
			result = temp
		}
	}
}

