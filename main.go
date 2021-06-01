package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
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
	//faker.SetStringLang(faker.LangCHI)
	//fmt.Println(faker.Email())
	//fmt.Println(faker.Phonenumber())
	////faker.SetPoolGroup("aa", data)
	//fmt.Println(faker.IPv4())
	//fmt.Println(faker.IPv6())
	//fmt.Println(faker.Person{})
	//
	//a := SomeStruct{}
	//_ = faker.SetRandomStringLength(5)
	//_ = faker.FakeData(&a)
	//fmt.Printf("%+v", a)

	//for i := 0; i <= 10; i++ {
	//	fmt.Println(faker.Username())
	//	fmt.Println(faker.Cou)
	//	fmt.Println(faker.FreeEmail())
	//	fmt.Println(faker.AddressFull())
	//	fmt.Println(faker.Email())
	//	fmt.Println(faker.PhoneNumber())
	//	fmt.Println(faker.URL())
	//	fmt.Println(faker.LangName())
	//	fmt.Println("---------------------")
	//}
	//fmt.Println(faker.FakeData(&SomeStruct{}))
	//fmt.Println(faker.Phonenumber())
	//fmt.Println(faker.GetAddress())
	//fmt.Println(faker.Word())
	//fmt.Println(faker.Phonenumber())

	// 生成数据 数组
	// 随机获取
	// 格式化数据
	fmt.Println(gofakeit.FileExtension())
	fmt.Println(gofakeit.FileMimeType())
	fmt.Println(gofakeit.Username())
	fmt.Println(gofakeit.Gender())
	fmt.Println(gofakeit.SSN())
	fmt.Println(gofakeit.Address())
	fmt.Println(gofakeit.Country())
	fmt.Println(gofakeit.City())
	fmt.Println(gofakeit.CountryAbr())
	fmt.Println(gofakeit.StreetName())
	fmt.Println(gofakeit.Gamertag())
	fmt.Println(gofakeit.CarType())
	fmt.Println(gofakeit.Word())
	fmt.Println(gofakeit.Fruit())
	fmt.Println(gofakeit.Vegetable())
	fmt.Println(gofakeit.Dinner())
	fmt.Println(gofakeit.HTTPStatusCode())
	fmt.Println(gofakeit.HTTPMethod())
	fmt.Println(gofakeit.UserAgent())
	fmt.Println(gofakeit.HTTPStatusCodeSimple())
	fmt.Println(gofakeit.ChromeUserAgent())
	fmt.Println(gofakeit.Price(10.2, 20.3))
	fmt.Println(gofakeit.Company())
	fmt.Println(gofakeit.JobLevel())
	fmt.Println(gofakeit.JobTitle())
	fmt.Println(gofakeit.JobDescriptor())
	fmt.Println(gofakeit.BS())
	fmt.Println(gofakeit.AppName())
	fmt.Println(gofakeit.AppVersion())
	fmt.Println(gofakeit.Dog())
	fmt.Println(gofakeit.Emoji())
	fmt.Println(gofakeit.Language())
	fmt.Println(gofakeit.Date())

}
