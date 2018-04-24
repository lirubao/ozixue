package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int64
	Name string
	Age  int
}

func (stu Student) Hello() {
	fmt.Println("Hello")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o) //获取接口类型
	fmt.Println("TypeName", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("FieldsName:")

	//获取所有字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)                                   //字段
		val := v.Field(i).Interface()                     //字段值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //字段名称,字段类型,字段值
	}

	//获取所有方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type) //方法名称,方法类型
	}
}
func main() {
	stu := Student{1, "OK", 18}
	Info(stu)
}
