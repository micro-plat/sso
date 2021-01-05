package province

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	sqls "github.com/micro-plat/sso/common/dds/const/sql/province"
)

//GetProvince 获取第一级省市
func GetProvince(parentCode ...string) ([]types.XMap, error) {
	dbe := components.Def.DB().GetRegularDB()
	code := "QG"
	if parentCode != nil && len(parentCode) != 0 && parentCode[0] != "" {
		code = parentCode[0]
	}
	data, err := dbe.Query(sqls.GetProvince, map[string]interface{}{
		"parent_code": code,
	})
	if err != nil {
		return nil, fmt.Errorf("获取省信息出错: err:%+v", err)
	}
	return data, nil
}

//GetCitys 根据省获取市信息
func GetCitys(parentCode, grade string) ([]types.XMap, error) {
	dbe := components.Def.DB().GetRegularDB()
	data, err := dbe.Query(sqls.GetCityByProvinceID, map[string]interface{}{
		"parent_code": parentCode,
		"grade":       grade,
	})
	if err != nil {
		return nil, fmt.Errorf("获取市信息出错: err:%+v", err)
	}
	return data, nil
}

//GetAll 根据省获取市信息
func GetAll() ([]types.XMap, error) {
	dbe := components.Def.DB().GetRegularDB()
	data, err := dbe.Query(sqls.GetAll, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取市信息出错: err:%+v", err)
	}
	return data, nil
}
