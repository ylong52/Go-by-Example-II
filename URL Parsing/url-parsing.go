//Go 解析URL

package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	s := "postgres://user.pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

/*
output:
$ go run url-parsing.go
postgres
user.pass

host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
*/
