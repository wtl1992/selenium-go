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
	CloseWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Close Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window",
			CmdHttpRequestMethod: http.MethodDelete,
			CmdHttpRequestBody:   "{\n   \n    \"window\" : \"WINDOW_HANDLE\"\n    \n}",
		},
	}
	SwitchToWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Switch To Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n   \n    \"window\" : \"WINDOW_HANDLE\"\n    \n}",
		},
	}
	GetWindowHandleswHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Window Handles",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/handles",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	NewWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "New Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/new",
			CmdHttpRequestMethod: http.MethodPost,
			/**
			type: tab、window
			*/
			CmdHttpRequestBody: "{\n   \n    \"type\":\"NEW_WINDOW_TYPE\"\n}",
		},
	}
	SwitchToFrameHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Switch To Frame",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/frame",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n   \n    \"id\" : \"FRAME_ID\"\n}",
		},
	}
	SwitchToParentFrameHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Switch To Parent Frame",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/frame/parent",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	GetWindowRectHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Window Rect",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/rect",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	SetWindowRectHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Set Window Rect",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/rect",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"width\": WINDOW_RECT_WIDTH,\n    \"height\": WINDOW_RECT_HEIGHT,\n    \"x\": WINDOW_RECT_X\n    \"y\": WINDOW_RECT_Y\n}",
		},
	}
	MaximizeWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Maximize Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/maximize",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	MinimizeWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Minimize Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/minimize",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	FullscreenWindowHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Fullscreen Window",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/window/fullscreen",
			CmdHttpRequestMethod: http.MethodPost,
		},
	}
	FindElementHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Element",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element",
			CmdHttpRequestMethod: http.MethodPost,
			/**
			using:
			CSS selector	"css selector"
			Link text selector	"link text"
			Partial link text selector	"partial link text"
			Tag name	"tag name"
			XPath selector	"xpath"
			*/
			CmdHttpRequestBody: "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	FindElementsHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Elements",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/elements",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	FindElementFromElementHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Element From Element",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/element",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	FindElementsFromElementHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Elements From Element",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/elements",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	FindElementFromShadowRootHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Element From Shadow Root",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/shadow/SHADOW_ID/element",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	FindElementsFromShadowRootHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Find Elements From Shadow Root",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/shadow/SHADOW_ID/elements",
			CmdHttpRequestMethod: http.MethodPost,
			CmdHttpRequestBody:   "{\n    \"using\" :\"LOCATION_STRATEGY\",\n    \"value\":\"SELECT_VALUE\"\n}",
		},
	}
	GetActiveElementHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Active Element",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/active",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetElementShadowRootHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Element Shadow Root",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/shadow",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	IsElementSelectedHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Is Element Selected",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/selected",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetElementAttributeHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Element Attribute",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/attribute/NAME",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetElementPropertyHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Element Property",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/property/NAME",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
	GetElementCSSValueHttpRequestCmd = HttpRequestCmd{
		BaseCmd{
			CmdName:              "Get Element CSS Value",
			CmdHttpAddress:       "http://localhost:SELENIUM_LISTEN_PORT/session/SESSION_ID/element/ELEMENT_ID/css/PROPERTY_NAME",
			CmdHttpRequestMethod: http.MethodGet,
		},
	}
)
