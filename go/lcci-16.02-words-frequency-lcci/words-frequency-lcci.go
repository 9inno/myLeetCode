package leetcode

type WordsFrequency struct {
	hashMap map[string]int
}


func Constructor(book []string) WordsFrequency {
	tmp := WordsFrequency{hashMap:map[string]int{}}
	for _ , value := range book {
		tmp.hashMap[value] ++
	}
	return tmp
}


func (this *WordsFrequency) Get(word string) int {
	if _, ok := this.hashMap[word] ; ok {
		return this.hashMap[word]
	}
	return 0
}


/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
