package image

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/utility"
)

var imageExts = []string{".jpg", ".jpeg", ".gif", ".png"}

//ImageHandler 处理组件
type ImageHandler struct {
	container component.IContainer
	localDir  string
	baseURL   string
}

//NewImageHandler  创建处理组件
func NewImageHandler(dir string, url string) func(container component.IContainer) (c *ImageHandler) {
	return func(container component.IContainer) (c *ImageHandler) {
		return &ImageHandler{
			container: container,
			localDir:  dir,
			baseURL:   url,
		}
	}
}
func isImage(f string) bool {
	for _, i := range imageExts {
		if f == i {
			return true
		}
	}
	return false
}

//Handle 处理文件上传
func (ch *ImageHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--上传图片--")
	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("filename"); err != nil {
		return err
	}
	f, err := ctx.Request.Http.Get()
	if err != nil {
		return
	}
	ctx.Log.Info("2.检查图片格式")
	extName := filepath.Ext(ctx.Request.GetString("filename"))
	if !isImage(extName) {
		return fmt.Errorf("不是有效的图片格式：%v", imageExts)
	}
	uf, _, err := f.FormFile("file")
	if err != nil {
		err = fmt.Errorf("无法读取上传的文件:image(err:%v)", err)
		return
	}
	defer uf.Close()
	name := fmt.Sprintf("%s%s", utility.GetGUID(), extName)
	localPath := filepath.Join(ch.localDir, name)
	nf, err := os.Create(localPath)
	if err != nil {
		err = fmt.Errorf("保存文件失败:%s(err:%v)", localPath, err)
		return
	}
	defer nf.Close()
	_, err = io.Copy(nf, uf)
	if err != nil {
		return context.NewError(500, err)
	}
	return map[string]interface{}{
		"url": filepath.Join(localPath),
	}
}

