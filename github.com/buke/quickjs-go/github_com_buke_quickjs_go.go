package quickjs_go
import (
github_com_buke_quickjs_go "github.com/buke/quickjs-go"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/buke/quickjs-go"]; !ok {
   Api["github.com/buke/quickjs-go"] = map[string]interface{}{}
}
if _, ok := Api["github.com/buke/quickjs-go/new"]; !ok {
	Api["github.com/buke/quickjs-go/new"] = map[string]interface{}{}
}
Api["github.com/buke/quickjs-go"]["Atom"] = github_com_buke_quickjs_go.Atom{}
Api["github.com/buke/quickjs-go/new"]["Atom"] = func(args ...interface{}) interface{} {
					return new(github_com_buke_quickjs_go.Atom)
}
Api["github.com/buke/quickjs-go"]["Loop"] = github_com_buke_quickjs_go.Loop{}
Api["github.com/buke/quickjs-go/new"]["Loop"] = func(args ...interface{}) interface{} {
					return new(github_com_buke_quickjs_go.Loop)
}
Api["github.com/buke/quickjs-go"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
Api["github.com/buke/quickjs-go"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
}
