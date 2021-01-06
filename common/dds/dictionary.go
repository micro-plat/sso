package dds

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/dds/access/dictionary"
	"github.com/micro-plat/sso/common/dds/access/province"
)

var onceLock sync.Once

func init() {
	onceLock.Do(func() {
		app := hydra.S
		app.Micro("/dds/dictionary/get", getDictionaryHandler) //获取字典数据
		app.Micro("/dds/province/get", getProvinceHandler)     //获取省数据
		app.Micro("/dds/city/get", getCityHandler)             //获取市数据
		app.Micro("/dds/region/get", getRegionHandler)         //获取市数据
	})
}

//GetDictionaryHandler 获取某个类型下面的字典信息
func getDictionaryHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取字典信息---------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("dic_type"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("dic_type不能为空"))
	}

	ctx.Log().Info("2: 获取数据")
	data, err := dictionary.Get(ctx.Request().GetString("dic_type"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}

//getProvinceHandler 获取第一级省市
func getProvinceHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取第一级省市---------")
	parentCode := ctx.Request().GetString("parent_code")

	ctx.Log().Info("1: 获取数据")
	data, err := province.GetProvince(parentCode)
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 返回数据")
	return data
}

//getCityHandler 根据省获取市信息
func getCityHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取市信息---------")

	ctx.Log().Info("1: 验证参数")
	// if err := ctx.Request().Check("parent_code"); err != nil {
	// 	return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("parent_code不能为空"))
	// }

	ctx.Log().Info("2: 获取数据")
	data, err := province.GetCitys(ctx.Request().GetString("parent_code"), ctx.Request().GetString("grade"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}

//getRegionHandler 获取所有省市信息
func getRegionHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取所有省市信息---------")

	// ctx.Log().Info("1: 验证参数")
	// if err := ctx.Request().Check("parent_code"); err != nil {
	// 	return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("parent_code不能为空"))
	// }

	ctx.Log().Info("2: 获取数据")
	data, err := province.GetAll()
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}
