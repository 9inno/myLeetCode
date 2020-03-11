package leetcode


func CanThreePartsEqualSum(A []int) bool {

	sum := 0
	for _ , value := range A{
		sum += value
	}

	if sum % 3 != 0 {
		return false
	}
	tmp := 0
	count := 0
	for i:=0 ; i < len(A) ; i++ {
		
		tmp += A[i]
		if tmp == sum / 3 {
			count ++
			tmp = 0
		}
	}

	return count >= 3
}

func CanThreePartsEqualSum1(A []int) bool {
	sum := 0
	for _ , value := range A{
		sum += value
	}

	if sum % 3 != 0 {
		return false
	}
	left := 0
	leftSum := A[left]
	right := len(A) - 1
	rightSum := A[right]

	for left + 1 < right  {

		if leftSum == sum / 3 && rightSum == sum / 3 {
			return true
		}

		if leftSum != sum / 3 {
			left++
			leftSum += A[left]

		}

		if rightSum != sum / 3 {
			right --
			rightSum += A[right]
		}
	}

	return false
}
