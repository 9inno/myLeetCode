package leetcode


type MountainArray struct {
}
func (this *MountainArray) get(index int) int {}
func (this *MountainArray) length() int {}


func findInMountainArray(target int, mountainArr *MountainArray) int {
	result := mountainArr.length()
	recursion(mountainArr, 0, result - 1, target, &result)
	if result == mountainArr.length() {
		result = -1
	}
	return result
}

func recursion(mountainArr *MountainArray, start, end int, target int, tmpIndex *int)  {
	if start > end {
		return
	}

	mid := start + (end - start) / 2
	startValue := mountainArr.get(start)
	midValue := mountainArr.get(mid)
	midNextValue := mountainArr.get(mid + 1)
	endValue := mountainArr.get(end)
	if midValue == target {
		*tmpIndex = min(*tmpIndex, mid)
		return
	}
	if startValue == target {
		*tmpIndex = min(*tmpIndex, start)
		return
	}
	if endValue == target {
		*tmpIndex = min(*tmpIndex, end)
		return
	}

	if  midValue < midNextValue {
		if target > startValue && target < midValue {
			recursion(mountainArr, start + 1, mid - 1, target, tmpIndex)
		} else {
			recursion(mountainArr, mid + 1, end - 1, target, tmpIndex)
		}

	} else {
		if target < midValue && target > endValue {
			recursion(mountainArr, mid + 1, end - 1, target, tmpIndex)
		} else{
			recursion(mountainArr, mid + 1, end - 1, target, tmpIndex)
		}
		if startValue < midValue {
			recursion(mountainArr, start + 1, mid - 1, target, tmpIndex)
		}

	}
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}