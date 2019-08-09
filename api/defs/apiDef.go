package defs

// 用户登陆
type UserCredential struct {
	Username string `json: "user_name"`
	Pwd string `json: "pwd"`
}