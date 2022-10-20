package main

import "github.com/go-playground/validator"

var val = validator.New()

type RegisterRequest struct {
	UserName   string `validate :"gt=0"`           //>0
	PassWord   string `validate :"min=6,max=12"`   //6-12
	PassRepeat string `validate :eqfield=PassWord` //跨字段校验
	Email      string `validate :"email"`          //满足email格式
}
type crossStruct struct {
	PassWord string `validate :"eqcsfield = aaa.Pass"` //跨结构体相等校验
	Email    string `validate :"eqfiled = Pssword"`
	aaa      inner
}
type inner struct {
	pass  string `validate :"min=6,max=12"`
	Email string `validate : ""`
}

func main() {
	req := RegisterRequest{
		UserName:   "2",
		PassWord:   "2222222",
		PassRepeat: "2222222",
		Email:      "123.qq.com",
	}
	//struct返回的err有两种类型：
	val.Struct(req)

}
