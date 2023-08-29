package model

type sqlResult struct {
	Result []byte `db:"result"`
}

type SqlUser struct {
	Id         int64  `db:"id"`           // 用户ID
	State      int64  `db:"state"`        // 用户状态:0=正常,1=禁用
	Name       string `db:"name"`         // 姓名
	HeadImgUrl string `db:"head_img_url"` // 头像图片地址
	Mobile     string `db:"mobile"`       // 手机号码
	Deleted    int64  `db:"deleted"`      // 逻辑删除:0=未删除,1=已删除
}

type SqlGroup struct {
	Name     string     `db:"name"`    // 用户组名称
	Code     string     `db:"code"`    // 用户组CODE唯一代码
	Intro    string     `db:"intro"`   // 用户组介绍
	Created  int64      `db:"created"` // 创建时间
	Creator  SqlUser    `db:"creator"` // 创建人
	Edited   int64      `db:"edited"`  // 修改时间
	Editor   SqlUser    `db:"editor"`  // 修改人
	Level    int64      `db:"level"`   // 用户组层级
	Children []SqlGroup // 子级用户组
}
