package pq
import (
github_com_lib_pq "github.com/lib/pq"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/lib/pq"]; !ok {
   Api["github.com/lib/pq"] = map[string]interface{}{}
}
if _, ok := Api["github.com/lib/pq/new"]; !ok {
	Api["github.com/lib/pq/new"] = map[string]interface{}{}
}
Api["github.com/lib/pq"]["Efatal"] = github_com_lib_pq.Efatal
Api["github.com/lib/pq"]["Epanic"] = github_com_lib_pq.Epanic
Api["github.com/lib/pq"]["Ewarning"] = github_com_lib_pq.Ewarning
Api["github.com/lib/pq"]["Enotice"] = github_com_lib_pq.Enotice
Api["github.com/lib/pq"]["Edebug"] = github_com_lib_pq.Edebug
Api["github.com/lib/pq"]["Einfo"] = github_com_lib_pq.Einfo
Api["github.com/lib/pq"]["Elog"] = github_com_lib_pq.Elog
Api["github.com/lib/pq"]["ErrChannelAlreadyOpen"] = github_com_lib_pq.ErrChannelAlreadyOpen
Api["github.com/lib/pq"]["ErrChannelNotOpen"] = github_com_lib_pq.ErrChannelNotOpen
Api["github.com/lib/pq"]["Array"] = github_com_lib_pq.Array
Api["github.com/lib/pq"]["BufferQuoteIdentifier"] = github_com_lib_pq.BufferQuoteIdentifier
Api["github.com/lib/pq"]["ConnectorNoticeHandler"] = github_com_lib_pq.ConnectorNoticeHandler
Api["github.com/lib/pq"]["ConnectorNotificationHandler"] = github_com_lib_pq.ConnectorNotificationHandler
Api["github.com/lib/pq"]["CopyIn"] = github_com_lib_pq.CopyIn
Api["github.com/lib/pq"]["CopyInSchema"] = github_com_lib_pq.CopyInSchema
Api["github.com/lib/pq"]["DialOpen"] = github_com_lib_pq.DialOpen
Api["github.com/lib/pq"]["EnableInfinityTs"] = github_com_lib_pq.EnableInfinityTs
Api["github.com/lib/pq"]["FormatTimestamp"] = github_com_lib_pq.FormatTimestamp
Api["github.com/lib/pq"]["NoticeHandler"] = github_com_lib_pq.NoticeHandler
Api["github.com/lib/pq"]["Open"] = github_com_lib_pq.Open
Api["github.com/lib/pq"]["ParseTimestamp"] = github_com_lib_pq.ParseTimestamp
Api["github.com/lib/pq"]["ParseURL"] = github_com_lib_pq.ParseURL
Api["github.com/lib/pq"]["QuoteIdentifier"] = github_com_lib_pq.QuoteIdentifier
Api["github.com/lib/pq"]["QuoteLiteral"] = github_com_lib_pq.QuoteLiteral
Api["github.com/lib/pq"]["RegisterGSSProvider"] = github_com_lib_pq.RegisterGSSProvider
Api["github.com/lib/pq"]["SetNoticeHandler"] = github_com_lib_pq.SetNoticeHandler
Api["github.com/lib/pq"]["SetNotificationHandler"] = github_com_lib_pq.SetNotificationHandler
Api["github.com/lib/pq"]["Connector"] = github_com_lib_pq.Connector{}
Api["github.com/lib/pq/new"]["Connector"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.Connector)
}
Api["github.com/lib/pq"]["NewConnector"] = github_com_lib_pq.NewConnector
Api["github.com/lib/pq"]["Driver"] = github_com_lib_pq.Driver{}
Api["github.com/lib/pq/new"]["Driver"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.Driver)
}
Api["github.com/lib/pq"]["GenericArray"] = github_com_lib_pq.GenericArray{}
Api["github.com/lib/pq/new"]["GenericArray"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.GenericArray)
}
Api["github.com/lib/pq"]["NewDialListener"] = github_com_lib_pq.NewDialListener
Api["github.com/lib/pq"]["NewListener"] = github_com_lib_pq.NewListener
Api["github.com/lib/pq"]["ListenerConn"] = github_com_lib_pq.ListenerConn{}
Api["github.com/lib/pq/new"]["ListenerConn"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.ListenerConn)
}
Api["github.com/lib/pq"]["NewListenerConn"] = github_com_lib_pq.NewListenerConn
Api["github.com/lib/pq"]["ListenerEventConnected"] = github_com_lib_pq.ListenerEventConnected
Api["github.com/lib/pq"]["ListenerEventDisconnected"] = github_com_lib_pq.ListenerEventDisconnected
Api["github.com/lib/pq"]["ListenerEventReconnected"] = github_com_lib_pq.ListenerEventReconnected
Api["github.com/lib/pq"]["ListenerEventConnectionAttemptFailed"] = github_com_lib_pq.ListenerEventConnectionAttemptFailed
Api["github.com/lib/pq"]["NoticeHandlerConnector"] = github_com_lib_pq.NoticeHandlerConnector{}
Api["github.com/lib/pq/new"]["NoticeHandlerConnector"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.NoticeHandlerConnector)
}
Api["github.com/lib/pq"]["ConnectorWithNoticeHandler"] = github_com_lib_pq.ConnectorWithNoticeHandler
Api["github.com/lib/pq"]["Notification"] = github_com_lib_pq.Notification{}
Api["github.com/lib/pq/new"]["Notification"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.Notification)
}
Api["github.com/lib/pq"]["NotificationHandlerConnector"] = github_com_lib_pq.NotificationHandlerConnector{}
Api["github.com/lib/pq/new"]["NotificationHandlerConnector"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.NotificationHandlerConnector)
}
Api["github.com/lib/pq"]["ConnectorWithNotificationHandler"] = github_com_lib_pq.ConnectorWithNotificationHandler
Api["github.com/lib/pq"]["NullTime"] = github_com_lib_pq.NullTime{}
Api["github.com/lib/pq/new"]["NullTime"] = func(args ...interface{}) interface{} {
					return new(github_com_lib_pq.NullTime)
}
}
