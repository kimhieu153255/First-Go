package practiceleetcode

import (
	"sort"
	"strconv"
	"strings"
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
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]string, numRows)
	row, step := 0, 1
	for _, char := range s {
		result[row] += string(char)
		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}
		row += step
	}
	return strings.Join(result, "")
}

// 7. Reverse Integer
func Reverse(x int) int {
	var result string
	s := strconv.Itoa(x)
	if x < 0 {
		result += "-"
		s = strings.Split(s, "-")[1]
	}
	var temp string
	for _, c := range s {
		temp = string(c) + temp
	}
	tempInt, _ := strconv.Atoi(result + temp)
	if tempInt > 2147483647 || tempInt < -2147483648 {
		return 0
	}
	return tempInt
}

// 8. String to Integer (atoi)
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	var result string
	if s[0] == '-' || s[0] == '+' {
		result += string(s[0])
		s = s[1:]
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			break
		}
		result += string(c)
	}

	if len(result) == 0 || result == "-" || result == "+" {
		return 0
	}

	tempInt, _ := strconv.Atoi(result)
	if tempInt > 2147483647 {
		return 2147483647
	} else if tempInt < -2147483648 {
		return -2147483648
	}
	return tempInt
}

// 9. Palindrome Number
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	if len(s) <= 1 {
		return true
	}
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 10. Regular Expression Matching

