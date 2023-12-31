package go_chart
import (
github_com_wcharczuk_go_chart "github.com/wcharczuk/go-chart"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/wcharczuk/go-chart"]; !ok {
   Api["github.com/wcharczuk/go-chart"] = map[string]interface{}{}
}
if _, ok := Api["github.com/wcharczuk/go-chart/new"]; !ok {
	Api["github.com/wcharczuk/go-chart/new"] = map[string]interface{}{}
}
Api["github.com/wcharczuk/go-chart"]["DefaultChartHeight"] = github_com_wcharczuk_go_chart.DefaultChartHeight
Api["github.com/wcharczuk/go-chart"]["DefaultChartWidth"] = github_com_wcharczuk_go_chart.DefaultChartWidth
Api["github.com/wcharczuk/go-chart"]["DefaultStrokeWidth"] = github_com_wcharczuk_go_chart.DefaultStrokeWidth
Api["github.com/wcharczuk/go-chart"]["DefaultDotWidth"] = github_com_wcharczuk_go_chart.DefaultDotWidth
Api["github.com/wcharczuk/go-chart"]["DefaultSeriesLineWidth"] = github_com_wcharczuk_go_chart.DefaultSeriesLineWidth
Api["github.com/wcharczuk/go-chart"]["DefaultAxisLineWidth"] = github_com_wcharczuk_go_chart.DefaultAxisLineWidth
Api["github.com/wcharczuk/go-chart"]["DefaultDPI"] = github_com_wcharczuk_go_chart.DefaultDPI
Api["github.com/wcharczuk/go-chart"]["DefaultMinimumFontSize"] = github_com_wcharczuk_go_chart.DefaultMinimumFontSize
Api["github.com/wcharczuk/go-chart"]["DefaultFontSize"] = github_com_wcharczuk_go_chart.DefaultFontSize
Api["github.com/wcharczuk/go-chart"]["DefaultTitleFontSize"] = github_com_wcharczuk_go_chart.DefaultTitleFontSize
Api["github.com/wcharczuk/go-chart"]["DefaultAnnotationDeltaWidth"] = github_com_wcharczuk_go_chart.DefaultAnnotationDeltaWidth
Api["github.com/wcharczuk/go-chart"]["DefaultAnnotationFontSize"] = github_com_wcharczuk_go_chart.DefaultAnnotationFontSize
Api["github.com/wcharczuk/go-chart"]["DefaultAxisFontSize"] = github_com_wcharczuk_go_chart.DefaultAxisFontSize
Api["github.com/wcharczuk/go-chart"]["DefaultTitleTop"] = github_com_wcharczuk_go_chart.DefaultTitleTop
Api["github.com/wcharczuk/go-chart"]["DefaultBackgroundStrokeWidth"] = github_com_wcharczuk_go_chart.DefaultBackgroundStrokeWidth
Api["github.com/wcharczuk/go-chart"]["DefaultCanvasStrokeWidth"] = github_com_wcharczuk_go_chart.DefaultCanvasStrokeWidth
Api["github.com/wcharczuk/go-chart"]["DefaultLineSpacing"] = github_com_wcharczuk_go_chart.DefaultLineSpacing
Api["github.com/wcharczuk/go-chart"]["DefaultYAxisMargin"] = github_com_wcharczuk_go_chart.DefaultYAxisMargin
Api["github.com/wcharczuk/go-chart"]["DefaultXAxisMargin"] = github_com_wcharczuk_go_chart.DefaultXAxisMargin
Api["github.com/wcharczuk/go-chart"]["DefaultVerticalTickHeight"] = github_com_wcharczuk_go_chart.DefaultVerticalTickHeight
Api["github.com/wcharczuk/go-chart"]["DefaultHorizontalTickWidth"] = github_com_wcharczuk_go_chart.DefaultHorizontalTickWidth
Api["github.com/wcharczuk/go-chart"]["DefaultTickCount"] = github_com_wcharczuk_go_chart.DefaultTickCount
Api["github.com/wcharczuk/go-chart"]["DefaultTickCountSanityCheck"] = github_com_wcharczuk_go_chart.DefaultTickCountSanityCheck
Api["github.com/wcharczuk/go-chart"]["DefaultMinimumTickHorizontalSpacing"] = github_com_wcharczuk_go_chart.DefaultMinimumTickHorizontalSpacing
Api["github.com/wcharczuk/go-chart"]["DefaultMinimumTickVerticalSpacing"] = github_com_wcharczuk_go_chart.DefaultMinimumTickVerticalSpacing
Api["github.com/wcharczuk/go-chart"]["DefaultDateFormat"] = github_com_wcharczuk_go_chart.DefaultDateFormat
Api["github.com/wcharczuk/go-chart"]["DefaultDateHourFormat"] = github_com_wcharczuk_go_chart.DefaultDateHourFormat
Api["github.com/wcharczuk/go-chart"]["DefaultDateMinuteFormat"] = github_com_wcharczuk_go_chart.DefaultDateMinuteFormat
Api["github.com/wcharczuk/go-chart"]["DefaultFloatFormat"] = github_com_wcharczuk_go_chart.DefaultFloatFormat
Api["github.com/wcharczuk/go-chart"]["DefaultPercentValueFormat"] = github_com_wcharczuk_go_chart.DefaultPercentValueFormat
Api["github.com/wcharczuk/go-chart"]["DefaultBarSpacing"] = github_com_wcharczuk_go_chart.DefaultBarSpacing
Api["github.com/wcharczuk/go-chart"]["DefaultBarWidth"] = github_com_wcharczuk_go_chart.DefaultBarWidth
Api["github.com/wcharczuk/go-chart"]["ContentTypePNG"] = github_com_wcharczuk_go_chart.ContentTypePNG
Api["github.com/wcharczuk/go-chart"]["ContentTypeSVG"] = github_com_wcharczuk_go_chart.ContentTypeSVG
Api["github.com/wcharczuk/go-chart"]["DefaultMACDPeriodPrimary"] = github_com_wcharczuk_go_chart.DefaultMACDPeriodPrimary
Api["github.com/wcharczuk/go-chart"]["DefaultMACDPeriodSecondary"] = github_com_wcharczuk_go_chart.DefaultMACDPeriodSecondary
Api["github.com/wcharczuk/go-chart"]["DefaultMACDSignalPeriod"] = github_com_wcharczuk_go_chart.DefaultMACDSignalPeriod
Api["github.com/wcharczuk/go-chart"]["DefaultEMAPeriod"] = github_com_wcharczuk_go_chart.DefaultEMAPeriod
Api["github.com/wcharczuk/go-chart"]["DefaultSimpleMovingAveragePeriod"] = github_com_wcharczuk_go_chart.DefaultSimpleMovingAveragePeriod
Api["github.com/wcharczuk/go-chart"]["Disabled"] = github_com_wcharczuk_go_chart.Disabled
Api["github.com/wcharczuk/go-chart"]["AlternateColorPalette"] = github_com_wcharczuk_go_chart.AlternateColorPalette
Api["github.com/wcharczuk/go-chart"]["DefaultColorPalette"] = github_com_wcharczuk_go_chart.DefaultColorPalette
Api["github.com/wcharczuk/go-chart"]["FloatValueFormatter"] = github_com_wcharczuk_go_chart.FloatValueFormatter
Api["github.com/wcharczuk/go-chart"]["FloatValueFormatterWithFormat"] = github_com_wcharczuk_go_chart.FloatValueFormatterWithFormat
Api["github.com/wcharczuk/go-chart"]["GetAlternateColor"] = github_com_wcharczuk_go_chart.GetAlternateColor
Api["github.com/wcharczuk/go-chart"]["GetDefaultColor"] = github_com_wcharczuk_go_chart.GetDefaultColor
Api["github.com/wcharczuk/go-chart"]["GetDefaultFont"] = github_com_wcharczuk_go_chart.GetDefaultFont
Api["github.com/wcharczuk/go-chart"]["Jet"] = github_com_wcharczuk_go_chart.Jet
Api["github.com/wcharczuk/go-chart"]["PercentValueFormatter"] = github_com_wcharczuk_go_chart.PercentValueFormatter
Api["github.com/wcharczuk/go-chart"]["TimeDateValueFormatter"] = github_com_wcharczuk_go_chart.TimeDateValueFormatter
Api["github.com/wcharczuk/go-chart"]["TimeHourValueFormatter"] = github_com_wcharczuk_go_chart.TimeHourValueFormatter
Api["github.com/wcharczuk/go-chart"]["TimeMinuteValueFormatter"] = github_com_wcharczuk_go_chart.TimeMinuteValueFormatter
Api["github.com/wcharczuk/go-chart"]["TimeValueFormatter"] = github_com_wcharczuk_go_chart.TimeValueFormatter
Api["github.com/wcharczuk/go-chart"]["Viridis"] = github_com_wcharczuk_go_chart.Viridis
Api["github.com/wcharczuk/go-chart"]["AnnotationSeries"] = github_com_wcharczuk_go_chart.AnnotationSeries{}
Api["github.com/wcharczuk/go-chart/new"]["AnnotationSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.AnnotationSeries)
}
Api["github.com/wcharczuk/go-chart"]["LastValueAnnotation"] = github_com_wcharczuk_go_chart.LastValueAnnotation
Api["github.com/wcharczuk/go-chart"]["BarChart"] = github_com_wcharczuk_go_chart.BarChart{}
Api["github.com/wcharczuk/go-chart/new"]["BarChart"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.BarChart)
}
Api["github.com/wcharczuk/go-chart"]["BollingerBandsSeries"] = github_com_wcharczuk_go_chart.BollingerBandsSeries{}
Api["github.com/wcharczuk/go-chart/new"]["BollingerBandsSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.BollingerBandsSeries)
}
Api["github.com/wcharczuk/go-chart"]["NewBox"] = github_com_wcharczuk_go_chart.NewBox
Api["github.com/wcharczuk/go-chart"]["BoxCorners"] = github_com_wcharczuk_go_chart.BoxCorners{}
Api["github.com/wcharczuk/go-chart/new"]["BoxCorners"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.BoxCorners)
}
Api["github.com/wcharczuk/go-chart"]["Chart"] = github_com_wcharczuk_go_chart.Chart{}
Api["github.com/wcharczuk/go-chart/new"]["Chart"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Chart)
}
Api["github.com/wcharczuk/go-chart"]["ContinuousRange"] = github_com_wcharczuk_go_chart.ContinuousRange{}
Api["github.com/wcharczuk/go-chart/new"]["ContinuousRange"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.ContinuousRange)
}
Api["github.com/wcharczuk/go-chart"]["ContinuousSeries"] = github_com_wcharczuk_go_chart.ContinuousSeries{}
Api["github.com/wcharczuk/go-chart/new"]["ContinuousSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.ContinuousSeries)
}
Api["github.com/wcharczuk/go-chart"]["EMASeries"] = github_com_wcharczuk_go_chart.EMASeries{}
Api["github.com/wcharczuk/go-chart/new"]["EMASeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.EMASeries)
}
Api["github.com/wcharczuk/go-chart"]["GridLine"] = github_com_wcharczuk_go_chart.GridLine{}
Api["github.com/wcharczuk/go-chart/new"]["GridLine"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.GridLine)
}
Api["github.com/wcharczuk/go-chart"]["GenerateGridLines"] = github_com_wcharczuk_go_chart.GenerateGridLines
Api["github.com/wcharczuk/go-chart"]["HistogramSeries"] = github_com_wcharczuk_go_chart.HistogramSeries{}
Api["github.com/wcharczuk/go-chart/new"]["HistogramSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.HistogramSeries)
}
Api["github.com/wcharczuk/go-chart"]["ImageWriter"] = github_com_wcharczuk_go_chart.ImageWriter{}
Api["github.com/wcharczuk/go-chart/new"]["ImageWriter"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.ImageWriter)
}
Api["github.com/wcharczuk/go-chart"]["LinearRegressionSeries"] = github_com_wcharczuk_go_chart.LinearRegressionSeries{}
Api["github.com/wcharczuk/go-chart/new"]["LinearRegressionSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.LinearRegressionSeries)
}
Api["github.com/wcharczuk/go-chart"]["MACDLineSeries"] = github_com_wcharczuk_go_chart.MACDLineSeries{}
Api["github.com/wcharczuk/go-chart/new"]["MACDLineSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MACDLineSeries)
}
Api["github.com/wcharczuk/go-chart"]["MACDSeries"] = github_com_wcharczuk_go_chart.MACDSeries{}
Api["github.com/wcharczuk/go-chart/new"]["MACDSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MACDSeries)
}
Api["github.com/wcharczuk/go-chart"]["MACDSignalSeries"] = github_com_wcharczuk_go_chart.MACDSignalSeries{}
Api["github.com/wcharczuk/go-chart/new"]["MACDSignalSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MACDSignalSeries)
}
Api["github.com/wcharczuk/go-chart"]["MarketHoursRange"] = github_com_wcharczuk_go_chart.MarketHoursRange{}
Api["github.com/wcharczuk/go-chart/new"]["MarketHoursRange"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MarketHoursRange)
}
Api["github.com/wcharczuk/go-chart"]["MaxSeries"] = github_com_wcharczuk_go_chart.MaxSeries{}
Api["github.com/wcharczuk/go-chart/new"]["MaxSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MaxSeries)
}
Api["github.com/wcharczuk/go-chart"]["MinSeries"] = github_com_wcharczuk_go_chart.MinSeries{}
Api["github.com/wcharczuk/go-chart/new"]["MinSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.MinSeries)
}
Api["github.com/wcharczuk/go-chart"]["PieChart"] = github_com_wcharczuk_go_chart.PieChart{}
Api["github.com/wcharczuk/go-chart/new"]["PieChart"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.PieChart)
}
Api["github.com/wcharczuk/go-chart"]["Point"] = github_com_wcharczuk_go_chart.Point{}
Api["github.com/wcharczuk/go-chart/new"]["Point"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Point)
}
Api["github.com/wcharczuk/go-chart"]["PolynomialRegressionSeries"] = github_com_wcharczuk_go_chart.PolynomialRegressionSeries{}
Api["github.com/wcharczuk/go-chart/new"]["PolynomialRegressionSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.PolynomialRegressionSeries)
}
Api["github.com/wcharczuk/go-chart"]["Legend"] = github_com_wcharczuk_go_chart.Legend
Api["github.com/wcharczuk/go-chart"]["LegendLeft"] = github_com_wcharczuk_go_chart.LegendLeft
Api["github.com/wcharczuk/go-chart"]["LegendThin"] = github_com_wcharczuk_go_chart.LegendThin
Api["github.com/wcharczuk/go-chart"]["PNG"] = github_com_wcharczuk_go_chart.PNG
Api["github.com/wcharczuk/go-chart"]["SVG"] = github_com_wcharczuk_go_chart.SVG
Api["github.com/wcharczuk/go-chart"]["SMASeries"] = github_com_wcharczuk_go_chart.SMASeries{}
Api["github.com/wcharczuk/go-chart/new"]["SMASeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.SMASeries)
}
Api["github.com/wcharczuk/go-chart"]["StackedBar"] = github_com_wcharczuk_go_chart.StackedBar{}
Api["github.com/wcharczuk/go-chart/new"]["StackedBar"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.StackedBar)
}
Api["github.com/wcharczuk/go-chart"]["StackedBarChart"] = github_com_wcharczuk_go_chart.StackedBarChart{}
Api["github.com/wcharczuk/go-chart/new"]["StackedBarChart"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.StackedBarChart)
}
Api["github.com/wcharczuk/go-chart"]["Style"] = github_com_wcharczuk_go_chart.Style{}
Api["github.com/wcharczuk/go-chart/new"]["Style"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Style)
}
Api["github.com/wcharczuk/go-chart"]["StyleShow"] = github_com_wcharczuk_go_chart.StyleShow
Api["github.com/wcharczuk/go-chart"]["StyleTextDefaults"] = github_com_wcharczuk_go_chart.StyleTextDefaults
Api["github.com/wcharczuk/go-chart"]["TextHorizontalAlignUnset"] = github_com_wcharczuk_go_chart.TextHorizontalAlignUnset
Api["github.com/wcharczuk/go-chart"]["TextHorizontalAlignLeft"] = github_com_wcharczuk_go_chart.TextHorizontalAlignLeft
Api["github.com/wcharczuk/go-chart"]["TextHorizontalAlignCenter"] = github_com_wcharczuk_go_chart.TextHorizontalAlignCenter
Api["github.com/wcharczuk/go-chart"]["TextHorizontalAlignRight"] = github_com_wcharczuk_go_chart.TextHorizontalAlignRight
Api["github.com/wcharczuk/go-chart"]["TextStyle"] = github_com_wcharczuk_go_chart.TextStyle{}
Api["github.com/wcharczuk/go-chart/new"]["TextStyle"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.TextStyle)
}
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignUnset"] = github_com_wcharczuk_go_chart.TextVerticalAlignUnset
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignBaseline"] = github_com_wcharczuk_go_chart.TextVerticalAlignBaseline
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignBottom"] = github_com_wcharczuk_go_chart.TextVerticalAlignBottom
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignMiddle"] = github_com_wcharczuk_go_chart.TextVerticalAlignMiddle
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignMiddleBaseline"] = github_com_wcharczuk_go_chart.TextVerticalAlignMiddleBaseline
Api["github.com/wcharczuk/go-chart"]["TextVerticalAlignTop"] = github_com_wcharczuk_go_chart.TextVerticalAlignTop
Api["github.com/wcharczuk/go-chart"]["TextWrapUnset"] = github_com_wcharczuk_go_chart.TextWrapUnset
Api["github.com/wcharczuk/go-chart"]["TextWrapNone"] = github_com_wcharczuk_go_chart.TextWrapNone
Api["github.com/wcharczuk/go-chart"]["TextWrapWord"] = github_com_wcharczuk_go_chart.TextWrapWord
Api["github.com/wcharczuk/go-chart"]["TextWrapRune"] = github_com_wcharczuk_go_chart.TextWrapRune
Api["github.com/wcharczuk/go-chart"]["Tick"] = github_com_wcharczuk_go_chart.Tick{}
Api["github.com/wcharczuk/go-chart/new"]["Tick"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Tick)
}
Api["github.com/wcharczuk/go-chart"]["GenerateContinuousTicks"] = github_com_wcharczuk_go_chart.GenerateContinuousTicks
Api["github.com/wcharczuk/go-chart"]["TickPositionUnset"] = github_com_wcharczuk_go_chart.TickPositionUnset
Api["github.com/wcharczuk/go-chart"]["TickPositionBetweenTicks"] = github_com_wcharczuk_go_chart.TickPositionBetweenTicks
Api["github.com/wcharczuk/go-chart"]["TickPositionUnderTick"] = github_com_wcharczuk_go_chart.TickPositionUnderTick
Api["github.com/wcharczuk/go-chart"]["TimeSeries"] = github_com_wcharczuk_go_chart.TimeSeries{}
Api["github.com/wcharczuk/go-chart/new"]["TimeSeries"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.TimeSeries)
}
Api["github.com/wcharczuk/go-chart"]["Value"] = github_com_wcharczuk_go_chart.Value{}
Api["github.com/wcharczuk/go-chart/new"]["Value"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Value)
}
Api["github.com/wcharczuk/go-chart"]["Value2"] = github_com_wcharczuk_go_chart.Value2{}
Api["github.com/wcharczuk/go-chart/new"]["Value2"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.Value2)
}
Api["github.com/wcharczuk/go-chart"]["TimeValueFormatterWithFormat"] = github_com_wcharczuk_go_chart.TimeValueFormatterWithFormat
Api["github.com/wcharczuk/go-chart"]["XAxis"] = github_com_wcharczuk_go_chart.XAxis{}
Api["github.com/wcharczuk/go-chart/new"]["XAxis"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.XAxis)
}
Api["github.com/wcharczuk/go-chart"]["YAxis"] = github_com_wcharczuk_go_chart.YAxis{}
Api["github.com/wcharczuk/go-chart/new"]["YAxis"] = func(args ...interface{}) interface{} {
					return new(github_com_wcharczuk_go_chart.YAxis)
}
Api["github.com/wcharczuk/go-chart"]["YAxisPrimary"] = github_com_wcharczuk_go_chart.YAxisPrimary
Api["github.com/wcharczuk/go-chart"]["YAxisSecondary"] = github_com_wcharczuk_go_chart.YAxisSecondary
}
