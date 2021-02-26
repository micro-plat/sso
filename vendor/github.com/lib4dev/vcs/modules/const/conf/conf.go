package conf

//CacheName 缓存配置名称配置名称
var CacheName = "cache"

//RemoteName http/rpc请求客户端配置名称
var RemoteName = "http"

func Config(cache, remote string) {
	CacheName = cache
	RemoteName = remote
	return
}
