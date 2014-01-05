package pusher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
)

type UserAuthentication struct {
	key, secret string
}

type AuthInfo struct {
	Auth        string `json:"auth"`
	ChannelData string `json:"channel_data,omitempty"`
}

func (ua *UserAuthentication) Authenticate(channelName string, socketID string, data interface{}) (*AuthInfo, error) {

	stringToSign := socketID + ":" + channelName
	jsonStr := ""

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		jsonStr = string(b)
		stringToSign = stringToSign + ":" + jsonStr
	}

	hash := hmac.New(sha256.New, []byte(ua.secret))
	io.WriteString(hash, stringToSign)
	hexDigest := hex.EncodeToString(hash.Sum(nil))
	return &AuthInfo{ua.key + ":" + hexDigest, jsonStr}, nil
}
