package vcs

import "github.com/lib4dev/vcs/modules/const/conf"

//Option Option
type Option func()

//Config 设置配置
func Config(options ...Option) {
	for _, f := range options {
		f()
	}
}

//WithImgConfig 图形验证码配置
func WithImgConfig(c *conf.ImgCodeConf) Option {
	return func() {
		conf.ImgCodeSetting = c
	}
}

//WithSmsConfig 消息配置
func WithSmsConfig(c *conf.SmsCodeConf) Option {
	return func() {
		conf.SmsCodeSetting = c
	}
}

//WithCacheConfig 缓存配置
func WithCacheConfig(cacheName, httpName string) Option {
	return func() {
		conf.Config(cacheName, httpName)
	}
}

//WithSmsSendURL 消息发送地址
func WithSmsSendURL(url string) Option {
	return func() {
		conf.SmsCodeSetting.SmsCodeSendRequestURL = url
	}
}

//WithSmsCodeCache 短信验证码缓存设置
func WithSmsCodeCache(cacheTimeout, errlimit, errlimitTimeout int) Option {
	return func() {
		conf.SmsCodeSetting.SmsCodeCacheTimeout = cacheTimeout
		conf.SmsCodeSetting.SmsCodeErrorLimit = errlimit
		conf.SmsCodeSetting.SmsCodeErrorLimitTimeout = errlimitTimeout
	}
}

//WithImgcodeCache 图形验证码缓存设置
func WithImgcodeCache(timeout, errlimit int) Option {
	return func() {
		conf.ImgCodeSetting.ImgCodeCacheTimeout = timeout
		conf.ImgCodeSetting.ImgCodeErrorLimit = errlimit
	}
}
