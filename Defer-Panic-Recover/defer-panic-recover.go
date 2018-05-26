//Go 异常处理组合

package main

import (
	"fmt"
	"time"
)

/*
	Golang的异常处理组合panic，defer，recover。
	使用panic抛出异常，抛出异常后立即停止当前函数的执行并运行所有被defer的函数，
	然后将panic抛向上一层直至程序carsh。但是也可以使用被defer的recover函数来捕获异常阻止程序崩溃
	recover只有被defer后才有意义。

	note：
	1.defer需要放在panic之前定义，另外recover只有在defer调用的函数中才有效
	2.recover处理异常后，逻辑并不会恢复到panic那个点去，函数会跑到defer之后的那个点
	3.多个defer会形成defer栈，后定义的defer语句会被先调用
*/

func main() {
	//必须先声明defer，不然无法铺货panic异常
	//panic("BUG")之后来到这里
	defer func() {
		fmt.Println("2")
		//调用recover函数保证程序继续向下执行
		if err := recover(); err != nil {
			//这里的err就是panic传入的内容"BUG"
			fmt.Println(err)
		}
		fmt.Println("3")
	}()
	F()
}

//F 抛出异常
func F() {
	for {
		fmt.Println("1")
		//抛出异常
		panic("BUG")
		//panic后执行defer向上执行，后面的部分不执行
		fmt.Println("4") //这里不会运行
		time.Sleep(time.Second)
	}
}

/*
output:
$ go run defer-panic-recover.go
1
2
BUG
3
*/

/*
	通过输出结果我们能看到：输出"1",然后到了panic不会马上执行而是跑到了defer执行，
	输出"2",然后recover处理异常输出panic的输出"BUG",最后输出"3"
*/
