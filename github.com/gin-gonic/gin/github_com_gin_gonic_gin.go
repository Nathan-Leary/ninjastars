package gin
import (
github_com_gin_gonic_gin "github.com/gin-gonic/gin"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/gin-gonic/gin"]; !ok {
   Api["github.com/gin-gonic/gin"] = map[string]interface{}{}
}
if _, ok := Api["github.com/gin-gonic/gin/new"]; !ok {
	Api["github.com/gin-gonic/gin/new"] = map[string]interface{}{}
}
Api["github.com/gin-gonic/gin"]["MIMEJSON"] = github_com_gin_gonic_gin.MIMEJSON
Api["github.com/gin-gonic/gin"]["MIMEHTML"] = github_com_gin_gonic_gin.MIMEHTML
Api["github.com/gin-gonic/gin"]["MIMEXML"] = github_com_gin_gonic_gin.MIMEXML
Api["github.com/gin-gonic/gin"]["MIMEXML2"] = github_com_gin_gonic_gin.MIMEXML2
Api["github.com/gin-gonic/gin"]["MIMEPlain"] = github_com_gin_gonic_gin.MIMEPlain
Api["github.com/gin-gonic/gin"]["MIMEPOSTForm"] = github_com_gin_gonic_gin.MIMEPOSTForm
Api["github.com/gin-gonic/gin"]["MIMEMultipartPOSTForm"] = github_com_gin_gonic_gin.MIMEMultipartPOSTForm
Api["github.com/gin-gonic/gin"]["MIMEYAML"] = github_com_gin_gonic_gin.MIMEYAML
Api["github.com/gin-gonic/gin"]["MIMETOML"] = github_com_gin_gonic_gin.MIMETOML
Api["github.com/gin-gonic/gin"]["PlatformGoogleAppEngine"] = github_com_gin_gonic_gin.PlatformGoogleAppEngine
Api["github.com/gin-gonic/gin"]["PlatformCloudflare"] = github_com_gin_gonic_gin.PlatformCloudflare
Api["github.com/gin-gonic/gin"]["DebugMode"] = github_com_gin_gonic_gin.DebugMode
Api["github.com/gin-gonic/gin"]["ReleaseMode"] = github_com_gin_gonic_gin.ReleaseMode
Api["github.com/gin-gonic/gin"]["TestMode"] = github_com_gin_gonic_gin.TestMode
Api["github.com/gin-gonic/gin"]["DebugPrintRouteFunc"] = github_com_gin_gonic_gin.DebugPrintRouteFunc
Api["github.com/gin-gonic/gin"]["DefaultErrorWriter"] = github_com_gin_gonic_gin.DefaultErrorWriter
Api["github.com/gin-gonic/gin"]["DefaultWriter"] = github_com_gin_gonic_gin.DefaultWriter
Api["github.com/gin-gonic/gin"]["CreateTestContext"] = github_com_gin_gonic_gin.CreateTestContext
Api["github.com/gin-gonic/gin"]["Dir"] = github_com_gin_gonic_gin.Dir
Api["github.com/gin-gonic/gin"]["DisableBindValidation"] = github_com_gin_gonic_gin.DisableBindValidation
Api["github.com/gin-gonic/gin"]["DisableConsoleColor"] = github_com_gin_gonic_gin.DisableConsoleColor
Api["github.com/gin-gonic/gin"]["EnableJsonDecoderDisallowUnknownFields"] = github_com_gin_gonic_gin.EnableJsonDecoderDisallowUnknownFields
Api["github.com/gin-gonic/gin"]["EnableJsonDecoderUseNumber"] = github_com_gin_gonic_gin.EnableJsonDecoderUseNumber
Api["github.com/gin-gonic/gin"]["ForceConsoleColor"] = github_com_gin_gonic_gin.ForceConsoleColor
Api["github.com/gin-gonic/gin"]["IsDebugging"] = github_com_gin_gonic_gin.IsDebugging
Api["github.com/gin-gonic/gin"]["Mode"] = github_com_gin_gonic_gin.Mode
Api["github.com/gin-gonic/gin"]["SetMode"] = github_com_gin_gonic_gin.SetMode
Api["github.com/gin-gonic/gin"]["Context"] = github_com_gin_gonic_gin.Context{}
Api["github.com/gin-gonic/gin/new"]["Context"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.Context)
}
Api["github.com/gin-gonic/gin"]["CreateTestContextOnly"] = github_com_gin_gonic_gin.CreateTestContextOnly
Api["github.com/gin-gonic/gin"]["Engine"] = github_com_gin_gonic_gin.Engine{}
Api["github.com/gin-gonic/gin/new"]["Engine"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.Engine)
}
Api["github.com/gin-gonic/gin"]["Default"] = github_com_gin_gonic_gin.Default
Api["github.com/gin-gonic/gin"]["New"] = github_com_gin_gonic_gin.New
Api["github.com/gin-gonic/gin"]["Error"] = github_com_gin_gonic_gin.Error{}
Api["github.com/gin-gonic/gin/new"]["Error"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.Error)
}
Api["github.com/gin-gonic/gin"]["ErrorTypeBind"] = github_com_gin_gonic_gin.ErrorTypeBind
Api["github.com/gin-gonic/gin"]["ErrorTypeRender"] = github_com_gin_gonic_gin.ErrorTypeRender
Api["github.com/gin-gonic/gin"]["ErrorTypePrivate"] = github_com_gin_gonic_gin.ErrorTypePrivate
Api["github.com/gin-gonic/gin"]["ErrorTypePublic"] = github_com_gin_gonic_gin.ErrorTypePublic
Api["github.com/gin-gonic/gin"]["ErrorTypeAny"] = github_com_gin_gonic_gin.ErrorTypeAny
Api["github.com/gin-gonic/gin"]["ErrorTypeNu"] = github_com_gin_gonic_gin.ErrorTypeNu
Api["github.com/gin-gonic/gin"]["BasicAuth"] = github_com_gin_gonic_gin.BasicAuth
Api["github.com/gin-gonic/gin"]["BasicAuthForRealm"] = github_com_gin_gonic_gin.BasicAuthForRealm
Api["github.com/gin-gonic/gin"]["Bind"] = github_com_gin_gonic_gin.Bind
Api["github.com/gin-gonic/gin"]["CustomRecovery"] = github_com_gin_gonic_gin.CustomRecovery
Api["github.com/gin-gonic/gin"]["CustomRecoveryWithWriter"] = github_com_gin_gonic_gin.CustomRecoveryWithWriter
Api["github.com/gin-gonic/gin"]["ErrorLogger"] = github_com_gin_gonic_gin.ErrorLogger
Api["github.com/gin-gonic/gin"]["ErrorLoggerT"] = github_com_gin_gonic_gin.ErrorLoggerT
Api["github.com/gin-gonic/gin"]["Logger"] = github_com_gin_gonic_gin.Logger
Api["github.com/gin-gonic/gin"]["LoggerWithConfig"] = github_com_gin_gonic_gin.LoggerWithConfig
Api["github.com/gin-gonic/gin"]["LoggerWithFormatter"] = github_com_gin_gonic_gin.LoggerWithFormatter
Api["github.com/gin-gonic/gin"]["LoggerWithWriter"] = github_com_gin_gonic_gin.LoggerWithWriter
Api["github.com/gin-gonic/gin"]["Recovery"] = github_com_gin_gonic_gin.Recovery
Api["github.com/gin-gonic/gin"]["RecoveryWithWriter"] = github_com_gin_gonic_gin.RecoveryWithWriter
Api["github.com/gin-gonic/gin"]["WrapF"] = github_com_gin_gonic_gin.WrapF
Api["github.com/gin-gonic/gin"]["WrapH"] = github_com_gin_gonic_gin.WrapH
Api["github.com/gin-gonic/gin"]["LogFormatterParams"] = github_com_gin_gonic_gin.LogFormatterParams{}
Api["github.com/gin-gonic/gin/new"]["LogFormatterParams"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.LogFormatterParams)
}
Api["github.com/gin-gonic/gin"]["LoggerConfig"] = github_com_gin_gonic_gin.LoggerConfig{}
Api["github.com/gin-gonic/gin/new"]["LoggerConfig"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.LoggerConfig)
}
Api["github.com/gin-gonic/gin"]["Negotiate"] = github_com_gin_gonic_gin.Negotiate{}
Api["github.com/gin-gonic/gin/new"]["Negotiate"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.Negotiate)
}
Api["github.com/gin-gonic/gin"]["Param"] = github_com_gin_gonic_gin.Param{}
Api["github.com/gin-gonic/gin/new"]["Param"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.Param)
}
Api["github.com/gin-gonic/gin"]["RouteInfo"] = github_com_gin_gonic_gin.RouteInfo{}
Api["github.com/gin-gonic/gin/new"]["RouteInfo"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.RouteInfo)
}
Api["github.com/gin-gonic/gin"]["RouterGroup"] = github_com_gin_gonic_gin.RouterGroup{}
Api["github.com/gin-gonic/gin/new"]["RouterGroup"] = func(args ...interface{}) interface{} {
					return new(github_com_gin_gonic_gin.RouterGroup)
}
}
