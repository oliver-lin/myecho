package boot

import (
	"echo/controllers"
	"echo/middle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRouter() *echo.Echo {
	//实例化echo对象。
	e := echo.New()

	var U controllers.User
	// post表单操作
	e.POST("/postjson", U.PostJsonUser)
	e.POST("/postform", U.PostFormUser)
	e.POST("/postformvalue", U.PostFormValue)

	// get query操作
	e.GET("/users/:name/:age", U.GetUser)
	e.GET("/users", U.GetUser)

	// 模板的使用
	var gs controllers.Person
	e.GET("/tpldemo", gs.Tpldemo)

	// Any 可以接收 Get 或 Post 请求
	reqAny(e)

	// 中间件的使用，起拦截作用
	// Pre ：在路由之前执行，不可使用 echo.Context
	// log中可以看到 http://127.0.0.1:8080/tpldemo 请求地址变为："uri":"/api/user/add/ ,尾巴 / 会报错，路由解析会有问题
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())

	// Use：路由执行完，执行，可使用 echo.Context
	// 在终端打印输出具体请求
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middle.Count)

	// cookie 操作
	e.GET("/setcookie", controllers.SCookie)
	e.GET("/readcookie", controllers.RCookie)

	// session 操作
	e.Use(session.Middleware(controllers.Store))
	e.GET("/setsession", controllers.SSession)
	e.GET("/getsession", controllers.RSession)

	// e.StdLogger.Panic("sss")
	// e.Use(middleware.Recover()) // 记录panic

	e.GET("/html", controllers.Displayhtml)
	e.GET("/ret", controllers.RetJson)
	e.GET("/json", controllers.Displayjson)

	e.GET("/routes", func(c echo.Context) error {
		fmt.Println(111, c.Response().Header().Get("X-Request-Id"))
		return c.JSON(http.StatusOK, e.Routes())
	})
	// 加入各个分组
	setUserApi(e)

	return e
}

// 分组
func setUserApi(e *echo.Echo) {
	g := e.Group("/api/user")
	//注册一个Get请求, 路由地址为: /max  并且绑定一个控制器函数, 这里使用的是闭包函数。
	g.GET("/add", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.String(http.StatusOK, "api/user/add")
	})

}

func reqAny(e *echo.Echo) {

	// http://127.0.0.1:8080/postform?uid=333&uid=444 结果：[333 444]
	e.Any("/queryformjson", func(c echo.Context) error {
		body := struct {
			Name string `form:"name" json:"name"`
			Sex  string `form:"sex" json:"sex"`
		}{}
		r := c.Request()
		// query,post 请求
		// r.ParseForm()
		r.ParseMultipartForm(32 << 20)
		if len(r.Form["name"]) != 0 {
			fmt.Printf("%+v,%+v,%+v\n", r.Form["name"], r.Form["sex"][0], r.FormValue("name")) // 此方法只能打印出 json的值，会把key丢失
			return nil
		}

		// json 数据请求
		// Request: { "name":"xuxinhua", "sex": "male", "age": 20 }
		// 根据上面定义的body结构体取值，没有定义则不取
		jstr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		isjson := json.Valid(jstr)
		if isjson {
			err = json.Unmarshal(jstr, &body)
			if err != nil {
				return err
			}
		}
		fmt.Printf("%+v,%+v\n", body, string(jstr))
		return nil
	})
}
