package filecfg

import (
	"encoding/json"

	"github.com/gopyai/err"
)

// JsonMarshal is a helper for Configurator's Marshal() implementation
func JsonMarshal(v interface{}) []byte {
	b, e := json.MarshalIndent(v, "", "\t")
	err.Panic(e)
	return b
}

// JsonUnmarshal is a helper for Configurator's Unmarshal() implementation
func JsonUnmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
