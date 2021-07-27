package main

import (
	"fmt"
	"reflect"
)

// 关于反射，有人说过：清晰优于聪明，必要的时候再用反射
// 今天我们通过反射来实现一个通用的方法，针对不同的对象，生成不同的insert into sql语句，对象类型名称与mysql的表对应

type order struct {
	orderId    int
	customerid int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func create_insert_sql(obj interface{}) string {
	kind := reflect.ValueOf(obj).Kind()
	if kind == reflect.Struct {
		name := reflect.TypeOf(obj).Name()
		query := fmt.Sprintf("insert into %s values(", name)
		v := reflect.ValueOf(obj)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			switch field.Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, field.Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, field.Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, field.String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, field.String())
				}
			default:
				return "unsupported type."
			}
		}
		return fmt.Sprintf("%s%s", query, ")")
	} else {
		return "unsupported type."
	}
}

func main() {
	ord := order{
		orderId:    234,
		customerid: 11111111111,
	}
	empl := employee{
		name:    "chenzuoli",
		id:      0,
		address: "Haidian district",
		salary:  90000,
		country: "China",
	}
	fmt.Println(create_insert_sql(ord))
	fmt.Println(create_insert_sql(empl))
}
