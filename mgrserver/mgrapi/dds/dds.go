package dds

import (
	"github.com/micro-plat/hydra"
	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/const/conf"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/db"
)

//GetEnums 获取枚举信息
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//params-- 查询省份type=province 查询城市type=city 查询地区type=region
func GetEnums(c interface{}, params types.XMap) (data []types.XMap, err error) {
	tp := params.GetString("dic_type")
	if tp == "province" {
		return GetProvince(c, params.GetString("parent_code"))
	}
	if tp == "city" {
		return GetCity(c, params.GetString("parent_code"))
	}
	if tp == "region" {
		return GetRegion(c)
	}
	return GetDictionary(c, tp)
}

//GetDictionary 获取字典信息
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//dicType:字典类型
func GetDictionary(c interface{}, dicType string) (data []types.XMap, err error) {
	xdb, err := getDB(c)
	if err != nil {
		return nil, err
	}
	return db.Get(xdb, dicType)
}

//GetProvince 获取第一级省市
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//parentCode:父级code 可为空
func GetProvince(c interface{}, parentCode string) (data []types.XMap, err error) {
	xdb, err := getDB(c)
	if err != nil {
		return nil, err
	}
	return db.GetProvince(xdb, parentCode)
}

//GetCity 获取市信息
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//parentCode:父级code 可为空
func GetCity(c interface{}, parentCode string) (data []types.XMap, err error) {
	xdb, err := getDB(c)
	if err != nil {
		return nil, err
	}
	return db.GetCitys(xdb, parentCode, "")
}

//GetRegion 获取所有省市信息
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
func GetRegion(c interface{}) (data []types.XMap, err error) {
	xdb, err := getDB(c)
	if err != nil {
		return nil, err
	}
	return db.GetAll(xdb)
}

//---------------------------------内部函数-----------------------------------
func getDB(c interface{}) (ldb.IDBExecuter, error) {
	switch v := c.(type) {
	case ldb.IDB:
		return v, nil
	case ldb.IDBTrans:
		return v, nil
	default:
		db, err := hydra.C.DB().GetDB(conf.DbName)
		return db, err
	}
}
