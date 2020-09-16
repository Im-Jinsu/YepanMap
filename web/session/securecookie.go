package session

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/securecookie"
)

// CookieHandler :
var CookieHandler *securecookie.SecureCookie

// RedirectLogin : UserPortal에서 리다이렉트 시
var RedirectLogin = make(map[string]string)

// SaveKeyBlock : SIGUSR2 발생 시 Key 정보 저장
func SaveKeyBlock() {
	jsKeyInfo, _ := json.Marshal(KeyInfoMap)
	ioutil.WriteFile("keyinfo.key", jsKeyInfo, 0600)
}

// KeyInfoMap :
var KeyInfoMap = make(map[string][]byte)

// SetKeyBlock : SIGUSR2 발생 시 Key 불러오기
func SetKeyBlock() {
	fileBody, err := ioutil.ReadFile("keyinfo.key")
	if err != nil || fileBody == nil {
		KeyInfoMap["hashkey"] = securecookie.GenerateRandomKey(64)
		KeyInfoMap["blockkey"] = securecookie.GenerateRandomKey(32)
	} else {
		json.Unmarshal(fileBody, &KeyInfoMap)
	}
	CookieHandler = securecookie.New(KeyInfoMap["hashkey"], KeyInfoMap["blockkey"])
	SaveKeyBlock()
}

// CheckPermissions : check permissions for nil
func CheckPermissions(info map[string]interface{}) map[string]interface{} {
	if info["cloudwaf"] == nil {
		info["cloudwaf"] = "0"
	}
	if info["sitechecker"] == nil {
		info["sitechecker"] = "0"
	}
	if info["advanset"] == nil {
		info["advanset"] = "0"
	}
	if info["alarmset"] == nil {
		info["alarmset"] = "0"
	}
	if info["auditlog"] == nil {
		info["auditlog"] = "0"
	}
	if info["modelplan"] == nil || info["modelplan"] == "" {
		info["modelplan"] = "whitelabel"
	}
	if info["is_postpay"] == nil {
		info["is_postpay"] = "0"
	}
	if info["payment_method"] == nil {
		info["payment_method"] = "paypal"
	}
	if info["user_group"] == nil {
		info["user_group"] = "1"
	}
	if info["isinvited"] == nil {
		info["isinvited"] = "0"
	}
	if info["mother_account"] == nil {
		info["mother_account"] = ""
	}
	if info["permission_country"] == nil {
		info["permission_country"] = ""
	}
	if info["org_id"] == nil {
		info["org_id"] = "0"
	}
	if info["last_service"] == nil {
		info["last_service"] = "waf"
	}
	if info["use_swg"] == nil {
		info["use_swg"] = "0"
	}
	return info
}
