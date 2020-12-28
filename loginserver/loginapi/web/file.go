package web

import (
	"os"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/common/archive"
)

//Archive 归档文件
var Archive = "./static.tar.gz"

func init() {
	hydra.OnReady(func() error {
		return archive.OnReady(Archive, AssetNames, Asset)
	})
	hydra.G.AddCloser(func() error {
		return os.Remove(Archive)
	})
}
