package main

func main() {
	findMedianSortedArrays([]int{2}, []int{})
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	minN := 0
	maxN := n

	for minN <= maxN {
		i := (minN + maxN) / 2
		j := (n + m + 1) / 2 - i
		if j != 0 && i != n && nums2[j - 1]  > nums1[i] {
			minN = i + 1
		} else if i != 0 && j != m && nums1[i - 1] > nums2[j] {
			maxN = i - 1
		} else {
			var maxLeft int
			var minRight int
			if i == 0 {
				maxLeft = nums2[j - 1]
			} else if j == 0 {
				maxLeft = nums1[i - 1]
			} else {
				maxLeft = max(nums1[i - 1], nums2[j - 1])
			}
			if (n + m ) % 2 == 1 {
				return float64(maxLeft)
			}

			if i == n {
				minRight = nums2[j]
			} else if j == m{
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft + minRight) /2
		}
	}
	return 0
}

func max(a , b int) int  {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}