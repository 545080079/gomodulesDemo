package test

import "fmt"

func PrintAdd(a, b uint32) string {
	return fmt.Sprintf("a:%v+b:%v", a, b)
}


type SumTest struct {
}

func (*SumTest)PrintSum(a, b uint32) string {
	return fmt.Sprintf("a:%v+b:%v", a, b)
}

