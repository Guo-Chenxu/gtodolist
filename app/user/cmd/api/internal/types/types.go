// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"user_name"`
	CreateAt int64  `json:"create_at"`
}

type Data struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type LoginReq struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResp struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}
