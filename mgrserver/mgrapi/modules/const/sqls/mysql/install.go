package mysql

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls/mysql/data"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls/mysql/scheme"
)

func init() {
	hydra.OnReadying(func() {
		schemeNames := scheme.AssetNames()
		for i := range schemeNames {
			bytes, _ := scheme.Asset(schemeNames[i])
			hydra.Installer.DB.AddSQL(string(bytes))
		}
	})

	hydra.OnReady(func() {
		dataNames := data.AssetNames()
		for i := range dataNames {
			bytes, _ := data.Asset(dataNames[i])
			hydra.Installer.DB.AddSQL(string(bytes))
		}
	})
}
