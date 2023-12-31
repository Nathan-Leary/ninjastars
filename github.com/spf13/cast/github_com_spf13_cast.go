package cast
import (
github_com_spf13_cast "github.com/spf13/cast"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/spf13/cast"]; !ok {
   Api["github.com/spf13/cast"] = map[string]interface{}{}
}
if _, ok := Api["github.com/spf13/cast/new"]; !ok {
	Api["github.com/spf13/cast/new"] = map[string]interface{}{}
}
Api["github.com/spf13/cast"]["StringToDate"] = github_com_spf13_cast.StringToDate
Api["github.com/spf13/cast"]["StringToDateInDefaultLocation"] = github_com_spf13_cast.StringToDateInDefaultLocation
Api["github.com/spf13/cast"]["ToBool"] = github_com_spf13_cast.ToBool
Api["github.com/spf13/cast"]["ToBoolE"] = github_com_spf13_cast.ToBoolE
Api["github.com/spf13/cast"]["ToBoolSlice"] = github_com_spf13_cast.ToBoolSlice
Api["github.com/spf13/cast"]["ToBoolSliceE"] = github_com_spf13_cast.ToBoolSliceE
Api["github.com/spf13/cast"]["ToDuration"] = github_com_spf13_cast.ToDuration
Api["github.com/spf13/cast"]["ToDurationE"] = github_com_spf13_cast.ToDurationE
Api["github.com/spf13/cast"]["ToDurationSlice"] = github_com_spf13_cast.ToDurationSlice
Api["github.com/spf13/cast"]["ToDurationSliceE"] = github_com_spf13_cast.ToDurationSliceE
Api["github.com/spf13/cast"]["ToFloat32"] = github_com_spf13_cast.ToFloat32
Api["github.com/spf13/cast"]["ToFloat32E"] = github_com_spf13_cast.ToFloat32E
Api["github.com/spf13/cast"]["ToFloat64"] = github_com_spf13_cast.ToFloat64
Api["github.com/spf13/cast"]["ToFloat64E"] = github_com_spf13_cast.ToFloat64E
Api["github.com/spf13/cast"]["ToInt"] = github_com_spf13_cast.ToInt
Api["github.com/spf13/cast"]["ToInt16"] = github_com_spf13_cast.ToInt16
Api["github.com/spf13/cast"]["ToInt16E"] = github_com_spf13_cast.ToInt16E
Api["github.com/spf13/cast"]["ToInt32"] = github_com_spf13_cast.ToInt32
Api["github.com/spf13/cast"]["ToInt32E"] = github_com_spf13_cast.ToInt32E
Api["github.com/spf13/cast"]["ToInt64"] = github_com_spf13_cast.ToInt64
Api["github.com/spf13/cast"]["ToInt64E"] = github_com_spf13_cast.ToInt64E
Api["github.com/spf13/cast"]["ToInt8"] = github_com_spf13_cast.ToInt8
Api["github.com/spf13/cast"]["ToInt8E"] = github_com_spf13_cast.ToInt8E
Api["github.com/spf13/cast"]["ToIntE"] = github_com_spf13_cast.ToIntE
Api["github.com/spf13/cast"]["ToIntSlice"] = github_com_spf13_cast.ToIntSlice
Api["github.com/spf13/cast"]["ToIntSliceE"] = github_com_spf13_cast.ToIntSliceE
Api["github.com/spf13/cast"]["ToSlice"] = github_com_spf13_cast.ToSlice
Api["github.com/spf13/cast"]["ToSliceE"] = github_com_spf13_cast.ToSliceE
Api["github.com/spf13/cast"]["ToString"] = github_com_spf13_cast.ToString
Api["github.com/spf13/cast"]["ToStringE"] = github_com_spf13_cast.ToStringE
Api["github.com/spf13/cast"]["ToStringMap"] = github_com_spf13_cast.ToStringMap
Api["github.com/spf13/cast"]["ToStringMapBool"] = github_com_spf13_cast.ToStringMapBool
Api["github.com/spf13/cast"]["ToStringMapBoolE"] = github_com_spf13_cast.ToStringMapBoolE
Api["github.com/spf13/cast"]["ToStringMapE"] = github_com_spf13_cast.ToStringMapE
Api["github.com/spf13/cast"]["ToStringMapInt"] = github_com_spf13_cast.ToStringMapInt
Api["github.com/spf13/cast"]["ToStringMapInt64"] = github_com_spf13_cast.ToStringMapInt64
Api["github.com/spf13/cast"]["ToStringMapInt64E"] = github_com_spf13_cast.ToStringMapInt64E
Api["github.com/spf13/cast"]["ToStringMapIntE"] = github_com_spf13_cast.ToStringMapIntE
Api["github.com/spf13/cast"]["ToStringMapString"] = github_com_spf13_cast.ToStringMapString
Api["github.com/spf13/cast"]["ToStringMapStringE"] = github_com_spf13_cast.ToStringMapStringE
Api["github.com/spf13/cast"]["ToStringMapStringSlice"] = github_com_spf13_cast.ToStringMapStringSlice
Api["github.com/spf13/cast"]["ToStringMapStringSliceE"] = github_com_spf13_cast.ToStringMapStringSliceE
Api["github.com/spf13/cast"]["ToStringSlice"] = github_com_spf13_cast.ToStringSlice
Api["github.com/spf13/cast"]["ToStringSliceE"] = github_com_spf13_cast.ToStringSliceE
Api["github.com/spf13/cast"]["ToTime"] = github_com_spf13_cast.ToTime
Api["github.com/spf13/cast"]["ToTimeE"] = github_com_spf13_cast.ToTimeE
Api["github.com/spf13/cast"]["ToTimeInDefaultLocation"] = github_com_spf13_cast.ToTimeInDefaultLocation
Api["github.com/spf13/cast"]["ToTimeInDefaultLocationE"] = github_com_spf13_cast.ToTimeInDefaultLocationE
Api["github.com/spf13/cast"]["ToUint"] = github_com_spf13_cast.ToUint
Api["github.com/spf13/cast"]["ToUint16"] = github_com_spf13_cast.ToUint16
Api["github.com/spf13/cast"]["ToUint16E"] = github_com_spf13_cast.ToUint16E
Api["github.com/spf13/cast"]["ToUint32"] = github_com_spf13_cast.ToUint32
Api["github.com/spf13/cast"]["ToUint32E"] = github_com_spf13_cast.ToUint32E
Api["github.com/spf13/cast"]["ToUint64"] = github_com_spf13_cast.ToUint64
Api["github.com/spf13/cast"]["ToUint64E"] = github_com_spf13_cast.ToUint64E
Api["github.com/spf13/cast"]["ToUint8"] = github_com_spf13_cast.ToUint8
Api["github.com/spf13/cast"]["ToUint8E"] = github_com_spf13_cast.ToUint8E
Api["github.com/spf13/cast"]["ToUintE"] = github_com_spf13_cast.ToUintE
}
