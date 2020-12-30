package imgcode

import (
	"fmt"
	"io"

	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/lib4dev/vcs/modules/util"
	"github.com/micro-plat/hydra/components"
)

//Code
type Code struct {
	cfg   *conf.ImgCodeConf
	cache ICodeCache
}

//NewImgcode
func NewCode() (*Code, error) {

	cfg := conf.ImgCodeSetting
	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	return &Code{
		cfg:   cfg,
		cache: NewCodeCache(),
	}, nil
}

//Get 生成图形验证码
func (s *Code) Get(w io.Writer, ident, account, platName string) (err error) {

	//1.获取图形验证码
	codeByte, validCode := util.GetImgCode()

	//2.设置返回流
	_, err = util.NewImage(codeByte, 100, 40).WriteTo(w)
	if err != nil {
		return fmt.Errorf("设置图形验证码到返回的字节流中出错,err:%+v", err)
	}

	//3.获取缓存数据操作对象
	sdkCache, err := components.Def.Cache().GetCache(conf.CacheName)
	if err != nil {
		return err
	}

	//4.保存到缓存
	err = s.cache.Save(sdkCache, platName, ident, account, validCode)
	if err != nil {
		return err
	}

	//5.账号错误次数清除
	return s.cache.ResetErrLimit(sdkCache, platName, ident, account)

}

//Verify 校验图形验证码
func (s *Code) Verify(ident, account, code, platName string) (err error) {

	//1.获取缓存数据操作对象
	sdkCache, err := components.Def.Cache().GetCache(conf.CacheName)
	if err != nil {
		return
	}

	//2.保存code到缓存中
	return s.cache.Verify(sdkCache, platName, ident, account, code)

}
