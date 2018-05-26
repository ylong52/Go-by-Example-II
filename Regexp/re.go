//Go 正则表达式

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p("1", match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	p("2", r.MatchString("peach"))
	p("3", r.FindString("peach punch"))
	p("4", r.FindStringIndex("peach punch"))
	p("5", r.FindStringSubmatch("peach punch"))
	p("6", r.FindStringSubmatchIndex("peach punch"))
	p("7", r.FindAllString("peach punch pinch", -1))
	p("8", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	p("9", r.FindAllString("peach punch pinch", 2))
	p("10", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	p("11", r)
	p("12", r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p("13", string(out))
}

/*
outut:
$ go run re.go
1 true
2 true
3 peach
4 [0 5]
5 [peach ea]
6 [0 5 1 3]
7 [peach punch pinch]
8 [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
9 [peach punch]
10 true
11 p([a-z]+)ch
12 a <fruit>
13 a PEACH
*/
