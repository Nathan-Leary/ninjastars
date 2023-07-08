package fiber
import (
github_com_gofiber_fiber "github.com/gofiber/fiber"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/gofiber/fiber"]; !ok {
   Api["github.com/gofiber/fiber"] = map[string]interface{}{}
}
if _, ok := Api["github.com/gofiber/fiber/new"]; !ok {
	Api["github.com/gofiber/fiber/new"] = map[string]interface{}{}
}
Api["github.com/gofiber/fiber"]["MethodGet"] = github_com_gofiber_fiber.MethodGet
Api["github.com/gofiber/fiber"]["MethodHead"] = github_com_gofiber_fiber.MethodHead
Api["github.com/gofiber/fiber"]["MethodPost"] = github_com_gofiber_fiber.MethodPost
Api["github.com/gofiber/fiber"]["MethodPut"] = github_com_gofiber_fiber.MethodPut
Api["github.com/gofiber/fiber"]["MethodPatch"] = github_com_gofiber_fiber.MethodPatch
Api["github.com/gofiber/fiber"]["MethodDelete"] = github_com_gofiber_fiber.MethodDelete
Api["github.com/gofiber/fiber"]["MethodConnect"] = github_com_gofiber_fiber.MethodConnect
Api["github.com/gofiber/fiber"]["MethodOptions"] = github_com_gofiber_fiber.MethodOptions
Api["github.com/gofiber/fiber"]["MethodTrace"] = github_com_gofiber_fiber.MethodTrace
Api["github.com/gofiber/fiber"]["MIMETextXML"] = github_com_gofiber_fiber.MIMETextXML
Api["github.com/gofiber/fiber"]["MIMETextHTML"] = github_com_gofiber_fiber.MIMETextHTML
Api["github.com/gofiber/fiber"]["MIMETextPlain"] = github_com_gofiber_fiber.MIMETextPlain
Api["github.com/gofiber/fiber"]["MIMEApplicationXML"] = github_com_gofiber_fiber.MIMEApplicationXML
Api["github.com/gofiber/fiber"]["MIMEApplicationJSON"] = github_com_gofiber_fiber.MIMEApplicationJSON
Api["github.com/gofiber/fiber"]["MIMEApplicationJavaScript"] = github_com_gofiber_fiber.MIMEApplicationJavaScript
Api["github.com/gofiber/fiber"]["MIMEApplicationForm"] = github_com_gofiber_fiber.MIMEApplicationForm
Api["github.com/gofiber/fiber"]["MIMEOctetStream"] = github_com_gofiber_fiber.MIMEOctetStream
Api["github.com/gofiber/fiber"]["MIMEMultipartForm"] = github_com_gofiber_fiber.MIMEMultipartForm
Api["github.com/gofiber/fiber"]["MIMETextXMLCharsetUTF8"] = github_com_gofiber_fiber.MIMETextXMLCharsetUTF8
Api["github.com/gofiber/fiber"]["MIMETextHTMLCharsetUTF8"] = github_com_gofiber_fiber.MIMETextHTMLCharsetUTF8
Api["github.com/gofiber/fiber"]["MIMETextPlainCharsetUTF8"] = github_com_gofiber_fiber.MIMETextPlainCharsetUTF8
Api["github.com/gofiber/fiber"]["MIMEApplicationXMLCharsetUTF8"] = github_com_gofiber_fiber.MIMEApplicationXMLCharsetUTF8
Api["github.com/gofiber/fiber"]["MIMEApplicationJSONCharsetUTF8"] = github_com_gofiber_fiber.MIMEApplicationJSONCharsetUTF8
Api["github.com/gofiber/fiber"]["MIMEApplicationJavaScriptCharsetUTF8"] = github_com_gofiber_fiber.MIMEApplicationJavaScriptCharsetUTF8
Api["github.com/gofiber/fiber"]["StatusContinue"] = github_com_gofiber_fiber.StatusContinue
Api["github.com/gofiber/fiber"]["StatusSwitchingProtocols"] = github_com_gofiber_fiber.StatusSwitchingProtocols
Api["github.com/gofiber/fiber"]["StatusProcessing"] = github_com_gofiber_fiber.StatusProcessing
Api["github.com/gofiber/fiber"]["StatusEarlyHints"] = github_com_gofiber_fiber.StatusEarlyHints
Api["github.com/gofiber/fiber"]["StatusOK"] = github_com_gofiber_fiber.StatusOK
Api["github.com/gofiber/fiber"]["StatusCreated"] = github_com_gofiber_fiber.StatusCreated
Api["github.com/gofiber/fiber"]["StatusAccepted"] = github_com_gofiber_fiber.StatusAccepted
Api["github.com/gofiber/fiber"]["StatusNonAuthoritativeInfo"] = github_com_gofiber_fiber.StatusNonAuthoritativeInfo
Api["github.com/gofiber/fiber"]["StatusNoContent"] = github_com_gofiber_fiber.StatusNoContent
Api["github.com/gofiber/fiber"]["StatusResetContent"] = github_com_gofiber_fiber.StatusResetContent
Api["github.com/gofiber/fiber"]["StatusPartialContent"] = github_com_gofiber_fiber.StatusPartialContent
Api["github.com/gofiber/fiber"]["StatusMultiStatus"] = github_com_gofiber_fiber.StatusMultiStatus
Api["github.com/gofiber/fiber"]["StatusAlreadyReported"] = github_com_gofiber_fiber.StatusAlreadyReported
Api["github.com/gofiber/fiber"]["StatusIMUsed"] = github_com_gofiber_fiber.StatusIMUsed
Api["github.com/gofiber/fiber"]["StatusMultipleChoices"] = github_com_gofiber_fiber.StatusMultipleChoices
Api["github.com/gofiber/fiber"]["StatusMovedPermanently"] = github_com_gofiber_fiber.StatusMovedPermanently
Api["github.com/gofiber/fiber"]["StatusFound"] = github_com_gofiber_fiber.StatusFound
Api["github.com/gofiber/fiber"]["StatusSeeOther"] = github_com_gofiber_fiber.StatusSeeOther
Api["github.com/gofiber/fiber"]["StatusNotModified"] = github_com_gofiber_fiber.StatusNotModified
Api["github.com/gofiber/fiber"]["StatusUseProxy"] = github_com_gofiber_fiber.StatusUseProxy
Api["github.com/gofiber/fiber"]["StatusTemporaryRedirect"] = github_com_gofiber_fiber.StatusTemporaryRedirect
Api["github.com/gofiber/fiber"]["StatusPermanentRedirect"] = github_com_gofiber_fiber.StatusPermanentRedirect
Api["github.com/gofiber/fiber"]["StatusBadRequest"] = github_com_gofiber_fiber.StatusBadRequest
Api["github.com/gofiber/fiber"]["StatusUnauthorized"] = github_com_gofiber_fiber.StatusUnauthorized
Api["github.com/gofiber/fiber"]["StatusPaymentRequired"] = github_com_gofiber_fiber.StatusPaymentRequired
Api["github.com/gofiber/fiber"]["StatusForbidden"] = github_com_gofiber_fiber.StatusForbidden
Api["github.com/gofiber/fiber"]["StatusNotFound"] = github_com_gofiber_fiber.StatusNotFound
Api["github.com/gofiber/fiber"]["StatusMethodNotAllowed"] = github_com_gofiber_fiber.StatusMethodNotAllowed
Api["github.com/gofiber/fiber"]["StatusNotAcceptable"] = github_com_gofiber_fiber.StatusNotAcceptable
Api["github.com/gofiber/fiber"]["StatusProxyAuthRequired"] = github_com_gofiber_fiber.StatusProxyAuthRequired
Api["github.com/gofiber/fiber"]["StatusRequestTimeout"] = github_com_gofiber_fiber.StatusRequestTimeout
Api["github.com/gofiber/fiber"]["StatusConflict"] = github_com_gofiber_fiber.StatusConflict
Api["github.com/gofiber/fiber"]["StatusGone"] = github_com_gofiber_fiber.StatusGone
Api["github.com/gofiber/fiber"]["StatusLengthRequired"] = github_com_gofiber_fiber.StatusLengthRequired
Api["github.com/gofiber/fiber"]["StatusPreconditionFailed"] = github_com_gofiber_fiber.StatusPreconditionFailed
Api["github.com/gofiber/fiber"]["StatusRequestEntityTooLarge"] = github_com_gofiber_fiber.StatusRequestEntityTooLarge
Api["github.com/gofiber/fiber"]["StatusRequestURITooLong"] = github_com_gofiber_fiber.StatusRequestURITooLong
Api["github.com/gofiber/fiber"]["StatusUnsupportedMediaType"] = github_com_gofiber_fiber.StatusUnsupportedMediaType
Api["github.com/gofiber/fiber"]["StatusRequestedRangeNotSatisfiable"] = github_com_gofiber_fiber.StatusRequestedRangeNotSatisfiable
Api["github.com/gofiber/fiber"]["StatusExpectationFailed"] = github_com_gofiber_fiber.StatusExpectationFailed
Api["github.com/gofiber/fiber"]["StatusTeapot"] = github_com_gofiber_fiber.StatusTeapot
Api["github.com/gofiber/fiber"]["StatusMisdirectedRequest"] = github_com_gofiber_fiber.StatusMisdirectedRequest
Api["github.com/gofiber/fiber"]["StatusUnprocessableEntity"] = github_com_gofiber_fiber.StatusUnprocessableEntity
Api["github.com/gofiber/fiber"]["StatusLocked"] = github_com_gofiber_fiber.StatusLocked
Api["github.com/gofiber/fiber"]["StatusFailedDependency"] = github_com_gofiber_fiber.StatusFailedDependency
Api["github.com/gofiber/fiber"]["StatusTooEarly"] = github_com_gofiber_fiber.StatusTooEarly
Api["github.com/gofiber/fiber"]["StatusUpgradeRequired"] = github_com_gofiber_fiber.StatusUpgradeRequired
Api["github.com/gofiber/fiber"]["StatusPreconditionRequired"] = github_com_gofiber_fiber.StatusPreconditionRequired
Api["github.com/gofiber/fiber"]["StatusTooManyRequests"] = github_com_gofiber_fiber.StatusTooManyRequests
Api["github.com/gofiber/fiber"]["StatusRequestHeaderFieldsTooLarge"] = github_com_gofiber_fiber.StatusRequestHeaderFieldsTooLarge
Api["github.com/gofiber/fiber"]["StatusUnavailableForLegalReasons"] = github_com_gofiber_fiber.StatusUnavailableForLegalReasons
Api["github.com/gofiber/fiber"]["StatusInternalServerError"] = github_com_gofiber_fiber.StatusInternalServerError
Api["github.com/gofiber/fiber"]["StatusNotImplemented"] = github_com_gofiber_fiber.StatusNotImplemented
Api["github.com/gofiber/fiber"]["StatusBadGateway"] = github_com_gofiber_fiber.StatusBadGateway
Api["github.com/gofiber/fiber"]["StatusServiceUnavailable"] = github_com_gofiber_fiber.StatusServiceUnavailable
Api["github.com/gofiber/fiber"]["StatusGatewayTimeout"] = github_com_gofiber_fiber.StatusGatewayTimeout
Api["github.com/gofiber/fiber"]["StatusHTTPVersionNotSupported"] = github_com_gofiber_fiber.StatusHTTPVersionNotSupported
Api["github.com/gofiber/fiber"]["StatusVariantAlsoNegotiates"] = github_com_gofiber_fiber.StatusVariantAlsoNegotiates
Api["github.com/gofiber/fiber"]["StatusInsufficientStorage"] = github_com_gofiber_fiber.StatusInsufficientStorage
Api["github.com/gofiber/fiber"]["StatusLoopDetected"] = github_com_gofiber_fiber.StatusLoopDetected
Api["github.com/gofiber/fiber"]["StatusNotExtended"] = github_com_gofiber_fiber.StatusNotExtended
Api["github.com/gofiber/fiber"]["StatusNetworkAuthenticationRequired"] = github_com_gofiber_fiber.StatusNetworkAuthenticationRequired
Api["github.com/gofiber/fiber"]["HeaderAuthorization"] = github_com_gofiber_fiber.HeaderAuthorization
Api["github.com/gofiber/fiber"]["HeaderProxyAuthenticate"] = github_com_gofiber_fiber.HeaderProxyAuthenticate
Api["github.com/gofiber/fiber"]["HeaderProxyAuthorization"] = github_com_gofiber_fiber.HeaderProxyAuthorization
Api["github.com/gofiber/fiber"]["HeaderWWWAuthenticate"] = github_com_gofiber_fiber.HeaderWWWAuthenticate
Api["github.com/gofiber/fiber"]["HeaderAge"] = github_com_gofiber_fiber.HeaderAge
Api["github.com/gofiber/fiber"]["HeaderCacheControl"] = github_com_gofiber_fiber.HeaderCacheControl
Api["github.com/gofiber/fiber"]["HeaderClearSiteData"] = github_com_gofiber_fiber.HeaderClearSiteData
Api["github.com/gofiber/fiber"]["HeaderExpires"] = github_com_gofiber_fiber.HeaderExpires
Api["github.com/gofiber/fiber"]["HeaderPragma"] = github_com_gofiber_fiber.HeaderPragma
Api["github.com/gofiber/fiber"]["HeaderWarning"] = github_com_gofiber_fiber.HeaderWarning
Api["github.com/gofiber/fiber"]["HeaderAcceptCH"] = github_com_gofiber_fiber.HeaderAcceptCH
Api["github.com/gofiber/fiber"]["HeaderAcceptCHLifetime"] = github_com_gofiber_fiber.HeaderAcceptCHLifetime
Api["github.com/gofiber/fiber"]["HeaderContentDPR"] = github_com_gofiber_fiber.HeaderContentDPR
Api["github.com/gofiber/fiber"]["HeaderDPR"] = github_com_gofiber_fiber.HeaderDPR
Api["github.com/gofiber/fiber"]["HeaderEarlyData"] = github_com_gofiber_fiber.HeaderEarlyData
Api["github.com/gofiber/fiber"]["HeaderSaveData"] = github_com_gofiber_fiber.HeaderSaveData
Api["github.com/gofiber/fiber"]["HeaderViewportWidth"] = github_com_gofiber_fiber.HeaderViewportWidth
Api["github.com/gofiber/fiber"]["HeaderWidth"] = github_com_gofiber_fiber.HeaderWidth
Api["github.com/gofiber/fiber"]["HeaderETag"] = github_com_gofiber_fiber.HeaderETag
Api["github.com/gofiber/fiber"]["HeaderIfMatch"] = github_com_gofiber_fiber.HeaderIfMatch
Api["github.com/gofiber/fiber"]["HeaderIfModifiedSince"] = github_com_gofiber_fiber.HeaderIfModifiedSince
Api["github.com/gofiber/fiber"]["HeaderIfNoneMatch"] = github_com_gofiber_fiber.HeaderIfNoneMatch
Api["github.com/gofiber/fiber"]["HeaderIfUnmodifiedSince"] = github_com_gofiber_fiber.HeaderIfUnmodifiedSince
Api["github.com/gofiber/fiber"]["HeaderLastModified"] = github_com_gofiber_fiber.HeaderLastModified
Api["github.com/gofiber/fiber"]["HeaderVary"] = github_com_gofiber_fiber.HeaderVary
Api["github.com/gofiber/fiber"]["HeaderConnection"] = github_com_gofiber_fiber.HeaderConnection
Api["github.com/gofiber/fiber"]["HeaderKeepAlive"] = github_com_gofiber_fiber.HeaderKeepAlive
Api["github.com/gofiber/fiber"]["HeaderAccept"] = github_com_gofiber_fiber.HeaderAccept
Api["github.com/gofiber/fiber"]["HeaderAcceptCharset"] = github_com_gofiber_fiber.HeaderAcceptCharset
Api["github.com/gofiber/fiber"]["HeaderAcceptEncoding"] = github_com_gofiber_fiber.HeaderAcceptEncoding
Api["github.com/gofiber/fiber"]["HeaderAcceptLanguage"] = github_com_gofiber_fiber.HeaderAcceptLanguage
Api["github.com/gofiber/fiber"]["HeaderCookie"] = github_com_gofiber_fiber.HeaderCookie
Api["github.com/gofiber/fiber"]["HeaderExpect"] = github_com_gofiber_fiber.HeaderExpect
Api["github.com/gofiber/fiber"]["HeaderMaxForwards"] = github_com_gofiber_fiber.HeaderMaxForwards
Api["github.com/gofiber/fiber"]["HeaderSetCookie"] = github_com_gofiber_fiber.HeaderSetCookie
Api["github.com/gofiber/fiber"]["HeaderAccessControlAllowCredentials"] = github_com_gofiber_fiber.HeaderAccessControlAllowCredentials
Api["github.com/gofiber/fiber"]["HeaderAccessControlAllowHeaders"] = github_com_gofiber_fiber.HeaderAccessControlAllowHeaders
Api["github.com/gofiber/fiber"]["HeaderAccessControlAllowMethods"] = github_com_gofiber_fiber.HeaderAccessControlAllowMethods
Api["github.com/gofiber/fiber"]["HeaderAccessControlAllowOrigin"] = github_com_gofiber_fiber.HeaderAccessControlAllowOrigin
Api["github.com/gofiber/fiber"]["HeaderAccessControlExposeHeaders"] = github_com_gofiber_fiber.HeaderAccessControlExposeHeaders
Api["github.com/gofiber/fiber"]["HeaderAccessControlMaxAge"] = github_com_gofiber_fiber.HeaderAccessControlMaxAge
Api["github.com/gofiber/fiber"]["HeaderAccessControlRequestHeaders"] = github_com_gofiber_fiber.HeaderAccessControlRequestHeaders
Api["github.com/gofiber/fiber"]["HeaderAccessControlRequestMethod"] = github_com_gofiber_fiber.HeaderAccessControlRequestMethod
Api["github.com/gofiber/fiber"]["HeaderOrigin"] = github_com_gofiber_fiber.HeaderOrigin
Api["github.com/gofiber/fiber"]["HeaderTimingAllowOrigin"] = github_com_gofiber_fiber.HeaderTimingAllowOrigin
Api["github.com/gofiber/fiber"]["HeaderXPermittedCrossDomainPolicies"] = github_com_gofiber_fiber.HeaderXPermittedCrossDomainPolicies
Api["github.com/gofiber/fiber"]["HeaderDNT"] = github_com_gofiber_fiber.HeaderDNT
Api["github.com/gofiber/fiber"]["HeaderTk"] = github_com_gofiber_fiber.HeaderTk
Api["github.com/gofiber/fiber"]["HeaderContentDisposition"] = github_com_gofiber_fiber.HeaderContentDisposition
Api["github.com/gofiber/fiber"]["HeaderContentEncoding"] = github_com_gofiber_fiber.HeaderContentEncoding
Api["github.com/gofiber/fiber"]["HeaderContentLanguage"] = github_com_gofiber_fiber.HeaderContentLanguage
Api["github.com/gofiber/fiber"]["HeaderContentLength"] = github_com_gofiber_fiber.HeaderContentLength
Api["github.com/gofiber/fiber"]["HeaderContentLocation"] = github_com_gofiber_fiber.HeaderContentLocation
Api["github.com/gofiber/fiber"]["HeaderContentType"] = github_com_gofiber_fiber.HeaderContentType
Api["github.com/gofiber/fiber"]["HeaderForwarded"] = github_com_gofiber_fiber.HeaderForwarded
Api["github.com/gofiber/fiber"]["HeaderVia"] = github_com_gofiber_fiber.HeaderVia
Api["github.com/gofiber/fiber"]["HeaderXForwardedFor"] = github_com_gofiber_fiber.HeaderXForwardedFor
Api["github.com/gofiber/fiber"]["HeaderXForwardedHost"] = github_com_gofiber_fiber.HeaderXForwardedHost
Api["github.com/gofiber/fiber"]["HeaderXForwardedProto"] = github_com_gofiber_fiber.HeaderXForwardedProto
Api["github.com/gofiber/fiber"]["HeaderXForwardedProtocol"] = github_com_gofiber_fiber.HeaderXForwardedProtocol
Api["github.com/gofiber/fiber"]["HeaderXForwardedSsl"] = github_com_gofiber_fiber.HeaderXForwardedSsl
Api["github.com/gofiber/fiber"]["HeaderXUrlScheme"] = github_com_gofiber_fiber.HeaderXUrlScheme
Api["github.com/gofiber/fiber"]["HeaderLocation"] = github_com_gofiber_fiber.HeaderLocation
Api["github.com/gofiber/fiber"]["HeaderFrom"] = github_com_gofiber_fiber.HeaderFrom
Api["github.com/gofiber/fiber"]["HeaderHost"] = github_com_gofiber_fiber.HeaderHost
Api["github.com/gofiber/fiber"]["HeaderReferer"] = github_com_gofiber_fiber.HeaderReferer
Api["github.com/gofiber/fiber"]["HeaderReferrerPolicy"] = github_com_gofiber_fiber.HeaderReferrerPolicy
Api["github.com/gofiber/fiber"]["HeaderUserAgent"] = github_com_gofiber_fiber.HeaderUserAgent
Api["github.com/gofiber/fiber"]["HeaderAllow"] = github_com_gofiber_fiber.HeaderAllow
Api["github.com/gofiber/fiber"]["HeaderServer"] = github_com_gofiber_fiber.HeaderServer
Api["github.com/gofiber/fiber"]["HeaderAcceptRanges"] = github_com_gofiber_fiber.HeaderAcceptRanges
Api["github.com/gofiber/fiber"]["HeaderContentRange"] = github_com_gofiber_fiber.HeaderContentRange
Api["github.com/gofiber/fiber"]["HeaderIfRange"] = github_com_gofiber_fiber.HeaderIfRange
Api["github.com/gofiber/fiber"]["HeaderRange"] = github_com_gofiber_fiber.HeaderRange
Api["github.com/gofiber/fiber"]["HeaderContentSecurityPolicy"] = github_com_gofiber_fiber.HeaderContentSecurityPolicy
Api["github.com/gofiber/fiber"]["HeaderContentSecurityPolicyReportOnly"] = github_com_gofiber_fiber.HeaderContentSecurityPolicyReportOnly
Api["github.com/gofiber/fiber"]["HeaderCrossOriginResourcePolicy"] = github_com_gofiber_fiber.HeaderCrossOriginResourcePolicy
Api["github.com/gofiber/fiber"]["HeaderExpectCT"] = github_com_gofiber_fiber.HeaderExpectCT
Api["github.com/gofiber/fiber"]["HeaderFeaturePolicy"] = github_com_gofiber_fiber.HeaderFeaturePolicy
Api["github.com/gofiber/fiber"]["HeaderPublicKeyPins"] = github_com_gofiber_fiber.HeaderPublicKeyPins
Api["github.com/gofiber/fiber"]["HeaderPublicKeyPinsReportOnly"] = github_com_gofiber_fiber.HeaderPublicKeyPinsReportOnly
Api["github.com/gofiber/fiber"]["HeaderStrictTransportSecurity"] = github_com_gofiber_fiber.HeaderStrictTransportSecurity
Api["github.com/gofiber/fiber"]["HeaderUpgradeInsecureRequests"] = github_com_gofiber_fiber.HeaderUpgradeInsecureRequests
Api["github.com/gofiber/fiber"]["HeaderXContentTypeOptions"] = github_com_gofiber_fiber.HeaderXContentTypeOptions
Api["github.com/gofiber/fiber"]["HeaderXDownloadOptions"] = github_com_gofiber_fiber.HeaderXDownloadOptions
Api["github.com/gofiber/fiber"]["HeaderXFrameOptions"] = github_com_gofiber_fiber.HeaderXFrameOptions
Api["github.com/gofiber/fiber"]["HeaderXPoweredBy"] = github_com_gofiber_fiber.HeaderXPoweredBy
Api["github.com/gofiber/fiber"]["HeaderXXSSProtection"] = github_com_gofiber_fiber.HeaderXXSSProtection
Api["github.com/gofiber/fiber"]["HeaderLastEventID"] = github_com_gofiber_fiber.HeaderLastEventID
Api["github.com/gofiber/fiber"]["HeaderNEL"] = github_com_gofiber_fiber.HeaderNEL
Api["github.com/gofiber/fiber"]["HeaderPingFrom"] = github_com_gofiber_fiber.HeaderPingFrom
Api["github.com/gofiber/fiber"]["HeaderPingTo"] = github_com_gofiber_fiber.HeaderPingTo
Api["github.com/gofiber/fiber"]["HeaderReportTo"] = github_com_gofiber_fiber.HeaderReportTo
Api["github.com/gofiber/fiber"]["HeaderTE"] = github_com_gofiber_fiber.HeaderTE
Api["github.com/gofiber/fiber"]["HeaderTrailer"] = github_com_gofiber_fiber.HeaderTrailer
Api["github.com/gofiber/fiber"]["HeaderTransferEncoding"] = github_com_gofiber_fiber.HeaderTransferEncoding
Api["github.com/gofiber/fiber"]["HeaderSecWebSocketAccept"] = github_com_gofiber_fiber.HeaderSecWebSocketAccept
Api["github.com/gofiber/fiber"]["HeaderSecWebSocketExtensions"] = github_com_gofiber_fiber.HeaderSecWebSocketExtensions
Api["github.com/gofiber/fiber"]["HeaderSecWebSocketKey"] = github_com_gofiber_fiber.HeaderSecWebSocketKey
Api["github.com/gofiber/fiber"]["HeaderSecWebSocketProtocol"] = github_com_gofiber_fiber.HeaderSecWebSocketProtocol
Api["github.com/gofiber/fiber"]["HeaderSecWebSocketVersion"] = github_com_gofiber_fiber.HeaderSecWebSocketVersion
Api["github.com/gofiber/fiber"]["HeaderAcceptPatch"] = github_com_gofiber_fiber.HeaderAcceptPatch
Api["github.com/gofiber/fiber"]["HeaderAcceptPushPolicy"] = github_com_gofiber_fiber.HeaderAcceptPushPolicy
Api["github.com/gofiber/fiber"]["HeaderAcceptSignature"] = github_com_gofiber_fiber.HeaderAcceptSignature
Api["github.com/gofiber/fiber"]["HeaderAltSvc"] = github_com_gofiber_fiber.HeaderAltSvc
Api["github.com/gofiber/fiber"]["HeaderDate"] = github_com_gofiber_fiber.HeaderDate
Api["github.com/gofiber/fiber"]["HeaderIndex"] = github_com_gofiber_fiber.HeaderIndex
Api["github.com/gofiber/fiber"]["HeaderLargeAllocation"] = github_com_gofiber_fiber.HeaderLargeAllocation
Api["github.com/gofiber/fiber"]["HeaderLink"] = github_com_gofiber_fiber.HeaderLink
Api["github.com/gofiber/fiber"]["HeaderPushPolicy"] = github_com_gofiber_fiber.HeaderPushPolicy
Api["github.com/gofiber/fiber"]["HeaderRetryAfter"] = github_com_gofiber_fiber.HeaderRetryAfter
Api["github.com/gofiber/fiber"]["HeaderServerTiming"] = github_com_gofiber_fiber.HeaderServerTiming
Api["github.com/gofiber/fiber"]["HeaderSignature"] = github_com_gofiber_fiber.HeaderSignature
Api["github.com/gofiber/fiber"]["HeaderSignedHeaders"] = github_com_gofiber_fiber.HeaderSignedHeaders
Api["github.com/gofiber/fiber"]["HeaderSourceMap"] = github_com_gofiber_fiber.HeaderSourceMap
Api["github.com/gofiber/fiber"]["HeaderUpgrade"] = github_com_gofiber_fiber.HeaderUpgrade
Api["github.com/gofiber/fiber"]["HeaderXDNSPrefetchControl"] = github_com_gofiber_fiber.HeaderXDNSPrefetchControl
Api["github.com/gofiber/fiber"]["HeaderXPingback"] = github_com_gofiber_fiber.HeaderXPingback
Api["github.com/gofiber/fiber"]["HeaderXRequestID"] = github_com_gofiber_fiber.HeaderXRequestID
Api["github.com/gofiber/fiber"]["HeaderXRequestedWith"] = github_com_gofiber_fiber.HeaderXRequestedWith
Api["github.com/gofiber/fiber"]["HeaderXRobotsTag"] = github_com_gofiber_fiber.HeaderXRobotsTag
Api["github.com/gofiber/fiber"]["HeaderXUACompatible"] = github_com_gofiber_fiber.HeaderXUACompatible
Api["github.com/gofiber/fiber"]["App"] = github_com_gofiber_fiber.App{}
Api["github.com/gofiber/fiber/new"]["App"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.App)
}
Api["github.com/gofiber/fiber"]["Cookie"] = github_com_gofiber_fiber.Cookie{}
Api["github.com/gofiber/fiber/new"]["Cookie"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Cookie)
}
Api["github.com/gofiber/fiber"]["Ctx"] = github_com_gofiber_fiber.Ctx{}
Api["github.com/gofiber/fiber/new"]["Ctx"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Ctx)
}
Api["github.com/gofiber/fiber"]["NewError"] = github_com_gofiber_fiber.NewError
Api["github.com/gofiber/fiber"]["Group"] = github_com_gofiber_fiber.Group{}
Api["github.com/gofiber/fiber/new"]["Group"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Group)
}
Api["github.com/gofiber/fiber"]["Range"] = github_com_gofiber_fiber.Range{}
Api["github.com/gofiber/fiber/new"]["Range"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Range)
}
Api["github.com/gofiber/fiber"]["Route"] = github_com_gofiber_fiber.Route{}
Api["github.com/gofiber/fiber/new"]["Route"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Route)
}
Api["github.com/gofiber/fiber"]["Static"] = github_com_gofiber_fiber.Static{}
Api["github.com/gofiber/fiber/new"]["Static"] = func(args ...interface{}) interface{} {
					return new(github_com_gofiber_fiber.Static)
}
}
