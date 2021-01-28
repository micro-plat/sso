package dds

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/services"
)

func init() {
	hydra.OnReady(func() {
		hydra.S.Micro("/dds/dictionary/get", services.GetDictionaryHandler) //获取字典数据
		hydra.S.Micro("/dds/province/get", services.GetProvinceHandler)     //获取省数据
		hydra.S.Micro("/dds/city/get", services.GetCityHandler)             //获取市数据
		hydra.S.Micro("/dds/region/get", services.GetRegionHandler)         //获取省市数据
	})
}
