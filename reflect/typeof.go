package main

import (
	"fmt"
	"reflect"
)

func main() {
	getStructRelation()
}
func get_type() {
	type User struct {
		name int
		age  string
	}
	//获取类型，返回值是reflect.type
	a := reflect.TypeOf(1)
	b := reflect.TypeOf("哈哈哈")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.Kind() == reflect.Int)
	fmt.Println(b.Kind() == reflect.String)
	user := User{}
	typeUser := reflect.TypeOf(user)
	//结构体
	fmt.Println(reflect.TypeOf(typeUser))
	fmt.Println(typeUser.Kind() == reflect.Struct)

	//获取指针
	user2 := new(User)
	typeUser2 := reflect.TypeOf(user2)
	fmt.Println(reflect.TypeOf(typeUser2))
	fmt.Println(typeUser.Kind() == reflect.Ptr)

	//指针解析
	typeUser3 := typeUser2.Elem()
	fmt.Println(typeUser3.Kind() == reflect.Struct)

	fmt.Println("--------------------")
	fmt.Println(typeUser.PkgPath())
	fmt.Println(typeUser.Name())
	fmt.Println(typeUser.Size())
}

// 获取成员变量
func getField() {
	type User struct {
		name int
		age  string
	}
	typeUser := reflect.TypeOf(User{})
	filedNum := typeUser.NumField()
	for i := 0; i < filedNum; i++ {
		field := typeUser.Field(i)
		fmt.Printf("index值%d,名称%s\n", i, field.Name)
	}
}
func add(a int) int {
	return a
}
func getFn() {
	fnType := reflect.TypeOf(add)
	fmt.Println(fnType.Kind() == reflect.Func)
	//输入参数个数
	fmt.Println(fnType.NumIn())
	//输出参数个数
	fmt.Println(fnType.NumOut())

}
func getStructRelation() {
	type C struct {
		sex string
	}
	type A struct {
		name string
		C
	}
	type B struct {
		age string
		C
	}
	ctype := reflect.TypeOf(C{})
	atype := reflect.TypeOf(A{})
	btype := reflect.TypeOf(B{})
	//类型不一样，没有严格意义上的继承
	fmt.Println(atype.AssignableTo(ctype))
	fmt.Println(btype.ConvertibleTo(ctype))
}
