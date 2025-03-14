package constants

var (
	JsExtensions   = []string{".js", ".jsx", ".mjs", ".cjs"}
	TsExtensions   = []string{".ts", ".tsx", ".mts", ".cts"}
	NodeExtensions = append(JsExtensions, TsExtensions...)
)
