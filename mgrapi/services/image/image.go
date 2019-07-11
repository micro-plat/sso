package image

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrapi/modules/util"
)

//ImageHandler 处理组件
type ImageHandler struct {
	container component.IContainer
	localDir  string
}

//NewImageHandler  创建处理组件
func NewImageHandler(dir string) func(container component.IContainer) (c *ImageHandler) {
	return func(container component.IContainer) (c *ImageHandler) {
		return &ImageHandler{
			container: container,
			localDir:  dir,
		}
	}
}

//PostHandle 处理图片上传
func (ch *ImageHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--上传图片--")
	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("filename"); err != nil {
		return err
	}
	f, _ := ctx.Request.Http.Get()

	ctx.Log.Info("2.检查图片格式")
	extName := filepath.Ext(ctx.Request.GetString("filename"))
	if !util.IsImage(extName) {
		return fmt.Errorf("不是有效的图片格式：%v", util.ImageExts)
	}
	uf, _, err := f.FormFile("file")
	if err != nil {
		return fmt.Errorf("无法读取上传的文件:image(err:%v)", err)
	}
	defer uf.Close()

	name := fmt.Sprintf("%s%s", utility.GetGUID(), extName)
	localPath := filepath.Join(ch.localDir, name)
	nf, err := os.Create(localPath)

	if err != nil {
		return fmt.Errorf("保存文件失败:%s(err:%v)", localPath, err)
	}
	defer nf.Close()
	_, err = io.Copy(nf, uf)
	if err != nil {
		return context.NewError(500, err)
	}
	ctx.Log.Info("3.返回数据")
	return map[string]interface{}{
		"url": model.GetConf(ch.container).GetWebHostName() + "/static/img/" + name,
	}
}
