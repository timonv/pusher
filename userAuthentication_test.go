package pusher

import (
	"encoding/json"
	"testing"
)

func TestPrivateChannelAuthentication(t *testing.T) {
	ua := &UserAuthentication{"278d425bdf160c739803", "7ad3773142a6692b25b8"}

	authExpected := "278d425bdf160c739803:58df8b0c36d6982b82c3ecf6b4662e34fe8c25bba48f5369f135bf843651c3a4"
	authInfo, _ := ua.Authenticate("private-foobar", "1234.1234", nil)

	if authExpected != authInfo.Auth {
		t.Errorf("Authenticate(): Expected %s, got %s", authExpected, authInfo.Auth)
	}

	if "" != authInfo.ChannelData {
		t.Errorf("Authenticate(): Expected %s, got %s", nil, authInfo.ChannelData)
	}
}

func TestPresenceChannelAuthentication(t *testing.T) {
	ua := &UserAuthentication{"278d425bdf160c739803", "7ad3773142a6692b25b8"}

	authExpected := "278d425bdf160c739803:afaed3695da2ffd16931f457e338e6c9f2921fa133ce7dac49f529792be6304c"
	channelDataExpected := "{\"user_id\":10,\"user_info\":{\"name\":\"Mr. Pusher\"}}"

	dataStr := channelDataExpected
	var data interface{}
	json.Unmarshal([]byte(dataStr), &data)

	authInfo, _ := ua.Authenticate("presence-foobar", "1234.1234", data)

	if authExpected != authInfo.Auth {
		t.Errorf("Authenticate(): Expected %s, got %s", authExpected, authInfo.Auth)
	}

	if channelDataExpected != authInfo.ChannelData {
		t.Errorf("Authenticate(): Expected %s, got %s", channelDataExpected, authInfo.ChannelData)
	}
}

func TestAuthInfoJsonWithData(t *testing.T) {
	authInfo := &AuthInfo{"auth", "{\"key\": \"value\"}"}

	jsonExpected := "{\"auth\":\"auth\",\"channel_data\":\"{\\\"key\\\": \\\"value\\\"}\"}"

	b, _ := json.Marshal(authInfo)
	res := string(b)

	if res != jsonExpected {
		t.Errorf("AuthInfo json: Expected %s, got %s", jsonExpected, res)
	}
}

func TestAuthInfoJsonWithoutData(t *testing.T) {
	authInfo := &AuthInfo{"auth", ""}

	jsonExpected := "{\"auth\":\"auth\"}"

	b, _ := json.Marshal(authInfo)
	res := string(b)

	if res != jsonExpected {
		t.Errorf("AuthInfo json: Expected %s, got %s", jsonExpected, res)
	}
}
