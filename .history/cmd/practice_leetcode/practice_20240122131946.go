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
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	result := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(s[i:j]) > len(result) {
				result = s[i:j]
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[len(s)-1-i] != s[i] {
			return false
		}
	}
	return true
}

// 6. ZigZag Conversion