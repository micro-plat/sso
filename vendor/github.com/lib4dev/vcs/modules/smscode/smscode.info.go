package smscode

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type SendRequest struct {
	ReqID        string `json:"req_id" m2s:"req_id" valid:"required"`
	Ident        string `json:"ident" m2s:"ident" valid:"required"`
	PhoneNo      string `json:"phone_no" m2s:"phone_no" valid:"required"`
	TemplateID   string `json:"template_id" m2s:"template_id" valid:"required"`
	Keywords     string `json:"keywords" m2s:"keywords" valid:"required"`
	DeliveryTime string `json:"delivery_time,omitempty" m2s:"delivery_time"` //格式：yyyy-mm-dd hh:mm:ss
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
