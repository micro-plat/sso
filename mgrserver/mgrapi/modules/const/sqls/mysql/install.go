package mysql

import (
	"fmt"

	"github.com/micro-plat/hydra"
)

func init() {

	names := AssetNames()
	for i := range names {
		bytes, _ := Asset(names[i])
		fmt.Println(names[i], string(bytes))
		hydra.Installer.DB.AddSQL(string(bytes))
	}
	/*
		hydra.Installer.DB.AddSQL(
			dds_area_info,
			dds_dictionary_info,
			sso_data_permission,
			sso_operate_log,
			sso_role_datapermission,
			sso_role_info,
			sso_role_menu,
			sso_system_info,
			sso_system_menu,
			sso_user_info,
			sso_user_role,
		)
	*/
}
