package sso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
)

//DataPermission 数据权限
type DataPermission struct {
	conf *Config
}

// NewDataPermissionLogic xx
func newDataPermission(conf *Config) *DataPermission {
	return &DataPermission{
		conf: conf,
	}
}

func (d *DataPermission) getUserDataPermission(userID int64, tableName string, opt ...PermissionOption) (string, error) {
	configs, err := d.getConfigFromAPI(userID, tableName, getOperateAction(opt...))
	if err != nil {
		return "", err
	}
	if len(configs) == 0 {
		return " 1=1 ", nil
	}
	params := d.constructParmas(userID, opt...)
	tempSQL, err := d.GenerateSQL(configs, tableName, params, opt...)
	return d.replaceSpecial(userID, tempSQL)
}

//GenerateSQL 生成sql
func (d *DataPermission) GenerateSQL(configs []GetPermissionConfigRes, tableName string, params map[string]interface{}, opt ...PermissionOption) (string, error) {
	alias := d.getReallyTableName(tableName, opt...)
	var sqlsArray []string
	for _, item := range configs {
		configSQL, err := d.convertConfigToSQL(item.RuleConfigs, alias, params)
		if err != nil {
			return "", err
		}
		sqlsArray = append(sqlsArray, configSQL)
	}
	return fmt.Sprintf("(%s)", strings.Join(sqlsArray, ") and (")), nil
}

//convertConfigToSQL 将一组json转换成sql语句
func (d *DataPermission) convertConfigToSQL(configJSON, alias string, params map[string]interface{}) (string, error) {
	var rules []PermissionConfig
	err := json.Unmarshal([]byte(configJSON), &rules)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	for _, item := range rules {
		transformValue := transform.Translate(item.Value, params)
		valueT := transformValue
		if strings.EqualFold(item.FieldType, "string") {
			valueT = fmt.Sprintf("'%s'", transformValue)
		}
		if strings.EqualFold(item.CompareSymbol, "in") || strings.EqualFold(item.CompareSymbol, "not in") {
			if strings.EqualFold(item.FieldType, "string") {
				valueT = fmt.Sprintf("('%s')", strings.Replace(strings.Replace(transformValue, " ", "", -1), ",", "','", -1))
			} else {
				valueT = fmt.Sprintf("(%s)", transformValue)
			}
		} else if strings.EqualFold(item.CompareSymbol, "like") {
			valueT = fmt.Sprintf("'%s%%'", transformValue)
		}

		buffer.WriteString(fmt.Sprintf(" %s %s.%s %s %s ", item.ConlinkSymbol, alias, item.FieldName, item.CompareSymbol, valueT))
	}
	return buffer.String(), nil
}

//replaceParams 替换sql中的参数信息
func (d *DataPermission) replaceSpecial(userID int64, tempSQL string) (string, error) {
	if strings.Contains(tempSQL, "@role") {
		userIds, err := newUser(d.conf).getRoleUsers(userID)
		if err != nil {
			return "", err
		}
		return transform.Translate(tempSQL, map[string]interface{}{
			"role": userIds,
		}), nil
	}
	return tempSQL, nil
}

//getConfigFromAPI 远程获取用户【数据权限】的配置信息
func (d *DataPermission) getConfigFromAPI(userID int64, tableName, operateAction string) (r []GetPermissionConfigRes, err error) {
	cfg := d.conf
	values := net.NewValues()
	values.Set("table_name", tableName)
	values.Set("operate_action", operateAction)
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))
	values.Set("user_id", types.GetString(userID))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	var res []GetPermissionConfigRes
	_, err = remoteRequest(cfg.host, permissionConfig, values.Join("=", "&"), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//GetReallyTableName 获取表真正名字
func (d *DataPermission) getReallyTableName(tableName string, opt ...PermissionOption) string {
	alias := getAlias(opt...)
	if strings.EqualFold(alias, "") {
		return tableName
	}
	return alias
}

//constructParmas 构造自定义参数
func (d *DataPermission) constructParmas(userID int64, opt ...PermissionOption) map[string]interface{} {
	params := types.NewXMapByMap(map[string]interface{}{
		"userid": userID,
	})
	businessParam := getParams(opt...)
	if businessParam != nil {
		params.MergeMap(businessParam)
	}
	return params.ToMap()
}
