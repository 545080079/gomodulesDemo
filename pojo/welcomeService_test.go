package pojo

import (
	mock_pojo "demo/mockPojo"
	"github.com/golang/mock/gomock"
	"testing"
)

//正常测试
func TestCompany_Meeting(t *testing.T) {
	talker := NewPerson("小微", "语音服务助手")
	company := NewCompany(talker)
	t.Log(company.Meeting("罗宇韬", "学生"))

}

//通过Mock测试
func TestCompany_Meeting2(t *testing.T) {

	//新建Mock控制器
	ctrl := gomock.NewController(t)
	//新建Mock对象-Talker
	talker := mock_pojo.NewMockTalker(ctrl)

	//断言
	talker.EXPECT().SayHello(gomock.Eq("震天嚎"), gomock.Eq("学生")).Return("Hello Faker(身份：学生), welcome to GoLand IDE. My name is 震天嚎")
	//mock对象传入方法
	company := NewCompany(talker)

	//Pass的例子
	t.Log(company.Meeting("震天嚎", "学生"))

	//报错的例子
	//t.Log(company.Meeting("小白", "学生"))

}
