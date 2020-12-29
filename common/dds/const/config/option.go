package config

type opts map[string]interface{}

//Option 配置函数
type Option func(opts)

//WithPrex url增加前缀
func WithPrex(prex string) Option {
	return func(o opts) {
		o["prex"] = prex
	}
}
