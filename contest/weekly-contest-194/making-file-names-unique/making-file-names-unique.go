package main

import "strconv"

func main() {
	a := []string{"gta","gta(1)","gta","avalon"}
	getFolderNames(a)
}
func getFolderNames(names []string) []string {
	hashMap := map[string]int{}

	result := []string{}
	for _ , value := range names {
		if v , ok := hashMap[value] ; ok {
			i := v + 1
			for {
				num := strconv.Itoa(i)
				tmp := value +"(" +num +")"
				if _, ok1 := hashMap[tmp] ; ok1 {
					i ++
					continue
				} else {
					result = append(result, tmp)
					hashMap[tmp] = 0
					hashMap[value] = i
					break
				}
			}

		} else  {
			hashMap[value] = 0
			result = append(result, value)
		}
	}
	return result

}
