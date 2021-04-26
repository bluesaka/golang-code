/**
单元测试

goland中，右键 -> Generate -> Tests for ... 生成xxx_test.go文件
*/
package test

func echo(s interface{}) interface{} {
	return s
}

func say(s string) string {
	return s
}
