package filecfg

import (
	"encoding/json"

	"github.com/gopyai/go-err"
	"gopkg.in/yaml.v2"
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

// YamlMarshal is a helper for Configurator's Marshal() implementation
func YamlMarshal(v interface{}) []byte {
	b, e := yaml.Marshal(v)
	err.Panic(e)
	return b
}

// YamlUnmarshal is a helper for Configurator's Unmarshal() implementation
func YamlUnmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}
