package main

// () 因试分解关键词 适用于const、var 和 type 的声明或定义
import (
	"bytes"
	"fmt"
	fm "fmt" // 别名
	rd "math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// 定义已存在类型别名，拥有更多特性
type IT int

type (
	name   string
	age    int
	weight float32
)

type intArray []int

// 常量变量类型
const (
	a = "C"
	b = "d"
)

var v1 int = 10

type typea struct{}


// 特殊函数，首先执行
func init() {
	fm.Println("version1.0")
}

func main() {
	fm.Println("aa")
	fmt.Println("bb")

	fm.Println(firstFunction(10, 10.2))

	var a IT = 10
	fmt.Println(a)
	fmt.Println(b)

	// 简短形式赋值自动推断类型不需要显示指定类型 只能用于函数体内
	v2 := 10
	var v3 string = "10"
	var v4 = "122"
	var (
		v5 = "dsa"
		v6 = false
	)
	v6 = true

	fm.Println(v2, v3, v4, v5, v6)

	var v7 uint16 = 10

	// 指针地址 == != 比较必须类型一致
	fm.Println(&v2, v7, v4 == v3)

	var v8 int16 = 10
	var v10 int32 = 20

	fm.Println(v8, v10, int32(v8)+v10)
	fm.Println(rd.Float32())

	// 字符串操作
	stringOperate()

	// 指针操作
	pointerOperate()

	// 控制结构
	controlOperate()

	// 函数
	d := 10
	fmt.Println("d:%s address: %s", d, &d)
	doSomthing(&d)
	multi(10, 20, &d)
	fmt.Println("d:%s address: %s", d, &d)

	min, max := sort(10, 20)
	fmt.Println(min, max)

	lognParam(10, 20, "sre0", true, 12.43)

	deferTest()

	add1 := func(a int, b int) int {
		return a + b
	}
	fmt.Println(funcParam(10, 20, add1))

	adder := funcReturn(10, 20)
	fmt.Println(adder(10, 30))

	// 闭包调试函数
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line)
	}
	where()

	// 数组 slice map
	dataStruct()

	// 结构体接口
	structOperate()

	// 错误处理恢复
	errorOperate()

}

// 多个返回值函数判断是否执行成功
func firstFunction(a int16, b float32) (i int, error string) {
	return 0, "no error"
}

// 字符串操作和转换
func stringOperate() {
	a := "a b c de "
	fmt.Println(a[0:5])
	fmt.Println(strings.Split(a, " "))
	fmt.Println(strings.Count(a, " "))
	fmt.Println(strings.Contains(a, "de"))
	fmt.Println(strings.Index(a, "b"))
	fmt.Println(strings.ToLower(a), strings.ToUpper(a))
	fmt.Println(strings.ReplaceAll(a, "b", "aa"), strings.Repeat(a, 4))
	fmt.Println(strings.TrimLeft(a, "a"), "|", strings.TrimSpace(a))
	fmt.Println(strings.Join(strings.Split(a, " "), ","))

	// str to int | int to str
	i, err := strconv.Atoi("20")
	i1 := strconv.Itoa(20)

	if err == nil {
		fmt.Println("convert success")
	}

	// 解析错误 给类型默认值
	i2, err1 := strconv.ParseInt("10.102", 10, 64)
	i3, err2 := strconv.ParseFloat("10.102", 64)
	i4, err3 := strconv.ParseBool("truea")

	fmt.Println(i, err, i1, i2, i3, i4, err1, err2, err3)

	// T type to string
	fmt.Println(strconv.FormatInt(203, 10))

}

func timeOperate() {

}

// 指针（golang也是一种特殊数据类型 *type）就是地址，函数参数传递使用指针可以大大减少值传递带来的内存损耗
// &var 获取地址 *(&var) 获取该地址所存储的值 var v1 string 值;   var p *string= &v1 地址; *p 地址值， 函数传递直接传*p（引用传递）
func pointerOperate() {
	var i int = 10
	var intP *int
	intP = &i
	fmt.Println("memort address: %s value: %s", i, &i)
	fmt.Println("memort address: %s value: %s", intP, *intP)

	// 恒等于 反向引用
	fmt.Println(i == *(&i))
}

// if-else switch select for(range)
func controlOperate() {

	// if
	if runtime.GOOS == "windows" {
		fmt.Println("it is windows")
	} else {
		fmt.Println("it is unix")
	}

	a := 10
	c := -1
	if b := 15; (a <= b && a > 0) || c > 0 {
		fmt.Println("ok")
	} else if c < 0 {
		fmt.Println("fail")
	} else {
		fmt.Println("fail1")
	}

	// comma,ok 模式（pattern）
	if _, err := strconv.Atoi("das"); err != nil {
		fmt.Println(err.Error())
	}

	if _, ok := okPattern(); ok {
		fmt.Println("it is ok")
	}

	// switch
	switch a, b := 90, 91; {
	case a == 90 && b == 90:
		fmt.Println("90")
	case b == 91:
		fmt.Println("91")
	default:
		fmt.Println("default")
	}

	// for break continue
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i1 := 5
	for i1 > 0 {
		i1--
		fmt.Println("i1: %d", i1)
	}
	for {
		i1++
		if i1 == 3 {
			continue
		}
		if i1 >= 5 {
			break
		}
		fmt.Println("i11: %d", i1)

	}

	s := strings.Split("dsa sad das", " ")
	for ix, value := range s {
		fmt.Println(ix, value)
	}

}

// ok-pattern
func okPattern() (int, bool) {
	return 10, true
}

// 引用传递
func doSomthing(a *int) {
	b := a
	c := *a
	fmt.Println(b, c)
}

// 引用传递直接改变地址值
func multi(a, b int, reply *int) {
	*reply = a*b + *reply
}

// 带有形参的返回值 只需要赋值 最后return
func sort(a int, b int) (min int, max int) {
	if a > b {
		min = b
		max = a
	} else if b > a {
		max = b
		min = a
	} else {
		max = 0
		min = 0
	}
	return
}

// 传递长参数，scala里的语法糖
func lognParam(a int, who ...interface{}) {
	for a, value := range who {
		fmt.Println(a, value)
	}
}

// defer return之后执行 类似于java finally 代码块用于最后释放资源/调试时打印出参入参 多个defer存在时类似于栈后进先出
func deferTest() {
	fmt.Println("begin ...")

	defer func(a int) {
		fmt.Println("first ...")
		fmt.Println(a)
	}(10) // 最后括号为函数调用

	defer deferTest1(20)

	fmt.Println("finally ...")
}

func deferTest1(a int) {
	fmt.Println(a)
}

// 函数作为参数
func funcParam(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// 返回函数
func funcReturn(a int, b int) func(c int, d int) int {
	return func(c int, d int) int {
		return a + b + c + d
	}
}

// 数组 切片 map
func dataStruct() {

	// 数组在函数中为值传递 var identifier [len]type
	var arr [4]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var arr1 = [3]int{10, 20, 30}
	var arr2 = [3]string{0: "dsa", 2: "dsa"}
	fmt.Println(arr1, arr2, arr2[1], arr2[2], arr2[0:2])

	// 切片（slice 引用类型 ）是对数组一个连续片段的引用 长度可变 [] var identifier []type

	// 数组生成切片
	arr3 := &arr // 引用数组生成切片
	arr4 := arr[:]
	arr5 := arr[0:2]
	arr6 := []int{0, 1, 2}

	fmt.Println(*arr3, arr4, arr5, arr6)
	// make创建切片(python列表)
	slice := make([]string, 10, 20)
	slice1 := make([]interface{}, 20)
	slice1[0] = 10
	slice1[len(slice1)-2] = "dass"

	fmt.Println(slice, slice1)

	// buffer 类似于java的stringBuilder 减少内存cpu
	slice2 := make([]string, 10)
	slice2[0] = "das"
	slice2[2] = "dsad"
	var buffer bytes.Buffer
	for _, v := range slice2 {
		buffer.WriteString(v)
	}
	fmt.Println(slice2, buffer.String())

	// 追加复制
	copy(slice, slice2)
	slice2 = append(slice2, "aaaa")
	fmt.Println(slice2)

	var map1 map[string]string
	map1 = map[string]string{"one": "1"}

	var map2 = map[string]int{"one": 1, "two": 2, "three": 3}

	map3 := make(map[string]float32)
	map3["cc"] = 30.2

	map4 := make(map[string][]int, 100)
	map4["aa"] = []int{10, 20, 30}
	slice4 := make([]int, 10)
	slice4[0] = 10
	map4["bb"] = slice4

	//map1["aa"] = "bb"
	fmt.Println(len(map1), map2, map3, map3["cc"], map4, map4["bb"])

	k, v := map3["sdsa"]
	if v {
		fmt.Println("key exists ")
	} else {
		fmt.Println("key not exists ")
	}
	fmt.Println(k, v)
	for k, v := range map3 {
		fmt.Println(k, v)
	}

}

// 结构体
func structOperate() {

	// 值引用
	var person Person = Person{}
	person.Name = "a"

	person1 := Person{"aa", 10}
	person1.Age = 20

	// 指针引用 返回指针（&[底层也是使用new]/new） 编译时做了优化
	var person2 *Person = &Person{}
	person2.Name = "dsad"
	(*person2).Age = 20
	fmt.Println(person2, person2.Age)

	person3 := new(Person)
	person3.Name = "dsa"
	(*person3).Age = 30
	fmt.Println(person3, person3.Name)

	// 返回指针类型 *取值
	p := newPerson("aa", 30)
	fmt.Print(*p, p, (*p).Age, p.Age)

	p1 := &home{10, Person{"aa", 20}}
	fmt.Println(p1.number, p1.Age, p1.Person.Name)

	p2 := &home1{10, Person{"aa", 10}, Person1{"bb", 30}, 30.21, 40}
	fmt.Println(p2.Person.Name, p2.Person1.Name, p2.Age, p2.Age, p2.hight)

	// 结构体方法
	p1.settNumber(100)
	fmt.Println(p1.getterNumber())

	fmt.Println(intArray{10, 20, 30}.sum())

	// 方法多重继承
	p3 := &area{*p1, *p2}
	p3.resetHight()
	fmt.Println(p3, p3.hight)

	// 垃圾回收
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println(m.Alloc/1024, m.Frees/1024)
	runtime.GC()
	fmt.Println(m.Alloc/1024, m.Frees/1024)

	// 接口
	// 不能有字段必须实现所有方法才算实现接口，接口不能实列化但是可以指向实现改接口的实列
	fmt.Println("---------------interface{}--------")
	cat := &cat{}
	dog := &dog{}

	var animal animal
	animal = cat
	animal.Walk()
	animal = dog
	animal.Walk()

	cat.Walk()
	dog.Walk()
	fmt.Println(cat.Eat("aa"), dog.Eat("aa"))

	// 接口断言
	var a interface{}
	a = 10
	if v, ok := a.(int); ok {
		fmt.Println(" a is int type: ", v)
	}

	map1 := make(map[string]interface{})
	map1["a"] = "aaa"
	map1["b"] = 100
	map1["c"] = 10.323

	for _, v := range map1 {
		switch v.(type) {
		case string:
			fmt.Println("string type: ", v)
		case float32:
			fmt.Println("float type ", v)
		case int:
			fmt.Println("int type ", v)
		}
	}

	// 反射
	c := 10.321
	var d IT = 10
	fmt.Println(reflect.TypeOf(d), reflect.ValueOf(d), reflect.Kind(d), reflect.TypeOf(c))

}

// 结构体 pojo
type Person struct {
	Name string
	Age  int
}

// 结构体 pojo
type Person1 struct {
	Name  string
	hight int
}

// person别名
type pr Person

// 工厂方法返回实列指针
func newPerson(name string, age int) *Person {
	if age > 0 {
		return &Person{name, age}
	}
	return nil
}

// 继承（使用匿名字段只给出类型，通过obj.type 获取匿名字段数据） 重载（相同名字相同类型） 字段外层会覆盖内层名字,

type home struct {
	number int
	Person
}

type home1 struct {
	number int
	Person
	Person1
	Age   float64
	hight int
}

// 结构体方法（类方法）
func (ho *home) settNumber(value int) {
	ho.number = value
}

func (ho *home) getterNumber() int {
	return ho.number
}

func (arr intArray) sum() (s int) {
	for _, value := range arr {
		s += value
	}
	return s
}

// 方法多重继承
func (ho *home1) resetHight() {
	ho.hight = 10
}

type area struct {
	home
	home1
}

// 接口实现 （隐式实现 多态）

type animal interface {
	Walk()
	Eat(food string) string
}

type cat struct{}

type dog struct{}

func (cat *cat) Walk() {
	fmt.Println(" cat walk")
}

func (dog *dog) Walk() {
	fmt.Println(" dog walk ")
}

func (cat *cat) Eat(food string) string {
	return "cat " + food
}

func (dog *dog) Eat(food string) string {
	return "dog " + food
}

// 接口嵌套
type word interface {
	area() string
	animal
}

// 错误处理 defer-panic-and-recover
func errorOperate() {

}
