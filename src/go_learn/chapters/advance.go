package chapters

import (
	"fmt"
	"reflect"
)

func Reflect_Test() {

	// 获取类型信息
	var x float64 = 3.4
	fmt.Println("type ", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// 获取值信息
	p := reflect.ValueOf(&x)
	fmt.Println("type of p ", p.Type())
	fmt.Println("settability of p ", p.CanSet())

	y := p.Elem()
	fmt.Println("settability of v ", y.CanSet())

	y.SetFloat(7.1)
	fmt.Println(y.Interface())
	fmt.Println(x)

	// 对机构的反射操作
	type T struct {
		A int
		B string
	}

	t := T{111, "kkkk"}
	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n",
			i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}