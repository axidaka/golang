package chapters

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	Authors string
	Publisher string
	IsPublished bool
	Price float64
}

var g_bresult []byte

func Json_Marshall() {

	gobook := Book{
		"Go语言编程",
		"XuShiwei HughLv Pandaman GuaguaSong HanTuo BertYuan XuDaoli",
		"ituring.com.cn",
		true,
		9.99}

    var err error
	g_bresult, err = json.Marshal(gobook)
	if err != nil {
		fmt.Println("json marshall fail", err)
	}

	fmt.Println("struct book[" , gobook, "] json::marshall result:", g_bresult)
}

func Json_Unmarshall() {

	var test Book;
	err := json.Unmarshal(g_bresult, &test)
	if err != nil {
		fmt.Println("json unmarshall fail", err)
	}

	fmt.Println("byte[", g_bresult, "] json:unmarshall reuslt:", test)

	var r interface{}
	err = json.Unmarshal(g_bresult, &r)
	if err != nil {
		fmt.Println("json unmarshall to interface fail", err)
	}

	fmt.Println("interface:", r)
}