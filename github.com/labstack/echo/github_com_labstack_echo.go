package echo
import (
github_com_labstack_echo "github.com/labstack/echo"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/labstack/echo"]; !ok {
   Api["github.com/labstack/echo"] = map[string]interface{}{}
}
if _, ok := Api["github.com/labstack/echo/new"]; !ok {
	Api["github.com/labstack/echo/new"] = map[string]interface{}{}
}
Api["github.com/labstack/echo"]["CONNECT"] = github_com_labstack_echo.CONNECT
Api["github.com/labstack/echo"]["DELETE"] = github_com_labstack_echo.DELETE
Api["github.com/labstack/echo"]["GET"] = github_com_labstack_echo.GET
Api["github.com/labstack/echo"]["HEAD"] = github_com_labstack_echo.HEAD
Api["github.com/labstack/echo"]["OPTIONS"] = github_com_labstack_echo.OPTIONS
Api["github.com/labstack/echo"]["PATCH"] = github_com_labstack_echo.PATCH
Api["github.com/labstack/echo"]["POST"] = github_com_labstack_echo.POST
Api["github.com/labstack/echo"]["PUT"] = github_com_labstack_echo.PUT
Api["github.com/labstack/echo"]["TRACE"] = github_com_labstack_echo.TRACE
Api["github.com/labstack/echo"]["PROPFIND"] = github_com_labstack_echo.PROPFIND
Api["github.com/labstack/echo"]["DefaultBinder"] = github_com_labstack_echo.DefaultBinder{}
Api["github.com/labstack/echo/new"]["DefaultBinder"] = func(args ...interface{}) interface{} {
					return new(github_com_labstack_echo.DefaultBinder)
}
Api["github.com/labstack/echo"]["Echo"] = github_com_labstack_echo.Echo{}
Api["github.com/labstack/echo/new"]["Echo"] = func(args ...interface{}) interface{} {
					return new(github_com_labstack_echo.Echo)
}
Api["github.com/labstack/echo"]["HTTPError"] = github_com_labstack_echo.HTTPError{}
Api["github.com/labstack/echo/new"]["HTTPError"] = func(args ...interface{}) interface{} {
					return new(github_com_labstack_echo.HTTPError)
}
Api["github.com/labstack/echo"]["NewHTTPError"] = github_com_labstack_echo.NewHTTPError
Api["github.com/labstack/echo"]["WrapHandler"] = github_com_labstack_echo.WrapHandler
Api["github.com/labstack/echo"]["WrapMiddleware"] = github_com_labstack_echo.WrapMiddleware
Api["github.com/labstack/echo"]["Response"] = github_com_labstack_echo.Response{}
Api["github.com/labstack/echo/new"]["Response"] = func(args ...interface{}) interface{} {
					return new(github_com_labstack_echo.Response)
}
Api["github.com/labstack/echo"]["NewResponse"] = github_com_labstack_echo.NewResponse
Api["github.com/labstack/echo"]["Router"] = github_com_labstack_echo.Router{}
Api["github.com/labstack/echo/new"]["Router"] = func(args ...interface{}) interface{} {
					return new(github_com_labstack_echo.Router)
}
Api["github.com/labstack/echo"]["NewRouter"] = github_com_labstack_echo.NewRouter
}
