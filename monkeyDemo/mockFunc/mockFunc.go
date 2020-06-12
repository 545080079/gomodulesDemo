package mockFunc

import (
	"demo/monkeyDemo/test"
	"fmt"
)

//测试用的Mock函数PrintAdd
func PrintAdd(a, b uint32) string {
	return fmt.Sprintf("a:%v+b:%v===========%v", a, b, a+b)
}

//测试用的Mock函数 对应test文件夹下的PrintSum
func PrintSum(_ *test.SumTest, a, b uint32) string {
	return fmt.Sprintf("a:%v+b:%v=========%v", a, b,a+b)
}
