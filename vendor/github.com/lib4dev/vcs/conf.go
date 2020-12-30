package vcs

import "github.com/lib4dev/vcs/modules/const/conf"

type Option func()

//SetConfig 设置配置
func SetConfig(options ...Option) {
	for _, f := range options {
		f()
	}
}

func WithSmsConfig(c *conf.ImgCodeConf) Option {
	return func() {
		conf.ImgCodeSetting = c
	}
}

func WithImgConfig(c *conf.SmsCodeConf) Option {
	return func() {
		conf.SmsCodeSetting = c
	}
}

func WithCacheConfig(cacheName, httpName string) Option {
	return func() {
		conf.Config(cacheName, httpName)
	}
}

func WithSmsSendUrl(url string) Option {
	return func() {
		conf.SmsCodeSetting.SmsCodeSendRequestURL = url
	}
}

func WithSmsCodeCache(cacheTimeout, errlimit, errlimitTimeout int) Option {
	return func() {
		conf.SmsCodeSetting.SmsCodeCacheTimeout = cacheTimeout
		conf.SmsCodeSetting.SmsCodeErrorLimit = errlimit
		conf.SmsCodeSetting.SmsCodeErrorLimitTimeout = errlimitTimeout
	}
}

func WithImgcodeCache(timeout, errlimit int) Option {
	return func() {
		conf.ImgCodeSetting.ImgCodeCacheTimeout = timeout
		conf.ImgCodeSetting.ImgCodeErrorLimit = errlimit
	}
}
