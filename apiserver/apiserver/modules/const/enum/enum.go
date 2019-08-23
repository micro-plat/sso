package enum

// MaxFailCnt 失败次数
const MaxFailCnt = 5

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

const (
	SystemDisable = 0
	SystemNormal  = 1
)