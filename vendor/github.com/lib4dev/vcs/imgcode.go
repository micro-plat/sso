package vcs

import (
	"io"

	"github.com/lib4dev/vcs/modules/imgcode"
)

//GetImgCode 获取图形验证码
//w-->ctx.Request.Http.GetResponse()并且Header().Set("Content-Type", "image/png")
//ident-->系统标识,account-->账号
func GetImgCode(w io.Writer, ident, account, platName string) (err error) {

	obj, err := imgcode.NewCode()
	if err != nil {
		return err
	}
	return obj.Get(w, ident, account, platName)
}

//VerifyImgCode 验证图形验证码
//platName-->平台名,ident-->系统标识,account-->账号,code-->验证码
func VerifyImgCode(ident, account, code, platName string) (err error) {

	obj, err := imgcode.NewCode()
	if err != nil {
		return err
	}
	return obj.Verify(ident, account, code, platName)
}
