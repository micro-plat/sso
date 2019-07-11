package user

//UserKeyResp 通过key还回用户数据实体
type UserKeyResp struct {
	UserName string `json:"user_name"` //用户名
	UserId   int    `json:"user_id"`   //用户标识
}
