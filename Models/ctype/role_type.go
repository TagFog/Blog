package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 //管理
	PermissionUser        Role = 2 //用户
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //被禁止的人
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.string())
}

func (r Role) string() string {
	var str string
	switch r {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁言的用户"
	default:
		str = "其他"
	}
	return str
}
