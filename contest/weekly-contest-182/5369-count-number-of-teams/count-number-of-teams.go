package leetcode


func numTeams(rating []int) int {
	result :=0
	for i:=0; i < len(rating) - 2 ;i ++{
		for j:= i +1 ; j < len(rating) -1 ; j++ {
			if rating[i] > rating[j] {
				result += greater(rating,j,false)
			}else {
				result +=greater(rating,j,true)
			}
		}
	}
	return result
}

func greater(rating []int,index int, larger bool) int {
	sum := 0
	cur := rating[index]
	if larger {
		for i := index; i < len(rating); i++ {
			if rating[i] > cur{
				sum++
			}
		}
	} else {
		for i := index; i < len(rating); i++ {
			if rating[i] < cur{
				sum++
			}
		}
	}
	return sum
}