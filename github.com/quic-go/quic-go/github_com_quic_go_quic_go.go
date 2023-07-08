package quic_go
import (
github_com_quic_go_quic_go "github.com/quic-go/quic-go"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/quic-go/quic-go"]; !ok {
   Api["github.com/quic-go/quic-go"] = map[string]interface{}{}
}
if _, ok := Api["github.com/quic-go/quic-go/new"]; !ok {
	Api["github.com/quic-go/quic-go/new"] = map[string]interface{}{}
}
Api["github.com/quic-go/quic-go"]["NoError"] = github_com_quic_go_quic_go.NoError
Api["github.com/quic-go/quic-go"]["InternalError"] = github_com_quic_go_quic_go.InternalError
Api["github.com/quic-go/quic-go"]["ConnectionRefused"] = github_com_quic_go_quic_go.ConnectionRefused
Api["github.com/quic-go/quic-go"]["FlowControlError"] = github_com_quic_go_quic_go.FlowControlError
Api["github.com/quic-go/quic-go"]["StreamLimitError"] = github_com_quic_go_quic_go.StreamLimitError
Api["github.com/quic-go/quic-go"]["StreamStateError"] = github_com_quic_go_quic_go.StreamStateError
Api["github.com/quic-go/quic-go"]["FinalSizeError"] = github_com_quic_go_quic_go.FinalSizeError
Api["github.com/quic-go/quic-go"]["FrameEncodingError"] = github_com_quic_go_quic_go.FrameEncodingError
Api["github.com/quic-go/quic-go"]["TransportParameterError"] = github_com_quic_go_quic_go.TransportParameterError
Api["github.com/quic-go/quic-go"]["ConnectionIDLimitError"] = github_com_quic_go_quic_go.ConnectionIDLimitError
Api["github.com/quic-go/quic-go"]["ProtocolViolation"] = github_com_quic_go_quic_go.ProtocolViolation
Api["github.com/quic-go/quic-go"]["InvalidToken"] = github_com_quic_go_quic_go.InvalidToken
Api["github.com/quic-go/quic-go"]["ApplicationErrorErrorCode"] = github_com_quic_go_quic_go.ApplicationErrorErrorCode
Api["github.com/quic-go/quic-go"]["CryptoBufferExceeded"] = github_com_quic_go_quic_go.CryptoBufferExceeded
Api["github.com/quic-go/quic-go"]["KeyUpdateError"] = github_com_quic_go_quic_go.KeyUpdateError
Api["github.com/quic-go/quic-go"]["AEADLimitReached"] = github_com_quic_go_quic_go.AEADLimitReached
Api["github.com/quic-go/quic-go"]["NoViablePathError"] = github_com_quic_go_quic_go.NoViablePathError
Api["github.com/quic-go/quic-go"]["VersionDraft29"] = github_com_quic_go_quic_go.VersionDraft29
Api["github.com/quic-go/quic-go"]["Version1"] = github_com_quic_go_quic_go.Version1
Api["github.com/quic-go/quic-go"]["Version2"] = github_com_quic_go_quic_go.Version2
Api["github.com/quic-go/quic-go"]["ConnectionTracingKey"] = github_com_quic_go_quic_go.ConnectionTracingKey
Api["github.com/quic-go/quic-go"]["Err0RTTRejected"] = github_com_quic_go_quic_go.Err0RTTRejected
Api["github.com/quic-go/quic-go"]["ErrServerClosed"] = github_com_quic_go_quic_go.ErrServerClosed
Api["github.com/quic-go/quic-go"]["QUICVersionContextKey"] = github_com_quic_go_quic_go.QUICVersionContextKey
Api["github.com/quic-go/quic-go"]["OptimizeConn"] = github_com_quic_go_quic_go.OptimizeConn
Api["github.com/quic-go/quic-go"]["ClientHelloInfo"] = github_com_quic_go_quic_go.ClientHelloInfo{}
Api["github.com/quic-go/quic-go/new"]["ClientHelloInfo"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.ClientHelloInfo)
}
Api["github.com/quic-go/quic-go"]["ClientToken"] = github_com_quic_go_quic_go.ClientToken{}
Api["github.com/quic-go/quic-go/new"]["ClientToken"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.ClientToken)
}
Api["github.com/quic-go/quic-go"]["DialAddr"] = github_com_quic_go_quic_go.DialAddr
Api["github.com/quic-go/quic-go"]["ConnectionIDFromBytes"] = github_com_quic_go_quic_go.ConnectionIDFromBytes
Api["github.com/quic-go/quic-go"]["ConnectionState"] = github_com_quic_go_quic_go.ConnectionState{}
Api["github.com/quic-go/quic-go/new"]["ConnectionState"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.ConnectionState)
}
Api["github.com/quic-go/quic-go"]["EarlyConnection"] = new(github_com_quic_go_quic_go.EarlyConnection)
Api["github.com/quic-go/quic-go"]["DialAddrEarly"] = github_com_quic_go_quic_go.DialAddrEarly
Api["github.com/quic-go/quic-go"]["DialEarly"] = github_com_quic_go_quic_go.DialEarly
Api["github.com/quic-go/quic-go"]["EarlyListener"] = github_com_quic_go_quic_go.EarlyListener{}
Api["github.com/quic-go/quic-go/new"]["EarlyListener"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.EarlyListener)
}
Api["github.com/quic-go/quic-go"]["ListenAddrEarly"] = github_com_quic_go_quic_go.ListenAddrEarly
Api["github.com/quic-go/quic-go"]["ListenEarly"] = github_com_quic_go_quic_go.ListenEarly
Api["github.com/quic-go/quic-go"]["Listener"] = github_com_quic_go_quic_go.Listener{}
Api["github.com/quic-go/quic-go/new"]["Listener"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.Listener)
}
Api["github.com/quic-go/quic-go"]["Listen"] = github_com_quic_go_quic_go.Listen
Api["github.com/quic-go/quic-go"]["ListenAddr"] = github_com_quic_go_quic_go.ListenAddr
Api["github.com/quic-go/quic-go"]["StreamError"] = github_com_quic_go_quic_go.StreamError{}
Api["github.com/quic-go/quic-go/new"]["StreamError"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.StreamError)
}
Api["github.com/quic-go/quic-go"]["NewLRUTokenStore"] = github_com_quic_go_quic_go.NewLRUTokenStore
Api["github.com/quic-go/quic-go"]["Transport"] = github_com_quic_go_quic_go.Transport{}
Api["github.com/quic-go/quic-go/new"]["Transport"] = func(args ...interface{}) interface{} {
					return new(github_com_quic_go_quic_go.Transport)
}
}
