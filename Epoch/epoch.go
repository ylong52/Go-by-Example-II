//Go 时间戳

package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	//单位秒
	secs := now.Unix()
	//单位纳秒
	nanos := now.UnixNano()
	fmt.Println(now)

	//UnixMillis不存在，必须手动转化纳秒为毫秒
	millis := nanos / 1000000

	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//整数秒转化到相应时间
	fmt.Println(time.Unix(secs, 0))

	//纳秒转化到相应时间
	fmt.Println(time.Unix(0, nanos))
}

/*
output:
$ go run epoch.go
2018-05-26 14:16:39.6625148 +0800 CST m=+0.006000401
1527315399
1527315399662
1527315399662514800
2018-05-26 14:16:39 +0800 CST
2018-05-26 14:16:39.6625148 +0800 CST
*/
