// Code generated by goctl. DO NOT EDIT.
package types

type GetGroupReq struct {
	Id       int64 `json:"id"`
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
}

type User struct {
	Id         int64  `json:"id"`
	State      int64  `json:"state"`
	Name       string `json:"name"`
	HeadImgUrl string `json:"headImgUrl"`
	Mobile     string `json:"mobile"`
	Deleted    int64  `json:"deleted"`
}

type Group struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Intro    string  `json:"intro"`
	Created  int64   `json:"created"`
	Creator  User    `json:"creator"`
	Edited   int64   `json:"edited"`
	Editor   User    `json:"editor"`
	Level    int64   `json:"level"` // 用户组层级
	Children []Group `json:"children"`
}

type GroupList struct {
	Total int64   `json:"total"`
	List  []Group `json:"list"`
}

type ContentReq struct {
	Id           int64  `json:"id"`
	ContentClass string `json:"content_class"`
}

type ContentsReq struct {
	ContentClass string `json:"content_class"`
	PageNum      int64  `json:"pageNum"`
	PageSize     int64  `json:"pageSize"`
}

type Content struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	ImgUrl       string `json:"img_url"`
	Path         string `json:"path"`
	Percent      int64  `json:"percent"`
	ContentClass string `json:"content_class"`
	ContentType  int64  `json:"content_type"`
	ContentText  string `json:"content_text"`
	ContentImg   string `json:"content_img"`
	Grade        int64  `json:"grade"`
	Created      int64  `json:"created"`
	Edited       int64  `json:"edited"`
	IsEdit       int64  `json:"is_edit"`
	Deleted      int64  `json:"deleted"`
}

type ContentList struct {
	Total int64     `json:"total"`
	List  []Content `json:"list"`
}

type Ok struct {
	Ok bool `json:"ok"`
}