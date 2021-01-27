package data

import (
	"fmt"
	"strings"
)

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {

		return []byte(f), nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

var _bindata = map[string]string{
	"defaultData.sql": defaultData,
}
