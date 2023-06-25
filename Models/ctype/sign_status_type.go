package ctype

import "encoding/json"

type SignStatus int

const (
	QQ    SignStatus = 1
	Gitee SignStatus = 2
	Email SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case QQ:
		return "QQ"
	case Gitee:
		return "Gitee"
	case Email:
		return "Email"
	default:
		str = "其他"
	}
	return str
}
