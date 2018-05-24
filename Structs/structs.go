//Go 结构体

package main

import (
	"fmt"
	"strings"
)

/*
	结构体申明方式：
	type identifier struct {
		field1 type1
		field2 type2
	}

	type T struct {a, b int}也是合法的语法，它更适用于简单的结构体
	var t *T
	t = new(T)
	变量t是一个指向T的指针，此时结构体字段的值是他们所属类型的零值，
	使用new函数给一个新的结构体变量分配内存，他返回指向分配内存的指针

	Go 没有class ，结构体代替了class的位置，但是他们的作用却不尽相同
*/

//Person 定义一个 Person 结构体
type Person struct {
	firstName string
	lastName  string
}

//upPerson 将名字全部转换成大写
//传入参数：指向Person结构体的指针
func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {

	//结构体作变量
	//pers1 结构体类型变量
	var pers1 Person
	pers1.firstName = "Bean"
	pers1.lastName = "Wei"
	upPerson(&pers1)
	fmt.Printf("the name is %s %s\n", pers1.firstName, pers1.lastName)

	//结构体作指针
	pers2 := new(Person)
	/*
		等价于：
		var pers2 *Person
		pers2 = new(Person)
	*/
	pers2.firstName = "Bean"
	pers2.lastName = "Wei"
	/*
		也可以这样定义:
		(*pers2).firstName = "Bean"
		(*pers2).lastName = "Wei"
	*/
	upPerson(pers2)
	fmt.Printf("the name is %s %s\n", pers2.firstName, pers2.lastName)

	//结构体作序列
	pers3 := &Person{"Bean", "Wei"}
	upPerson(pers3)
	fmt.Printf("the name is %s %s\n", pers3.firstName, pers3.lastName)
}

/*
output:
$ go run structs.go
the name is BEAN WEI
the name is BEAN WEI
the name is BEAN WEI
*/
