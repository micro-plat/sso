package queue

import (
	"errors"
	"fmt"
	"strings"
)

var Nil = errors.New("nil")

type IQueue interface {
	Push(key string, value string) error
	Pop(key string) (string, error)
	Count(key string) (int64, error)
	Close() error
}

//IQueueResover 定义配置文件转换方法
type IQueueResover interface {
	Resolve(address []string, conf string) (IQueue, error)
}

var queueResolvers = make(map[string]IQueueResover)

//Register 注册配置文件适配器
func Register(proto string, resolver IQueueResover) {
	if resolver == nil {
		panic("queue: Register adapter is nil")
	}
	if _, ok := queueResolvers[proto]; ok {
		panic("queue: Register called twice for adapter " + proto)
	}
	queueResolvers[proto] = resolver
}

//NewQueue 根据适配器名称及参数返回配置处理器
func NewQueue(address string, conf string) (IQueue, error) {
	proto, addrs, err := getNames(address)
	if err != nil {
		return nil, err
	}
	resolver, ok := queueResolvers[proto]
	if !ok {
		return nil, fmt.Errorf("queue: unknown adapter name %q (forgotten import?)", proto)
	}
	return resolver.Resolve(addrs, conf)
}

func getNames(address string) (proto string, raddr []string, err error) {
	addr := strings.Split(address, "://")
	if len(addr) > 2 {
		return "", nil, fmt.Errorf("MQ地址配置错误%s，格式:redis://192.168.0.1:61613", addr)
	}
	if len(addr[0]) == 0 {
		return "", nil, fmt.Errorf("MQ地址配置错误%s，格式:redis://192.168.0.1:61613", addr)
	}
	proto = addr[0]
	if len(addr) > 1 {
		raddr = strings.Split(addr[1], ",")
	} else {
		raddr = append(raddr, "")
	}
	return
}
