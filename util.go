package gin

import "encoding/json"

// ToJson ...
func ToJson(obj interface{}) string {

	data, _ := json.Marshal(obj)

	return string(data)
}
