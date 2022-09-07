package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("kind: ", t.Kind(), v.Kind())
	if v.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}
	tableName := t.Name() // 通过结构体类型提取出SQL的表名
	sql := fmt.Sprintf("INSERT iNTO %s ", tableName)
	columns := "("
	values := "VALUES ("
	fmt.Println("NumField: ", t.NumField(), v.NumField())
	for i := 0; i < v.NumField(); i++ {
		// reflect.Value 也实现了 NumField、Kind这些方法
		// 这里的v.Field(i).Kind() 等价于 t.Field(i).Type().Kind()
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)
				values += fmt.Sprintf("%d", v.Field(i).Int())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", %d", v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)
				values += fmt.Sprintf("'%s'", v.Field(i).String())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", '%s'", v.Field(i).String())
			}
		}
	}
	columns += "); "
	values += "); "
	sql += columns + values
	fmt.Println(sql)
	return sql
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
}
