package sso

import "github.com/micro-plat/lib4go/types"

//GetPermissionConfigRes 获取数据配置的返回结构
type GetPermissionConfigRes struct {
	ID          string `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	RuleConfigs string `form:"rules" json:"rules"` //规则配置json串
}

//对应关系
// field_name => f
// compare_symbol => c
// field_type => t
// conlink_symbol => s
// value => v

//PermissionConfig 规则配置结构
type PermissionConfig struct {
	ID            int64  `form:"id" json:"id"`
	FieldName     string `form:"f" json:"f"` //字段名称
	FieldType     string `form:"t" json:"t"` //字段类型
	CompareSymbol string `form:"c" json:"c"` //比较符("<,>,=,>=,<=,<>,in")
	ConlinkSymbol string `form:"s" json:"s"` //条件链接符(and, or)
	Value         string `form:"v" json:"v"` //值(比较的值)
}

type opts map[string]interface{}

//PermissionOption 数据权限配置函数
type PermissionOption func(opts)

//WithAlias 表别名
func WithAlias(alias string) PermissionOption {
	return func(o opts) {
		o["alias"] = alias
	}
}

//WithOperateAction 包含操作动作
func WithOperateAction(operateAction string) PermissionOption {
	return func(o opts) {
		o["operate_action"] = operateAction
	}
}

//WithCustomParams 包含自定义参数(@userid, @role　这两个除外)
func WithCustomParams(params map[string]interface{}) PermissionOption {
	return func(o opts) {
		o["with_custom_params"] = params
	}
}

//getAlias 获取表别名
func getAlias(options ...PermissionOption) string {
	args := getOption(options...)
	value, flag := args["alias"]
	if flag {
		return types.GetString(value)
	}
	return ""
}

//getOperateAction 获取操作动作
func getOperateAction(options ...PermissionOption) string {
	args := getOption(options...)
	value, flag := args["operate_action"]
	if flag {
		return types.GetString(value)
	}
	return ""
}

//getParams 获取外部传入的自定义参数
func getParams(options ...PermissionOption) map[string]interface{} {
	args := getOption(options...)
	value, flag := args["with_custom_params"]
	if flag {
		return value.(map[string]interface{})
	}
	return nil
}

//getOption 获取配置信息
func getOption(options ...PermissionOption) map[string]interface{} {
	args := make(map[string]interface{})
	for _, opt := range options {
		opt(args)
	}
	return args
}
