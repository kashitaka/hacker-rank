package _33_number_of_recent_calls

type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	var que []int
	return RecentCounter{Queue: que}
}


func (this *RecentCounter) Ping(t int) int {
	this.Queue = append(this.Queue, t)
	cnt := 0
	for len(this.Queue) > 0 {
		pop := this.Queue[0]
		if pop < t-3000 {
			this.Queue = this.Queue[1:]
			continue
		}
		if t-3000<=pop && pop <=t {
			cnt = len(this.Queue)
			break
		}
	}
	return cnt
}