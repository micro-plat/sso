package smscode

import (
	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
)

type Code struct {
	cfg   *conf.SmsCodeConf
	cache ICodeCache
}

//NewCode
func NewCode() (*Code, error) {

	cfg := conf.SmsCodeSetting
	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	return &Code{
		cfg:   cfg,
		cache: NewCodeCache(),
	}, nil

}

func (s *Code) Send(info *SendRequest, platName string) (result types.XMap, err error) {

	//1.验证参数
	if err = info.Valid(); err != nil {
		return
	}

	//2.获取缓存数据操作对象
	sdkCache, err := components.Def.Cache().GetCache(conf.CacheName)
	if err != nil {
		return nil, err
	}

	//3.请求发送短信
	r, err := s.SendRequest(info)
	if err != nil {
		return
	}

	//4 保存验证码到缓存
	err = s.cache.Save(sdkCache, platName, info.Ident, info.PhoneNo, info.Keywords)
	if err != nil {
		return
	}

	//5 返回值
	return types.XMap{
		"record_id": r.RecordID,
	}, nil
}

func (s *Code) Validate(ident, phone, code, platName string) (err error) {

	//1.获取缓存数据操作对象
	sdkCache, err := components.Def.Cache().GetCache(conf.CacheName)
	if err != nil {
		return err
	}

	//2 校验错误次数
	errCount, err := s.cache.CheckErrorLimit(sdkCache, platName, ident, phone)
	if err != nil {
		return
	}

	//3. 校验验证码
	err = s.cache.Verify(sdkCache, platName, ident, phone, code, errCount)
	if err != nil {
		return
	}

	return
}
