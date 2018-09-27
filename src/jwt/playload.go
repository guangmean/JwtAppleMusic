package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type JwtPlayload struct {
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

func (p JwtPlayload) Base64Content() (string, error) {

	playloadByte, err := json.Marshal(p)

	if nil != err {
		return "", err
	}

	return strings.TrimRight(base64.URLEncoding.EncodeToString(playloadByte), "="), nil
}
