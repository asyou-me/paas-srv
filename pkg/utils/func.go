package utils

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
)

// 创建字符串uuid
func StrUUID() string {
	uid := uuid.NewUUID()
	str := base64.RawURLEncoding.EncodeToString([]byte(uid))
	str = str[:len(str)-2]
	return str
}
