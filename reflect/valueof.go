package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		name int
	}
	//valueof操作
	ivalue := reflect.ValueOf(1)
	iv2 := reflect.ValueOf("哈哈哈")
	fmt.Println(ivalue, iv2)
	itype := ivalue.Type()
	ivwtype := iv2.Type()
	fmt.Println(itype, ivwtype)
	fmt.Println(itype.Kind() == ivalue.Kind())
	//指针解析和获取的转换
	svalue := reflect.ValueOf(&user{
		name: 23,
	})
	//解析指针
	svalue2 := svalue.Elem()
	fmt.Println(svalue2.Kind() == reflect.Struct)
	//获取指针
	svalue3 := svalue2.Addr()
	fmt.Println(svalue3.Kind() == reflect.Ptr)

	//interface,转换成interface类型
	//转换成interface在转换成int,直接int()
	fmt.Println(ivalue.Interface().(int), ivalue.Int())

	//struct指针
	fmt.Println(svalue.Interface().(*user))
	var a interface{}
	fmt.Println(reflect.ValueOf(a).IsValid(), reflect.ValueOf(a).Kind() == reflect.Invalid)

	//调用isnil之前必须保证isValid为true,isZero同理
	var u *user = nil
	if reflect.ValueOf(u).IsValid() {
		s := reflect.ValueOf(u)
		fmt.Println(s.IsNil())
	}
	assemblyptr()
	seeptr()
	changemap()
}
func assemblyptr() {
	va1 := reflect.ValueOf(1)
	var x int
	v2 := reflect.ValueOf(x)
	v3 := reflect.ValueOf(&x)
	v4 := v3.Elem()

	//寻址,valueof转换再用Elem()寻址
	fmt.Println(va1.CanAddr(), v2.CanAddr(), v3.CanAddr(), v4.CanAddr())

	//切片寻址,切片不可寻址，某个元素可以寻址
	s1 := make([]int, 3, 5)
	s3 := reflect.ValueOf(s1)
	s2 := s3.Index(0)
	fmt.Println(s3.CanAddr(), s2.CanAddr())

	//map
}
func seeptr() {
	i := 10
	ivalue := reflect.ValueOf(&i)
	if ivalue.Elem().CanAddr() {
		ivalue.Elem().SetInt(222)
	}
	fmt.Println(i)
}
func changeslice() {
	type user struct {
		name int
	}
	sl := make([]*user, 3, 5)
	sl[0] = &user{
		name: 2,
	}
	//修改值
	value := reflect.ValueOf(sl)
	if value.Len() > 0 {
		value.Index(0).Elem().FieldByName("name").SetInt(90)
	}
	//替换0位置的元素
	value.Index(0).Set(reflect.ValueOf(&user{
		name: 77,
	}))
	//设置cap，不能升
	value.Elem().SetCap(9)
}
func changemap() {
	mp := make(map[int]bool, 10)
	mp[0] = false
	mvalue := reflect.ValueOf(mp)
	mapType := mvalue.Type()
	keyType := mapType.Key()    //获取key的类型
	valueType := mapType.Elem() //获取value的类型
	fmt.Println(mapType, keyType, valueType)
	mvalue.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf(false))

}
