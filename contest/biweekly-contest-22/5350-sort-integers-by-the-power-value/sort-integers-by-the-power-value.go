package main

import "sort"

func main() {
	getKth(12 ,15, 2)
}
func getKth(lo int, hi int, k int) int {
	hashMap := map[int] []int{}
	var q []int
	var result int
	for i := lo; i <= hi; i ++ {
		tmp := 0
		tmpI := i
		for tmpI != 1 {
			if tmpI % 2 == 0 {
				tmpI = tmpI / 2
			} else {
				tmpI = 3 * tmpI + 1
			}
			tmp ++
		}
		if _ , ok := hashMap[tmp] ; !ok{
			q = append(q, tmp)
		}
		hashMap[tmp] = append(hashMap[tmp] , i)
	}
	sort.Ints(q)
	for _ , key := range q {
		if value , ok := hashMap[key]; ok  {
			if len(value) >= k {
				result = value[ k - 1]
				break
			} else {
				k = k - len(value)
			}
		}
	}
	return result

}