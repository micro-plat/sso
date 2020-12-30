package middleware

import (
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ginCtx struct {
	*gin.Context
	once    sync.Once
	service string
}

func NewGinCtx(c *gin.Context) *ginCtx {
	return &ginCtx{Context: c}
}

func (g *ginCtx) load() {
	g.once.Do(func() {
		if g.Context.ContentType() == binding.MIMEMultipartPOSTForm {
			g.Context.Request.ParseMultipartForm(32 << 20)
		}
	})
}
func (g *ginCtx) GetParams() map[string]interface{} {
	params := make(map[string]interface{})
	for _, v := range g.Context.Params {
		params[v.Key] = v.Value
	}
	return params
}
func (g *ginCtx) GetRouterPath() string {

	return g.Context.FullPath()
}

func (g *ginCtx) GetService() string {
	return g.service
}
func (g *ginCtx) Service(service string) {
	g.service = service
}
func (g *ginCtx) GetBody() io.ReadCloser {
	g.load()
	return g.Request.Body
}
func (g *ginCtx) GetMethod() string {
	return g.Request.Method
}
func (g *ginCtx) GetURL() *url.URL {
	return g.Request.URL
}
func (g *ginCtx) GetHeaders() http.Header {
	return g.Request.Header
}

func (g *ginCtx) GetCookies() []*http.Cookie {
	return g.Request.Cookies()
}
func (g *ginCtx) Find(path string) bool {
	return true

}
func (g *ginCtx) Next() {
	g.Context.Next()

}
func (g *ginCtx) GetRawForm() map[string]interface{} {
	g.load()
	fm := make(map[string]interface{})
	for k, v := range g.Request.PostForm {
		fm[k] = v
	}
	return fm
}
func (g *ginCtx) GetPostForm() url.Values {
	g.load()
	return g.Request.PostForm
}
func (g *ginCtx) WStatus(s int) {
	g.Writer.WriteHeader(s)
}
func (g *ginCtx) Status() int {
	return g.Writer.Status()
}
func (g *ginCtx) Written() bool {
	return g.Writer.Written()
}

func (g *ginCtx) WHeaders() http.Header {
	return g.Writer.Header()
}

func (g *ginCtx) WHeader(k string) string {
	return g.Writer.Header().Get(k)
}

//GetUploadFile 获取上传文件
func (g *ginCtx) GetFile(fileKey string) (string, io.ReadCloser, int64, error) {
	g.load()
	header, err := g.FormFile(fileKey)
	if err != nil {
		return "", nil, 0, err
	}
	f, err := header.Open()
	if err != nil {
		return "", nil, 0, err
	}
	return header.Filename, f, header.Size, nil
}
