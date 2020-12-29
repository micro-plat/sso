package archive

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type AssetNamesFunc func() []string

type AssetFunc func(string) ([]byte, error)

func OnReady(archive string, namesFunc AssetNamesFunc, assetFunc AssetFunc) (isAuto bool, err error) {
	//处理网址程序
	_, err = os.Stat(archive)
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		return
	}
	isAuto = true
	for _, v := range namesFunc() {
		err = os.MkdirAll(filepath.Dir(v), 0777)
		if err != nil {
			return
		}
		buff, err1 := assetFunc(v)
		if err1 != nil {
			err = err1
			return
		}
		err = ioutil.WriteFile(v, buff, 0777)
		if err != nil {
			return
		}
	}
	return
}
