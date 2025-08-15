package _81_time_based_kvs

type Item struct {
	ts  int
	val string
}

type TimeMap struct {
	hash map[string][]*Item
}

func Constructor() TimeMap {
	return TimeMap{
		hash: make(map[string][]*Item),
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := t.hash[key]; !ok {
		t.hash[key] = []*Item{}
	}
	t.hash[key] = append(t.hash[key], &Item{
		ts: timestamp, val: value,
	})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	list, ok := t.hash[key]
	if !ok {
		return ""
	}
	l, r := 0, len(list)-1
	if timestamp < list[l].ts {
		return ""
	}
	if list[r].ts <= timestamp {
		return list[r].val
	}

	for l+1 < r {
		mid := (l + r) / 2
		if timestamp < list[mid].ts {
			// search left
			r = mid
		} else {
			// search right
			l = mid
		}
	}
	return list[l].val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
