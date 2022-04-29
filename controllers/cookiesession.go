package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/sessions"
)

func SCookie(ec echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "clwan-domain-name"
	cookie.Value = "clwan.com"
	cookie.Path = "/"
	cookie.MaxAge = 3600
	ec.SetCookie(cookie)
	return ec.String(http.StatusOK, "cookie操作")
}

func RCookie(ec echo.Context) error {
	cookie, _ := ec.Cookie("clwan-domain-name")
	fmt.Println(cookie.Domain)
	cookies := ec.Cookies()
	for _, cookie := range cookies {
		fmt.Printf("%+v,%+v,%+v\n", cookie.Name, cookie.Value, cookie.MaxAge)
	}
	return ec.String(http.StatusOK, "读取cookie")
}

var sessionPath = "/tmp/session_data"
var sessionKey = "Onxuh20a2ihhh2"
var Store = sessions.NewFilesystemStore(sessionPath, []byte(sessionKey))

func SSession(ec echo.Context) error {
	sess, err := Store.Get(ec.Request(), "user")
	if err != nil {
		return err
	}
	sess.Values["name"] = "dj"
	sess.Values["age"] = 18
	err = sess.Save(ec.Request(), ec.Response().Writer)
	if err != nil {
		return err
	}
	return nil
}

func RSession(ec echo.Context) error {
	sess, _ := Store.Get(ec.Request(), "user")
	fmt.Fprintf(ec.Response().Writer, "name:%s age:%d\n", sess.Values["name"], sess.Values["age"])
	return nil
}
