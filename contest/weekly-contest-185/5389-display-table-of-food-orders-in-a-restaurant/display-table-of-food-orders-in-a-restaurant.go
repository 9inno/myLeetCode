package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	var foodItem []string
	var table []int
	foodMap := map[string]bool{}
	tableMap := map[string]bool{}
	hashMap := map[string]map[string]int{}

	for _ , row := range orders {
		if _, ok := foodMap[row[2]]; !ok {
			foodItem = append(foodItem , row[2])
			foodMap[row[2]] = true
		}
		if _, ok := tableMap[row[1]]; !ok {
			tmpTable, _:=strconv.Atoi(row[1])
			table = append(table ,tmpTable)
			tableMap[row[1]]= true
		}

		if _, ok := hashMap[row[1]] ; ok {
			hashMap[row[1]][row[2]] ++
		} else  {
			hashMap[row[1]] = map[string]int{row[2] : 1}
		}
	}
	sort.Strings(foodItem)
	sort.Ints(table)
	result := [][]string{}
	title := []string{"Table"}
	for _ , value := range foodItem {
		title = append(title ,value)
	}
	result = append(result, title)

	for _ , tableNumber := range table {
		tmp := []string{strconv.Itoa(tableNumber)}
		for _ , food := range  foodItem {
			if count ,ok := hashMap[strconv.Itoa(tableNumber)][food]; ok{
				tmp = append(tmp, strconv.Itoa(count))
			} else {
				tmp = append(tmp, "0")
			}
		}
		result = append(result, tmp)
	}
	return result
}
