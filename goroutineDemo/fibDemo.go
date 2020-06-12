package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func fib(x int) int{
	if x < 2 {
		return x
	}
	return fib(x-1)+fib(x-2)
}


func spinner(delay time.Duration) {
	for {
		s := "-\\|/"
		for _,r := range s{
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

var cp = flag.String("cp", "", "write to file")
func main(){
	//go spinner(100 * time.Millisecond)
	//const n = 45
	//fibN := fib(n)
	//fmt.Printf("Fib(%d) = %d\n", n, fibN)

	flag.Parse()
	if *cp != "" {
		fmt.Println("zcx")
		f, err := os.Create(*cp)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	} else {
		fmt.Println("null")

	}

	fib(12)


}
