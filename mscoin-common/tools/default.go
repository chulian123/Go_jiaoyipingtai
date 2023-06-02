package tools

import (
	"errors"
	"reflect"
)

//总体上，该代码利用反射机制遍历结构体的字段，并根据字段类型设置默认值。
//在每个字段类型对应的函数中，创建一个具有默认值的变量，然后使用reflect.ValueOf()将其转换为反射值，
//并返回该反射值，以便在Default函数中设置结构体字段的默认值。

func Default(data any) error {
	// 获取传入数据的类型和值
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)

	// 检查传入的数据是否为指针类型
	if typeOf.Kind() != reflect.Ptr {
		return errors.New("must be pointer") // 必须是指针类型
	}

	// 获取指针指向的元素类型
	ele := typeOf.Elem()
	valueEle := valueOf.Elem()

	// 遍历结构体字段
	for i := 0; i < ele.NumField(); i++ {
		field := ele.Field(i)
		// field.Tag.Get("defaulnt")
		value := valueEle.Field(i)
		kind := field.Type.Kind()

		// 根据字段类型设置默认值
		if kind == reflect.Int {
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

// 设置字符串类型字段的默认值
func defaultString() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

// 设置整型字段的默认值
func defaultInt() reflect.Value {
	var i int = -1
	return reflect.ValueOf(i)
}

// 设置Int32类型字段的默认值
func defaultInt32() reflect.Value {
	var i int32 = -1
	return reflect.ValueOf(i)
}

// 设置Int64类型字段的默认值
func defaultInt64() reflect.Value {
	var i int64 = -1
	return reflect.ValueOf(i)
}

// 设置Float64类型字段的默认值
func defaultFloat64() reflect.Value {
	var i float64 = -1
	return reflect.ValueOf(i)
}

// 设置Float32类型字段的默认值
func defaultFloat32() reflect.Value {
	var i float32 = -1
	return reflect.ValueOf(i)
}
