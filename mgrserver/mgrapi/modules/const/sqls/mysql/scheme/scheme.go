package scheme

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
	"dds_area_info.sql":           dds_area_info,
	"dds_dictionary_info.sql":     dds_dictionary_info,
	"sso_data_permission.sql":     sso_data_permission,
	"sso_operate_log.sql":         sso_operate_log,
	"sso_role_datapermission.sql": sso_role_datapermission,
	"sso_role_info.sql":           sso_role_info,
	"sso_role_menu.sql":           sso_role_menu,
	"sso_system_info.sql":         sso_system_info,
	"sso_system_menu.sql":         sso_system_menu,
	"sso_user_info.sql":           sso_user_info,
	"sso_user_role.sql":           sso_user_role,
}
