package enum

// MaxFailCnt 失败次数
const MaxFailCnt = 5

const UserDefaultPassword = `1qaz2wsx`

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

const (
	SystemDisable = 0
	SystemNormal  = 1
)

const (
	Normal   = 0
	Locked   = 1
	Disabled = 2
	Unlock   = 11
)
