package main

func main() {
	//canConstruct("annabelle",2)
	//canConstruct("leetcode",2)
	print(canConstruct("cxayrgpcctwlfupgzirmazszgfiusitvzsnngmivctprcotcuutfxdpbrdlqukhxkrmpwqqwdxxrptaftpnilfzcmknqljgbfkzuhksxzplpoozablefndimqnffrqfwgaixsovmmilicjwhppikryerkdidupvzdmoejzczkbdpfqkgpbxcrxphhnxfazovxbvaxyxhgqxcxirjsryqxtreptusvupsstylpjniezyfokbowpbgxbtsemzsvqzkbrhkvzyogkuztgfmrprz",5))
}

func canConstruct(s string, k int) bool {
	length := len(s)
	if length < k {
		return false
	} else if length == k {
		return true
	}

	hashMap := map[string]int{}

	for _ , value := range s {
		hashMap[string(value)] ++
	}
	singular := 0
	for _ , count := range hashMap {
		if count %2 ==  1 {
			singular ++
		}
	}
	return singular <= k
}