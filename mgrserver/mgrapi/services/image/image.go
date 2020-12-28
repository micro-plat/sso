package image

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/util"
)

//ImageHandler 处理组件
type ImageHandler struct {
	localDir string
}

//NewImageHandler  创建处理组件
func NewImageHandler(dir string) func() (c *ImageHandler) {
	return func() (c *ImageHandler) {
		return &ImageHandler{
			localDir: dir,
		}
	}
}

//Handle 图片上传
func (ch *ImageHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--上传图片--")

	ctx.Log().Info("1.检查参数")
	if err := ctx.Request().Check("filename"); err != nil {
		return err
	}

	ctx.Log().Info("2.检查图片格式")
	filename, err := ctx.Request().GetFileName("filename")
	if err != nil {
		return err
	}
	extName := filepath.Ext(filename)
	if !util.IsImage(extName) {
		return fmt.Errorf("不是有效的图片格式：%v", util.ImageExts)
	}

	ctx.Log().Info("3.保存图片")
	url, err := ch.saveImg(extName, ctx)
	if err != nil {
		return err
	}
	return url
}

//saveImg 保存图片
func (ch *ImageHandler) saveImg(extName string, ctx hydra.IContext) (picURL string, err error) {
	uf, err := ctx.Request().GetFileBody("filename")
	if err != nil {
		return "", fmt.Errorf("无法读取上传的文件:image(err:%v)", err)
	}
	defer uf.Close()

	name := fmt.Sprintf("%s%s", components.Def.UUID().ToString(), extName)
	localPath := filepath.Join(ch.localDir, name)
	nf, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("保存文件失败:%s(err:%v)", localPath, err)
	}
	defer nf.Close()

	_, err = io.Copy(nf, uf)
	if err != nil {
		return "", errs.NewError(500, err)
	}

	return model.GetConf().GetWebHostName() + "/sso/" + name, nil
}
