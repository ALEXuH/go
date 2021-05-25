package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// 标准库时间和json编码解码
func main() {
	//timeDemo()
	jsonDemo()
}

const (
	STAND = "2006-01-02 15:04:05" // yyyy-mm-dd HH:mm:ss
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now, now.Month(), now.String(), now.Unix())

	// 格式化输出
	t := now.Format(time.RFC3339)
	fmt.Println(t)

	// 格式化时间字符串
	timeStr := "2021-02-21 12:12:12"
	t1,err := time.Parse(STAND, timeStr)
	if err == nil {
		fmt.Println(t1, t1.Format(STAND))
	} else {
		fmt.Println(err)
	}
	// 格式化时间戳
	unixTime := now.Unix()
	t2 := time.Unix(unixTime, 0)
	fmt.Println(unixTime, t2, t2.Format(STAND))
}


func jsonDemo() {

	// 编码 结构字段要大写
	h := &human{"aa", 10, true}
	h1 := &human{"bb", 30, false}
	h2 := &human{"cc", 50, true}
	//h3 := human{"cc", 50, true}
	humans := []*human{h, h1, h2}
	bed := map[string]string{"aa":"dsad", "bb": "dsa", "cc": "dsa"}
	house := &House{
		Humans: humans,
		Color: "dsa",
		Tree: []string{"ds", "vv"},
		Bed: bed,
	}
	js, err := json.Marshal(house)
	js1, err := jsoniter.Marshal(house)
	fmt.Println(err)
	fmt.Println(house)
	fmt.Println("==========")
	fmt.Printf("%s", js)
	fmt.Printf("%s", js1)

	// 解码json数据 类型断言
	fmt.Println("---------")
	var data string = "{\"Humans\":[{\"Name\":\"aa\"},{\"Name\":\"bb\"},{\"Name\":\"cc\"}],\"Color\":\"dsa\",\"Tree\":[\"ds\",\"vv\"],\"Bed\":{\"aa\":\"dsad\",\"bb\":\"dsa\",\"cc\":\"dsa\"}}"
	var f interface{}
	err = json.Unmarshal([]byte(data), &f)
	m := f.(map[string]interface{})

	for k,v := range m {
		switch v.(type) {
		case string:
			fmt.Println("string type", k, v)
		case int:
			fmt.Println("int type ", k, v)
		case bool:
			fmt.Println("bool type ",k, v)
		case []*human:
			fmt.Println("human type", k, v)
		case []string:
			fmt.Println("[]string type", k ,v)
		case map[string]interface{}:
			fmt.Println("map type ", k, v)
		default:
			fmt.Println("none type", k, v)
		}
	}

	// 解码到结构体
	fmt.Println("----------------")
	var m1 House
	var m2 House
	err = json.Unmarshal([]byte(data), &m1)
	err = jsoniter.Unmarshal([]byte(data), &m1)
	fmt.Println("==========")
	fmt.Println(m1.Bed, m1.Color, m1.Humans, m1.Tree)
	fmt.Println(m2.Bed, m2.Color, m2.Humans, m2.Tree)
	// 遍历指针值
	for k,v := range m1.Humans {
		fmt.Println(k, *v)
	}

	for k,v := range m2.Humans {
		fmt.Println(k, *v)
	}
}


type human struct {
	Name string
	age int
	flag bool
}

type House struct {
	Humans []*human
	Color string
	Tree []string
	Bed map[string]string
}