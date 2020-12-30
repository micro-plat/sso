package web

import (
	"os"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/common/archive"
)

//Archive 归档文件
var Archive = "./mgr.static.tar.gz"

func init() {
	isAutoArchiveFile := false
	hydra.OnReady(func() (err error) {
		isAutoArchiveFile, err = archive.OnReady(Archive, AssetNames, Asset)
		return
	})
	hydra.G.AddCloser(func() error {
		if isAutoArchiveFile {
			return os.Remove(Archive)
		}
		return nil
	})
}
