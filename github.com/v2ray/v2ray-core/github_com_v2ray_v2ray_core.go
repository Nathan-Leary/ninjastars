package v2ray_core
import (
github_com_v2ray_v2ray_core "github.com/v2ray/v2ray-core"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/v2ray/v2ray-core"]; !ok {
   Api["github.com/v2ray/v2ray-core"] = map[string]interface{}{}
}
if _, ok := Api["github.com/v2ray/v2ray-core/new"]; !ok {
	Api["github.com/v2ray/v2ray-core/new"] = map[string]interface{}{}
}
Api["github.com/v2ray/v2ray-core"]["AddInboundHandler"] = github_com_v2ray_v2ray_core.AddInboundHandler
Api["github.com/v2ray/v2ray-core"]["AddOutboundHandler"] = github_com_v2ray_v2ray_core.AddOutboundHandler
Api["github.com/v2ray/v2ray-core"]["CreateObject"] = github_com_v2ray_v2ray_core.CreateObject
Api["github.com/v2ray/v2ray-core"]["Dial"] = github_com_v2ray_v2ray_core.Dial
Api["github.com/v2ray/v2ray-core"]["DialUDP"] = github_com_v2ray_v2ray_core.DialUDP
Api["github.com/v2ray/v2ray-core"]["RegisterConfigLoader"] = github_com_v2ray_v2ray_core.RegisterConfigLoader
Api["github.com/v2ray/v2ray-core"]["RequireFeatures"] = github_com_v2ray_v2ray_core.RequireFeatures
Api["github.com/v2ray/v2ray-core"]["ServerType"] = github_com_v2ray_v2ray_core.ServerType
Api["github.com/v2ray/v2ray-core"]["Version"] = github_com_v2ray_v2ray_core.Version
Api["github.com/v2ray/v2ray-core"]["VersionStatement"] = github_com_v2ray_v2ray_core.VersionStatement
Api["github.com/v2ray/v2ray-core"]["Annotation"] = github_com_v2ray_v2ray_core.Annotation{}
Api["github.com/v2ray/v2ray-core/new"]["Annotation"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.Annotation)
}
Api["github.com/v2ray/v2ray-core"]["Config"] = github_com_v2ray_v2ray_core.Config{}
Api["github.com/v2ray/v2ray-core/new"]["Config"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.Config)
}
Api["github.com/v2ray/v2ray-core"]["LoadConfig"] = github_com_v2ray_v2ray_core.LoadConfig
Api["github.com/v2ray/v2ray-core"]["ConfigFormat"] = github_com_v2ray_v2ray_core.ConfigFormat{}
Api["github.com/v2ray/v2ray-core/new"]["ConfigFormat"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.ConfigFormat)
}
Api["github.com/v2ray/v2ray-core"]["InboundHandlerConfig"] = github_com_v2ray_v2ray_core.InboundHandlerConfig{}
Api["github.com/v2ray/v2ray-core/new"]["InboundHandlerConfig"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.InboundHandlerConfig)
}
Api["github.com/v2ray/v2ray-core"]["Instance"] = github_com_v2ray_v2ray_core.Instance{}
Api["github.com/v2ray/v2ray-core/new"]["Instance"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.Instance)
}
Api["github.com/v2ray/v2ray-core"]["FromContext"] = github_com_v2ray_v2ray_core.FromContext
Api["github.com/v2ray/v2ray-core"]["MustFromContext"] = github_com_v2ray_v2ray_core.MustFromContext
Api["github.com/v2ray/v2ray-core"]["StartInstance"] = github_com_v2ray_v2ray_core.StartInstance
Api["github.com/v2ray/v2ray-core"]["OutboundHandlerConfig"] = github_com_v2ray_v2ray_core.OutboundHandlerConfig{}
Api["github.com/v2ray/v2ray-core/new"]["OutboundHandlerConfig"] = func(args ...interface{}) interface{} {
					return new(github_com_v2ray_v2ray_core.OutboundHandlerConfig)
}
}
