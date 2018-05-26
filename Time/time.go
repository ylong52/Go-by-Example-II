//Go 时间

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	//方法sub返回一个Duration(时间段)来表示两个时间点的间隔时间
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

/*
output:
$ go run time.go
2018-05-26 14:04:41.8264569 +0800 CST m=+0.006000301
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
74673h29m43.175069663s
74673.49532640824
4.480409719584495e+06
2.688245831750697e+08
268824583175069663
2018-05-26 06:04:41.8264569 +0000 UTC
2001-05-12 11:05:15.476317574 +0000 UTC
*/
