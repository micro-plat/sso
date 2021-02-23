package smscode

import (
	"fmt"
	"github.com/micro-plat/lib4go/types"

	"github.com/asaskevich/govalidator"
)

type SendRequest struct {
	ReqID        string     `json:"req_id" m2s:"req_id" valid:"required"`
	Ident        string     `json:"ident" m2s:"ident" valid:"required"`
	UserAccount  string     `json:"user_account" m2s:"user_account" valid:"required"`
	TemplateID   string     `json:"template_id" m2s:"template_id" valid:"required"`
	Keywords     string     `json:"keywords" m2s:"keywords" valid:"required"`
	DeliveryTime string     `json:"delivery_time,omitempty" m2s:"delivery_time"` //格式：yyyy-mm-dd hh:mm:ss
	ExtParams    types.XMap `json:"ext_params,omitempty" m2s:"ext_params,omitempty"  `
}

type SendResult struct {
	RecordID string `json:"record_id"`
}

func (c *SendRequest) Valid() error {
	if b, err := govalidator.ValidateStruct(c); !b {
		return fmt.Errorf("AppConf 配置有误:%v", err)
	}
	return nil
}
