package algodaily

import (
	"fmt"
	"strings"
)

type count struct {
	number int
	index  int
}

func FirstNonRepeating(s string) {
	ss := strings.Split(s, "")
	string_count := make(map[string]*count)
	for i, ch := range ss {
		val, ok := string_count[ch]
		if ok {
			val.number++
		} else {
			string_count[ch] = &count{number: 1, index: i}
		}
	}
	// we can write a reduce funciton here that takes the map and another function by with it can reduce it
	var nch string = ""
	for k, v := range string_count {
		if v.number > 1 {
			continue
		} else if nch != "" {
			if string_count[nch].index > v.index {
				nch = k
			}
		} else if nch == "" {
			nch = k
		}
	}

	if nch == "" {
		fmt.Println("doesn't exist")
	} else {
		fmt.Println(nch)
	}
}

func CheckFirstNonRepeating() {
	FirstNonRepeating("aaaaabbbbs%%%%ddd")
}
