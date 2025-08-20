package _55_design_twitter

import "container/heap"

type Tweet struct {
	id int
	ts int
}

type MinHeap []*Tweet

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].ts < m[j].ts }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*Tweet))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type Twitter struct {
	Time     int
	Follower map[int]map[int]struct{}
	Posts    map[int][]*Tweet
}

func Constructor() Twitter {
	return Twitter{
		Time:     0,
		Follower: make(map[int]map[int]struct{}),
		Posts:    make(map[int][]*Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Posts[userId]; !ok {
		this.Posts[userId] = []*Tweet{}
	}
	this.Posts[userId] = append(this.Posts[userId], &Tweet{
		id: tweetId,
		ts: this.Time,
	})
	this.Time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	userHash := this.Follower[userId]
	h := &MinHeap{}
	heap.Init(h)
	for uId, _ := range userHash {
		tweets := this.Posts[uId]
		for _, tweet := range tweets {
			heap.Push(h, tweet)
			if h.Len() > 10 {
				heap.Pop(h)
			}
		}
	}
	userTweets := this.Posts[userId]
	for _, tweet := range userTweets {
		heap.Push(h, tweet)
		if h.Len() > 10 {
			heap.Pop(h)
		}
	}
	res := make([]int, h.Len())
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(*Tweet).id
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Follower[followerId]; !ok {
		this.Follower[followerId] = make(map[int]struct{})
	}
	this.Follower[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.Follower[followerId]; !ok {
		return
	}
	delete(this.Follower[followerId], followeeId)
}
