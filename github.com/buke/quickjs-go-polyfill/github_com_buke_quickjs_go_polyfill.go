package quickjs_go_polyfill
import (
github_com_buke_quickjs_go_polyfill "github.com/buke/quickjs-go-polyfill"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/buke/quickjs-go-polyfill"]; !ok {
   Api["github.com/buke/quickjs-go-polyfill"] = map[string]interface{}{}
}
if _, ok := Api["github.com/buke/quickjs-go-polyfill/new"]; !ok {
	Api["github.com/buke/quickjs-go-polyfill/new"] = map[string]interface{}{}
}
Api["github.com/buke/quickjs-go-polyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
}
