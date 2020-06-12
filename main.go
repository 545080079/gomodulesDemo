package main

import (
	"bou.ke/monkey"
	"demo/monkeyDemo/mockFunc"
	"demo/monkeyDemo/test"
	"fmt"
	"reflect"
	"sort"
)

type Dto struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Like string `json:"like"`
}


func Fun1() {
	monkey.Patch(test.PrintAdd, func(a, b uint32) string{
		return mockFunc.PrintAdd(a, b)
	})
	p := test.PrintAdd(1, 2)
	fmt.Println(p)
	monkey.UnpatchAll() //解除所有替换
	p = test.PrintAdd(1, 2)
	fmt.Println(p)
}

func Fun2() {
	structSum := &test.SumTest{}
	//para1:获取实例的反射类型,para2:被替换的方法名,para3:替换方法
	monkey.PatchInstanceMethod(reflect.TypeOf(structSum), "PrintSum", mockFunc.PrintSum)
	p := structSum.PrintSum(1, 2)
	fmt.Println(p)
	monkey.UnpatchAll() //解除所有替换
	p = structSum.PrintSum(1, 2)
	fmt.Println(p)

}

func Fun3(s string) string{
	return s + "f"
}

func helper(nums []int) {
	res := make([][]int, 0)
	res = append(res, []int{1})
	res = append(res, []int{3,4})
	fmt.Println(res)
	sort.Ints(nums)
}


func insert(nums []int, start int, value int) []int{
	if len(nums) < 2 {
		return nums
	}
	for i:= len(nums) - 2; i >= start; i-- {
		nums[i + 1] = nums[i]
	}
	return nums
}


const (
	Address = "127.0.0.1:9090"
)

func main() {
}

//
//func main() {
//	monkey.Patch(Fun3, func(s string) string{
//		return s + "a"
//	})
//
//	fmt.Println(Fun3("luoyutao"))
//
//
//	dd := &Dto{
//		Name: "xiaorui",
//		Addr: "rfyiamcool@163.com",
//		Like: "Python & Golang",
//	}
//	v, err := json.Marshal(dd)
//	fmt.Println(string(v), err)
//	monkey.Patch(json.Marshal, func(v interface{}) ([]byte, error) {
//		fmt.Println("use jsoniter")
//		return jjson.Marshal(v)
//	})
//
//	monkey.Patch(json.Unmarshal, func(data []byte, v interface{}) error {
//		fmt.Println("use jsoniter")
//		return jjson.Unmarshal(data, v)
//	})
//
//	resDto := &Dto{}
//
//	v, err = json.Marshal(dd)
//	fmt.Println(string(v), err)
//
//	errDe := json.Unmarshal(v, resDto)
//	fmt.Println(resDto, errDe)
//
//	fmt.Println("test end")
//}