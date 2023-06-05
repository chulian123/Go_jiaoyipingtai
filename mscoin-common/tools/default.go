package tools

import (
	"errors"
	"reflect"
)

//总体上，该代码利用反射机制遍历结构体的字段，并根据字段类型设置默认值。
//在每个字段类型对应的函数中，创建一个具有默认值的变量，然后使用reflect.ValueOf()将其转换为反射值，
//并返回该反射值，以便在Default函数中设置结构体字段的默认值。

func Default(data any) error {
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)
	if typeOf.Kind() != reflect.Pointer {
		return errors.New("must be pointer")
	}
	//member
	ele := typeOf.Elem()
	valueEle := valueOf.Elem()
	for i := 0; i < ele.NumField(); i++ {
		field := ele.Field(i)
		value := valueEle.Field(i)
		//field.Tag.Get("default")
		kind := field.Type.Kind()
		if kind == reflect.Int {
			//根据设置的tag进行值的设置
			value.Set(defaultInt())
		}
		if kind == reflect.Int32 {
			value.Set(defaultInt32())
		}
		if kind == reflect.Int64 {
			value.Set(defaultInt64())
		}
		if kind == reflect.String {
			value.Set(defaultString())
		}
		if kind == reflect.Float64 {
			value.Set(defaultFloat64())
		}
		if kind == reflect.Float32 {
			value.Set(defaultFloat32())
		}
	}
	return nil
}

func defaultString() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

//bug 记录 这里返回值不能为负数 不然会导致gorm框架注册不能正常insert数据

func defaultInt() reflect.Value {
	var i int = 0
	return reflect.ValueOf(i)
}

func defaultInt32() reflect.Value {
	var i int32 = 0
	return reflect.ValueOf(i)
}
func defaultInt64() reflect.Value {
	var i int64 = 0
	return reflect.ValueOf(i)
}

func defaultFloat64() reflect.Value {
	var i float64 = 0
	return reflect.ValueOf(i)
}
func defaultFloat32() reflect.Value {
	var i float32 = 0
	return reflect.ValueOf(i)
}
