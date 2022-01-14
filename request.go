package oceanengine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func EncodeQuery(req map[string]interface{}) string {
	values := &url.Values{}
	for k, v := range req {
		switch value := v.(type) {
		case string:
			values.Set(k, value)
		case []byte:
			values.Set(k, string(value))
		case int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64:
			values.Set(k, fmt.Sprintf("%v", value))
		default:
			b, _ := json.Marshal(value)
			values.Set(k, string(b))
		}
	}

	return values.Encode()
}

func EncodeBody(req interface{}) []byte {
	r, _ := json.Marshal(req)
	return r
}
