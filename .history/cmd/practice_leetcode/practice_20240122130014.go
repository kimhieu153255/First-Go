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

//5. Longest Palindromic Substring
func longestPalindrome(s string) string {
		var result string;
    for i, _ := range s {
			if(len(result)>2){
				
			}
		}
}
