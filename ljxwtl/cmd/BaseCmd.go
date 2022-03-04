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
	NewSessionHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "New Session",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	DeleteSessionHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Delete Session",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID",
			CmdHttpRequestMethod: http.MethodDelete,
		},
	}
	StatusHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Status",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/status",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetTimeoutsHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Timeouts",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/timeouts",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	SetTimeoutsHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Set Timeouts",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/timeouts",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n   \n    \"implicit\": IMPLICIT,\n    \"pageLoad\": PAGE_LOAD,\n    \"script\": SCRIPT\n    \n}",
		},
	}
	NavigateToHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Navigate To",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/url",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n   \n    \"url\" : \"Navigate_To_URL\"\n    \n}",
		},
	}
	GetCurrentURLHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Current URL",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/url",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	BackHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Back",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/back",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	ForwardHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Forward",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/forward",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	RefreshHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Refresh",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/refresh",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	GetTitleHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Title",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/title",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetWindowHandleHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Window Handle",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	RefreshHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Refresh",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/refresh",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
)
