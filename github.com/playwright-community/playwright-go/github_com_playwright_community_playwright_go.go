package playwright_go
import (
github_com_playwright_community_playwright_go "github.com/playwright-community/playwright-go"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/playwright-community/playwright-go"]; !ok {
   Api["github.com/playwright-community/playwright-go"] = map[string]interface{}{}
}
if _, ok := Api["github.com/playwright-community/playwright-go/new"]; !ok {
	Api["github.com/playwright-community/playwright-go/new"] = map[string]interface{}{}
}
Api["github.com/playwright-community/playwright-go"]["TimeoutError"] = github_com_playwright_community_playwright_go.TimeoutError
Api["github.com/playwright-community/playwright-go"]["Float"] = github_com_playwright_community_playwright_go.Float
Api["github.com/playwright-community/playwright-go"]["Install"] = github_com_playwright_community_playwright_go.Install
Api["github.com/playwright-community/playwright-go"]["Int"] = github_com_playwright_community_playwright_go.Int
Api["github.com/playwright-community/playwright-go"]["IntSlice"] = github_com_playwright_community_playwright_go.IntSlice
Api["github.com/playwright-community/playwright-go"]["StringSlice"] = github_com_playwright_community_playwright_go.StringSlice
Api["github.com/playwright-community/playwright-go"]["APIRequestContextDeleteOptions"] = github_com_playwright_community_playwright_go.APIRequestContextDeleteOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextDeleteOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextDeleteOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextFetchOptions"] = github_com_playwright_community_playwright_go.APIRequestContextFetchOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextFetchOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextFetchOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextGetOptions"] = github_com_playwright_community_playwright_go.APIRequestContextGetOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextGetOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextGetOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextHeadOptions"] = github_com_playwright_community_playwright_go.APIRequestContextHeadOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextHeadOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextHeadOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextPatchOptions"] = github_com_playwright_community_playwright_go.APIRequestContextPatchOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextPatchOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextPatchOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextPostOptions"] = github_com_playwright_community_playwright_go.APIRequestContextPostOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextPostOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextPostOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextPutOptions"] = github_com_playwright_community_playwright_go.APIRequestContextPutOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextPutOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextPutOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestContextStorageStateOptions"] = github_com_playwright_community_playwright_go.APIRequestContextStorageStateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestContextStorageStateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestContextStorageStateOptions)
}
Api["github.com/playwright-community/playwright-go"]["APIRequestNewContextOptions"] = github_com_playwright_community_playwright_go.APIRequestNewContextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["APIRequestNewContextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.APIRequestNewContextOptions)
}
Api["github.com/playwright-community/playwright-go"]["BindingSource"] = github_com_playwright_community_playwright_go.BindingSource{}
Api["github.com/playwright-community/playwright-go/new"]["BindingSource"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BindingSource)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextAddCookiesOptions"] = github_com_playwright_community_playwright_go.BrowserContextAddCookiesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextAddCookiesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextAddCookiesOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextAddInitScriptOptions"] = github_com_playwright_community_playwright_go.BrowserContextAddInitScriptOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextAddInitScriptOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextAddInitScriptOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextCookiesOptions"] = github_com_playwright_community_playwright_go.BrowserContextCookiesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextCookiesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextCookiesOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextExpectConsoleMessageOptions"] = github_com_playwright_community_playwright_go.BrowserContextExpectConsoleMessageOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextExpectConsoleMessageOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextExpectConsoleMessageOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextExpectEventOptions"] = github_com_playwright_community_playwright_go.BrowserContextExpectEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextExpectEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextExpectEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextExpectPageOptions"] = github_com_playwright_community_playwright_go.BrowserContextExpectPageOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextExpectPageOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextExpectPageOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextExposeBindingOptions"] = github_com_playwright_community_playwright_go.BrowserContextExposeBindingOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextExposeBindingOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextExposeBindingOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextGrantPermissionsOptions"] = github_com_playwright_community_playwright_go.BrowserContextGrantPermissionsOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextGrantPermissionsOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextGrantPermissionsOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextRouteFromHAROptions"] = github_com_playwright_community_playwright_go.BrowserContextRouteFromHAROptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextRouteFromHAROptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextRouteFromHAROptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextRouteOptions"] = github_com_playwright_community_playwright_go.BrowserContextRouteOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextRouteOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextRouteOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextStorageStateOptions"] = github_com_playwright_community_playwright_go.BrowserContextStorageStateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextStorageStateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextStorageStateOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextUnrouteOptions"] = github_com_playwright_community_playwright_go.BrowserContextUnrouteOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextUnrouteOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextUnrouteOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserContextWaitForEventOptions"] = github_com_playwright_community_playwright_go.BrowserContextWaitForEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserContextWaitForEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserContextWaitForEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserNewContextOptions"] = github_com_playwright_community_playwright_go.BrowserNewContextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserNewContextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserNewContextOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserNewPageOptions"] = github_com_playwright_community_playwright_go.BrowserNewPageOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserNewPageOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserNewPageOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserStartTracingOptions"] = github_com_playwright_community_playwright_go.BrowserStartTracingOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserStartTracingOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserStartTracingOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserTypeConnectOptions"] = github_com_playwright_community_playwright_go.BrowserTypeConnectOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserTypeConnectOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserTypeConnectOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserTypeConnectOverCDPOptions"] = github_com_playwright_community_playwright_go.BrowserTypeConnectOverCDPOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserTypeConnectOverCDPOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserTypeConnectOverCDPOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserTypeLaunchOptions"] = github_com_playwright_community_playwright_go.BrowserTypeLaunchOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserTypeLaunchOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserTypeLaunchOptions)
}
Api["github.com/playwright-community/playwright-go"]["BrowserTypeLaunchPersistentContextOptions"] = github_com_playwright_community_playwright_go.BrowserTypeLaunchPersistentContextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["BrowserTypeLaunchPersistentContextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.BrowserTypeLaunchPersistentContextOptions)
}
Api["github.com/playwright-community/playwright-go"]["ConsoleMessageLocation"] = github_com_playwright_community_playwright_go.ConsoleMessageLocation{}
Api["github.com/playwright-community/playwright-go/new"]["ConsoleMessageLocation"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ConsoleMessageLocation)
}
Api["github.com/playwright-community/playwright-go"]["DeviceDescriptor"] = github_com_playwright_community_playwright_go.DeviceDescriptor{}
Api["github.com/playwright-community/playwright-go/new"]["DeviceDescriptor"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.DeviceDescriptor)
}
Api["github.com/playwright-community/playwright-go"]["DialogAcceptOptions"] = github_com_playwright_community_playwright_go.DialogAcceptOptions{}
Api["github.com/playwright-community/playwright-go/new"]["DialogAcceptOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.DialogAcceptOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleCheckOptions"] = github_com_playwright_community_playwright_go.ElementHandleCheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleCheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleCheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleClickOptions"] = github_com_playwright_community_playwright_go.ElementHandleClickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleClickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleClickOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleDblclickOptions"] = github_com_playwright_community_playwright_go.ElementHandleDblclickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleDblclickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleDblclickOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleDispatchEventOptions"] = github_com_playwright_community_playwright_go.ElementHandleDispatchEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleDispatchEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleDispatchEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleEvalOnSelectorAllOptions"] = github_com_playwright_community_playwright_go.ElementHandleEvalOnSelectorAllOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleEvalOnSelectorAllOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleEvalOnSelectorAllOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleEvalOnSelectorOptions"] = github_com_playwright_community_playwright_go.ElementHandleEvalOnSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleEvalOnSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleEvalOnSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleFillOptions"] = github_com_playwright_community_playwright_go.ElementHandleFillOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleFillOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleFillOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleHoverOptions"] = github_com_playwright_community_playwright_go.ElementHandleHoverOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleHoverOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleHoverOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleInputValueOptions"] = github_com_playwright_community_playwright_go.ElementHandleInputValueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleInputValueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleInputValueOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandlePressOptions"] = github_com_playwright_community_playwright_go.ElementHandlePressOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandlePressOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandlePressOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleScreenshotOptions"] = github_com_playwright_community_playwright_go.ElementHandleScreenshotOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleScreenshotOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleScreenshotOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleScrollIntoViewIfNeededOptions"] = github_com_playwright_community_playwright_go.ElementHandleScrollIntoViewIfNeededOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleScrollIntoViewIfNeededOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleScrollIntoViewIfNeededOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleSelectOptionOptions"] = github_com_playwright_community_playwright_go.ElementHandleSelectOptionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleSelectOptionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleSelectOptionOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleSelectTextOptions"] = github_com_playwright_community_playwright_go.ElementHandleSelectTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleSelectTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleSelectTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleSetCheckedOptions"] = github_com_playwright_community_playwright_go.ElementHandleSetCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleSetCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleSetCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleSetInputFilesOptions"] = github_com_playwright_community_playwright_go.ElementHandleSetInputFilesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleSetInputFilesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleSetInputFilesOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleTapOptions"] = github_com_playwright_community_playwright_go.ElementHandleTapOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleTapOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleTapOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleTypeOptions"] = github_com_playwright_community_playwright_go.ElementHandleTypeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleTypeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleTypeOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleUncheckOptions"] = github_com_playwright_community_playwright_go.ElementHandleUncheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleUncheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleUncheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleWaitForElementStateOptions"] = github_com_playwright_community_playwright_go.ElementHandleWaitForElementStateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleWaitForElementStateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleWaitForElementStateOptions)
}
Api["github.com/playwright-community/playwright-go"]["ElementHandleWaitForSelectorOptions"] = github_com_playwright_community_playwright_go.ElementHandleWaitForSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["ElementHandleWaitForSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ElementHandleWaitForSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["FileChooserSetFilesOptions"] = github_com_playwright_community_playwright_go.FileChooserSetFilesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FileChooserSetFilesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FileChooserSetFilesOptions)
}
Api["github.com/playwright-community/playwright-go"]["Frame"] = new(github_com_playwright_community_playwright_go.Frame)
Api["github.com/playwright-community/playwright-go"]["FrameAddScriptTagOptions"] = github_com_playwright_community_playwright_go.FrameAddScriptTagOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameAddScriptTagOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameAddScriptTagOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameAddStyleTagOptions"] = github_com_playwright_community_playwright_go.FrameAddStyleTagOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameAddStyleTagOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameAddStyleTagOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameCheckOptions"] = github_com_playwright_community_playwright_go.FrameCheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameCheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameCheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameClickOptions"] = github_com_playwright_community_playwright_go.FrameClickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameClickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameClickOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameDblclickOptions"] = github_com_playwright_community_playwright_go.FrameDblclickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameDblclickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameDblclickOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameDispatchEventOptions"] = github_com_playwright_community_playwright_go.FrameDispatchEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameDispatchEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameDispatchEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameDragAndDropOptions"] = github_com_playwright_community_playwright_go.FrameDragAndDropOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameDragAndDropOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameDragAndDropOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameEvalOnSelectorAllOptions"] = github_com_playwright_community_playwright_go.FrameEvalOnSelectorAllOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameEvalOnSelectorAllOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameEvalOnSelectorAllOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameEvalOnSelectorOptions"] = github_com_playwright_community_playwright_go.FrameEvalOnSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameEvalOnSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameEvalOnSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameEvaluateHandleOptions"] = github_com_playwright_community_playwright_go.FrameEvaluateHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameEvaluateHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameEvaluateHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameEvaluateOptions"] = github_com_playwright_community_playwright_go.FrameEvaluateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameEvaluateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameEvaluateOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameFillOptions"] = github_com_playwright_community_playwright_go.FrameFillOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameFillOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameFillOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameFocusOptions"] = github_com_playwright_community_playwright_go.FrameFocusOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameFocusOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameFocusOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetAttributeOptions"] = github_com_playwright_community_playwright_go.FrameGetAttributeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetAttributeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetAttributeOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByAltTextOptions"] = github_com_playwright_community_playwright_go.FrameGetByAltTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByAltTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByAltTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByLabelOptions"] = github_com_playwright_community_playwright_go.FrameGetByLabelOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByLabelOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByLabelOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByPlaceholderOptions"] = github_com_playwright_community_playwright_go.FrameGetByPlaceholderOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByPlaceholderOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByPlaceholderOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByRoleOptions"] = github_com_playwright_community_playwright_go.FrameGetByRoleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByRoleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByRoleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByTextOptions"] = github_com_playwright_community_playwright_go.FrameGetByTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGetByTitleOptions"] = github_com_playwright_community_playwright_go.FrameGetByTitleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGetByTitleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGetByTitleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameGotoOptions"] = github_com_playwright_community_playwright_go.FrameGotoOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameGotoOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameGotoOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameHoverOptions"] = github_com_playwright_community_playwright_go.FrameHoverOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameHoverOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameHoverOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameInnerHTMLOptions"] = github_com_playwright_community_playwright_go.FrameInnerHTMLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameInnerHTMLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameInnerHTMLOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameInnerTextOptions"] = github_com_playwright_community_playwright_go.FrameInnerTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameInnerTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameInnerTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameInputValueOptions"] = github_com_playwright_community_playwright_go.FrameInputValueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameInputValueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameInputValueOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsCheckedOptions"] = github_com_playwright_community_playwright_go.FrameIsCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsDisabledOptions"] = github_com_playwright_community_playwright_go.FrameIsDisabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsDisabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsDisabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsEditableOptions"] = github_com_playwright_community_playwright_go.FrameIsEditableOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsEditableOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsEditableOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsEnabledOptions"] = github_com_playwright_community_playwright_go.FrameIsEnabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsEnabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsEnabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsHiddenOptions"] = github_com_playwright_community_playwright_go.FrameIsHiddenOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsHiddenOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsHiddenOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameIsVisibleOptions"] = github_com_playwright_community_playwright_go.FrameIsVisibleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameIsVisibleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameIsVisibleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByAltTextOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByAltTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByAltTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByAltTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByLabelOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByLabelOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByLabelOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByLabelOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByPlaceholderOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByPlaceholderOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByPlaceholderOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByPlaceholderOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByRoleOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByRoleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByRoleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByRoleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByTextOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorGetByTitleOptions"] = github_com_playwright_community_playwright_go.FrameLocatorGetByTitleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorGetByTitleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorGetByTitleOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameLocatorOptions"] = github_com_playwright_community_playwright_go.FrameLocatorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameLocatorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameLocatorOptions)
}
Api["github.com/playwright-community/playwright-go"]["FramePressOptions"] = github_com_playwright_community_playwright_go.FramePressOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FramePressOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FramePressOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameQuerySelectorOptions"] = github_com_playwright_community_playwright_go.FrameQuerySelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameQuerySelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameQuerySelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameReceivedPayload"] = github_com_playwright_community_playwright_go.FrameReceivedPayload{}
Api["github.com/playwright-community/playwright-go/new"]["FrameReceivedPayload"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameReceivedPayload)
}
Api["github.com/playwright-community/playwright-go"]["FrameSelectOptionOptions"] = github_com_playwright_community_playwright_go.FrameSelectOptionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameSelectOptionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameSelectOptionOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameSentPayload"] = github_com_playwright_community_playwright_go.FrameSentPayload{}
Api["github.com/playwright-community/playwright-go/new"]["FrameSentPayload"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameSentPayload)
}
Api["github.com/playwright-community/playwright-go"]["FrameSetCheckedOptions"] = github_com_playwright_community_playwright_go.FrameSetCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameSetCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameSetCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameSetContentOptions"] = github_com_playwright_community_playwright_go.FrameSetContentOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameSetContentOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameSetContentOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameSetInputFilesOptions"] = github_com_playwright_community_playwright_go.FrameSetInputFilesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameSetInputFilesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameSetInputFilesOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameTapOptions"] = github_com_playwright_community_playwright_go.FrameTapOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameTapOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameTapOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameTextContentOptions"] = github_com_playwright_community_playwright_go.FrameTextContentOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameTextContentOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameTextContentOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameTypeOptions"] = github_com_playwright_community_playwright_go.FrameTypeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameTypeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameTypeOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameUncheckOptions"] = github_com_playwright_community_playwright_go.FrameUncheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameUncheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameUncheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameWaitForFunctionOptions"] = github_com_playwright_community_playwright_go.FrameWaitForFunctionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameWaitForFunctionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameWaitForFunctionOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameWaitForLoadStateOptions"] = github_com_playwright_community_playwright_go.FrameWaitForLoadStateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameWaitForLoadStateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameWaitForLoadStateOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameWaitForNavigationOptions"] = github_com_playwright_community_playwright_go.FrameWaitForNavigationOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameWaitForNavigationOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameWaitForNavigationOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameWaitForSelectorOptions"] = github_com_playwright_community_playwright_go.FrameWaitForSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameWaitForSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameWaitForSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["FrameWaitForURLOptions"] = github_com_playwright_community_playwright_go.FrameWaitForURLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["FrameWaitForURLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.FrameWaitForURLOptions)
}
Api["github.com/playwright-community/playwright-go"]["Geolocation"] = github_com_playwright_community_playwright_go.Geolocation{}
Api["github.com/playwright-community/playwright-go/new"]["Geolocation"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.Geolocation)
}
Api["github.com/playwright-community/playwright-go"]["HttpCredentials"] = github_com_playwright_community_playwright_go.HttpCredentials{}
Api["github.com/playwright-community/playwright-go/new"]["HttpCredentials"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.HttpCredentials)
}
Api["github.com/playwright-community/playwright-go"]["InputFile"] = github_com_playwright_community_playwright_go.InputFile{}
Api["github.com/playwright-community/playwright-go/new"]["InputFile"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.InputFile)
}
Api["github.com/playwright-community/playwright-go"]["JSHandleEvaluateHandleOptions"] = github_com_playwright_community_playwright_go.JSHandleEvaluateHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["JSHandleEvaluateHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.JSHandleEvaluateHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["JSHandleEvaluateOptions"] = github_com_playwright_community_playwright_go.JSHandleEvaluateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["JSHandleEvaluateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.JSHandleEvaluateOptions)
}
Api["github.com/playwright-community/playwright-go"]["KeyboardPressOptions"] = github_com_playwright_community_playwright_go.KeyboardPressOptions{}
Api["github.com/playwright-community/playwright-go/new"]["KeyboardPressOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.KeyboardPressOptions)
}
Api["github.com/playwright-community/playwright-go"]["KeyboardTypeOptions"] = github_com_playwright_community_playwright_go.KeyboardTypeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["KeyboardTypeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.KeyboardTypeOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocalStorageEntry"] = github_com_playwright_community_playwright_go.LocalStorageEntry{}
Api["github.com/playwright-community/playwright-go/new"]["LocalStorageEntry"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocalStorageEntry)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeAttachedOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeAttachedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeAttachedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeAttachedOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeCheckedOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeDisabledOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeDisabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeDisabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeDisabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeEditableOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeEditableOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeEditableOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeEditableOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeEmptyOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeEmptyOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeEmptyOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeEmptyOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeEnabledOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeEnabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeEnabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeEnabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeFocusedOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeFocusedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeFocusedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeFocusedOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeHiddenOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeHiddenOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeHiddenOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeHiddenOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeInViewportOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeInViewportOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeInViewportOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeInViewportOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToBeVisibleOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToBeVisibleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToBeVisibleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToBeVisibleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToContainTextOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToContainTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToContainTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToContainTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveAttributeOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveAttributeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveAttributeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveAttributeOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveCSSOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveCSSOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveCSSOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveCSSOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveClassOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveClassOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveClassOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveClassOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveCountOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveCountOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveCountOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveCountOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveIdOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveIdOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveIdOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveIdOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveJSPropertyOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveJSPropertyOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveJSPropertyOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveJSPropertyOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveTextOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveValueOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveValueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveValueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveValueOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorAssertionsToHaveValuesOptions"] = github_com_playwright_community_playwright_go.LocatorAssertionsToHaveValuesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorAssertionsToHaveValuesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorAssertionsToHaveValuesOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorBlurOptions"] = github_com_playwright_community_playwright_go.LocatorBlurOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorBlurOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorBlurOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorBoundingBoxOptions"] = github_com_playwright_community_playwright_go.LocatorBoundingBoxOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorBoundingBoxOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorBoundingBoxOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorCheckOptions"] = github_com_playwright_community_playwright_go.LocatorCheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorCheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorCheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorClearOptions"] = github_com_playwright_community_playwright_go.LocatorClearOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorClearOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorClearOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorClickOptions"] = github_com_playwright_community_playwright_go.LocatorClickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorClickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorClickOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorDblclickOptions"] = github_com_playwright_community_playwright_go.LocatorDblclickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorDblclickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorDblclickOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorDispatchEventOptions"] = github_com_playwright_community_playwright_go.LocatorDispatchEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorDispatchEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorDispatchEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorDragToOptions"] = github_com_playwright_community_playwright_go.LocatorDragToOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorDragToOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorDragToOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorElementHandleOptions"] = github_com_playwright_community_playwright_go.LocatorElementHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorElementHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorElementHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorEvaluateAllOptions"] = github_com_playwright_community_playwright_go.LocatorEvaluateAllOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorEvaluateAllOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorEvaluateAllOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorEvaluateHandleOptions"] = github_com_playwright_community_playwright_go.LocatorEvaluateHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorEvaluateHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorEvaluateHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorEvaluateOptions"] = github_com_playwright_community_playwright_go.LocatorEvaluateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorEvaluateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorEvaluateOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorFillOptions"] = github_com_playwright_community_playwright_go.LocatorFillOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorFillOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorFillOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorFilterOptions"] = github_com_playwright_community_playwright_go.LocatorFilterOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorFilterOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorFilterOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorFocusOptions"] = github_com_playwright_community_playwright_go.LocatorFocusOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorFocusOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorFocusOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetAttributeOptions"] = github_com_playwright_community_playwright_go.LocatorGetAttributeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetAttributeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetAttributeOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByAltTextOptions"] = github_com_playwright_community_playwright_go.LocatorGetByAltTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByAltTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByAltTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByLabelOptions"] = github_com_playwright_community_playwright_go.LocatorGetByLabelOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByLabelOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByLabelOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByPlaceholderOptions"] = github_com_playwright_community_playwright_go.LocatorGetByPlaceholderOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByPlaceholderOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByPlaceholderOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByRoleOptions"] = github_com_playwright_community_playwright_go.LocatorGetByRoleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByRoleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByRoleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByTextOptions"] = github_com_playwright_community_playwright_go.LocatorGetByTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorGetByTitleOptions"] = github_com_playwright_community_playwright_go.LocatorGetByTitleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorGetByTitleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorGetByTitleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorHoverOptions"] = github_com_playwright_community_playwright_go.LocatorHoverOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorHoverOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorHoverOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorInnerHTMLOptions"] = github_com_playwright_community_playwright_go.LocatorInnerHTMLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorInnerHTMLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorInnerHTMLOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorInnerTextOptions"] = github_com_playwright_community_playwright_go.LocatorInnerTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorInnerTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorInnerTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorInputValueOptions"] = github_com_playwright_community_playwright_go.LocatorInputValueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorInputValueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorInputValueOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsCheckedOptions"] = github_com_playwright_community_playwright_go.LocatorIsCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsDisabledOptions"] = github_com_playwright_community_playwright_go.LocatorIsDisabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsDisabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsDisabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsEditableOptions"] = github_com_playwright_community_playwright_go.LocatorIsEditableOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsEditableOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsEditableOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsEnabledOptions"] = github_com_playwright_community_playwright_go.LocatorIsEnabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsEnabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsEnabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsHiddenOptions"] = github_com_playwright_community_playwright_go.LocatorIsHiddenOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsHiddenOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsHiddenOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorIsVisibleOptions"] = github_com_playwright_community_playwright_go.LocatorIsVisibleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorIsVisibleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorIsVisibleOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorLocatorOptions"] = github_com_playwright_community_playwright_go.LocatorLocatorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorLocatorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorLocatorOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorPressOptions"] = github_com_playwright_community_playwright_go.LocatorPressOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorPressOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorPressOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorScreenshotOptions"] = github_com_playwright_community_playwright_go.LocatorScreenshotOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorScreenshotOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorScreenshotOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorScrollIntoViewIfNeededOptions"] = github_com_playwright_community_playwright_go.LocatorScrollIntoViewIfNeededOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorScrollIntoViewIfNeededOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorScrollIntoViewIfNeededOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorSelectOptionOptions"] = github_com_playwright_community_playwright_go.LocatorSelectOptionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorSelectOptionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorSelectOptionOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorSelectTextOptions"] = github_com_playwright_community_playwright_go.LocatorSelectTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorSelectTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorSelectTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorSetCheckedOptions"] = github_com_playwright_community_playwright_go.LocatorSetCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorSetCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorSetCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorSetInputFilesOptions"] = github_com_playwright_community_playwright_go.LocatorSetInputFilesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorSetInputFilesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorSetInputFilesOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorTapOptions"] = github_com_playwright_community_playwright_go.LocatorTapOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorTapOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorTapOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorTextContentOptions"] = github_com_playwright_community_playwright_go.LocatorTextContentOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorTextContentOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorTextContentOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorTypeOptions"] = github_com_playwright_community_playwright_go.LocatorTypeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorTypeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorTypeOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorUncheckOptions"] = github_com_playwright_community_playwright_go.LocatorUncheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorUncheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorUncheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["LocatorWaitForOptions"] = github_com_playwright_community_playwright_go.LocatorWaitForOptions{}
Api["github.com/playwright-community/playwright-go/new"]["LocatorWaitForOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.LocatorWaitForOptions)
}
Api["github.com/playwright-community/playwright-go"]["Margin"] = github_com_playwright_community_playwright_go.Margin{}
Api["github.com/playwright-community/playwright-go/new"]["Margin"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.Margin)
}
Api["github.com/playwright-community/playwright-go"]["MouseClickOptions"] = github_com_playwright_community_playwright_go.MouseClickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["MouseClickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.MouseClickOptions)
}
Api["github.com/playwright-community/playwright-go"]["MouseDblclickOptions"] = github_com_playwright_community_playwright_go.MouseDblclickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["MouseDblclickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.MouseDblclickOptions)
}
Api["github.com/playwright-community/playwright-go"]["MouseDownOptions"] = github_com_playwright_community_playwright_go.MouseDownOptions{}
Api["github.com/playwright-community/playwright-go/new"]["MouseDownOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.MouseDownOptions)
}
Api["github.com/playwright-community/playwright-go"]["MouseMoveOptions"] = github_com_playwright_community_playwright_go.MouseMoveOptions{}
Api["github.com/playwright-community/playwright-go/new"]["MouseMoveOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.MouseMoveOptions)
}
Api["github.com/playwright-community/playwright-go"]["MouseUpOptions"] = github_com_playwright_community_playwright_go.MouseUpOptions{}
Api["github.com/playwright-community/playwright-go/new"]["MouseUpOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.MouseUpOptions)
}
Api["github.com/playwright-community/playwright-go"]["NameValue"] = github_com_playwright_community_playwright_go.NameValue{}
Api["github.com/playwright-community/playwright-go/new"]["NameValue"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.NameValue)
}
Api["github.com/playwright-community/playwright-go"]["OptionalCookie"] = github_com_playwright_community_playwright_go.OptionalCookie{}
Api["github.com/playwright-community/playwright-go/new"]["OptionalCookie"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.OptionalCookie)
}
Api["github.com/playwright-community/playwright-go"]["OptionalStorageState"] = github_com_playwright_community_playwright_go.OptionalStorageState{}
Api["github.com/playwright-community/playwright-go/new"]["OptionalStorageState"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.OptionalStorageState)
}
Api["github.com/playwright-community/playwright-go"]["OriginsState"] = github_com_playwright_community_playwright_go.OriginsState{}
Api["github.com/playwright-community/playwright-go/new"]["OriginsState"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.OriginsState)
}
Api["github.com/playwright-community/playwright-go"]["Page"] = new(github_com_playwright_community_playwright_go.Page)
Api["github.com/playwright-community/playwright-go"]["PageAddInitScriptOptions"] = github_com_playwright_community_playwright_go.PageAddInitScriptOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageAddInitScriptOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageAddInitScriptOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageAddScriptTagOptions"] = github_com_playwright_community_playwright_go.PageAddScriptTagOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageAddScriptTagOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageAddScriptTagOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageAddStyleTagOptions"] = github_com_playwright_community_playwright_go.PageAddStyleTagOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageAddStyleTagOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageAddStyleTagOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageAssertionsToHaveTitleOptions"] = github_com_playwright_community_playwright_go.PageAssertionsToHaveTitleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageAssertionsToHaveTitleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageAssertionsToHaveTitleOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageAssertionsToHaveURLOptions"] = github_com_playwright_community_playwright_go.PageAssertionsToHaveURLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageAssertionsToHaveURLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageAssertionsToHaveURLOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageCheckOptions"] = github_com_playwright_community_playwright_go.PageCheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageCheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageCheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageClickOptions"] = github_com_playwright_community_playwright_go.PageClickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageClickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageClickOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageCloseOptions"] = github_com_playwright_community_playwright_go.PageCloseOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageCloseOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageCloseOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageDblclickOptions"] = github_com_playwright_community_playwright_go.PageDblclickOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageDblclickOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageDblclickOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageDispatchEventOptions"] = github_com_playwright_community_playwright_go.PageDispatchEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageDispatchEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageDispatchEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageDragAndDropOptions"] = github_com_playwright_community_playwright_go.PageDragAndDropOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageDragAndDropOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageDragAndDropOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageEmulateMediaOptions"] = github_com_playwright_community_playwright_go.PageEmulateMediaOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageEmulateMediaOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageEmulateMediaOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageEvalOnSelectorAllOptions"] = github_com_playwright_community_playwright_go.PageEvalOnSelectorAllOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageEvalOnSelectorAllOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageEvalOnSelectorAllOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageEvalOnSelectorOptions"] = github_com_playwright_community_playwright_go.PageEvalOnSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageEvalOnSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageEvalOnSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageEvaluateHandleOptions"] = github_com_playwright_community_playwright_go.PageEvaluateHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageEvaluateHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageEvaluateHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageEvaluateOptions"] = github_com_playwright_community_playwright_go.PageEvaluateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageEvaluateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageEvaluateOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectConsoleMessageOptions"] = github_com_playwright_community_playwright_go.PageExpectConsoleMessageOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectConsoleMessageOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectConsoleMessageOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectDownloadOptions"] = github_com_playwright_community_playwright_go.PageExpectDownloadOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectDownloadOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectDownloadOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectEventOptions"] = github_com_playwright_community_playwright_go.PageExpectEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectFileChooserOptions"] = github_com_playwright_community_playwright_go.PageExpectFileChooserOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectFileChooserOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectFileChooserOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectPopupOptions"] = github_com_playwright_community_playwright_go.PageExpectPopupOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectPopupOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectPopupOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectRequestFinishedOptions"] = github_com_playwright_community_playwright_go.PageExpectRequestFinishedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectRequestFinishedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectRequestFinishedOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectWebSocketOptions"] = github_com_playwright_community_playwright_go.PageExpectWebSocketOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectWebSocketOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectWebSocketOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExpectWorkerOptions"] = github_com_playwright_community_playwright_go.PageExpectWorkerOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExpectWorkerOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExpectWorkerOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageExposeBindingOptions"] = github_com_playwright_community_playwright_go.PageExposeBindingOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageExposeBindingOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageExposeBindingOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageFillOptions"] = github_com_playwright_community_playwright_go.PageFillOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageFillOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageFillOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageFocusOptions"] = github_com_playwright_community_playwright_go.PageFocusOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageFocusOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageFocusOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageFrameOptions"] = github_com_playwright_community_playwright_go.PageFrameOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageFrameOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageFrameOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetAttributeOptions"] = github_com_playwright_community_playwright_go.PageGetAttributeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetAttributeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetAttributeOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByAltTextOptions"] = github_com_playwright_community_playwright_go.PageGetByAltTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByAltTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByAltTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByLabelOptions"] = github_com_playwright_community_playwright_go.PageGetByLabelOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByLabelOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByLabelOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByPlaceholderOptions"] = github_com_playwright_community_playwright_go.PageGetByPlaceholderOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByPlaceholderOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByPlaceholderOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByRoleOptions"] = github_com_playwright_community_playwright_go.PageGetByRoleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByRoleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByRoleOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByTextOptions"] = github_com_playwright_community_playwright_go.PageGetByTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGetByTitleOptions"] = github_com_playwright_community_playwright_go.PageGetByTitleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGetByTitleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGetByTitleOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGoBackOptions"] = github_com_playwright_community_playwright_go.PageGoBackOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGoBackOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGoBackOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGoForwardOptions"] = github_com_playwright_community_playwright_go.PageGoForwardOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGoForwardOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGoForwardOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageGotoOptions"] = github_com_playwright_community_playwright_go.PageGotoOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageGotoOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageGotoOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageHoverOptions"] = github_com_playwright_community_playwright_go.PageHoverOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageHoverOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageHoverOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageInnerHTMLOptions"] = github_com_playwright_community_playwright_go.PageInnerHTMLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageInnerHTMLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageInnerHTMLOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageInnerTextOptions"] = github_com_playwright_community_playwright_go.PageInnerTextOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageInnerTextOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageInnerTextOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageInputValueOptions"] = github_com_playwright_community_playwright_go.PageInputValueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageInputValueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageInputValueOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsCheckedOptions"] = github_com_playwright_community_playwright_go.PageIsCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsDisabledOptions"] = github_com_playwright_community_playwright_go.PageIsDisabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsDisabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsDisabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsEditableOptions"] = github_com_playwright_community_playwright_go.PageIsEditableOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsEditableOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsEditableOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsEnabledOptions"] = github_com_playwright_community_playwright_go.PageIsEnabledOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsEnabledOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsEnabledOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsHiddenOptions"] = github_com_playwright_community_playwright_go.PageIsHiddenOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsHiddenOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsHiddenOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageIsVisibleOptions"] = github_com_playwright_community_playwright_go.PageIsVisibleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageIsVisibleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageIsVisibleOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageLocatorOptions"] = github_com_playwright_community_playwright_go.PageLocatorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageLocatorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageLocatorOptions)
}
Api["github.com/playwright-community/playwright-go"]["PagePdfOptions"] = github_com_playwright_community_playwright_go.PagePdfOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PagePdfOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PagePdfOptions)
}
Api["github.com/playwright-community/playwright-go"]["PagePressOptions"] = github_com_playwright_community_playwright_go.PagePressOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PagePressOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PagePressOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageQuerySelectorOptions"] = github_com_playwright_community_playwright_go.PageQuerySelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageQuerySelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageQuerySelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageReloadOptions"] = github_com_playwright_community_playwright_go.PageReloadOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageReloadOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageReloadOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageRouteFromHAROptions"] = github_com_playwright_community_playwright_go.PageRouteFromHAROptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageRouteFromHAROptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageRouteFromHAROptions)
}
Api["github.com/playwright-community/playwright-go"]["PageRouteOptions"] = github_com_playwright_community_playwright_go.PageRouteOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageRouteOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageRouteOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageScreenshotOptions"] = github_com_playwright_community_playwright_go.PageScreenshotOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageScreenshotOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageScreenshotOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageSelectOptionOptions"] = github_com_playwright_community_playwright_go.PageSelectOptionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageSelectOptionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageSelectOptionOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageSetCheckedOptions"] = github_com_playwright_community_playwright_go.PageSetCheckedOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageSetCheckedOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageSetCheckedOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageSetContentOptions"] = github_com_playwright_community_playwright_go.PageSetContentOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageSetContentOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageSetContentOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageSetInputFilesOptions"] = github_com_playwright_community_playwright_go.PageSetInputFilesOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageSetInputFilesOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageSetInputFilesOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageTapOptions"] = github_com_playwright_community_playwright_go.PageTapOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageTapOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageTapOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageTextContentOptions"] = github_com_playwright_community_playwright_go.PageTextContentOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageTextContentOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageTextContentOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageTypeOptions"] = github_com_playwright_community_playwright_go.PageTypeOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageTypeOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageTypeOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageUncheckOptions"] = github_com_playwright_community_playwright_go.PageUncheckOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageUncheckOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageUncheckOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageUnrouteOptions"] = github_com_playwright_community_playwright_go.PageUnrouteOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageUnrouteOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageUnrouteOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForEventOptions"] = github_com_playwright_community_playwright_go.PageWaitForEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForFunctionOptions"] = github_com_playwright_community_playwright_go.PageWaitForFunctionOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForFunctionOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForFunctionOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForLoadStateOptions"] = github_com_playwright_community_playwright_go.PageWaitForLoadStateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForLoadStateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForLoadStateOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForNavigationOptions"] = github_com_playwright_community_playwright_go.PageWaitForNavigationOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForNavigationOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForNavigationOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForRequestOptions"] = github_com_playwright_community_playwright_go.PageWaitForRequestOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForRequestOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForRequestOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForResponseOptions"] = github_com_playwright_community_playwright_go.PageWaitForResponseOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForResponseOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForResponseOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForSelectorOptions"] = github_com_playwright_community_playwright_go.PageWaitForSelectorOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForSelectorOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForSelectorOptions)
}
Api["github.com/playwright-community/playwright-go"]["PageWaitForURLOptions"] = github_com_playwright_community_playwright_go.PageWaitForURLOptions{}
Api["github.com/playwright-community/playwright-go/new"]["PageWaitForURLOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PageWaitForURLOptions)
}
Api["github.com/playwright-community/playwright-go"]["Playwright"] = github_com_playwright_community_playwright_go.Playwright{}
Api["github.com/playwright-community/playwright-go/new"]["Playwright"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.Playwright)
}
Api["github.com/playwright-community/playwright-go"]["NewPlaywrightAssertions"] = github_com_playwright_community_playwright_go.NewPlaywrightAssertions
Api["github.com/playwright-community/playwright-go"]["PlaywrightDriver"] = github_com_playwright_community_playwright_go.PlaywrightDriver{}
Api["github.com/playwright-community/playwright-go/new"]["PlaywrightDriver"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.PlaywrightDriver)
}
Api["github.com/playwright-community/playwright-go"]["NewDriver"] = github_com_playwright_community_playwright_go.NewDriver
Api["github.com/playwright-community/playwright-go"]["Position"] = github_com_playwright_community_playwright_go.Position{}
Api["github.com/playwright-community/playwright-go/new"]["Position"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.Position)
}
Api["github.com/playwright-community/playwright-go"]["RecordVideo"] = github_com_playwright_community_playwright_go.RecordVideo{}
Api["github.com/playwright-community/playwright-go/new"]["RecordVideo"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RecordVideo)
}
Api["github.com/playwright-community/playwright-go"]["RecordVideoSize"] = github_com_playwright_community_playwright_go.RecordVideoSize{}
Api["github.com/playwright-community/playwright-go/new"]["RecordVideoSize"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RecordVideoSize)
}
Api["github.com/playwright-community/playwright-go"]["RequestFailure"] = github_com_playwright_community_playwright_go.RequestFailure{}
Api["github.com/playwright-community/playwright-go/new"]["RequestFailure"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RequestFailure)
}
Api["github.com/playwright-community/playwright-go"]["RequestSizesResult"] = github_com_playwright_community_playwright_go.RequestSizesResult{}
Api["github.com/playwright-community/playwright-go/new"]["RequestSizesResult"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RequestSizesResult)
}
Api["github.com/playwright-community/playwright-go"]["RequestTimingResult"] = github_com_playwright_community_playwright_go.RequestTimingResult{}
Api["github.com/playwright-community/playwright-go/new"]["RequestTimingResult"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RequestTimingResult)
}
Api["github.com/playwright-community/playwright-go"]["ResourceTiming"] = github_com_playwright_community_playwright_go.ResourceTiming{}
Api["github.com/playwright-community/playwright-go/new"]["ResourceTiming"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ResourceTiming)
}
Api["github.com/playwright-community/playwright-go"]["ResponseSecurityDetailsResult"] = github_com_playwright_community_playwright_go.ResponseSecurityDetailsResult{}
Api["github.com/playwright-community/playwright-go/new"]["ResponseSecurityDetailsResult"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ResponseSecurityDetailsResult)
}
Api["github.com/playwright-community/playwright-go"]["ResponseServerAddrResult"] = github_com_playwright_community_playwright_go.ResponseServerAddrResult{}
Api["github.com/playwright-community/playwright-go/new"]["ResponseServerAddrResult"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ResponseServerAddrResult)
}
Api["github.com/playwright-community/playwright-go"]["RouteAbortOptions"] = github_com_playwright_community_playwright_go.RouteAbortOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RouteAbortOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RouteAbortOptions)
}
Api["github.com/playwright-community/playwright-go"]["RouteContinueOptions"] = github_com_playwright_community_playwright_go.RouteContinueOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RouteContinueOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RouteContinueOptions)
}
Api["github.com/playwright-community/playwright-go"]["RouteFallbackOptions"] = github_com_playwright_community_playwright_go.RouteFallbackOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RouteFallbackOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RouteFallbackOptions)
}
Api["github.com/playwright-community/playwright-go"]["RouteFetchOptions"] = github_com_playwright_community_playwright_go.RouteFetchOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RouteFetchOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RouteFetchOptions)
}
Api["github.com/playwright-community/playwright-go"]["RouteFulfillOptions"] = github_com_playwright_community_playwright_go.RouteFulfillOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RouteFulfillOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RouteFulfillOptions)
}
Api["github.com/playwright-community/playwright-go"]["RunOptions"] = github_com_playwright_community_playwright_go.RunOptions{}
Api["github.com/playwright-community/playwright-go/new"]["RunOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.RunOptions)
}
Api["github.com/playwright-community/playwright-go"]["ScreenSize"] = github_com_playwright_community_playwright_go.ScreenSize{}
Api["github.com/playwright-community/playwright-go/new"]["ScreenSize"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ScreenSize)
}
Api["github.com/playwright-community/playwright-go"]["SelectOptionValues"] = github_com_playwright_community_playwright_go.SelectOptionValues{}
Api["github.com/playwright-community/playwright-go/new"]["SelectOptionValues"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.SelectOptionValues)
}
Api["github.com/playwright-community/playwright-go"]["SelectorsRegisterOptions"] = github_com_playwright_community_playwright_go.SelectorsRegisterOptions{}
Api["github.com/playwright-community/playwright-go/new"]["SelectorsRegisterOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.SelectorsRegisterOptions)
}
Api["github.com/playwright-community/playwright-go"]["StorageState"] = github_com_playwright_community_playwright_go.StorageState{}
Api["github.com/playwright-community/playwright-go/new"]["StorageState"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.StorageState)
}
Api["github.com/playwright-community/playwright-go"]["TracingStartChunkOptions"] = github_com_playwright_community_playwright_go.TracingStartChunkOptions{}
Api["github.com/playwright-community/playwright-go/new"]["TracingStartChunkOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.TracingStartChunkOptions)
}
Api["github.com/playwright-community/playwright-go"]["TracingStartOptions"] = github_com_playwright_community_playwright_go.TracingStartOptions{}
Api["github.com/playwright-community/playwright-go/new"]["TracingStartOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.TracingStartOptions)
}
Api["github.com/playwright-community/playwright-go"]["TracingStopChunkOptions"] = github_com_playwright_community_playwright_go.TracingStopChunkOptions{}
Api["github.com/playwright-community/playwright-go/new"]["TracingStopChunkOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.TracingStopChunkOptions)
}
Api["github.com/playwright-community/playwright-go"]["TracingStopOptions"] = github_com_playwright_community_playwright_go.TracingStopOptions{}
Api["github.com/playwright-community/playwright-go/new"]["TracingStopOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.TracingStopOptions)
}
Api["github.com/playwright-community/playwright-go"]["ViewportSize"] = github_com_playwright_community_playwright_go.ViewportSize{}
Api["github.com/playwright-community/playwright-go/new"]["ViewportSize"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.ViewportSize)
}
Api["github.com/playwright-community/playwright-go"]["WebSocketExpectEventOptions"] = github_com_playwright_community_playwright_go.WebSocketExpectEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["WebSocketExpectEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.WebSocketExpectEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["WebSocketWaitForEventOptions"] = github_com_playwright_community_playwright_go.WebSocketWaitForEventOptions{}
Api["github.com/playwright-community/playwright-go/new"]["WebSocketWaitForEventOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.WebSocketWaitForEventOptions)
}
Api["github.com/playwright-community/playwright-go"]["WorkerEvaluateHandleOptions"] = github_com_playwright_community_playwright_go.WorkerEvaluateHandleOptions{}
Api["github.com/playwright-community/playwright-go/new"]["WorkerEvaluateHandleOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.WorkerEvaluateHandleOptions)
}
Api["github.com/playwright-community/playwright-go"]["WorkerEvaluateOptions"] = github_com_playwright_community_playwright_go.WorkerEvaluateOptions{}
Api["github.com/playwright-community/playwright-go/new"]["WorkerEvaluateOptions"] = func(args ...interface{}) interface{} {
					return new(github_com_playwright_community_playwright_go.WorkerEvaluateOptions)
}
}
