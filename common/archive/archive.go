package archive

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type AssetNamesFunc func() []string

type AssetFunc func(string) ([]byte, error)

func OnReady(archive string, namesFunc AssetNamesFunc, assetFunc AssetFunc) error {
	//处理网址程序
	_, err := os.Stat(archive)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	for _, v := range namesFunc() {
		err := os.MkdirAll(filepath.Dir(v), 0777)
		if err != nil {
			return err
		}
		buff, err := assetFunc(v)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(v, buff, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
