// Code generated by goctl. DO NOT EDIT.
package types

type CreateReq struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Status  string `form:"status"`
}

type CreateResp struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type Task struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      int    `json:"view"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type ListReq struct {
	Limit string `form:"limit"`
	Start string `form:"start"`
}

type Data struct {
	Task  []Task `json:"item"`
	Total int    `json:"total"`
}

type ListResp struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type ShowReq struct {
	Id int64 `form:"id"`
}

type ShowResp struct {
	Status  int    `json:"status"`
	Data    Task   `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type UpdateReq struct {
	Id      string `form:"id"`
	Title   string `form:"title"`
	Content string `form:"content"`
	Status  string `form:"status"`
}

type UpdateResp struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type SearchReq struct {
	Info string `form:"info"`
}

type SearchResp struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type DeleteReq struct {
	Id string `form:"id"`
}

type DeleteResp struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}
