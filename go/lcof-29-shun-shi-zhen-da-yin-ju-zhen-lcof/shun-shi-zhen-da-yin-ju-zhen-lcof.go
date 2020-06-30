package main

func main() {
	a := [][]int {
		[]int{1,2,3,4},
		[]int{5,6,7,8},
		[]int{9,10,11,12},
	}

	spiralOrder(a)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0{
		return []int{}
	}
	step:=0
	size:=len(matrix)*len(matrix[0])
	top,bottom,right,left:=0,len(matrix)-1,len(matrix[0])-1,0
	result :=make([]int,size)
	for step<size{
		for i:=left;i<=right && step<size;i++{
			result[step]=matrix[top][i]
			step++
		}
		top++
		for i:=top;i<=bottom && step<size;i++{
			result[step]=matrix[i][right]
			step++
		}
		right--
		for i:=right;i>=left && step<size;i--{
			result[step]=matrix[bottom][i]
			step++
		}
		bottom--
		for i:=bottom;i>=top && step<size;i--{
			result[step]=matrix[i][left]
			step++
		}
		left++
	}
	return result

}
