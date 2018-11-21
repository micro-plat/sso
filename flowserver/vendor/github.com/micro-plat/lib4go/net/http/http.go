package http

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"time"

	"github.com/micro-plat/lib4go/encoding"
)

type OptionConf struct {
	ConnectionTimeout time.Duration
	RequestTimeout    time.Duration
}

//Option 配置选项
type Option func(*OptionConf)

//WithConnTimeout 设置请求超时时长
func WithConnTimeout(tm time.Duration) Option {
	return func(o *OptionConf) {
		o.ConnectionTimeout = tm
	}
}

//WithRequestTimeout 设置请求超时时长
func WithRequestTimeout(tm time.Duration) Option {
	return func(o *OptionConf) {
		o.RequestTimeout = tm
	}
}

// HTTPClient HTTP客户端
type HTTPClient struct {
	*OptionConf
	client   *http.Client
	Response *http.Response
}

// HTTPClientRequest  http请求
type HTTPClientRequest struct {
	headers  map[string]string
	client   *http.Client
	method   string
	url      string
	params   string
	encoding string
}

// NewHTTPClientCert 根据pem证书初始化httpClient
func NewHTTPClientCert(certFile string, keyFile string, caFile string) (client *HTTPClient, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}
	caData, err := ioutil.ReadFile(caFile)
	if err != nil {
		return
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)
	ssl := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	ssl.Rand = rand.Reader
	client = &HTTPClient{}
	client.client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   ssl,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 0)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   0,
			ResponseHeaderTimeout: 0,
		},
	}
	return
}

// NewHTTPClientCert2 根据ca证书来初始化httpClient
func NewHTTPClientCert2(caFile string) (client *HTTPClient, err error) {
	pool := x509.NewCertPool()
	client = &HTTPClient{}
	caData, err := ioutil.ReadFile(caFile)
	if err != nil {
		return
	}
	pool.AppendCertsFromPEM(caData)

	client.client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{RootCAs: pool},
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 0)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   0,
			ResponseHeaderTimeout: 0,
			DisableCompression:    true,
		},
	}
	return
}

// NewHTTPClientCert1 根据ca证书来初始化httpClient
func NewHTTPClientCert1(caFile string) (client *HTTPClient, err error) {

	pool := x509.NewCertPool()
	caData, err := ioutil.ReadFile(caFile)
	if err != nil {
		return
	}
	pool.AppendCertsFromPEM(caData)
	client = &HTTPClient{}
	client.client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{RootCAs: pool},
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 0)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   0,
			ResponseHeaderTimeout: 0,
			DisableCompression:    true,
		},
	}

	return
}

// NewHTTPClient 构建HTTP客户端，用于发送GET POST等请求
func NewHTTPClient(opts ...Option) (client *HTTPClient) {
	client = &HTTPClient{}
	client.OptionConf = &OptionConf{ConnectionTimeout: time.Second * 3, RequestTimeout: time.Second * 20}
	for _, opt := range opts {
		opt(client.OptionConf)
	}
	client.client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, client.OptionConf.ConnectionTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(client.OptionConf.RequestTimeout))
				return c, nil
			},

			MaxIdleConnsPerHost:   0,
			ResponseHeaderTimeout: 0,
		},
	}
	return
}

// NewHTTPClientProxy 根据代理服务器地址创建httpClient
func NewHTTPClientProxy(proxy string) (client *HTTPClient) {
	client = &HTTPClient{}
	client.client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(proxy) //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
			},
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 0)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   0,
			ResponseHeaderTimeout: 0,
		},
	}
	return
}

// Download 发送http请求, method:http请求方法包括:get,post,delete,put等 url: 请求的HTTP地址,不包括参数,params:请求参数,
// header,http请求头多个用/n分隔,每个键值之前用=号连接
func (c *HTTPClient) Download(method string, url string, params string, header map[string]string) (body []byte, status int, err error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(params))
	if err != nil {
		return
	}
	req.Close = true
	for i, v := range header {
		req.Header.Set(i, v)
	}
	c.Response, err = c.client.Do(req)
	if c.Response != nil {
		defer c.Response.Body.Close()
	}
	if err != nil {
		return
	}
	status = c.Response.StatusCode
	body, err = ioutil.ReadAll(c.Response.Body)
	return
}

// Save 发送http请求, method:http请求方法包括:get,post,delete,put等 url: 请求的HTTP地址,不包括参数,params:请求参数,
// header,http请求头多个用/n分隔,每个键值之前用=号连接
func (c *HTTPClient) Save(method string, url string, params string, header map[string]string, path string) (status int, err error) {
	body, status, err := c.Download(method, url, params, header)
	if err != nil {
		return
	}
	fl, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return
	}
	defer fl.Close()
	n, err := fl.Write(body)
	if err == nil && n < len(body) {
		err = io.ErrShortWrite
	}
	return
}

// Request 发送http请求, method:http请求方法包括:get,post,delete,put等 url: 请求的HTTP地址,不包括参数,params:请求参数,
// header,http请求头多个用/n分隔,每个键值之前用=号连接
func (c *HTTPClient) Request(method string, url string, params string, charset string, header map[string]string) (content string, status int, err error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(params))
	if err != nil {
		return
	}
	req.Close = true
	for i, v := range header {
		req.Header.Set(i, v)
	}
	c.Response, err = c.client.Do(req)
	if c.Response != nil {
		defer c.Response.Body.Close()
	}
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		return
	}
	status = c.Response.StatusCode
	ct, err := encoding.DecodeBytes(body, charset)
	content = string(ct)
	return
}

// Get http get请求
func (c *HTTPClient) Get(url string, args ...string) (content string, status int, err error) {
	charset := getEncoding(args...)
	c.Response, err = c.client.Get(url)
	if c.Response != nil {
		defer c.Response.Body.Close()
	}
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		return
	}
	status = c.Response.StatusCode
	ct, err := encoding.DecodeBytes(body, charset)
	content = string(ct)
	return
}

// Post http Post请求
func (c *HTTPClient) Post(url string, params string, args ...string) (content string, status int, err error) {
	charset := getEncoding(args...)
	c.Response, err = c.client.Post(url, fmt.Sprintf("application/x-www-form-urlencoded;charset=%s", charset), encoding.GetEncodeReader([]byte(params), charset))
	if c.Response != nil {
		defer c.Response.Body.Close()
	}
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		return
	}
	status = c.Response.StatusCode
	rcontent, err := encoding.DecodeBytes(body, charset)
	content = string(rcontent)
	return
}

//Upload 文件上传
func (c *HTTPClient) Upload(url string, params map[string]string, files map[string]string, args ...string) (content string, status int, err error) {
	charset := getEncoding(args...)
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	//字段处理
	for k, v := range params {
		err = bodyWriter.WriteField(k, v)
		if err != nil {
			return "", 0, fmt.Errorf("设置字段失败:%s(%v)", k, v)
		}
	}

	//文件流处理
	for k, v := range files {
		fw1, err := bodyWriter.CreateFormFile(k, v)
		if err != nil {
			return "", 0, fmt.Errorf("无法创建文件流:%v", v)
		}
		f1, err := os.Open(v)
		if err != nil {
			return "", 0, fmt.Errorf("无法读取文件:%s", v)
		}
		defer f1.Close()
		io.Copy(fw1, f1)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	//发送POST请求
	c.Response, err = c.client.Post(url, contentType, encoding.GetEncodeReader(bodyBuffer.Bytes(), charset))
	if err != nil {
		return
	}
	defer c.Response.Body.Close()

	//处理响应包
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		return
	}
	status = c.Response.StatusCode
	rcontent, err := encoding.DecodeBytes(body, charset)
	content = string(rcontent)
	return
}

func getEncoding(params ...string) (encoding string) {
	if len(params) > 0 {
		encoding = strings.ToUpper(params[0])
		return
	}
	return "UTF-8"
}
