package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User2 struct {
	Id   int    // 必须大写，否则私有变量
	Name string `json:"name"` // 最后展示需小写，可以这样注明
}

func Displayjson(c echo.Context) error {
	u := User2{2, "joson"}
	return c.JSON(http.StatusOK, u)
}

func Displayhtml(c echo.Context) error {
	//以网页形式返回html代码，c.HTML语法: c.HTML(http状态码, "html内容")
	html := "<html><head><title>tizi365.com</title></head><body>欢迎访问tizi365.com</body></html>"
	return c.HTML(http.StatusOK, html)
}

func RetJson(c echo.Context) error {
	// TMap := map[string]interface{}{"code": 0, "msg": "111", "data": "222"}
	TMap := "ssss"
	return c.JSON(http.StatusForbidden, TMap)
}
