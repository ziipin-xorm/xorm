package xorm

import "xorm.io/xorm/internal/json"

type JSONInterface = json.Interface

func SetDefaultJSONHandler(jsonHandler json.Interface) {
	json.DefaultJSONHandler = jsonHandler
}

func GetDefaultJSONHandler() json.Interface {
	return json.DefaultJSONHandler
}
