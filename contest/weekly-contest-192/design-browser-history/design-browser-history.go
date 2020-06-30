package main

import "fmt"

type BrowserHistory struct {
	urls []string
	currentIndex int
}


func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		urls:         []string{homepage},
		currentIndex: 0,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	this.urls = append(this.urls[:this.currentIndex + 1], url)
	this.currentIndex ++
}


func (this *BrowserHistory) Back(steps int) string {
	if steps > this.currentIndex {
		this.currentIndex = 0
		return this.urls[0]
	}
	this.currentIndex = this.currentIndex - steps
	return this.urls[this.currentIndex]
}


func (this *BrowserHistory) Forward(steps int) string {
	if steps + this.currentIndex > len(this.urls) - 1 {
		this.currentIndex = len(this.urls) - 1
		return this.urls[len(this.urls) - 1]
	}
	this.currentIndex = this.currentIndex + steps
	return this.urls[this.currentIndex]
}

func main() {
	 browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")      // 你原本在浏览 "leetcode.com" 。访问 "google.com"
	browserHistory.Visit("facebook.com")     // 你原本在浏览 "google.com" 。访问 "facebook.com"
	browserHistory.Visit("youtube.com")      // 你原本在浏览 "facebook.com" 。访问 "youtube.com"
	fmt.Print(browserHistory.Back(1))                   // 你原本在浏览 "youtube.com" ，后退到 "facebook.com" 并返回 "facebook.com"
	fmt.Print(browserHistory.Back(1))                   // 你原本在浏览 "facebook.com" ，后退到 "google.com" 并返回 "google.com"
	fmt.Print(browserHistory.Forward(1))               // 你原本在浏览 "google.com" ，前进到 "facebook.com" 并返回 "facebook.com"
	browserHistory.Visit("linkedin.com")    // 你原本在浏览 "facebook.com" 。 访问 "linkedin.com"
	fmt.Print(browserHistory.Forward(2))               // 你原本在浏览 "linkedin.com" ，你无法前进任何步数。
	fmt.Print(browserHistory.Back(2))                 // 你原本在浏览 "linkedin.com" ，后退两步依次先到 "facebook.com" ，然后到 "google.com" ，并返回 "google.com"
	fmt.Print(browserHistory.Back(7))
}
/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
