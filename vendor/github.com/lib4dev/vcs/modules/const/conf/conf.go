package conf

//CacheName 缓存配置名称配置名称
var CacheName = "cache"

//HTTPName http请求客户端配置名称
var HTTPName = "http"

func Config(cache, http string) {
	CacheName = cache
	HTTPName = http
	return
}
