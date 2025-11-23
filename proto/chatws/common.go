package chatws

import (
	"github.com/tidwall/gjson"
)

func GetName(data []byte) string {
	result := gjson.GetBytes(data, "name")
	if result.Exists() {
		return result.String()
	}
	return ""
}

func GetEvent(data []byte) string {
	result := gjson.GetBytes(data, "event")
	if result.Exists() {
		return result.String()
	}
	return ""
}

func GetID(data []byte) int64 {
	result := gjson.GetBytes(data, "id")
	if result.Exists() {
		return result.Int()
	}
	return -1
}
