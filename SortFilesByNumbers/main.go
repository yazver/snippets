package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func nextSpan(s string) (rest string, span string, isNum bool) {
	if s == "" {
		return
	}
	for i, r := range s {
		if i == 0 {
			isNum = unicode.IsNumber(r)
		} else if isNum != unicode.IsNumber(r) {
			return s[i:], s[:i], isNum
		}
	}
	return "", s, isNum
}

// func Nums(s string) []uint64 {
// 	ss := strings.FieldsFunc(s, func(r rune) bool {
// 		return !unicode.IsNumber(r)
// 	})
// 	nn := []uint64{}
// 	for _, v := range ss {
// 		if n, err := strconv.ParseUint(v, 10, 64); err == nil {
// 			nn = append(nn, n)
// 		}
// 	}
// 	return nn
// }

func Less(s1, s2 string) bool {
	for {
		var span1, span2 string
		var isNum1, isNum2 bool
		s1, span1, isNum1 = nextSpan(s1)
		s2, span2, isNum2 = nextSpan(s2)
		fmt.Printf("Rest1 %s, Span1 %s, IsNum1 %v\n", s1, span1, isNum1)
		fmt.Printf("Rest2 %s, Span2 %s, IsNum2 %v\n", s2, span2, isNum2)
		if span1 == "" && span2 == "" {
			return false
		}
		if isNum1 && isNum2 {
			if n1, err := strconv.ParseUint(span1, 10, 64); err == nil {
				if n2, err := strconv.ParseUint(span2, 10, 64); err == nil {
					if n1 != n2 {
						return n1 < n2
					} else {
						continue
					}
				}
			}
		}
		if span1 != span2 {
			return span1 < span2
		}
	}
}

func main() {
	s := []string{
		"10www10",
		"10www2",
		"3www10",
		"21www25",
		"21www4",
	}
	sort.Slice(s, func(i, j int) bool {
		return Less(s[i], s[j])
	})
	fmt.Printf("%v", s)
}
