package shop

import (
	sneakerq "github.com/ukurysheva/sneaker-q"
)

func ParseMenu(shopInfo sneakerq.Shop) []*MenuItem {

	var ParseMenuCall map[string][]*MenuItem = map[string][]*MenuItem{
		"nike": NikeParseMenu(shopInfo),
	}
	return ParseMenuCall[shopInfo.ClassName]
}

func ParseModels(shopInfo sneakerq.Shop, menus []*MenuItem) []*sneakerq.Model {

	var ParseModelsCall map[string][]*sneakerq.Model = map[string][]*sneakerq.Model{
		"nike": NikeParseModels(shopInfo, menus),
	}

	// TODO : call func by reflect
	// funcName := shopInfo.ClassName + "ParseModels"
	// f := reflect.ValueOf(funcName)
	// params:=[]
	// in := make([]reflect.Value, len(params))
	// for k, param := range params {
	// 	in[k] = reflect.ValueOf(param)
	// }
	// var res []reflect.Value
	// res = f.Call(in)
	// // result = res[0].Interface().(*Parser)
	// // result = res[0].Interface().([]*sneakerq.Model)
	// result := res.Interface().(float64)
	// return

	return ParseModelsCall[shopInfo.ClassName]
}

type MenuItem struct {
	Title string
	Link  string
}
