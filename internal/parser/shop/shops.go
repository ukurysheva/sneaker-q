package shop

import (
	"reflect"

	sneakerq "github.com/ukurysheva/sneaker-q"
)

var ShopsModelsParsers map[string]interface{} = map[string]interface{}{
	"nike": nikeParseModels,
}

func Call(funcName string, params ...interface{}) (result []*sneakerq.Model, err error) {
	f := reflect.ValueOf(ShopsModelsParsers[funcName])
	// if len(params) != f.Type().NumIn() {
	// 	// err = errors.New("The number of params is out of index.")
	// 	return
	// }

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface().([]*sneakerq.Model)
	// result := res.Interface().(float64)
	return
}

// func loop(slice []interface{}) {
// 	for _, elem := range slice {
// 		switch elemTyped := elem.(type) {
// 		case int:
// 			fmt.Println("int:", elemTyped)
// 		case string:
// 			fmt.Println("string:", elemTyped)
// 		case []string:
// 			fmt.Println("[]string:", elemTyped)
// 		case interface{}:
// 			fmt.Println("map:", elemTyped)
// 		}
// 	}
// }
