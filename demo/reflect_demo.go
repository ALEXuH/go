package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (person Person) Say() {
	fmt.Println("say hello")
}

func (person Person) Talk(v string) (string, int) {
	v = fmt.Sprintf("say %s", v)
	fmt.Println(v)
	return v, 300
}

// 反射
func main() {

	fl := 12.321
	var fl1 = &struct {
		name string
		age  int
	}{}
	fl1.age = 10
	fl1.name = "aaa"
	// TypeOf返回类型， valueOf返回值 kind返回数据底层类型 struct为ptr
	a := reflect.TypeOf(fl)
	b := reflect.ValueOf(fl)
	c := reflect.TypeOf(fl1)
	d := reflect.ValueOf(fl1)
	fmt.Println(c)
	fmt.Println(d, c.Kind(), d.Kind())
	fmt.Println(a, b, a.Kind(), b.Kind())
	fmt.Println(reflect.Int, reflect.Float64, reflect.Ptr)

	// 根据类型获取值 通过指针修改值
	fmt.Println(b.Interface().(float64), b.Float())
	//b.SetFloat(30.012) 值传递拷贝创建不可修改 必须传递指针TypeOf(&a)
	fmt.Println(b, b.CanSet())
	e := reflect.ValueOf(&fl)
	fmt.Println(e.CanSet())
	e = e.Elem()
	e.SetFloat(3213.32)
	fmt.Println(e, e.CanSet())

	// 结构体反射 修改值(字段明必须大写且反射指针并且使用Elem)
	p := Person{"aa", 20}
	rp := reflect.ValueOf(&p).Elem()
	fmt.Println(rp.CanSet())
	//rp.Field(0).SetString("ddd")
	fmt.Println(rp.FieldByName("name"), rp.FieldByName("sas"), rp.NumField())
	fmt.Println(rp.Field(1))
	for i := 0; i < rp.NumField(); i++ {
		fmt.Println(rp.Field(i))
	}
	rp.Field(0).SetString("dsd")
	fmt.Println(rp.Field(0))

	// 反射调用方法 首字母必须大写
	rp.MethodByName("Say").Call(nil)
	s := reflect.ValueOf(" zc.xu")
	s1 := []reflect.Value{s}
	rp.MethodByName("Say").Call(nil)
	result := rp.MethodByName("Talk").Call(s1)
	fmt.Println(rp.NumMethod(), result, result[0], result[1])
}
