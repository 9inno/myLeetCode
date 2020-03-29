package main

import (
	"fmt"
)

type UndergroundSystem struct {
	idMap map[int]struct{
		start string
		end string
		t int
	}
	stationMap map[string]struct{
		time float64
		count float64
	}
}


func Constructor() UndergroundSystem {
	return UndergroundSystem{
		idMap: map[int]struct {
			start string
			end   string
			t     int
		}{},
		stationMap: map[string]struct {
			time  float64
			count float64
		}{},
	}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	if _ , ok := this.idMap[id]; !ok {
		this.idMap[id] = struct {
			start string
			end   string
			t     int
		}{start:stationName , end: "" , t:t }
	}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	if _ , ok := this.idMap[id]; ok {
		s := this.idMap[id]
		s.end = stationName
		if _,ok := this.stationMap[s.start + "-" +s.end]; ok {
			station := this.stationMap[s.start + "-" +s.end]
			station.count = station.count + float64(1)
			station.time += float64(t) - float64(s.t)
			this.stationMap[s.start + "-" +s.end] = station
		} else {
			this.stationMap[s.start + "-" +s.end]= struct {
				time  float64
				count float64
			}{time: float64(t) -float64(s.t), count:float64(1) }
		}
		delete(this.idMap,id)
	}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if value , ok := this.stationMap[startStation + "-" + endStation]; ok{
		return value.time /value.count
	}else {
		return float64(0)
	}
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */

func main() {
	obj := Constructor()
	obj.CheckIn(45, "Leyton", 3)
	obj.CheckIn(32, "Paradise", 8)
	obj.CheckIn(28, "Leyton", 10)
	obj.CheckOut(45, "Waterloo", 15)
	obj.CheckOut(28, "Waterloo", 21)
	obj.CheckOut(32, "Cambridge", 22)
	fmt.Print(obj.GetAverageTime("Paradise", "Cambridge"))
	obj.GetAverageTime("Leyton", "Waterloo")
	obj.CheckIn(10, "Leyton", 24)
	obj.GetAverageTime("Leyton", "Waterloo")
	obj.CheckOut(10, "Waterloo", 38)
	obj.GetAverageTime("Leyton", "Waterloo")
}