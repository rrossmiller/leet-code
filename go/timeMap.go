package main

import "fmt"

type tuple struct {
	value     string
	timestamp int
}
type TimeMap struct {
	store map[string][]tuple
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]tuple)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// key doesn't exist make the slice
	if _, exists := this.store[key]; !exists {
		this.store[key] = []tuple{tuple{value, timestamp}}
	} else {
		s := this.store[key]
		s = append(s, tuple{value, timestamp})
		this.store[key] = s
	}

}

func (this *TimeMap) Get(key string, timestamp int) string {
	rtn := ""
	if vals, exists := this.store[key]; exists {
		// bin search to find timestamp
		l, r := 0, len(vals)-1

		for l <= r {

			m := (l + r) / 2
			if vals[m].timestamp == timestamp {
				return vals[m].value
			} else if vals[m].timestamp > timestamp {
				//search left
				r = m - 1
			} else {
				rtn = vals[m].value
				//search right
				l = m + 1
			}
		}

	}
	return rtn
}

