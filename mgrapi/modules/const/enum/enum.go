package enum

const UserDefaultPassword = `123456`

const (
	Normal   = 0
	Locked   = 1
	Disabled = 2
	Unlock   = 11
)

const (
	UserNormal int = iota
	UserLock
	UserDisable
)
