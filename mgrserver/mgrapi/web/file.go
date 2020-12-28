package web

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/micro-plat/hydra"
)

//Archive 归档文件
var Archive = "./static.tar.gz"

func init() {
	hydra.OnReady(func() error {
		//处理网址程序
		_, err := os.Stat(Archive)
		if err == nil {
			return nil
		}
		if !os.IsNotExist(err) {
			return err
		}
		for _, v := range AssetNames() {
			err := os.MkdirAll(filepath.Dir(v), 0777)
			if err != nil {
				return err
			}
			buff, err := Asset(v)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(v, buff, 0777)
			if err != nil {
				return err
			}
		}
		return nil
	})

}
