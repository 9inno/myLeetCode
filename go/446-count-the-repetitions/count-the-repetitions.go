package leetcode

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	count1,count2:=0,0
	pos2:=0
	for count1<n1 {
		//还有足够的s1用
		for pos1:=0;pos1<len(s1);pos1+=1 {
			if s1[pos1]==s2[pos2]{
				//当前字符能匹配
				if pos2==len(s2)-1 {
					//匹配到s2的最后一个字符了,pos2回到开头，count2加1表示匹配上一个s2了
					pos2=0
					count2+=1
				}else{
					//没到s2最后一个字符
					pos2+=1
				}
			}
			//无论如何，pos1在往前推进
		}
		//这个循环每结束一次代表扫完了一次s1
		count1+=1
	}
	//循环结束后，count2表示在整个[s1,n1]中,可以获得1个[s2,count2],那么[s2,n2]自然就可以获得count2/n2个
	return count2/n2
}
