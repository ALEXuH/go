package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
)

type SomeStruct struct {
	StringENG string `faker:"lang=eng"`
	StringCHI string `faker:"lang=chi"`
	StringRUS string `faker:"lang=rus"`
}

func main() {
	//var data = faker.PoolGroup{
	//	"name": {"Chinese"},
	//	"code": {"ca"},
	//}
	faker.SetStringLang(faker.LangCHI)
	fmt.Println(faker.Email())
	fmt.Println(faker.Phonenumber())
	//faker.SetPoolGroup("aa", data)
	fmt.Println(faker.IPv4())
	fmt.Println(faker.IPv6())
	fmt.Println(faker.Person{})

	a := SomeStruct{}
	_ = faker.SetRandomStringLength(5)
	_ = faker.FakeData(&a)
	fmt.Printf("%+v", a)

	//for i := 0; i <= 10; i++ {
	//	fmt.Println(faker.CountryName())
	//	fmt.Println(faker.Username())
	//	fmt.Println(faker.FreeEmail())
	//	fmt.Println(faker.AddressFull())
	//	fmt.Println(faker.Email())
	//	fmt.Println(faker.PhoneNumber())
	//	fmt.Println(faker.URL())
	//	fmt.Println(faker.LangName())
	//	fmt.Println("---------------------")
	//}
	fmt.Println(faker.FakeData(&SomeStruct{}))

	// 生成数据 数组
	// 随机获取
	// 格式化数据

}
