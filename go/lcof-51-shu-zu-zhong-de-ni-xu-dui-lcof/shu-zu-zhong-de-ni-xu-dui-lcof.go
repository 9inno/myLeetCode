package leetcode



func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result, _ := mergeSort(nums, 0, len(nums) - 1)
	return result
}

func mergeSort(nums []int, start, end int) (int, []int) {

	if start == end {
		return 0 , []int{nums[start]}
	}
	mid := start + (end - start) / 2
	countLeft, left := mergeSort(nums, start, mid)
	countRight, right := mergeSort(nums, mid + 1, end)

	count := countLeft + countRight
	i := 0
	j := 0
	var tmp []int
	for i < len(left) || j < len(right) {
		if i == len(left) {
			tmp = append(tmp, right[j:] ...)
			break
		}
		if j == len(right) {
			tmp = append(tmp, left[i:]...)
			break
		}
		if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
		} else  {
			tmp = append(tmp, right[j])
			count += len(left) - i
			j++
		}
	}
	return count,tmp
}

func mergeSort1(nums []int, start int ,end int ) int {
	if start >= end {
		return 0
	}
	mid := start + (end - start)/2
	count := mergeSort1(nums, start, mid) + mergeSort1(nums, mid + 1, end)
	var tmp []int
	i, j := start, mid + 1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			count += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		count += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i - start]
	}
	return count
}