package mysql

import (
	"github.com/micro-plat/hydra"
)

func init() {

	names := AssetNames()
	for i := range names {
		bytes, _ := Asset(names[i])
		hydra.Installer.DB.AddSQL(string(bytes))
	}
}
