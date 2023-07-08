package goja
import (
github_com_dop251_goja "github.com/dop251/goja"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/dop251/goja"]; !ok {
   Api["github.com/dop251/goja"] = map[string]interface{}{}
}
if _, ok := Api["github.com/dop251/goja/new"]; !ok {
	Api["github.com/dop251/goja/new"] = map[string]interface{}{}
}
Api["github.com/dop251/goja"]["_"] = github_com_dop251_goja._
Api["github.com/dop251/goja"]["IsInfinity"] = github_com_dop251_goja.IsInfinity
Api["github.com/dop251/goja"]["IsNaN"] = github_com_dop251_goja.IsNaN
Api["github.com/dop251/goja"]["IsNull"] = github_com_dop251_goja.IsNull
Api["github.com/dop251/goja"]["IsUndefined"] = github_com_dop251_goja.IsUndefined
Api["github.com/dop251/goja"]["Parse"] = github_com_dop251_goja.Parse
Api["github.com/dop251/goja"]["StartProfile"] = github_com_dop251_goja.StartProfile
Api["github.com/dop251/goja"]["StopProfile"] = github_com_dop251_goja.StopProfile
Api["github.com/dop251/goja"]["ArrayBuffer"] = github_com_dop251_goja.ArrayBuffer{}
Api["github.com/dop251/goja/new"]["ArrayBuffer"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.ArrayBuffer)
}
Api["github.com/dop251/goja"]["AssertFunction"] = github_com_dop251_goja.AssertFunction
Api["github.com/dop251/goja"]["CompilerError"] = github_com_dop251_goja.CompilerError{}
Api["github.com/dop251/goja/new"]["CompilerError"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.CompilerError)
}
Api["github.com/dop251/goja"]["CompilerReferenceError"] = github_com_dop251_goja.CompilerReferenceError{}
Api["github.com/dop251/goja/new"]["CompilerReferenceError"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.CompilerReferenceError)
}
Api["github.com/dop251/goja"]["CompilerSyntaxError"] = github_com_dop251_goja.CompilerSyntaxError{}
Api["github.com/dop251/goja/new"]["CompilerSyntaxError"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.CompilerSyntaxError)
}
Api["github.com/dop251/goja"]["AssertConstructor"] = github_com_dop251_goja.AssertConstructor
Api["github.com/dop251/goja"]["ConstructorCall"] = github_com_dop251_goja.ConstructorCall{}
Api["github.com/dop251/goja/new"]["ConstructorCall"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.ConstructorCall)
}
Api["github.com/dop251/goja"]["Exception"] = github_com_dop251_goja.Exception{}
Api["github.com/dop251/goja/new"]["Exception"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Exception)
}
Api["github.com/dop251/goja"]["FieldNameMapper"] = new(github_com_dop251_goja.FieldNameMapper)
Api["github.com/dop251/goja"]["TagFieldNameMapper"] = github_com_dop251_goja.TagFieldNameMapper
Api["github.com/dop251/goja"]["UncapFieldNameMapper"] = github_com_dop251_goja.UncapFieldNameMapper
Api["github.com/dop251/goja"]["FLAG_NOT_SET"] = github_com_dop251_goja.FLAG_NOT_SET
Api["github.com/dop251/goja"]["FLAG_FALSE"] = github_com_dop251_goja.FLAG_FALSE
Api["github.com/dop251/goja"]["FLAG_TRUE"] = github_com_dop251_goja.FLAG_TRUE
Api["github.com/dop251/goja"]["ToFlag"] = github_com_dop251_goja.ToFlag
Api["github.com/dop251/goja"]["FunctionCall"] = github_com_dop251_goja.FunctionCall{}
Api["github.com/dop251/goja/new"]["FunctionCall"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.FunctionCall)
}
Api["github.com/dop251/goja"]["InterruptedError"] = github_com_dop251_goja.InterruptedError{}
Api["github.com/dop251/goja/new"]["InterruptedError"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.InterruptedError)
}
Api["github.com/dop251/goja"]["Object"] = github_com_dop251_goja.Object{}
Api["github.com/dop251/goja/new"]["Object"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Object)
}
Api["github.com/dop251/goja"]["NewSharedDynamicArray"] = github_com_dop251_goja.NewSharedDynamicArray
Api["github.com/dop251/goja"]["NewSharedDynamicObject"] = github_com_dop251_goja.NewSharedDynamicObject
Api["github.com/dop251/goja"]["Program"] = github_com_dop251_goja.Program{}
Api["github.com/dop251/goja/new"]["Program"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Program)
}
Api["github.com/dop251/goja"]["Compile"] = github_com_dop251_goja.Compile
Api["github.com/dop251/goja"]["CompileAST"] = github_com_dop251_goja.CompileAST
Api["github.com/dop251/goja"]["MustCompile"] = github_com_dop251_goja.MustCompile
Api["github.com/dop251/goja"]["Promise"] = github_com_dop251_goja.Promise{}
Api["github.com/dop251/goja/new"]["Promise"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Promise)
}
Api["github.com/dop251/goja"]["PromiseRejectionReject"] = github_com_dop251_goja.PromiseRejectionReject
Api["github.com/dop251/goja"]["PromiseRejectionHandle"] = github_com_dop251_goja.PromiseRejectionHandle
Api["github.com/dop251/goja"]["PromiseStatePending"] = github_com_dop251_goja.PromiseStatePending
Api["github.com/dop251/goja"]["PromiseStateFulfilled"] = github_com_dop251_goja.PromiseStateFulfilled
Api["github.com/dop251/goja"]["PromiseStateRejected"] = github_com_dop251_goja.PromiseStateRejected
Api["github.com/dop251/goja"]["PropertyDescriptor"] = github_com_dop251_goja.PropertyDescriptor{}
Api["github.com/dop251/goja/new"]["PropertyDescriptor"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.PropertyDescriptor)
}
Api["github.com/dop251/goja"]["Proxy"] = github_com_dop251_goja.Proxy{}
Api["github.com/dop251/goja/new"]["Proxy"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Proxy)
}
Api["github.com/dop251/goja"]["ProxyTrapConfig"] = github_com_dop251_goja.ProxyTrapConfig{}
Api["github.com/dop251/goja/new"]["ProxyTrapConfig"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.ProxyTrapConfig)
}
Api["github.com/dop251/goja"]["Runtime"] = github_com_dop251_goja.Runtime{}
Api["github.com/dop251/goja/new"]["Runtime"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Runtime)
}
Api["github.com/dop251/goja"]["StackFrame"] = github_com_dop251_goja.StackFrame{}
Api["github.com/dop251/goja/new"]["StackFrame"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.StackFrame)
}
Api["github.com/dop251/goja"]["StackOverflowError"] = github_com_dop251_goja.StackOverflowError{}
Api["github.com/dop251/goja/new"]["StackOverflowError"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.StackOverflowError)
}
Api["github.com/dop251/goja"]["Symbol"] = github_com_dop251_goja.Symbol{}
Api["github.com/dop251/goja/new"]["Symbol"] = func(args ...interface{}) interface{} {
					return new(github_com_dop251_goja.Symbol)
}
Api["github.com/dop251/goja"]["NewSymbol"] = github_com_dop251_goja.NewSymbol
Api["github.com/dop251/goja"]["NaN"] = github_com_dop251_goja.NaN
Api["github.com/dop251/goja"]["NegativeInf"] = github_com_dop251_goja.NegativeInf
Api["github.com/dop251/goja"]["Null"] = github_com_dop251_goja.Null
Api["github.com/dop251/goja"]["PositiveInf"] = github_com_dop251_goja.PositiveInf
Api["github.com/dop251/goja"]["Undefined"] = github_com_dop251_goja.Undefined
}
