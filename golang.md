# Go

## 复合数据类型

## Array

#### 定义

数组是具有***固定长度***，拥有零个或者多个***相同数据类型***的元素序列

#### 初始化

1.数组通过***索引***来访问元素，索引从0到数组长度-1，使用***len函数***返回数组的长度

2.新数组的元素初始值为元素类型的***零值***，例如int类型，为0

3.省略号***...***出现在数组的长度位置，那么数组长度将由***初始化数组元素的个数***决定

```go
var a [3]int

//output：0,第一个元素
fmt.Printf("%d\n",a[0])

//output：末尾的元素
fmt.Printf("%d\n",a[len(a)-1])

//0值,len=100，最后一个元素为1，其他为0
r := [...]int{99:1}

//输出索引和元素
for k,v := range a {
    fmt.Printf("%d,%d\n",k,v)
}

//只输出元素
for _,v := range a {
    fmt.Printf("%d,%d\n",k,v)
}

//...使用，由后面元素动态决定
var sx := [...]int{1,2,3,4,5}//len=5
```

#### 类型

数组的类型包括***数组长度***，[3]int 和[4]int 是不同的类型，编译时候确定

```go
q :=[3]int{1,2,3}
q = [4]int{1,2,3,4}
```

#### 比较

数组类型是可比较的，可以***==***比较，两边的值是否完全相同，反之***!=***

```go
d :=[3]int{1,2,3}
c :=[3]int{1,2,3}
fmt.Printf("b",d == c)
```

#### 作为参数

当调用一个函数的时候，每个传入的参数都会创建一个***副本***，然后赋值给***对应的函数变量***，所以函数***接受的是一个副本，而不是原始参数***，消耗性能

显示传递一个数组的***指针***，这样对数组的修改可以反映到原始数组上

```go
arr := [...]int{1,2,3,4,5}
change(&arr)
func change(arr *[len(arr)]int)
```

#### 循环

```go
a := [...]int{124, 12, 66, 34, 12}
	for k, v := range a {
		fmt.Printf("键%d，值%d\n", k, v)
}
```

## Slice

#### 定义

slice是一个拥有相同元素类型的***可变长度***序列，类似没有长度的数组

slice是一个轻量的数据结构，可以访问数组的部分或者全部的元素，这个数组称为slice***底层数组***

slice包括三个部分：***容量，长度，底层数组***

容量：起始元素直底层数组的最后一个元素的个数

#### 创建

```go
months :=string{1:"january",12:"december"}
```

slice[i:j]，（0<=i<=j<=cap(s)），***引用***着i到j-1个元素，即长度为j-i个

slice[i:]，引用着i到len(s)-1，最后一个元素，相当于slice[i:len(s)]

slice[:]，引用着整个数组

因为slice包含了***指向数组元素的指针***，将一个slice传递给函数的时候，可以在***函数内部修改底层数组的元素***

```go
var arr = [...]int{1,2,3,4,5}
func reverse (s []int){
	for i,j :=0,len(s)-1;i<j;i,j = i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}
//output:5,4,3,2,1
reverse(arr[:])
```

#### 区别

arr :=[3]int{1,2,3}

s :=[]int{1,2,3}

初始化slice和初始化数组，slice***没有指定长度***，创建有固定的长度的数组和创建指向数组的slice

#### 比较

不能用==来比较两个slice是否有两个相同的元素

如果底层数组元素改变，同一个slice在不同时间有不同的元素

slice唯一允许的比较操作是和nil作比较

make函数

创建一个无名数组并返回了一个它的slice

```go
make(T[],len)
make(T[],len,cap)
```

### append函数

append函数用来将元素追加到slice后面

对于任何函数，只要有可能改变slice的长度或者容量，或使slice指向不同的底层数组，都需要更新slice的变量

```go
var runes []rune
	for _, v := range "hello,你好" {
		runes = append(runes, v)
	}
	fmt.Printf("%q\n", runes)
```

```go
func main() {
	s := make([]int, 3)
	fmt.Printf("原切片%v,%d,%d，%p\n", s, len(s), cap(s), &s)
	appendSlice(s)
	fmt.Printf("原切片%v,%d,%d,%p\n", s, len(s), cap(s), &s)
}
func appendSlice(s []int) {
	s[0] = 9
	s = append(s, 1)
	//fmt.Printf("%v", s)
	fmt.Printf("子切片%v,%d,%d,%p\n", s, len(s), cap(s), &s)
}
//append追加元素后，不影响母切片

```

##### 方法一

```go
func demo() {
	data := []string{"one", "", "three"}
	res := noempty(data)
	fmt.Printf("%q\n", res)  //"one","three"
	fmt.Printf("%q\n", data) //"one", "three", "three"

}

func noempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings
}
```

##### 方法二

```go
func useAppend() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data) //"one","","three"
	res := noempty1(data)
	fmt.Printf("%q\n", res)  //"one","three","three"
	fmt.Printf("%q\n", data) //"one", "three", "three"
}
func noempty1(strings []string) []string {
    //相当于应用原始的slice的新的零长度的slice
	out := strings[:0]
	for _, v := range strings {
		if v != "" {
			out = append(out, v)
		}
	}
	return strings
}
```

## Map

#### 定义

拥有一对键值对的***无序***元素集合，***键的值***是***唯一***的，值可以***通过键***来获取、更新或者移除

map是对散列表的引用，m[k]v（m[键类型]值类型）

#### 创建

可以使用make函数来创建  
