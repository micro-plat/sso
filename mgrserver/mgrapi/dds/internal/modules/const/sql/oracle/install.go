package oracle

import "github.com/micro-plat/hydra"

func init() {
	hydra.OnReadying(func() {
		hydra.Installer.DB.AddSQL(dds_area_info, dds_dictionary_info)
	})

	hydra.OnReady(func() {
		hydra.Installer.DB.AddSQL(dds_area_info_data)
	})
}
