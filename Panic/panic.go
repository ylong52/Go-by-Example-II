//Go 异常处理之panic

package main

/*
	Golang中提供panic用于错误处理，当调用panic时，正常的执行流程将停止
	panic是用来表示非常严重的不可恢复的错误的，是Go语言中的内置函数，panic的作用就像我们平常接触的异常。
	不过Go没有try···catch(except),所以panic一般会导致程序主动退出除非recover

	panic与defer
	即使函数执行的时候panic了，函数不往下走了，运行的时候并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了panic再
	向上传递。所以这个时候有点雷士tyr···catch(except)···finally中的finally

	panic是抛出的真正意义上的异常
*/

func main() {
	panic("A Problem")
	//程序会打印出调用栈
}

/*
output:
$ go run panic.go
panic: A Problem

goroutine 1 [running]:
main.main()
        E:/WebProject/Code/Go/src/github.com/BeanWei/Go-by-Example/Panic/panic.go:10 +0x40
exit status 2
*/
