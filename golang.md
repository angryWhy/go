# Go

# 基础数据类型



```go
bool		 1字节		默认值：false
byte		 1字节		默认值：uint8，取值范围[0,255]
rune		 4字节		unicode ponit code int32
int，uint     4或8字节   取决于操作系统，32或64
int8,uint8	  1字节	 -128~127，0~255
int16，uint16 2字节	
float32
float64
complex64				  对应的实部或者虚部：float64
complex128				  对应的实部或者虚部：float128
uintptr
```

## string

string每个元素称为字符：

byte：1个字节，ASCII码的一个字符

rune:4个字节，代表一个utf-8的字符，***一个汉字可以用一个rune表示***

string底层是***byte数组***，可以转换成***[]byte或者[]rune***类型

```go
len(str)//默认的是[]byte类型
for循环默认是[]rune类型
```

#### API

```go
len(str)			求长度
strings.Split		分割
strings.Contains	判断是否包含
strings.HasPrefix	前后缀判断
strings.index，lastindex	子串出现位置

func split_str() {
	s := "我是,123"
	arr := strings.Split(s, ",")
	fmt.Println(arr, arr[0])
}
func contains_str() {
	s := "hello world"
	fmt.Printf("%t\n", strings.Contains(s, "e"))
}
func before_after() {
	s := "hello world"
	fmt.Printf("%t\n", strings.HasPrefix(s, "he"))
	fmt.Printf("%t\n", strings.HasSuffix(s, "ld"))
}
func str_func(){
	s := "hello world"
    ss := "lo"
    position :=strings.Index(s,ss)
}
```

#### 拼接

加号连接

func fmt.Sprintf( format string,a ...interface{}) string

func strings.Join(elem []string,step) string

strings.builder，大量连接，***效率最高***

```go
func join_str() {
	s1 := "a"
	s2 := "b"
	s3 := "c"

	str := s1 + " " + s2 + " " + s3
	str2 := fmt.Sprintf("%s %s %s", s1, s2, s3)
	str3 := strings.Join([]string{s1, s2, s3}, " ")
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	s4 := sb.String()
	fmt.Println(str, str2, str3, s4)
}
```

# 自定义类型

```go
type singal int8
type flo float64
type add func(a,b int)int
```

# 复合数据类型

```go
array		默认值：无		值类型
struct		默认值：无		值类型
string		默认值：""		 utf-8字符串
slice		默认值：nil		 引用类型
map			默认值：nil		 引用类型
channel		默认值：nil		 引用类型
interface	默认值：nil		  接口
function	默认值：nil		  函数
```

### Array

#### 定义

数组是具有***固定长度***，拥有零个或者多个***相同数据类型***的元素序列

数组是连续的内存空间，声明的时候必须指定长度，***长度不能改变***，声明的时候把内存空间分配好

数组的***地址就是首元素地址***

#### 初始化

1.数组通过***索引***来访问元素，索引从0到数组长度-1，使用***len函数***返回数组的长度

2.新数组的元素初始值为元素类型的***零值***，例如int类型，为0

3.省略号***...***出现在数组的长度位置，那么数组长度将由***初始化数组元素的个数***决定

```go
var a [3]int
//根据index赋值
var b = [3]int{2:99}
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
for i :=0;i<len(a);i++{
	
}
```

### Slice

#### 定义

slice是一个拥有相同元素类型的***可变长度***序列，类似没有长度的数组

slice是一个轻量的数据结构，可以访问数组的部分或者全部的元素，这个数组称为slice***底层数组***

slice包括三个部分：***容量，长度，底层数组***

容量：起始元素直底层数组的最后一个元素的个数

***切片地址和数组首元素地址是两码事***

```go
type slice struct{
	arr unsafe.Pointer
    len int
    cap int
}
```

#### 创建

```go
var s []int
s = []int{}				//len=cap=0
s = make([]int,3)		//初始化，len=3，cap=3
s = make([]int,3,5)		//len=3,cap=5
s = []int{1,2,3,4,5}	//len=cap=5
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

#### make函数

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

//append超出容量
s := make([]int, 3, 5)
	s = append(s, 1)
	fmt.Printf("添加后的地址%p\n", s)
	s = append(s, 1)
	fmt.Printf("添加后的地址%p\n", s)
	s = append(s, 1)
	fmt.Printf("超出容量,添加后的地址%p\n", s)
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

### 截取 子切片

```go
s :=make([]int,3,5)
sub :=s[1:3]
//刚开始子切片和母切片共享内存空间，修改子切片会反映到母切片上，在子切片执行append方法会把新元素添加到母切片预留的内存空间上
//当子切片不断执行append，消耗完母切片的预留空间，子切片和母切片会发生内存分离，没有关系
```

### Map

#### 定义

拥有一对键值对的***无序***元素集合，***键的值***是***唯一***的，值可以***通过键***来获取、更新或者移除

go map 底层是实现的***hash table***，根据key查找value的时间复杂度是***O(1)***

key--->计算哈希值--->对槽位总数取模--->找到槽位

冲突太多会***自动扩容***，增加槽位数，***重新分配***槽位

map是对***散列表的引用***，***m[k]v（m[键类型]值类型）***

#### 创建

可以使用make函数来创建  

```go
ages :=make(map[int]string)
ages := map(int)int{1:2,2:3}
var m map[int]int
m = make(map[int]int) 		//cap=0
m = make(map[int]int,10) 	//cap=10,减少扩容概率
```

#### 访问

map通过下标访问值

通过delete删除

```go
ages :=make(map[int]string)
ages[1] = "2"
delete(ages,1)
```

即使对应的键不存在，就返回值的***零值***

map元素的地址***不能获取***，不是一个变量

for循环结合range关键字来遍历map中所有的键和对应的值,***顺序不固定***

```go
func main() {
	ages := make(map[string]int)
	names := map[int]int{1: 2}
	fmt.Println(ages, names)
}
func sortNames(names map[int]int) {
	sl := make([]int, 0, len(names))
	for _, v := range names {
		sl = append(sl, v)
	}
	sort.Ints(sl)
	for _, v := range sl {
		fmt.Println(v)
	}
}
```

map类型的***零值是nil***，也就是说，没有引用任何散列表

```go
var map map[int]int
map == nil
len(map) == 0
```

向nil的map直接设置值，会导致报错，***必须要初始化***

```go
var mapa map[int]int
mapa[1] = 1
//panic: assignment to entry in nil map 
```

通过下标访问，总会有值，如果没有则为对应类型的零值，安全判断

通过下标访问，会输出两个值，***第二个值是布尔值***，用来说明值是否存在

```go
names := map[int]int{1: 2}
value,ok :=names[1]
if ok{
	//执行操作
}
if value,ok := names[1];!ok{
	//执行操作操作
}
```

#### 比较

和slice一样***不可比较***,自己写一个比较的函数

```go
func equal(x, y map[int]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if value, ok := y[k]; !ok || value != v {
			return false
		}
	}
	return true
}
```

map的值本身可以是复合数据类型，例如map和slice

### Channel

#### 定义

channel底层是一个环形队列（先进先出），send(插入),和recv（取走）从同一个位置沿同一个方向顺序执行

#### 声明

```go
var ch chan int
ch = make(chan int,8)//初始化，滑行队列最多容纳8个元素
```

#### 访问

```go
var ch chan int
//放入
ch<-1
//取出
v := <-ch
```

#### 遍历

```go
var ch chan int
close(ch)		//遍历管道前必须关闭
for ele := range ch {
	
}
for循环可以不用关闭channel
range关闭管道
```

#### 同步与异步

关闭channel后，读操作会立即返回，如果channel没有元素，返回零值

不能重复关闭，不能关闭值为nil

不能向已关闭的channel写入

```go
ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for ele := range ch {
		fmt.Println(ele)
	}
	v := <-ch
	fmt.Println(v)
```

### Struct

结构体是将零个或者多个***任意类型的命名变量***组合在一起的***聚合***数据类型

相同类型的成员变量可以写在一行

#### 创建

结构体不可以定一个和自己一样的结构体作为自己的成员变量，***不能包含自己***

定义结构体但是赋值，那么是对应的零值

结构体的***零值***是有结构体成员的零值组成的，没有任何成员变量的结构体成为空结构体，struct{}

```go
type User struct{
	ID	    int
    name	string
    bir		time.Time
}
```

####  访问

struct每一个成员变量都通过  ***.***  来访问，或者获取成员变量的地址，通过指针来访问它

  ***.***  也可以应用到结构体的指针上

如果结构体变量名称是***大写***的，那么这个变量是可以***导出***的，结构体可以包含可导出和不可以导出的变量

```go
type User struct{
	ID	    int
    name	string
    bir		time.Time
}
var wang User
//.访问
wang.ID = 999
//变量地址
value := &wang.ID
value = 888
//结构体指针
var zhang *User = &wang 
```

#### 字面量

1.按照正确顺序，为每个成员变量指定一个值

2.指定变量的名称初始化赋值

```go
type Point struct{x,y int}
p :=Point{1,2}
p :=Point{x:2}
```

#### 指针

结构体一般是通过指针使用,一般防止值拷贝

```go
pp :=&Point{1,2}
//等价于
pp :=new(Point)
*pp = Point{1,2}
```

#### 比较

如果成员变量都可以比较，那么结构体是可以比较的

按成员变量的顺序比较

```go
type Point struct{x,y int}
p :=Point{1,2}
q :=Point{2,1}
//p!=q
```

#### 匿名和嵌套

结构体内允许定义***不带名称***的结构体成员，指定类型即可，叫做***匿名成员***

匿名成员：必须是一个命名类型或者指向命名类型的指针

```go
type Point struct {
	x, y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.x = 8
	w.Circle.Center.y = 9
}
//匿名
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}
w.X = 8
//初始化
w = Wheel{
    Circle:Circle{x:1,y:1},
    spokes:20
}
```

#### 方法

```go
func (u User) say(){
    fmt.Println(u.name)
}
func (_ User) say(){
    fmt.Println(u.name)
}
```

### JSON

json.Marshal将go对象转换成json数据格式

json.MarshalIndent格式化输出json，参数一：前缀字符串，参数二：缩进的字符串

json.Unmarshal将json数据格式转换成go对象

```go
type movie struct {
	Title string
    Year   int `json:"relased"`
    //omitempty，如果为零值或者为空，则不输出
    Color bool `json:"color,omitempty"`
}
var mov = []movie{
    {Title:"haha",Year:2000,Color:false}
}
//转换成JSON，json.Marshal
json.Marshal(mov)
```

# 函数

#### 定义

函数包含函数***名字***、一个***形参列表***，可选的返回列表及函数体

形参:指定了一组变量的参数名和参数类型

​		 形参是函数的***局部变量***

实参：***是按值传递的，每个函数接受的都是实参的副本，修改函数的形参变量不会影响到实参，除了引用类型(指针，		  slice、map、channel、)***

返回列表：指定了函数返回值的类型，没有返回值则省略

```go
func name(params)(result-list){
	body
}
//返回值
func add (x,y int)int{
	return x+y
}
func sub(x,y int)(z int){
	z = x+y
    return
}
```

##### 函数签名

两个函数拥有***相同的形参列表和返回列表***，认为这两个函数的类型或者签名是相同的，形参和返回值的名字不会影响函数类型

#### 递归

```go
func fibonacci(n int) int {
  if n < 2 {
   return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}
 
func main() {
    var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
```

#### 返回值

返回值，可以返回多个，使用返回值必须显示的赋值给变量

返回值的结果可以是调用另外函数的返回的结果

#### 函数变量

函数类型的***零值是nil***，调用一个空函数会导致出错

函数间***不可以进行比较***，***不能***相互比较或者作为***键值出现在map中***

strings.map对每个字符进行函数调用，并连接返回

```go
func add (item rune) rune{
	return item+1
}
strings.map(add,"HAL")//"IBM"
```

#### 匿名函数

匿名函数只能在包级别的作用域进行声明，函数字面量在使用的时候指定函数变量

go语言也有闭包

```go
func squares (){
	var x int
	return func()int{
		x++
		return x
	}
}
```

#### 可变参数

可变参数的函数，参数列表最后的类型名称之前用省略号声明可变参数

func( []int ) != func(...int)

```go
func sum (values ...int)int{
    total := 0
    for _,v := range values{
    	total+ = v
    }
    return total
}
//调用
//显示的声明一个数组，将实参复制给这个数组
values :=[]int{1,2,3}
sum(values...)
```

#### defer

延迟调用函数（函数返回值前，释放一些资源）

多个defer后注册的先执行

```go
func de() int {
	i := 9
	//这里是5，跟的是函数
	defer func() {
		fmt.Println(i)
	}()
    //这里是9
    defer func(i int) {
		fmt.Println(i)
	}(i)
	//这里是9，后面跟的是表达式，直接复制
	defer fmt.Println(i)
	return 5
}
```



# 方法

方法声明和普通函数的声明类似，在函数名称前多了一个参数，这个参数吧这个方法绑定到参数对应的类型上去。

在一个结构体内，声明一个叫做x方法会与字段x发生冲突，不允许的

```go
type Point struct{
	x,y int
}
func sum(p,q Point){
    z := p.x +q.x
}
//参数p成为方法的接收者
func(p Point) center(x int) int{
	return p.x + x
}
//方法调用
p.center(q)
```

指针传递变量的地址

```
r :=&Point{1,2}
func sum(r *Point){

}
```

# 接口

接口类型是对其他类型行为的概括与抽象，接口类型是抽象类型,只有方法

一个接口类型定义了一套方法，如果一个具体类型要实现该接口，必须实现接口上的所有方法

```go
type tran interface{
    move(int) string
}
```



#### 类型断言

```go
func types(){
	var i interface{}
    if v,ok :=i.(int),ok;ok{
        fmr.Println(v)
    }	
}
func swtype(){
    switch i.(type) {
    case int :
        //执行操作
    case int,byte://是interface{}类型
        //执行操作
    }
}
```

```go
func one(arr ...float64) (float64, error) {
	p := 1.0
	for _, v := range arr {
		p *= v
	}
	if p != 0 {
		return 1.0 / p, nil
	} else {
		return 0, errors.New("")
	}
}
```

# 控制流程

### if

```go
if 9>5{
	
}
//逻辑表达式成立，执行{}内容
//逻辑表达式不需要加（）
//{不能另起一行

if k,v :=2,3;v>k{
	
}
//初始化多个局部变量，分号后是逻辑判断
//逻辑表达式可以有变量或常量
//if语句仅包含一个分号，局部变量在if内有效
if-esle if
if c,d:=4,7;c>d{
		
}else if 5>10 {
		
}else if 5==5 {
		
}else {
		
}
```

### switch

switch和case后面可以跟常量、变量或者函数表达式，只要他们数据类型相同就可以

case后面可以跟多个值，只要一个值满足就可以

```go
color := "black"
	switch color {
	case "black":
		fmt.Println()
	case "red","black":
		fmt.Println("多个case，满足一个即可")
	default:
		fmt.Println()
	}
```

#### 空switch

switch后面直接跟{}

```go
switch{
	case add(5)>10:
		fmt.Println()
	default:
		fmt.Println()
	}
```

#### fallthrough

fallthrough,满足条件继续向下执行

```go
switch{
	case add(5)>10:
		fmt.Println()
    	fallthrough
	default:
		fmt.Println()
	}
```

#### switch-type

```go
func switch_type() {
	var num interface{} = 6.5
	switch value := num.(type) {
	case int:
		fmt.Printf("int类型-%d", value)
	case string:
		fmt.Printf("string类型-%s", value)
	}
}
```

### for

for初始化局部变量;条件表达式;后续操作

```go
for i:=0;i<100;i++
//死循环
for{}
//只有条件判断，前后的分号不要
```

#### range

for range拿到的是数组的拷贝

```go
//数组或者切片
for i,ele := rang s
//字符串
for i,ele := rang "哈哈"     	//rune类型
//遍历map
for i,ele := rang m			//不保证顺序
//遍历channel	，先close
for i,ele := rang ch
```

### break-continue

break-continue用于控制最外层最靠近自己的for循环

break:退出for循环，本轮break下面的代码不再执行

```go
func breakfn() {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		if i > 3 {
			fmt.Printf("%d", v)
		}
	}
}
```

continue：本轮continue代码不执行，直接进入下层循环

```go
func continuefn() {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		if i > 3 {
			continue
			//直接进入下一层循环
		}
		fmt.Printf("%d", v)
	}
}
fun
```

### label

简单的label

```go
func basic(){
	var a = 3
	me:
		a+=3
		goto me
}
```

```go
//当匹配到l1时候会继续执行下面的label
func mul() {
	i := 4
	if i%2 == 0 {
		goto l1
	} else {
		goto l2
	}
l1:
	fmt.Printf("这是l1\n")
	i += 3
	fmt.Printf("i的值%d\n", i)
l2:
	i *= 7
	fmt.Printf("这是l2\n")
	fmt.Printf("i的值%d\n", i)
}
```

# 反射

### 定义

***运行期间***探知对象的类型信息和内存结构、更新变量、调用他们的方法

函数的参数类型是interface{},需要在运行时进行判断针对不同类型采用不同的处理方式没比如，json.Marshal(v interface{})

在运行时候根据某些动态条件决定调用那些函数，根据配置文件执行相应的算子函数

### 弊端

代码难以维护

编译期间不能发现类型错误，测试覆盖念难度大

性能差，比代码慢一到二数量级

### typeof

```go
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
```

### valueof

```go
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

}
```

### 寻址

#### 普通变量

```go
va1 := reflect.ValueOf(1)
	var x int
	v2 := reflect.ValueOf(x)
	v3 := reflect.ValueOf(&x)
	v4 := v3.Elem()

	//寻址,valueof转换再用Elem()寻址
	fmt.Println(va1.CanAddr(), v2.CanAddr(), v3.CanAddr(), v4.CanAddr())
```

#### 切片

切片寻址,切片不可寻址，某个元素可以寻址

```go
//切片寻址,切片不可寻址，某个元素可以寻址
	s1 := make([]int, 3, 5)
	s3 := reflect.ValueOf(s1)
	s2 := s3.Index(0)
	fmt.Println(s3.CanAddr(), s2.CanAddr())
```

#### map

```go
mp :=make(map[int]bool,5)
v :=reflect.valueof(mp)
v.CanAddr()	//false
```

#### 设置值

```go
	i := 10
	ivalue := reflect.ValueOf(&i)
	if ivalue.Elem().CanAddr() {
		ivalue.Elem().SetInt(222)
	}
	fmt.Println(i)
```

#### 设置切片

```go
func changeslice() {
	type user struct {
		name int
	}
	sl := make([]*user, 3, 5)
	sl[0] = &user{
		name: 2,
	}
	value := reflect.ValueOf(sl)
	if value.Len() > 0 {
		value.Index(0).Elem().FieldByName("name").SetInt(90)
	}
}
```

#### 设置map

```go
mp := make(map[int]bool, 10)
	mp[0] = false
	mvalue := reflect.ValueOf(mp)
	mapType := mvalue.Type()
	keyType := mapType.Key()    //获取key的类型
	valueType := mapType.Elem() //获取value的类型
	//key，value
	fmt.Println(mapType, keyType, valueType)
	mvalue.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf(false))
```

#### 调函数

##### 普通函数

```go
func callfn() {
	//获取value
	valuefn := reflect.ValueOf(sub)
	//获取类型
	typefn := valuefn.Type()
	//获取参数个数
	argnum := typefn.NumIn()
	args := make([]reflect.Value, argnum)
	for i := 0; i < argnum; i++ {
		if typefn.In(i).Kind() == reflect.Int {
			args[i] = reflect.ValueOf(i)
		}
	}
	//输入切片，返回切片结果，调用函数
	results := valuefn.Call(args)[0]
	if typefn.Kind() == reflect.Int {
		i := results.Interface().(int)
		fmt.Println(i)
	}
}
```

##### 方法

```go
func callmethod() {
	u := &user{
		name: "name",
	}
	uValue := reflect.ValueOf(u)
	//不能获得未导出的方法
	sayMethod := uValue.MethodByName("say")
	fmt.Println(sayMethod)
	//没参数，传空切片
	sayMethod.Call([]reflect.Value{})
}
```

#### new()

```go
func newMethod() {
	u := reflect.TypeOf(user{})
	//转换成value
	value := reflect.New(u)
	value.Elem().FieldByName("name").SetString("哈哈哈")
	user := value.Interface().(*user)
	fmt.Println(user)
}
```

#### 创建slice

```go
func newSlice() {
	var s []int
	stype := reflect.TypeOf(s)
	//通过反射创建
	sliceValue := reflect.MakeSlice(stype, 3, 5)
	sliceValue.Index(0).Set(reflect.ValueOf(1))
	users := sliceValue.Interface().([]int)
	fmt.Println(users)
}
```

# 面向对象

#### 单例模式

```go
func newUser() *user {
	return &user{
		name: 1,
		age:  2,
	}
}

var (
	u        *user
    //不能公用，需要建立多个
	useronce sync.Once
)

func getInstance() *user {
	useronce.Do(func() {
		if u == nil {
			u = &user{}
		}
	})
	return u
}
```

#### 继承和重写

通过嵌入匿名结构体，变相实现"继承的功能"，因为访问匿名成员时可以跳过可以跳过成员直接访问它的内部成员

```go
type Plane struct {
	name string
	transport
}
type transport struct {
}
```

```go
type over struct {
}
type over2 struct {
}

func (o over) say() {

}
//重写
func (o over2) say() {

}
```

#### 组合

go语言不支持继承，支持组合

```go
type Plane struct {
	name string
	transport
}
type transport struct {
}
```

# 标准库

### 数学计算

### 时间函数

#### 时间格式化

```go
	time1 := "2022-10-10 10:10:10"
	now := time.Now()
	ts := now.Format(time1)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(time1, ts, loc)
	fmt.Println(t)
```

#### 时间计算

```

```

#### ticker-timer

```go
func tk() {
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-tk.C
	}
	tk.Stop()
}
func timer() {
	ti := time.NewTimer(1 * time.Second)
	<-ti.C
	ti.Stop()
}
```

### I/O操作

#### 控制台写入

```go
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(s1, s2)
```

#### 打开文件

```go
//打开文件,传入文件路径，通常来读
func os.Open(name string)(*os.File,error)
//写文件
fout,err :=os.OpenFile
```

#### 读文件

```go
func readFile() {
	if fin, err := os.Open("go.mod"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err == nil {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			} else {
				//读到末尾
				if err == io.EOF {
					if len(line) > 0 {
						fmt.Println(line)
					}
					break
				} else {
					fmt.Println(err)
					break
				}
			}
		}
	}
}
```

#### 写文件

```go
func writeFile() {
	if fo, err := os.OpenFile("a.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fo.Close()
		writer := bufio.NewWriter(fo)
		writer.WriteString("我是写文件")
		//强制清空缓存，把缓存写入磁盘
		writer.Flush()
	}
}
```

#### 遍历目录

```go
func rangeDir(path string) error {
	if subFiles, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, file := range subFiles {
			fmt.Println(file.Name())
			if file.IsDir() {
				if err := rangeDir(filepath.Join(path, file.Name())); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
```

### 编码

# 并发

进程内的多个线程共享系统资源，进程的创建和销毁切换比线程大很多

从进程到线程再到协程，不断减少切换成本

### Goroutine 

父协程结束后，子协程并不会结束

main协程结束后，所有协程都会结束

##### 创建

```go
//有名函数
func add(a, b int) {
	v := a + b
	fmt.Println(v)
}
go add(2, 3)
//匿名函数
go func(a, b int) {
		fmt.Println(a + b)
	}(2, 2)
```

##### sync

```go
//sync包
	wg := sync.WaitGroup{}
	//计数器：10个协程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(a, b int) {
			defer wg.Done()
			//do something
		}(i, i+1)
	}
	//等待为0
	wg.Wait()
```

sync.once

只加载一次

```go
var resource map[int]int
var loadResource sync.Once

func loadRe() {
	loadResource.Do(func() {
		resource[1] = 1
	})
}
```

##### panic

发生：主动调用panic，程序错误

panic执行什么：1.逆序执行当前的groutine的defer链（recover从这里介入）

​						   2.打印错误信息和调用堆栈

​						   3.调用exit(2)结束这个进程

### 并发安全

##### 资源竞争

多协程修改同一块内存，产生资源竞争

n++不是原子操作，多协程执行时会脏写

```go
	wg := sync.WaitGroup{}
	wg.Add(1000)
	n := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			n++
			//atomic.AddInt32(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Println(n)
//期望：1000
//实际：960左右

//方式一：原子操作
atomic.AddInt32(&n, 1)
//方式二：加锁
var lock sync.RWMutex
lock.Lock()
n++
lock.Unlock()
```

```go
//原子操作
func atomic.AddInt32(addr *int32,delta int32)(new int32)
func atomic.LoadInt32(addr *int32)(val int32)
```

```go
//读写锁
var lock sync.RWMutex
//写锁
lock.Lock()，lock.UnLock()
//读锁
lock.RLock() lock.RUnlock()
//任意时刻，只可以加一把写锁，且不能加读锁
//同时上读锁
```



### 多路复用

### 协程泄漏

### 协程管理
