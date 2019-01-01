package leetcode

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) == 0 && len(nums2) == 0 {
		panic("invalid input")
	}

	if len(nums1) == 0 {
		return findMedianSortedArray(nums2)
	}

	if len(nums2) == 0 {
		return findMedianSortedArray(nums1)
	}

	var (
		total          = len(nums1) + len(nums2)
		half           = total / 2
		start1, start2 = 0, 0
		end1, end2     = len(nums1), len(nums2)
		ix1            = middle(start1, end1)
		ix2            = middle(start2, end2)
	)

	var (
		firstRelation  = compareIntegers(nums1[ix1], nums2[ix2])
		actualRelation = firstRelation
	)
	fmt.Println(">")
	fmt.Println("start1", start1, "ix1", ix1, "end1", end1)
	fmt.Println("start2", start2, "ix2", ix2, "end2", end2)
	fmt.Println(">")

	for {
		if actualRelation == less {

			budget := len(nums1) - ix1 - 1 + ix2
			fmt.Println("less", budget)
			if budget <= 0 {
				fmt.Println("e6")
				break
			}
			ix1, start1 = shiftBorderRight(ix1, end1)
			ix2, end2 = shiftBorderLeft(ix2, start2)
		} else if actualRelation == greater {
			budget := len(nums2) - ix2 - 1 + ix1
			fmt.Println("greater", budget)
			if budget <= 0 {
				fmt.Println("e7")
				break
			}
			ix1, end1 = shiftBorderLeft(ix1, start1)
			ix2, start2 = shiftBorderRight(ix2, end2)
		} else {
			if total == 2 {
				ix2 -= 1
			}
			fmt.Println("e1")
			break
		}

		fmt.Println(">")
		fmt.Println("start1", start1, "ix1", ix1, "end1", end1)
		fmt.Println("start2", start2, "ix2", ix2, "end2", end2)
		fmt.Println(">")

		// case when the whole half of sample belongs to the only array
		if ix1 < 0 {
			ix2 = half - 1
			fmt.Println("e4")
			break
		} else if ix2 < 0 {
			ix1 = half - 1
			fmt.Println("e5")
			break
		}

		if actualRelation = compareIntegers(nums1[ix1], nums2[ix2]); actualRelation != firstRelation {
			break
		}

	}

	// border is determined
	fmt.Println("ix", ix1, ix2)

	// lookup for next value to count median

	// even number of elements
	if total%2 == 0 {
		right1 := rightBorder(nums1, ix1)
		right2 := rightBorder(nums2, ix2)
		right := minIntegers(right1, right2)

		left1 := leftBorder(nums1, ix1)
		left2 := leftBorder(nums2, ix2)
		left := maxIntegers(left1, left2)

		fmt.Println("left", left, "right", right)

		return (float64(left) + float64(right)) / 2
	}

	// odd number of elements
	right1 := rightBorder(nums1, ix1)
	right2 := rightBorder(nums2, ix2)
	return float64(minIntegers(right1, right2))
}

func middle(start, stop int) int {
	span := stop - start
	if span%2 == 0 {
		return start + span/2 - 1
	}
	return start + span/2
}

func shiftBorderRight(ixIn, end int) (ixOut, startOut int) {
	ixOut = middle(ixIn+1, end)
	startOut = ixIn
	return
}

func shiftBorderLeft(ixIn, start int) (ixOut, endOut int) {
	ixOut = middle(start, ixIn)
	endOut = ixIn
	return
}

type relation int8

const (
	less    relation = -1
	equal   relation = 0
	greater relation = 1
)

func compareIntegers(a, b int) relation {
	if a < b {
		return less
	}
	if a > b {
		return greater
	}
	return equal
}

func minIntegers(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxIntegers(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rightBorder(nums []int, ix int) int {
	right := int(math.MaxInt64)
	if ix < 0 {
		right = nums[0]
	} else if ix < len(nums)-1 {
		right = nums[ix+1]
	}
	return right
}

func leftBorder(nums []int, ix int) int {
	left := int(math.MinInt64)
	if ix >= 0 {
		left = nums[ix]
	}
	return left
}

func findMedianSortedArray(nums []int) float64 {
	size := len(nums)
	if size == 0 {
		panic("empty array")
	}

	if size == 1 {
		return float64(nums[0])
	}

	ix := size / 2
	if size%2 == 0 {
		return (float64(nums[ix-1]) + float64(nums[ix])) / 2
	}

	return float64(nums[ix])
}
