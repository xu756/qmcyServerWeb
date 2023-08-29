package xjwt

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	jwt := NewJwt(Jwt{
		SignKey: "u@n}8atwChc#+h@o1b.+RA-Q+w",
		Expire:  86400,
	})
	newtoken, err := jwt.NewJwt(1, []int64{1, 3}, "kkkk")
	if err != nil {
		log.Print("生成token错误", err)
		return
	}
	token, err := jwt.parseTokenString(newtoken)
	if err != nil {
		log.Print("解析token错误", err)
		return
	}
	log.Print(token)

}
