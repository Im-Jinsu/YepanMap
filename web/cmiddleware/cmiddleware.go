package cmiddleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/im-jinsu/yepanmap/shared/mongodb"
	"github.com/im-jinsu/yepanmap/web/cerror"
	"github.com/im-jinsu/yepanmap/web/session"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
)

// CustomContext :
type CustomContext struct {
	echo.Context
	UserInfo map[string]string
}

// MidRedirect :
func MidRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}

		ok, _ := cc.CheckAdminSession()
		log.Println(ok)
		if ok {
			// login page
			if c.Request().Method == "GET" && c.Request().RequestURI == "/login" {
				log.Println("test1")
				return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s/adminpage/memberlist", c.Request().Host))
			}

			// empty path
			if c.Request().Method == "GET" && (c.Request().RequestURI == "/" || c.Request().RequestURI == "") {
				return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s/adminpage/memberlist", c.Request().Host))
			}
		} else {
			if c.Request().Method == "GET" && (c.Request().RequestURI == "/" || c.Request().RequestURI == "") {
				return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s/adminpage/login", c.Request().Host))
			}
		}

		return next(cc)
	}
}

// MidWithAuth : Custom middleware for auth
func MidWithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}

		ok, userInfo := cc.CheckAdminSession()
		if !ok {
			return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s/adminpage/login", c.Request().Host))
		}

		cc.UserInfo = userInfo
		return next(cc)
	}
}

// RespErr : error response
func (c *CustomContext) RespErr(code int, msg string) error {
	return c.JSON(code, cerror.Message{Msg: msg})
}

// Resp : Response
func (c *CustomContext) Resp(code int, i interface{}) error {
	return c.JSON(code, i)
}

// RespStatus : Response
func (c *CustomContext) RespStatus(code int, status string, msg string) error {
	result := make(map[string]string)
	result["status"] = status
	result["msg"] = msg
	return c.JSON(code, result)
}

// CheckAdminSession : adminpage session check
func (c *CustomContext) CheckAdminSession() (ok bool, info map[string]string) {
	if cookie, err := c.Context.Cookie("adminsession"); err == nil {
		info = make(map[string]string)
		cookieValue := make(map[string]interface{})
		if err = session.CookieHandler.Decode("adminsession", cookie.Value, &cookieValue); err == nil {
			cookieValue = session.CheckPermissions(cookieValue)
			if cookieValue["isadmin"] == "10" {
				info["email"] = cookieValue["email"].(string)
				info["isadmin"] = cookieValue["isadmin"].(string)
				info["group"] = cookieValue["group"].(string)
				info["language"] = cookieValue["language"].(string)
				info["cloudwaf"] = cookieValue["cloudwaf"].(string)
				info["sitechecker"] = cookieValue["sitechecker"].(string)
				info["advanset"] = cookieValue["advanset"].(string)
				info["alarmset"] = cookieValue["alarmset"].(string)
				info["auditlog"] = cookieValue["auditlog"].(string)
				info["is_postpay"] = cookieValue["is_postpay"].(string)
				info["payment_method"] = cookieValue["payment_method"].(string)
				info["user_group"] = cookieValue["user_group"].(string)
				info["permission_country"] = cookieValue["permission_country"].(string)
				ok = true
				return ok, info
			}
			info["isadmin"] = cookieValue["isadmin"].(string)
			info["group"] = cookieValue["group"].(string)
			info["language"] = cookieValue["language"].(string)
			info["cloudwaf"] = cookieValue["cloudwaf"].(string)
			info["sitechecker"] = cookieValue["sitechecker"].(string)
			info["advanset"] = cookieValue["advanset"].(string)
			info["alarmset"] = cookieValue["alarmset"].(string)
			info["auditlog"] = cookieValue["auditlog"].(string)
			info["is_postpay"] = cookieValue["is_postpay"].(string)
			info["payment_method"] = cookieValue["payment_method"].(string)
			info["user_group"] = cookieValue["user_group"].(string)
			info["permission_country"] = cookieValue["permission_country"].(string)
			oldTime := cookieValue["ctime"].(int64)
			curTime := time.Now().Unix()
			diffTime := curTime - oldTime
			// for test
			if diffTime > 36000 {
				c.ClearSession()
				info["email"] = ""
				ok = false
			} else {
				cookieValue["ctime"] = curTime
				c.SetSession(cookieValue)
				info["email"] = cookieValue["email"].(string)
				ok = true
			}
		} else {
			log.Println(err)
			if cookie, err := c.Context.Cookie("sessionid"); err == nil {
				log.Println(cookie)
				var exist bool
				var email string
				log.Println(session.RedirectLogin)
				if email, exist = session.RedirectLogin[cookie.Value]; exist {
					log.Println(c.Context.Cookie("sessionid"))

					var result = make(map[string]interface{})
					var tmpData = make(map[string]interface{})

					client, err := mongodb.SetClient()
					if err != nil {
						ok = false
						log.Println(err)
					}

					coll := client.Database(mongodb.SetUserDBName(email)).Collection("user_info")
					err = coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
					if err != nil {
						log.Println(err)
					}

					result = session.CheckPermissions(result)

					tmpData["email"] = email
					tmpData["isadmin"] = result["isadmin"].(string)
					tmpData["group"] = result["group"].(string)
					tmpData["language"] = result["language"].(string)
					tmpData["cloudwaf"] = result["cloudwaf"].(string)
					tmpData["sitechecker"] = result["sitechecker"].(string)
					tmpData["advanset"] = result["advanset"].(string)
					tmpData["alarmset"] = result["alarmset"].(string)
					tmpData["auditlog"] = result["auditlog"].(string)
					tmpData["is_postpay"] = result["is_postpay"].(string)
					tmpData["payment_method"] = result["payment_method"].(string)
					tmpData["user_group"] = result["user_group"].(string)
					tmpData["permission_country"] = result["permission_country"].(string)
					c.SetAdminSession(tmpData)
					delete(session.RedirectLogin, cookie.Value)
					ok = true
				} else {
					ok = false
					log.Println(err)
				}
			}
		}
	} else if cookie, err := c.Context.Cookie("sessionid"); err == nil {
		var exist bool
		var email string
		if email, exist = session.RedirectLogin[cookie.Value]; exist {
			var result = make(map[string]interface{})
			var tmpData = make(map[string]interface{})

			client, err := mongodb.SetClient()
			if err != nil {
				ok = false
				log.Println(err)
			}

			coll := client.Database(mongodb.SetUserDBName(email)).Collection("user_info")
			err = coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
			if err != nil {
				log.Println(err)
			}

			result = session.CheckPermissions(result)

			tmpData["email"] = email
			tmpData["isadmin"] = result["isadmin"].(string)
			tmpData["group"] = result["group"].(string)
			tmpData["language"] = result["language"].(string)
			tmpData["cloudwaf"] = result["cloudwaf"].(string)
			tmpData["sitechecker"] = result["sitechecker"].(string)
			tmpData["advanset"] = result["advanset"].(string)
			tmpData["alarmset"] = result["alarmset"].(string)
			tmpData["auditlog"] = result["auditlog"].(string)
			tmpData["is_postpay"] = result["is_postpay"].(string)
			tmpData["payment_method"] = result["payment_method"].(string)
			tmpData["user_group"] = result["user_group"].(string)
			tmpData["permission_country"] = result["permission_country"].(string)
			c.SetAdminSession(tmpData)
			delete(session.RedirectLogin, cookie.Value)

			info = make(map[string]string)
			info["email"] = email
			info["isadmin"] = result["isadmin"].(string)
			info["group"] = result["group"].(string)
			info["language"] = result["language"].(string)
			info["cloudwaf"] = result["cloudwaf"].(string)
			info["sitechecker"] = result["sitechecker"].(string)
			info["advanset"] = result["advanset"].(string)
			info["alarmset"] = result["alarmset"].(string)
			info["auditlog"] = result["auditlog"].(string)
			info["is_postpay"] = result["is_postpay"].(string)
			info["payment_method"] = result["payment_method"].(string)
			info["user_group"] = result["user_group"].(string)
			info["permission_country"] = result["permission_country"].(string)
			ok = true
		} else {
			ok = false
			log.Println(err)
		}
	}
	return ok, info
}

// SetSession : Session Setting
func (c *CustomContext) SetSession(data map[string]interface{}) *http.Cookie {
	value := map[string]interface{}{
		"email":          data["email"],
		"isadmin":        data["isadmin"],
		"ctime":          time.Now().Unix(),
		"domainList":     data["domainList"],
		"group":          data["group"],
		"language":       data["language"],
		"cloudwaf":       data["cloudwaf"],
		"sitechecker":    data["sitechecker"],
		"advanset":       data["advanset"],
		"alarmset":       data["alarmset"],
		"auditlog":       data["auditlog"],
		"modelplan":      data["modelplan"],
		"targetdomain":   data["targetdomain"],
		"is_postpay":     data["is_postpay"],
		"payment_method": data["payment_method"],
		"user_group":     data["user_group"],
		"isinvited":      data["isinvited"],
		"mother_account": data["mother_account"],
		"f_name":         data["f_name"],
		"l_name":         data["l_name"],
		"org_id":         data["org_id"],
		"last_service":   data["last_service"],
		"use_swg":        data["use_swg"],
	}
	encoded, err := session.CookieHandler.Encode("session", value)
	if err != nil {
		log.Println("[ERR]", err)
		return nil
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	c.Context.SetCookie(cookie)
	return cookie
}

// ClearSession : Session Clear
func (c *CustomContext) ClearSession() {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	c.Context.SetCookie(cookie)
}

// SetAdminSession : Admin page Session Setting
func (c *CustomContext) SetAdminSession(data map[string]interface{}) *http.Cookie {
	value := map[string]interface{}{
		"email":              data["email"],
		"isadmin":            data["isadmin"],
		"ctime":              time.Now().Unix(),
		"group":              data["group"],
		"language":           data["language"],
		"cloudwaf":           data["cloudwaf"],
		"sitechecker":        data["sitechecker"],
		"advanset":           data["advanset"],
		"alarmset":           data["alarmset"],
		"auditlog":           data["auditlog"],
		"targetdomain":       data["targetdomain"],
		"is_postpay":         data["is_postpay"],
		"payment_method":     data["payment_method"],
		"user_group":         data["user_group"],
		"permission_country": data["permission_country"],
	}
	encoded, err := session.CookieHandler.Encode("adminsession", value)
	if err != nil {
		log.Println("[ERR]", err)
		return nil
	}
	cookie := &http.Cookie{
		Name:     "adminsession",
		Value:    encoded,
		Path:     "/adminpage",
		MaxAge:   0,
		HttpOnly: true,
	}
	c.Context.SetCookie(cookie)
	return cookie
}

// ClearAdminSession : adminpage Session Clear
func (c *CustomContext) ClearAdminSession() {
	cookie := &http.Cookie{
		Name:   "adminsession",
		Value:  "",
		Path:   "/adminpage",
		MaxAge: -1,
	}
	c.Context.SetCookie(cookie)
}
