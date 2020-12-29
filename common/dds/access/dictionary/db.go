package dictionary

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	sqls "github.com/micro-plat/sso/common/dds/const/sql/dictionary"
)

//Get 获取某个类型下面的字典信息
func Get(dicType string) ([]types.XMap, error) {

	dbe := components.Def.DB().GetRegularDB()
	data, q, a, err := dbe.Query(sqls.Get, map[string]interface{}{
		"type": dicType,
	})
	if err != nil {
		return nil, fmt.Errorf("获取某个类型下面的字典信息出错: sql:%s, args:%+v, err:%+v", q, a, err)
	}
	return data, nil
}
