package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	Breakfaset = iota
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
	Age   int    `json:"age" form:"age" query:"age"`
}

func (u *User) GetUser(c echo.Context) error {
	fmt.Println(c.ParamNames(), c.ParamValues(), c.QueryParams())

	slc := make([]string, 10)
	slc = append(slc, "http://127.0.0.1:8080/users/linning/20?p1=a&p2=b")
	slc_path_keys := c.ParamNames()
	slc_path_vals := c.ParamValues()
	map_query_params := c.QueryParams()
	if len(slc_path_keys) != 0 {
		slc = append(slc, "ParamName,切片形式，path/:name/:age上的路径参数名[name age]：")
		slc = append(slc, slc_path_keys...)
	}
	if len(slc_path_vals) != 0 {
		slc = append(slc, "ParamValues,切片形式，path/:name/:age上的路径参数名对应的值[linning 20]：")
		slc = append(slc, slc_path_vals...)
	}
	if len(map_query_params) != 0 {
		slc = append(slc, "QueryParams?后的key=value&key=value值 map[p1:[a] p2:[b]]:")
		for k, _ := range map_query_params {
			slc = append(slc, k+"=>"+c.QueryParam(k))
		}
		p1 := c.QueryParam("p1")
		p2 := c.QueryParam("p2")
		name := c.Param("name")
		age := c.Param("age")
		path := c.Path()
		slc = append(slc, "path="+path+"\n name="+name+"\n age="+age+"\n c.QueryParam(\"p1\")="+p1+"\n c.QueryParam(\"p2\")="+p2)
	}
	return c.JSON(http.StatusOK, slc)
}

// post json请求
//	curl --location --request POST 'http://127.0.0.1:8080/postjson' \
//	--header 'Content-Type: application/json' \
//	--data-raw '{
//		"name":"myfirstname",
//		"email":"myEmail"
//	}'
// Response:
// 	{ "name": "form_name", "email": "form_email" }
func (u *User) PostJsonUser(ec echo.Context) error {
	err := ec.Bind(u)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, u)
}

// Request
// 	curl --location --request POST 'http://127.0.0.1:8080/postform' \
// 	--form 'name="form_name"' \
// 	--form 'email="form_email"'
// Response : 二维数组
// 	map[email:[formemail] name:[formname]]
// 	map[email][0] == formemail
func (u *User) PostFormUser(ec echo.Context) error {
	urlvalues, err := ec.FormParams()
	if err != nil {
		return err
	}
	u.Name = urlvalues["name"][0]
	u.Email = urlvalues["email"][0]
	return ec.JSON(http.StatusOK, u)
}

// FormValue函数获取参数的值，数据类型都是String类型， 如果需要其他类型的数据，需要自己转换数据格式。
func (u *User) PostFormValue(ec echo.Context) error {
	u.Name = ec.FormValue("name")
	u.Email = ec.FormValue("email")
	u.Age, _ = strconv.Atoi(ec.FormValue("age"))
	// fmt.Printf("%+v", u)
	return ec.JSON(http.StatusOK, u)
}
