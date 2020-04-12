package wxa

type ExtConfig struct {
	AppId                          string                       `json:"extAppid"`
	Ext                            map[string]string            `json:"ext,omitempty"`
	Pages                          []string                     `json:"pages,omitempty"`
	ExtPages                       map[string]PageConfig        `json:"extPages,omitempty"`
	Sitemap                        *SitemapConfig               `json:"sitemap,omitempty"`
	Window                         *WindowConfig                `json:"window,omitempty"`
	TabBar                         *TabBarConfig                `json:"tabBar,omitempty"`
	NetworkTimeout                 *NetworkTimeoutConfig        `json:"networkTimeout,omitempty"`
	Debug                          bool                         `json:"debug,omitempty"`
	FunctionalPages                map[string]bool              `json:"functionalPages,omitempty"`
	Subpackages                    []SubpackageConfig           `json:"subpackages,omitempty"`
	Workers                        string                       `json:"workers,omitempty"`
	RequiredBackgroundModes        []string                     `json:"requiredBackgroundModes,omitempty"`
	Plugins                        map[string]PluginConfig      `json:"plugins,omitempty"`
	PreloadRule                    map[string]PreloadRuleConfig `json:"preloadRule,omitempty"`
	Resizable                      bool                         `json:"resizable,omitempty"`
	NavigateToMiniProgramAppIdList []string                     `json:"navigateToMiniProgramAppIdList,omitempty"`
	UsingComponents                map[string]string            `json:"usingComponents,omitempty"`
	permission                     map[string]struct {
		Desc string `json:"desc"`
	} `json:"permission,omitempty"`
	SitemapLocation string `json:"sitemapLocation,omitempty"`
	Style           string `json:"style,omitempty"`
	Enable          bool   `json:"extEnable,omitempty"`
}

type PageConfig struct {
	NavigationBarBackgroundColor string            `json:"navigationBarBackgroundColor,omitempty"`
	NavigationBarTextStyle       string            `json:"navigationBarTextStyle,omitempty"`
	NavigationBarTitleText       string            `json:"navigationBarTitleText,omitempty"`
	NavigationStyle              string            `json:"navigationStyle,omitempty"`
	BackgroundColor              string            `json:"backgroundColor,omitempty"`
	BackgroundTextStyle          string            `json:"backgroundTextStyle,omitempty"`
	BackgroundColorTop           string            `json:"backgroundColorTop,omitempty"`
	BackgroundColorBottom        string            `json:"backgroundColorBottom,omitempty"`
	EnablePullDownRefresh        bool              `json:"enablePullDownRefresh,omitempty"`
	OnReachBottomDistance        int               `json:"onReachBottomDistance,omitempty"`
	PageOrientation              PageOrientation   `json:"pageOrientation,omitempty"`
	DisableScroll                bool              `json:"disableScroll,omitempty"`
	DisableSwipeBack             bool              `json:"disableSwipeBack,omitempty"`
	UsingComponents              map[string]string `json:"usingComponents,omitempty"`
}

type SitemapConfig struct {
	Rules []SitemapRule `json:"rules,omitempty"`
}

type SitemapAction = string

const (
	ALLOW    SitemapAction = "allow"
	DISALLOW SitemapAction = "disallow"
)

type SitemapMatching = string

const (
	EXACT_MATCHING     SitemapMatching = "exact"
	INCLUSIVE_MATCHING SitemapMatching = "inclusive"
	EXCLUSIVE_MATCHING SitemapMatching = "exclusive"
	PARTIAL_MATCHING   SitemapMatching = "partial"
)

type SitemapRule struct {
	Action   SitemapAction   `json:"action,omitempty"`
	Page     string          `json:"page"`
	Params   []string        `json:"params"`
	Matching SitemapMatching `json:"matching,omitempty"`
	Priority uint            `json:"priority,omitempty"`
}

type PageOrientation = string

const (
	AUTO_ORIENTATION      PageOrientation = "auto"
	PORTRAIT_ORIENTATION  PageOrientation = "portrait"
	LANDSCAPE_ORIENTATION PageOrientation = "landscape"
)

type WindowConfig struct {
	NavigationBarBackgroundColor string          `json:"navigationBarBackgroundColor,omitempty"`
	NavigationBarTextStyle       string          `json:"navigationBarTextStyle,omitempty"`
	NavigationBarTitleText       string          `json:"navigationBarTitleText,omitempty"`
	NavigationStyle              string          `json:"navigationStyle,omitempty"`
	BackgroundColor              string          `json:"backgroundColor,omitempty"`
	BackgroundTextStyle          string          `json:"backgroundTextStyle,omitempty"`
	BackgroundColorTop           string          `json:"backgroundColorTop,omitempty"`
	BackgroundColorBottom        string          `json:"backgroundColorBottom,omitempty"`
	EnablePullDownRefresh        bool            `json:"enablePullDownRefresh,omitempty"`
	OnReachBottomDistance        int             `json:"onReachBottomDistance,omitempty"`
	PageOrientation              PageOrientation `json:"pageOrientation,omitempty"`
}

type BorderStyle = string

const (
	BLACK_BORDER BorderStyle = "black"
	WHITE_BORDER BorderStyle = "white"
)

type TabBarPosition = string

const (
	TOP_TABBAR    TabBarPosition = "top"
	BOTTOM_TABBAR TabBarPosition = "bottom"
)

type TabBarConfig struct {
	Color           string         `json:"color,omitempty"`
	SelectedColor   string         `json:"selectedColor,omitempty"`
	BackgroundColor string         `json:"backgroundColor,omitempty"`
	BorderStyle     BorderStyle    `json:"borderStyle,omitempty"`
	List            []TabBarItem   `json:"list,omitempty"`
	Positon         TabBarPosition `json:"position,omitempty"`
	Custom          bool           `json:"custom,omitempty"`
}

type TabBarItem struct {
	PagePath         string `json:"pagePath"`
	Text             string `json:"text"`
	IconPath         string `json:"iconPath,omitempty"`
	SelectedIconPath string `json:"selectedIconPath,omitempty"`
}

type NetworkTimeoutConfig struct {
	Request       int64 `json:"request,omitempty"`
	ConnectSocket int64 `json:"connectSocket,omitempty"`
	UploadFile    int64 `json:"uploadFile,omitempty"`
	DownloadFile  int64 `json:"downloadFile,omitempty"`
}

type PluginConfig struct {
	Version  string `json:"version"`
	Provider string `json:"provider"`
}

type SubpackageConfig struct {
	Root        string                  `json:"root"`
	Name        string                  `json:"name,omitempty"`
	Pages       []string                `json:"pages,omitempty"`
	Plugins     map[string]PluginConfig `json:"plugins,omitempty"`
	Independent bool                    `json:"independent,omitempty"`
}

type PreloadRuleNetwork = string

const (
	ALL_PRELOAD_RULE_NETWORK  PreloadRuleNetwork = "all"
	WIFI_PRELOAD_RULE_NETWORK PreloadRuleNetwork = "wifi"
)

type PreloadRuleConfig struct {
	Packages []string           `json:"packages"`
	Network  PreloadRuleNetwork `json:"network"`
}
