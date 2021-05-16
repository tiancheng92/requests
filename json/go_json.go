//go:build go_json
// +build go_json

package json

import gojson "github.com/goccy/go-json"

var (
	Marshal       = gojson.Marshal
	Unmarshal     = gojson.Unmarshal
	MarshalIndent = gojson.MarshalIndent
	NewDecoder    = gojson.NewDecoder
	NewEncoder    = gojson.NewEncoder
)
