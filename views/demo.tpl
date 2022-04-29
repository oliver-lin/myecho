<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        单独一个姓名 .Name
        <p>
        hello,{{.Name}}
        </p>
    </p>
    <p>
        单独 Email 循环
        {{range .Email}}
        <p>
            email: {{.}}
        </p>
        {{- end}}
    </p>
    <p>
        结构体 Pgs 内，含义另一结构属性（Gname,IsWin）
        {{with .Pgs}}
            {{- range .}}
            <p>
                <p>
                昵称： {{.Gname}},
                </p>
                <p>
                {{- if .IsWin}}
                恭喜，大吉大利，今晚吃鸡！
                {{- else}}
                遗憾，鸡被吃光了！
                {{- end}}
                </p>
            </p>
            {{- end}}
        {{end}}
    </p>
</body>
</html>
