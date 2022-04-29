package controllers

import (
	"fmt"
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
)

//这里定义一个struct类型的模版参数，实际上模版可以是任意类型数据
type GameStatus struct {
	Gname string
	IsWin bool
}

type Person struct {
	Name  string
	Email []string
	Pgs   []GameStatus
}

func (p *Person) Tpldemo(ec echo.Context) error {
	tmpsct := []GameStatus{
		{"魔兽", true},
		{"星际", false},
		{"红警", true},
	}
	// 加载模版代码，并且创建template对象t
	// template.ParseGlob 函数加载views目录下的所有tpl为后缀的模版文件
	// template.Must函数主要用于检测加载的模版有没有错误，有错误输出panic错误，并且结束程序。
	tpl, err := template.ParseFiles("./views/demo.tpl")
	if err != nil {
		fmt.Println(err)
	}
	t := template.Must(tpl, err)

	p.Name = "游戏大侠"
	p.Email = []string{"a@g.com", "b@g.cn"}
	p.Pgs = tmpsct

	//根据参数u, 渲染命名为demo的模板，并且将渲染结果打印到标准输出
	// err := t.ExecuteTemplate(os.Stdout, "demo", u)
	err = t.Execute(ec.Response().Writer, p)
	if err != nil {
		log.Println("executing template:", err)
	}
	return nil
}
