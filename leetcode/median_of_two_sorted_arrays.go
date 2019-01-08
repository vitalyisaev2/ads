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

	if len(nums1) == 1 && len(nums2) == 1 {
		return (float64(minIntegers(nums1[0], nums2[0])) + float64(maxIntegers(nums1[0], nums2[0]))) / 2
	}

	var (
		total = len(nums1) + len(nums2)

		arr1 = newArray(nums1)
		arr2 = newArray(nums2)

		inter1 = arr1.currentInterval()
		inter2 = arr2.currentInterval()

		firstOrd  = inter1.compare(inter2)
		actualOrd = firstOrd
	)

	for {
		if arr1.currentInterval().intersects(arr2.currentInterval()) {
			if total%2 == 0 && arr1.size()+arr2.size() < total/2 {
				right1 := arr1.rightItem()
				right2 := arr2.rightItem()
				if right1 > right2 {
					arr2.index++
				} else {
					arr1.index++
				}
			}

			if total%2 != 0 && arr1.size()+arr2.size() > total/2 {
				left1 := arr1.leftItem()
				left2 := arr2.leftItem()
				if left1 > left2 {
					arr1.index--
				} else {
					arr2.index--
				}
			}

			break
		}

		switch actualOrd {
		case less:
			if !arr1.onRightBorder() {
				inter1 = arr1.shiftIndexRight()
			}
			inter2 = arr2.shiftIndexLeft()
		case greater:
			inter1 = arr1.shiftIndexLeft()
			if arr1.currentInterval().intersects(arr2.currentInterval()) && arr1.size()+arr2.size() == total/2 {
				break
			}
			inter2 = arr2.shiftIndexRight()
		case equal:
			break
		default:
			panic("unexpected ord value")
		}

		if actualOrd = inter1.compare(inter2); actualOrd != firstOrd {
			break
		}
	}

	fmt.Println(">")
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println(">")

	// even number of elements
	if total%2 == 0 {
		right1 := arr1.rightItem()
		right2 := arr2.rightItem()
		right := minIntegers(right1, right2)

		left1 := arr1.leftItem()
		left2 := arr2.leftItem()
		left := maxIntegers(left1, left2)

		fmt.Println("left1", left1, "left2", left2, "right1", right1, "right2", right2)

		return (float64(left) + float64(right)) / 2
	}

	// odd number of elements
	right1 := arr1.rightItem()
	right2 := arr2.rightItem()
	fmt.Println("right1", right1, "right2", right2)
	return float64(minIntegers(right1, right2))

}

// array - helper struct that contains array of numbers
// with index of border
type array struct {
	start int
	end   int
	index int
	nums  []int
}

func (a array) size() int { return a.index + 1 }

func (a array) currentInterval() interval {
	if a.index == -1 {
		return interval{
			right:  a.nums[0],
			border: true,
		}
	} else if a.index == len(a.nums)-1 {
		return interval{
			left:   a.nums[len(a.nums)-1],
			border: true,
		}
	}
	return interval{
		left:  a.nums[a.index],
		right: a.nums[a.index+1],
	}
}

func (a *array) shiftIndexRight() interval {
	var (
		newIndex int
		newStart int
		result   interval
	)

	if a.index == len(a.nums)-1 {
		// right border reached
		panic("right border reached")
	}

	if a.end-a.start > 1 {
		newIndex = middle(a.index+1, a.end)
	} else {
		newIndex = a.end
	}

	if newIndex == len(a.nums)-1 {
		// last item before border
		result = interval{left: a.nums[newIndex], border: true}
	} else {
		result = interval{left: a.nums[newIndex], right: a.nums[newIndex+1]}
	}

	newStart = a.index + 1
	a.index = newIndex
	a.start = newStart
	return result
}

func (a *array) shiftIndexLeft() interval {
	var (
		newIndex int
		newEnd   int
		result   interval
	)

	if a.index == -1 {
		// left border reached
		panic("left border reached")
	}

	if a.index == 0 {
		// last item before border
		result = interval{right: a.nums[newIndex], border: true}
		newIndex = -1
	} else {

		if a.end-a.start > 1 {
			newIndex = middle(a.start, a.index)
		} else {
			newIndex = a.start
		}
		result = interval{left: a.nums[newIndex], right: a.nums[newIndex+1]}
	}

	newEnd = a.index
	a.index = newIndex
	a.end = newEnd
	return result
}

func middle(start, stop int) int {
	if start == stop {
		return start
	}

	span := stop - start
	if span%2 == 0 {
		return start + span/2 - 1
	}
	return start + span/2
}

func (a array) leftItem() int {
	left := int(math.MinInt64)
	if a.index >= 0 {
		left = a.nums[a.index]
	}
	return left
}

func (a array) rightItem() int {
	right := int(math.MaxInt64)
	if a.index < 0 {
		right = a.nums[0]
	} else if a.index < len(a.nums)-1 {
		right = a.nums[a.index+1]
	}
	return right
}

func (a array) onRightBorder() bool { return a.index == len(a.nums)-1 }

func (a array) onLeftBorder() bool { return a.index == -1 }

func (a array) String() string {
	return fmt.Sprintf("start: %d, index: %d, end: %d", a.start, a.index, a.end)
}

func newArray(nums []int) *array {
	result := &array{
		start: 0,
		end:   len(nums) - 1,
		nums:  nums,
	}
	result.index = middle(result.start, result.end)
	return result
}

type interval struct {
	left   int
	right  int
	border bool
}

func (i interval) compare(j interval) ord {
	if i.left < j.left {
		return less
	}
	if i.left > j.left {
		return greater
	}
	return equal
}

func (i interval) intersects(j interval) bool {
	if i.border && j.border {
		return intervalBorderBorderIntersection(i, j) || intervalBorderBorderIntersection(j, i)
	}
	if i.border || j.border {
		return intervalMiddleBorderIntersection(i, j) || intervalMiddleBorderIntersection(j, i)
	}
	return intervalMiddleMiddleIntersection(i, j) || intervalMiddleMiddleIntersection(j, i)
}

func intervalBorderBorderIntersection(i, j interval) bool {
	return i.left <= j.right
}

func intervalMiddleBorderIntersection(i, j interval) bool {
	return i.left <= j.left && j.left <= j.right ||
		j.left <= j.right && j.right <= i.right
}

func intervalMiddleMiddleIntersection(i, j interval) bool {
	return (i.left <= j.left && j.left <= i.right && i.right <= j.left) ||
		(i.left <= j.left && j.left <= j.right && j.right <= i.right)
}

func (i interval) String() string {
	return fmt.Sprintf("left: %d, right: %d, border: %v", i.left, i.right, i.border)
}

type ord int8

const (
	less    ord = -1
	equal   ord = 0
	greater ord = 1
)

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
