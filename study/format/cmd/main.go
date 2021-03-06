/**
占位符

一般占位符
%v	相应值的默认格式
%+v	在打印结构体时，默认格式，会添加字段名
%#v	相应值的 Go 语法表示
%T	相应值的类型的 Go 语法表示
%%	字面上的百分号，并非值的占位符

布尔占位符
%t	单词 true 或 false

整数占位符
%b	二进制表示
%c	相应 Unicode 码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由 Go 语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode 格式：U+1234，等同于 "U+%04X"

浮点数及其复合构成占位符
%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
%e	科学计数法，例如 -1234.456e+78
%E	科学计数法，例如 -1234.456E+78
%f	有小数点而无指数，例如 123.456
%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出

字符串与字节切片占位符
%s	字符串或切片的无解译字节
%q	双引号围绕的字符串，由 Go 语法安全地转义
%x	十六进制，小写字母，每字节两个字符
%X	十六进制，大写字母，每字节两个字符

指针
%p	十六进制表示，前缀 0x

*/

package main

import "fmt"

func main() {
	fmt.Printf("%T\n", 1)
	fmt.Printf("%#v\n", 1)
	fmt.Printf("%#v\n", "111")
	fmt.Printf("%v\n", "111")

	fmt.Printf("%t\n", true)
	fmt.Println(false)

	i := 111
	fmt.Printf("%p\n", &i)

	// 左对齐，6 30为宽度 (like gin)
	fmt.Printf("[DD]%-6s %-30s --> %s\n", "GET", "/api/user/info", "/web/controllers.UserInfo")
	fmt.Printf("[DD]%-6s %-30s --> %s\n", "POST", "/api/user/login", "/web/controllers.UserLogin")
}
