package service

import (
	"demo/controller"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestSession_Get(t *testing.T) {
	convey.Convey("Try to Get Session(123): \n", t, func() {
		var client controller.Client
		sess := Session{
			client,
		}
		id, err := sess.Get("123")
		convey.So(err, convey.ShouldBeNil)
		fmt.Println(id)
		convey.So(id, convey.ShouldEqual, 123)
	})

}

func TestSession_Set(t *testing.T) {

	convey.Convey("Try to Set Session(300): ", t, func() {

		var client controller.Client
		sess := Session{
			client,
		}
		token, err := sess.Set(123)
		convey.So(err, convey.ShouldBeNil)
		convey.So(token, convey.ShouldNotBeEmpty)

	})

}

func TestRunServiceOne_testing(t *testing.T) {
	tp := 5
	err := RunServiceOne(tp)
	if err != nil {
		t.Errorf("tp = %d, err = %v, RunServiceOne测试失败", tp, err)
	}
	tp = 9
	err = RunServiceOne(tp)
	if err != nil {
		t.Errorf("tp = %d, err = %v, RunServiceOne测试失败", tp, err)
	}

}

func TestRunServiceOne_table(t *testing.T) {
	tests := []struct {
		tp int
	}{
		// TODO: test cases
		{1}, {2}, {3}, {5}, {9},
	}
	for _, test := range tests {
		t.Run(string(test.tp), func(t *testing.T) {

			err := RunServiceOne(test.tp)
			if err != nil {
				t.Errorf("tp = %d, err = %v, RunServiceOne测试失败", test.tp, err)
			}
		})
	}
}

func ExampleDemo1() {
	for _, val := range rand.Perm(4) {
		fmt.Println(val)
	}
	//Unordered Output:
	//3
	//2
	//0
	//1
}
func ExampleDemo2() {
	nums := []int{50, 60, 80, 70}
	for _, val := range nums {
		fmt.Println(val)
	}
	//Output:
	//50
	//60
	//80
	//70

}
