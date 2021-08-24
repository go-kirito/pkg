/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/24 17:46
 */
package errors

import (
	"bytes"
	"errors"
	"html/template"
)

var errCodeDocPrefix = `# 错误码

！！{{.Name}}项目错误码列表，由命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中 {{.Symbol}}code{{.Symbol}} 字段不等于"0"，则表示调用 API 接口失败。例如：

{{.Symbol}}{{.Symbol}}{{.Symbol}}json
{
  "code": "100001",
  "message": "用户不存在"
}
{{.Symbol}}{{.Symbol}}{{.Symbol}}

上述返回中 {{.Symbol}}code{{.Symbol}} 表示错误码，{{.Symbol}}message{{.Symbol}} 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 404(NotFound)。

## 错误码列表

系统支持的错误码列表如下：

| HTTP Code |  Code | Description |
| --------- |  ---- | ----------- |{{ range $key, $value := .Codes }}
| {{$value.Code}} | {{$value.Reason}} | {{$value.Message}} |
{{- end }}
`

type content struct {
	Name   string
	Codes  map[string]*Error
	Symbol string
}

func GenErrCodeDoc(appName string) ([]byte, error) {
	if len(_codes) > 0 {
		c := content{
			Name:   appName,
			Codes:  _codes,
			Symbol: "`",
		}
		tmpl, err := template.New("doc").Parse(errCodeDocPrefix)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, c)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, errors.New("没有需要生成的错误码文档")
}
