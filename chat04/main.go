package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(findMedianSortedArrays2([]int{1, 3, 4}, []int{2, 5, 6})) // 3.5
	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2, 5, 6}))    // 3
	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2}))          // 2
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))       // 2.5

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	if len(nums1) == 0 {
		idx := len(nums2) / 2
		if len(nums2)%2 == 0 {
			return float64(nums2[idx]+nums2[idx-1]) / 2
		}
		return float64(nums2[idx])
	}

	if len(nums2) == 0 {
		idx := len(nums1) / 2
		if len(nums1)%2 == 0 {
			return float64(nums1[idx]+nums1[idx-1]) / 2
		}
		return float64(nums1[idx])
	}

	var center = (len(nums1) + len(nums2)) / 2
	var next = true
	if (len(nums1)+len(nums2))%2 != 0 {
		center = (len(nums1)+len(nums2))/2 + 1
		next = false
	}

	var s, i, j, n int
	for n <= center {
		n++
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				if (n == center) || (next && n == center+1) {
					s += nums1[i]
				}
				i++
			} else {
				if (n == center) || (next && n == center+1) {
					s += nums2[j]
				}
				j++
			}

		} else if i < len(nums1) {
			if (n == center) || (next && n == center+1) {
				s += nums1[i]
			}
			i++
		} else if j < len(nums2) {
			if (n == center) || (next && n == center+1) {
				s += nums2[j]
			}
			j++
		} else {
			break
		}
	}
	if next {
		return float64(s) / 2
	}
	return float64(s)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {

	set := make([]int, len(nums1)+len(nums2))
	copy(set, nums1)
	copy(set[len(nums1):], nums2)
	sort.Ints(set)
	if len(set)%2 == 0 {
		return float64(set[len(set)/2]+set[len(set)/2-1]) / 2.0
	}
	return float64(set[len(set)/2])
}
