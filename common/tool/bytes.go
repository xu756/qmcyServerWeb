package tool

import "encoding/json"

func ToBytes(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}
