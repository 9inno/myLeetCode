package main

import (
	"container/heap"
	"fmt"
)

type Twitter struct {
	twitter map[int][]*Item
	follow map[int]map[int]bool
	postCount int
}

type Item struct {
	value int
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		twitter: map[int][]*Item{},
		follow: map[int]map[int]bool{},
		postCount: 0,
	}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.twitter[userId] = append(this.twitter[userId], &Item{
		value:    tweetId,
		priority: this.postCount,
		index:    this.postCount,
	})
	this.postCount++
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var result []int
	idList := []int{userId}
	if value, ok := this.follow[userId]; ok {
		for key , _ := range value {
			idList = append(idList , key)
		}
	}
	var tmpTwitter PriorityQueue
	for _ , id := range idList {
		if  items , ok := this.twitter[id]; ok{
			tmpTwitter = append(tmpTwitter, items...)
		}
	}

	heap.Init(&tmpTwitter)
	for len(result) < 10 && len(tmpTwitter) > 0 {
		result = append(result, heap.Pop(&tmpTwitter).(*Item).value)
	}
	return result
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followeeId == followerId {
		return
	}
	if _, ok := this.follow[followerId]; !ok {
		this.follow[followerId] = map[int]bool{}
	}
	this.follow[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	delete(this.follow[followerId], followeeId)
}

func main() {
	obj := Constructor()
	obj.PostTweet(2,5)
	obj.PostTweet(1,3)
	obj.PostTweet(1,101)
	obj.PostTweet(2,13)
	obj.PostTweet(2,10)
	obj.PostTweet(1,2)
	obj.PostTweet(2,94)
	obj.PostTweet(2,505)
	obj.PostTweet(1,333)
	obj.PostTweet(1,22)
	fmt.Println(obj.GetNewsFeed(2))
	obj.Follow(2,1)
	fmt.Println(obj.GetNewsFeed(2))

}
/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */