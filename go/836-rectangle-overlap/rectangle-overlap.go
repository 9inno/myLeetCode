package leetcode

func IsRectangleOverlap(rec1 []int, rec2 []int) bool {

	return !(rec1[2] <= rec2[0] || rec1[3] <= rec2[1] || rec1[0] >= rec2[2] || rec1[1] >= rec2[3])

}

func IsRectangleOverlap1(rec1 []int, rec2 []int) bool {

	 x := !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0])
	 y := !(rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
	 // x轴 y轴映射  变成两个区间是否相交
	 return x && y
}