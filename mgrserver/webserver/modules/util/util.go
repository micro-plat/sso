package util

var ImageExts = []string{".jpg", ".jpeg", ".gif", ".png"}

// IsImage 判断是否是图片
func IsImage(f string) bool {
	for _, i := range ImageExts {
		if f == i {
			return true
		}
	}
	return false
}
