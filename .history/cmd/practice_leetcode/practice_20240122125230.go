package practiceleetcode

import (
	"sort"
)

// 4. Median of Two Sorted Arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortArr := append(nums1, nums2...)
	sort.Ints(sortArr)

	mid := len(sortArr) / 2
	if len(sortArr)%2 == 0 {
		return float64(sortArr[mid-1]+sortArr[mid]) / 2.0
	}
	return float64(sortArr[mid])
}
