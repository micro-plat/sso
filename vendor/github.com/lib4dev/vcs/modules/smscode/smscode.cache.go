package smscode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lib4dev/vcs/modules/const/cachekey"
	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/lib4dev/vcs/modules/const/errorcode"
	"github.com/micro-plat/hydra/components/caches"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

type ICodeCache interface {
	Save(c caches.ICache, platName, ident, phone, keywords string) (err error)
	CheckErrorLimit(c caches.ICache, platName, ident, phone string) (errCount string, err error)
	Verify(c caches.ICache, platName, ident, phone, code, errCount string) (err error)
}

type CodeCache struct {
}

func NewCodeCache() ICodeCache {
	return &CodeCache{}
}

func (s *CodeCache) Save(c caches.ICache, platName, ident, phone, keywords string) (err error) {

	key := types.Translate(cachekey.SmsCodeCachekey, "platname", platName, "ident", ident, "phone", phone)
	err = c.Set(key, keywords, conf.SmsCodeSetting.SmsCodeCacheTimeout)
	if err != nil {
		return fmt.Errorf("设置%s的值%s保存验证码到redis中失败,err%+v", key, keywords, err)
	}

	key = types.Translate(cachekey.SmsCodeErrorCountCacheKey, "platname", platName, "ident", ident, "phone", phone)
	err = c.Set(key, "0", conf.SmsCodeSetting.SmsCodeCacheTimeout)
	if err != nil {
		return fmt.Errorf("设置%s的值%s保存验证码到redis中失败,err%+v", key, "0", err)
	}

	return
}

func (s *CodeCache) CheckErrorLimit(c caches.ICache, platName, ident, phone string) (errCount string, err error) {

	key := types.Translate(cachekey.SmsCodeErrorCountCacheKey, "platname", platName, "ident", ident, "phone", phone)

	if !c.Exists(key) {
		return "", errs.NewError(errorcode.HTTPErrorKeyNotExistError, "消息验证码错误次数缓存不存在")
	}
	value, err := c.Get(key)
	if err != nil {
		err = fmt.Errorf("获取%s的缓存失败,err%+v", key, err)
		return
	}

	//校验次数
	if types.GetInt(value, 0) >= conf.SmsCodeSetting.SmsCodeErrorLimit {
		err = errs.NewError(errorcode.HTTPErrorFailedCodeErrorCountError, "消息验证码错误次数太多")
		return
	}

	return value, nil
}

func (s *CodeCache) Verify(c caches.ICache, platName, ident, phone, code, errCount string) (err error) {

	//1.获取验证码
	key := types.Translate(cachekey.SmsCodeCachekey, "platname", platName, "ident", ident, "phone", phone)

	if !c.Exists(key) {
		return errs.NewError(errorcode.HTTPErrorKeyNotExistError, "消息验证码缓存不存在")
	}

	cacheCode, err := c.Get(key)
	if err != nil {
		err = fmt.Errorf("VerifySmsCode:获取缓存中的验证码失败,err%+v", err)
		return
	}

	//2.比较验证码
	errCountKey := types.Translate(cachekey.SmsCodeErrorCountCacheKey, "platname", platName, "ident", ident, "phone", phone)
	if !strings.EqualFold(code, cacheCode) {
		//2.1增加错误次数
		err1 := c.Set(errCountKey, strconv.Itoa(types.GetInt(errCount, 0)+1), conf.SmsCodeSetting.SmsCodeErrorLimitTimeout)
		if err1 != nil {
			err = fmt.Errorf("VerifySmsCode:设置%s的缓存出错%+v", errCountKey, err1)
			return
		}
		return errs.NewError(errorcode.HTTPErrorFailedCodeCheckedError, "消息验证码不正确")
	}

	//3 删除验证码和错误次数
	c.Delete(key)
	c.Delete(errCountKey)

	return
}
