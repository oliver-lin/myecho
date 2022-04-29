package middle

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var total = 0

func Count(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//在这里处理拦截请求的逻辑
		total++
		fmt.Println(total)
		//执行下一个中间件或者执行控制器函数, 然后返回执行结果
		return next(c)
	}
}
