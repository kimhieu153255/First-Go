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
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if s[i:j] == reverse(s[i:j]) {
				return s[i:j]
			}
		}
	}

	return ""
}
