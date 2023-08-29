package tool

import (
	"github.com/satori/go.uuid"
)

func NewUid() string {
	uid := uuid.NewV4()
	return uid.String()
}
