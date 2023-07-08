package influxdb
import (
github_com_influxdata_influxdb "github.com/influxdata/influxdb"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/influxdata/influxdb"]; !ok {
   Api["github.com/influxdata/influxdb"] = map[string]interface{}{}
}
if _, ok := Api["github.com/influxdata/influxdb/new"]; !ok {
	Api["github.com/influxdata/influxdb/new"] = map[string]interface{}{}
}
Api["github.com/influxdata/influxdb"]["ErrFieldTypeConflict"] = github_com_influxdata_influxdb.ErrFieldTypeConflict
Api["github.com/influxdata/influxdb"]["ErrDatabaseNotFound"] = github_com_influxdata_influxdb.ErrDatabaseNotFound
Api["github.com/influxdata/influxdb"]["ErrRetentionPolicyNotFound"] = github_com_influxdata_influxdb.ErrRetentionPolicyNotFound
Api["github.com/influxdata/influxdb"]["IsAuthorizationError"] = github_com_influxdata_influxdb.IsAuthorizationError
Api["github.com/influxdata/influxdb"]["IsClientError"] = github_com_influxdata_influxdb.IsClientError
Api["github.com/influxdata/influxdb"]["Node"] = github_com_influxdata_influxdb.Node{}
Api["github.com/influxdata/influxdb/new"]["Node"] = func(args ...interface{}) interface{} {
					return new(github_com_influxdata_influxdb.Node)
}
Api["github.com/influxdata/influxdb"]["LoadNode"] = github_com_influxdata_influxdb.LoadNode
Api["github.com/influxdata/influxdb"]["NewNode"] = github_com_influxdata_influxdb.NewNode
}
