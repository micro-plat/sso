package model

/*DictionaryInfo 字典实体
 *ID 主键
 *Name 名称
 *Value 内容
 *Type 类型
 *SortNo 排序数字
 *Ident 系统标识
 *GroupCode 分组编号
 */
type DictionaryInfo struct {
	ID        int    `json:"id" m2s:"id"`
	Name      string `json:"name" m2s:"name" `
	Value     string `json:"value" m2s:"value"`
	Type      string `json:"type" m2s:"type"`
	Status    int    `json:"status" m2s:"status"`
	SortNo    int    `json:"sort_no" m2s:"sort_no"`
	GroupCode string `json:"group_code" m2s:"group_code"`
}
