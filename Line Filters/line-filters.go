//Go 行过滤器

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	一个行过滤器在读取标准输入流的输入，处理输出，然后将得到一些的结果输出到标准输出的程序中是常见 的一个功能。
	grep和sed是常见的行过滤器
*/

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
