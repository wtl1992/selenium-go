package cmd

import "net/http"

type BaseCmd struct {
	// cmd 的名称
	CmdName string
	// cmd 的执行参数
	CmdHttpAddress string
	//cmd 的http请求方法
	CmdHttpRequestMethod string
	// cmd 的http请求体
	CmdHttpRequestBody string
}

type HttpRequestCmd struct {
	BaseCmd
}

var (
	NewSessionHttpRequestCmd = HttpRequestCmd{BaseCmd{CmdName: "New Session", CmdHttpAddress: "http://localhost:SELENIUM_LISTEN_PORT/session", CmdHttpRequestMethod: http.MethodPost}}
)
