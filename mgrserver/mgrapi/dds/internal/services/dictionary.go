package services

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/const/conf"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/db"
)

//GetDictionaryHandler 获取某个类型下面的字典信息
func GetDictionaryHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取字典信息---------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("dic_type"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("dic_type不能为空"))
	}

	dbe, err := hydra.C.DB().GetDB(conf.DbName)
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 获取数据")
	data, err := db.Get(dbe, ctx.Request().GetString("dic_type"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}

//GetProvinceHandler 获取第一级省市
func GetProvinceHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取第一级省市---------")

	dbe, err := hydra.C.DB().GetDB(conf.DbName)
	if err != nil {
		return err
	}

	parentCode := ctx.Request().GetString("parent_code")

	ctx.Log().Info("1: 获取数据")
	data, err := db.GetProvince(dbe, parentCode)
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 返回数据")
	return data
}

//GetCityHandler 根据省获取市信息
func GetCityHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取市信息---------")

	ctx.Log().Info("1:获取db")
	dbe, err := hydra.C.DB().GetDB(conf.DbName)
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 获取数据")
	data, err := db.GetCitys(dbe, ctx.Request().GetString("parent_code"), ctx.Request().GetString("grade"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}

//GetRegionHandler 获取所有省市信息
func GetRegionHandler(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取所有省市信息---------")

	ctx.Log().Info("1:获取db")
	dbe, err := hydra.C.DB().GetDB(conf.DbName)
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 获取数据")
	data, err := db.GetAll(dbe)
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 返回数据")
	return data
}
