package geziyor
import (
github_com_geziyor_geziyor "github.com/geziyor/geziyor"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/geziyor/geziyor"]; !ok {
   Api["github.com/geziyor/geziyor"] = map[string]interface{}{}
}
if _, ok := Api["github.com/geziyor/geziyor/new"]; !ok {
	Api["github.com/geziyor/geziyor/new"] = map[string]interface{}{}
}
Api["github.com/geziyor/geziyor"]["Geziyor"] = github_com_geziyor_geziyor.Geziyor{}
Api["github.com/geziyor/geziyor/new"]["Geziyor"] = func(args ...interface{}) interface{} {
					return new(github_com_geziyor_geziyor.Geziyor)
}
Api["github.com/geziyor/geziyor"]["NewGeziyor"] = github_com_geziyor_geziyor.NewGeziyor
Api["github.com/geziyor/geziyor"]["Options"] = github_com_geziyor_geziyor.Options{}
Api["github.com/geziyor/geziyor/new"]["Options"] = func(args ...interface{}) interface{} {
					return new(github_com_geziyor_geziyor.Options)
}
}
