// +build !jsoniter

package json

import (
	"encoding/json"
)

var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
)
