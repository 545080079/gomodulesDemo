package main

import (
	"fmt"
	"net/http"
	"strings"
)

var pr = FPrintRequest

func FPrintRequest(r *http.Request) {
	fmt.Println("r.Form=", r.Form)         //这些信息是输出到服务器端的打印信息 , Get参数
	fmt.Println("r.PostForm=", r.PostForm) //Post参数
	fmt.Println("path=", r.URL.Path)
	fmt.Println("scheme=", r.URL.Scheme)
	fmt.Println("method=", r.Method) //获取请求的方法

	fmt.Println("Http Get参数列表 begin:")
	for k, v := range r.Form {
		fmt.Println("Http Get["+k+"]=", strings.Join(v, " ; "))
	}
	fmt.Println("Http Get参数列表 end:")

	fmt.Println("Http Post参数列表 begin:")
	for k, v := range r.PostForm {
		fmt.Println("Http Post["+k+"]=", strings.Join(v, " ; "))
	}
	fmt.Println("Http Post参数列表 end:")

	arraA := r.Form["a"]
	fmt.Println("r.Form['a']=", arraA)
	if len(arraA) > 0 {
		fmt.Println("r.Form['a'][0]=", arraA[0])
	}
}

func FSayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("******************************************sayhelloName")
	FPrintRequest(r)
	fmt.Fprintf(w, "<h1>Hello 世界!</h1>") //这个写入到w的是输出到客户端的
}

var Exec = func(id int) (string, error) {
	fmt.Println("Normal invoke.")
	return "Normal", nil
}

//定义：资源清理
var Process = func(id int) {

	//Do Clean
	fmt.Println("Clean ok.")
}

//func main() {
//
//	//为函数打桩，提供一个指定返回值
//	stubs := gostub.StubFunc(&Exec, "test", nil)
//	defer stubs.Reset()
//	fmt.Println(Exec(0))
//
//	//为过程（无返回值的函数）打桩
//	stubs = gostub.StubFunc(&Process)
//	defer stubs.Reset()
//	Process(0)
//
//	//用Stub()打桩
//	//stubs := gostub.Stub(&PrintRequest, func(r *http.Request) (string, error){
//	//	return "test", nil
//	//})
//
//	//使用StubFunc()更简单
//	//PrintRequest被代替
//	//stubs := gostub.StubFunc(&pr)
//	//http.HandleFunc("/", FSayhelloName)
//	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
//	//if err != nil {
//	//	log.Fatal("ListenAndServe: ", err)
//	//}
//	//defer stubs.Reset()
//}
