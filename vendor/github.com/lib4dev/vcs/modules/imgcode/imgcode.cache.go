package imgcode

import (
	"fmt"
	"strings"

	"github.com/lib4dev/vcs/modules/const/cachekey"
	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/lib4dev/vcs/modules/const/errorcode"
	"github.com/micro-plat/hydra/components/caches"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

type ICodeCache interface {
	Save(c caches.ICache, platName, ident, account, code string) (err error)
	Verify(c caches.ICache, platName, ident, account, code string) (err error)
	ResetErrLimit(c caches.ICache, platName, ident, account string) (err error)
}

type CodeCache struct {
}

func NewCodeCache() ICodeCache {
	return &CodeCache{}
}

func (s *CodeCache) Save(c caches.ICache, platName, ident, account, code string) (err error) {

	key := types.Translate(cachekey.ImageCodeCachekey, "platname", platName, "ident", ident, "account", account)
	err = c.Set(key, code, conf.ImgCodeSetting.ImgCodeCacheTimeout)

	if err != nil {
		err = fmt.Errorf("设置图形验证码到缓存中出错,err:%+v", err)
		return
	}
	return
}

func (s *CodeCache) ResetErrLimit(c caches.ICache, platName, ident, account string) (err error) {

	key := types.Translate(cachekey.ImageCodeErrorCountCachekey, "platname", platName, "ident", ident, "account", account)
	return c.Delete(key)
}

func (s *CodeCache) Verify(c caches.ICache, platName, ident, account, code string) (err error) {

	//1. 获取缓存
	key := types.Translate(cachekey.ImageCodeCachekey, "platname", platName, "ident", ident, "account", account)

	if !c.Exists(key) {
		return errs.NewError(errorcode.HTTPErrorKeyNotExistError, "图形验证码缓存不存在")
	}

	value, err := c.Get(key)
	if err != nil {
		return
	}

	//2 校验验证码
	errCountKey := types.Translate(cachekey.ImageCodeErrorCountCachekey, "platname", platName, "ident", ident, "account", account)

	if !strings.EqualFold(code, value) {
		curVal, err := c.Increment(errCountKey, 1)
		if err != nil {
			return err
		}

		c.Delay(errCountKey, conf.ImgCodeSetting.ImgCodeErrorLimitTimeout)
		//超过错误次数限制，清除图片验证码
		if curVal >= int64(conf.ImgCodeSetting.ImgCodeErrorLimit) {
			c.Delete(key)
			return errs.NewError(errorcode.HTTPErrorFailedIMGCodeErrorCountError, "图形验证码错误次数太多")
		}

		return errs.NewError(errorcode.HTTPErrorFailedIMGCodeCheckedError, fmt.Errorf("图形码验证失败"))
	}

	c.Delete(errCountKey)
	c.Delete(key)
	return nil
}
