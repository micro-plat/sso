package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/const/sql"
)

//GetProvince 获取第一级省市
func GetProvince(dbe db.IDBExecuter, parentCode string) ([]types.XMap, error) {
	data, err := dbe.Query(sql.GetProvince, map[string]interface{}{
		"parent_code": parentCode,
	})
	if err != nil {
		return nil, fmt.Errorf("获取省信息出错:err:%+v", err)
	}
	return data, nil
}

//GetCitys 根据省获取市信息
func GetCitys(dbe db.IDBExecuter, parentCode, grade string) ([]types.XMap, error) {
	data, err := dbe.Query(sql.GetCityByProvinceID, map[string]interface{}{
		"parent_code": parentCode,
		"grade":       grade,
	})
	if err != nil {
		return nil, fmt.Errorf("获取市信息出错:err:%+v", err)
	}
	return data, nil
}

//GetAll 根据省获取市信息
func GetAll(dbe db.IDBExecuter) ([]types.XMap, error) {
	data, err := dbe.Query(sql.GetAll, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取市信息出错: err:%+v", err)
	}
	return data, nil
}

//Get 获取某个类型下面的字典信息
func Get(dbe db.IDBExecuter, dicType string) ([]types.XMap, error) {
	data, err := dbe.Query(sql.Get, map[string]interface{}{
		"type": dicType,
	})
	if err != nil {
		return nil, fmt.Errorf("获取某个类型下面的字典信息出错:  err:%+v", err)
	}
	return data, nil
}
