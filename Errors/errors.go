//Go 异常(error)

package main

/*
	error类型本身就是一个预定义好的接口，里面定义了一个method
	type error interface {
		Error() string
	}
*/

import (
	"errors"
	"fmt"
)

//f1 采用errors包的New方法，返回一个error的类型
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//f2 采用自定义的方式实现一个error的一个鸭子类型
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

//f3 采用fmt.Errof 将string信息转换为error信息并返回
/*
	在这个方法的内部，先调用fmt包中的Sprintf方法把格式化的输入转换为字符串，在使用errors.New方法返回error类型
*/
func f3(arg int) (int, error) {
	if arg == 42 {
		return -1, fmt.Errorf("%s", "can't work")
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2  failed:", e)
		} else {
			fmt.Println("f2 workrd", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f3(i); e != nil {
			fmt.Println("f3 failed:", e)
		} else {
			fmt.Println("f3 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

/*
output:
$ go run errors.go
f1 worked: 10
f1 failed: can't work with 42
f2 workrd 10
f2  failed: 42 - can't work with it
f3 worked: 10
f3 failed: can't work
42
can't work with it
*/
