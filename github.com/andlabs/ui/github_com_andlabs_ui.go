package ui
import (
github_com_andlabs_ui "github.com/andlabs/ui"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/andlabs/ui"]; !ok {
   Api["github.com/andlabs/ui"] = map[string]interface{}{}
}
if _, ok := Api["github.com/andlabs/ui/new"]; !ok {
	Api["github.com/andlabs/ui/new"] = map[string]interface{}{}
}
Api["github.com/andlabs/ui"]["TableModelColumnNeverEditable"] = github_com_andlabs_ui.TableModelColumnNeverEditable
Api["github.com/andlabs/ui"]["TableModelColumnAlwaysEditable"] = github_com_andlabs_ui.TableModelColumnAlwaysEditable
Api["github.com/andlabs/ui"]["LibuiFreeText"] = github_com_andlabs_ui.LibuiFreeText
Api["github.com/andlabs/ui"]["Main"] = github_com_andlabs_ui.Main
Api["github.com/andlabs/ui"]["MsgBox"] = github_com_andlabs_ui.MsgBox
Api["github.com/andlabs/ui"]["MsgBoxError"] = github_com_andlabs_ui.MsgBoxError
Api["github.com/andlabs/ui"]["OnShouldQuit"] = github_com_andlabs_ui.OnShouldQuit
Api["github.com/andlabs/ui"]["OpenFile"] = github_com_andlabs_ui.OpenFile
Api["github.com/andlabs/ui"]["QueueMain"] = github_com_andlabs_ui.QueueMain
Api["github.com/andlabs/ui"]["Quit"] = github_com_andlabs_ui.Quit
Api["github.com/andlabs/ui"]["SaveFile"] = github_com_andlabs_ui.SaveFile
Api["github.com/andlabs/ui"]["AlignFill"] = github_com_andlabs_ui.AlignFill
Api["github.com/andlabs/ui"]["AlignStart"] = github_com_andlabs_ui.AlignStart
Api["github.com/andlabs/ui"]["AlignCenter"] = github_com_andlabs_ui.AlignCenter
Api["github.com/andlabs/ui"]["AlignEnd"] = github_com_andlabs_ui.AlignEnd
Api["github.com/andlabs/ui"]["Area"] = github_com_andlabs_ui.Area{}
Api["github.com/andlabs/ui/new"]["Area"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Area)
}
Api["github.com/andlabs/ui"]["NewArea"] = github_com_andlabs_ui.NewArea
Api["github.com/andlabs/ui"]["NewScrollingArea"] = github_com_andlabs_ui.NewScrollingArea
Api["github.com/andlabs/ui"]["AreaDrawParams"] = github_com_andlabs_ui.AreaDrawParams{}
Api["github.com/andlabs/ui/new"]["AreaDrawParams"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.AreaDrawParams)
}
Api["github.com/andlabs/ui"]["AreaKeyEvent"] = github_com_andlabs_ui.AreaKeyEvent{}
Api["github.com/andlabs/ui/new"]["AreaKeyEvent"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.AreaKeyEvent)
}
Api["github.com/andlabs/ui"]["AreaMouseEvent"] = github_com_andlabs_ui.AreaMouseEvent{}
Api["github.com/andlabs/ui/new"]["AreaMouseEvent"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.AreaMouseEvent)
}
Api["github.com/andlabs/ui"]["Leading"] = github_com_andlabs_ui.Leading
Api["github.com/andlabs/ui"]["Top"] = github_com_andlabs_ui.Top
Api["github.com/andlabs/ui"]["Trailing"] = github_com_andlabs_ui.Trailing
Api["github.com/andlabs/ui"]["Bottom"] = github_com_andlabs_ui.Bottom
Api["github.com/andlabs/ui"]["AttributedString"] = github_com_andlabs_ui.AttributedString{}
Api["github.com/andlabs/ui/new"]["AttributedString"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.AttributedString)
}
Api["github.com/andlabs/ui"]["NewAttributedString"] = github_com_andlabs_ui.NewAttributedString
Api["github.com/andlabs/ui"]["Box"] = github_com_andlabs_ui.Box{}
Api["github.com/andlabs/ui/new"]["Box"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Box)
}
Api["github.com/andlabs/ui"]["NewHorizontalBox"] = github_com_andlabs_ui.NewHorizontalBox
Api["github.com/andlabs/ui"]["NewVerticalBox"] = github_com_andlabs_ui.NewVerticalBox
Api["github.com/andlabs/ui"]["Button"] = github_com_andlabs_ui.Button{}
Api["github.com/andlabs/ui/new"]["Button"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Button)
}
Api["github.com/andlabs/ui"]["NewButton"] = github_com_andlabs_ui.NewButton
Api["github.com/andlabs/ui"]["Checkbox"] = github_com_andlabs_ui.Checkbox{}
Api["github.com/andlabs/ui/new"]["Checkbox"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Checkbox)
}
Api["github.com/andlabs/ui"]["NewCheckbox"] = github_com_andlabs_ui.NewCheckbox
Api["github.com/andlabs/ui"]["ColorButton"] = github_com_andlabs_ui.ColorButton{}
Api["github.com/andlabs/ui/new"]["ColorButton"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.ColorButton)
}
Api["github.com/andlabs/ui"]["NewColorButton"] = github_com_andlabs_ui.NewColorButton
Api["github.com/andlabs/ui"]["Combobox"] = github_com_andlabs_ui.Combobox{}
Api["github.com/andlabs/ui/new"]["Combobox"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Combobox)
}
Api["github.com/andlabs/ui"]["NewCombobox"] = github_com_andlabs_ui.NewCombobox
Api["github.com/andlabs/ui"]["ControlFromLibui"] = github_com_andlabs_ui.ControlFromLibui
Api["github.com/andlabs/ui"]["ControlBase"] = github_com_andlabs_ui.ControlBase{}
Api["github.com/andlabs/ui/new"]["ControlBase"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.ControlBase)
}
Api["github.com/andlabs/ui"]["NewControlBase"] = github_com_andlabs_ui.NewControlBase
Api["github.com/andlabs/ui"]["DateTimePicker"] = github_com_andlabs_ui.DateTimePicker{}
Api["github.com/andlabs/ui/new"]["DateTimePicker"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DateTimePicker)
}
Api["github.com/andlabs/ui"]["NewDatePicker"] = github_com_andlabs_ui.NewDatePicker
Api["github.com/andlabs/ui"]["NewDateTimePicker"] = github_com_andlabs_ui.NewDateTimePicker
Api["github.com/andlabs/ui"]["NewTimePicker"] = github_com_andlabs_ui.NewTimePicker
Api["github.com/andlabs/ui"]["DrawBrush"] = github_com_andlabs_ui.DrawBrush{}
Api["github.com/andlabs/ui/new"]["DrawBrush"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawBrush)
}
Api["github.com/andlabs/ui"]["DrawBrushTypeSolid"] = github_com_andlabs_ui.DrawBrushTypeSolid
Api["github.com/andlabs/ui"]["DrawBrushTypeLinearGradient"] = github_com_andlabs_ui.DrawBrushTypeLinearGradient
Api["github.com/andlabs/ui"]["DrawBrushTypeRadialGradient"] = github_com_andlabs_ui.DrawBrushTypeRadialGradient
Api["github.com/andlabs/ui"]["DrawBrushTypeImage"] = github_com_andlabs_ui.DrawBrushTypeImage
Api["github.com/andlabs/ui"]["DrawContext"] = github_com_andlabs_ui.DrawContext{}
Api["github.com/andlabs/ui/new"]["DrawContext"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawContext)
}
Api["github.com/andlabs/ui"]["DrawFillModeWinding"] = github_com_andlabs_ui.DrawFillModeWinding
Api["github.com/andlabs/ui"]["DrawFillModeAlternate"] = github_com_andlabs_ui.DrawFillModeAlternate
Api["github.com/andlabs/ui"]["DrawGradientStop"] = github_com_andlabs_ui.DrawGradientStop{}
Api["github.com/andlabs/ui/new"]["DrawGradientStop"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawGradientStop)
}
Api["github.com/andlabs/ui"]["DrawLineCapFlat"] = github_com_andlabs_ui.DrawLineCapFlat
Api["github.com/andlabs/ui"]["DrawLineCapRound"] = github_com_andlabs_ui.DrawLineCapRound
Api["github.com/andlabs/ui"]["DrawLineCapSquare"] = github_com_andlabs_ui.DrawLineCapSquare
Api["github.com/andlabs/ui"]["DrawLineJoinMiter"] = github_com_andlabs_ui.DrawLineJoinMiter
Api["github.com/andlabs/ui"]["DrawLineJoinRound"] = github_com_andlabs_ui.DrawLineJoinRound
Api["github.com/andlabs/ui"]["DrawLineJoinBevel"] = github_com_andlabs_ui.DrawLineJoinBevel
Api["github.com/andlabs/ui"]["DrawMatrix"] = github_com_andlabs_ui.DrawMatrix{}
Api["github.com/andlabs/ui/new"]["DrawMatrix"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawMatrix)
}
Api["github.com/andlabs/ui"]["DrawNewMatrix"] = github_com_andlabs_ui.DrawNewMatrix
Api["github.com/andlabs/ui"]["DrawPath"] = github_com_andlabs_ui.DrawPath{}
Api["github.com/andlabs/ui/new"]["DrawPath"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawPath)
}
Api["github.com/andlabs/ui"]["DrawNewPath"] = github_com_andlabs_ui.DrawNewPath
Api["github.com/andlabs/ui"]["DrawStrokeParams"] = github_com_andlabs_ui.DrawStrokeParams{}
Api["github.com/andlabs/ui/new"]["DrawStrokeParams"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawStrokeParams)
}
Api["github.com/andlabs/ui"]["DrawTextAlignLeft"] = github_com_andlabs_ui.DrawTextAlignLeft
Api["github.com/andlabs/ui"]["DrawTextAlignCenter"] = github_com_andlabs_ui.DrawTextAlignCenter
Api["github.com/andlabs/ui"]["DrawTextAlignRight"] = github_com_andlabs_ui.DrawTextAlignRight
Api["github.com/andlabs/ui"]["DrawTextLayout"] = github_com_andlabs_ui.DrawTextLayout{}
Api["github.com/andlabs/ui/new"]["DrawTextLayout"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawTextLayout)
}
Api["github.com/andlabs/ui"]["DrawNewTextLayout"] = github_com_andlabs_ui.DrawNewTextLayout
Api["github.com/andlabs/ui"]["DrawTextLayoutParams"] = github_com_andlabs_ui.DrawTextLayoutParams{}
Api["github.com/andlabs/ui/new"]["DrawTextLayoutParams"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.DrawTextLayoutParams)
}
Api["github.com/andlabs/ui"]["EditableCombobox"] = github_com_andlabs_ui.EditableCombobox{}
Api["github.com/andlabs/ui/new"]["EditableCombobox"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.EditableCombobox)
}
Api["github.com/andlabs/ui"]["NewEditableCombobox"] = github_com_andlabs_ui.NewEditableCombobox
Api["github.com/andlabs/ui"]["Entry"] = github_com_andlabs_ui.Entry{}
Api["github.com/andlabs/ui/new"]["Entry"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Entry)
}
Api["github.com/andlabs/ui"]["NewEntry"] = github_com_andlabs_ui.NewEntry
Api["github.com/andlabs/ui"]["NewPasswordEntry"] = github_com_andlabs_ui.NewPasswordEntry
Api["github.com/andlabs/ui"]["NewSearchEntry"] = github_com_andlabs_ui.NewSearchEntry
Api["github.com/andlabs/ui"]["Escape"] = github_com_andlabs_ui.Escape
Api["github.com/andlabs/ui"]["Insert"] = github_com_andlabs_ui.Insert
Api["github.com/andlabs/ui"]["Delete"] = github_com_andlabs_ui.Delete
Api["github.com/andlabs/ui"]["Home"] = github_com_andlabs_ui.Home
Api["github.com/andlabs/ui"]["End"] = github_com_andlabs_ui.End
Api["github.com/andlabs/ui"]["PageUp"] = github_com_andlabs_ui.PageUp
Api["github.com/andlabs/ui"]["PageDown"] = github_com_andlabs_ui.PageDown
Api["github.com/andlabs/ui"]["Up"] = github_com_andlabs_ui.Up
Api["github.com/andlabs/ui"]["Down"] = github_com_andlabs_ui.Down
Api["github.com/andlabs/ui"]["Left"] = github_com_andlabs_ui.Left
Api["github.com/andlabs/ui"]["Right"] = github_com_andlabs_ui.Right
Api["github.com/andlabs/ui"]["F1"] = github_com_andlabs_ui.F1
Api["github.com/andlabs/ui"]["F2"] = github_com_andlabs_ui.F2
Api["github.com/andlabs/ui"]["F3"] = github_com_andlabs_ui.F3
Api["github.com/andlabs/ui"]["F4"] = github_com_andlabs_ui.F4
Api["github.com/andlabs/ui"]["F5"] = github_com_andlabs_ui.F5
Api["github.com/andlabs/ui"]["F6"] = github_com_andlabs_ui.F6
Api["github.com/andlabs/ui"]["F7"] = github_com_andlabs_ui.F7
Api["github.com/andlabs/ui"]["F8"] = github_com_andlabs_ui.F8
Api["github.com/andlabs/ui"]["F9"] = github_com_andlabs_ui.F9
Api["github.com/andlabs/ui"]["F10"] = github_com_andlabs_ui.F10
Api["github.com/andlabs/ui"]["F11"] = github_com_andlabs_ui.F11
Api["github.com/andlabs/ui"]["F12"] = github_com_andlabs_ui.F12
Api["github.com/andlabs/ui"]["N0"] = github_com_andlabs_ui.N0
Api["github.com/andlabs/ui"]["N1"] = github_com_andlabs_ui.N1
Api["github.com/andlabs/ui"]["N2"] = github_com_andlabs_ui.N2
Api["github.com/andlabs/ui"]["N3"] = github_com_andlabs_ui.N3
Api["github.com/andlabs/ui"]["N4"] = github_com_andlabs_ui.N4
Api["github.com/andlabs/ui"]["N5"] = github_com_andlabs_ui.N5
Api["github.com/andlabs/ui"]["N6"] = github_com_andlabs_ui.N6
Api["github.com/andlabs/ui"]["N7"] = github_com_andlabs_ui.N7
Api["github.com/andlabs/ui"]["N8"] = github_com_andlabs_ui.N8
Api["github.com/andlabs/ui"]["N9"] = github_com_andlabs_ui.N9
Api["github.com/andlabs/ui"]["NDot"] = github_com_andlabs_ui.NDot
Api["github.com/andlabs/ui"]["NEnter"] = github_com_andlabs_ui.NEnter
Api["github.com/andlabs/ui"]["NAdd"] = github_com_andlabs_ui.NAdd
Api["github.com/andlabs/ui"]["NSubtract"] = github_com_andlabs_ui.NSubtract
Api["github.com/andlabs/ui"]["NMultiply"] = github_com_andlabs_ui.NMultiply
Api["github.com/andlabs/ui"]["NDivide"] = github_com_andlabs_ui.NDivide
Api["github.com/andlabs/ui"]["FontButton"] = github_com_andlabs_ui.FontButton{}
Api["github.com/andlabs/ui/new"]["FontButton"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.FontButton)
}
Api["github.com/andlabs/ui"]["NewFontButton"] = github_com_andlabs_ui.NewFontButton
Api["github.com/andlabs/ui"]["FontDescriptor"] = github_com_andlabs_ui.FontDescriptor{}
Api["github.com/andlabs/ui/new"]["FontDescriptor"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.FontDescriptor)
}
Api["github.com/andlabs/ui"]["Form"] = github_com_andlabs_ui.Form{}
Api["github.com/andlabs/ui/new"]["Form"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Form)
}
Api["github.com/andlabs/ui"]["NewForm"] = github_com_andlabs_ui.NewForm
Api["github.com/andlabs/ui"]["Grid"] = github_com_andlabs_ui.Grid{}
Api["github.com/andlabs/ui/new"]["Grid"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Grid)
}
Api["github.com/andlabs/ui"]["NewGrid"] = github_com_andlabs_ui.NewGrid
Api["github.com/andlabs/ui"]["NewGroup"] = github_com_andlabs_ui.NewGroup
Api["github.com/andlabs/ui"]["Image"] = github_com_andlabs_ui.Image{}
Api["github.com/andlabs/ui/new"]["Image"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Image)
}
Api["github.com/andlabs/ui"]["NewImage"] = github_com_andlabs_ui.NewImage
Api["github.com/andlabs/ui"]["Label"] = github_com_andlabs_ui.Label{}
Api["github.com/andlabs/ui/new"]["Label"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Label)
}
Api["github.com/andlabs/ui"]["NewLabel"] = github_com_andlabs_ui.NewLabel
Api["github.com/andlabs/ui"]["Ctrl"] = github_com_andlabs_ui.Ctrl
Api["github.com/andlabs/ui"]["Alt"] = github_com_andlabs_ui.Alt
Api["github.com/andlabs/ui"]["Shift"] = github_com_andlabs_ui.Shift
Api["github.com/andlabs/ui"]["Super"] = github_com_andlabs_ui.Super
Api["github.com/andlabs/ui"]["MultilineEntry"] = github_com_andlabs_ui.MultilineEntry{}
Api["github.com/andlabs/ui/new"]["MultilineEntry"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.MultilineEntry)
}
Api["github.com/andlabs/ui"]["NewMultilineEntry"] = github_com_andlabs_ui.NewMultilineEntry
Api["github.com/andlabs/ui"]["NewNonWrappingMultilineEntry"] = github_com_andlabs_ui.NewNonWrappingMultilineEntry
Api["github.com/andlabs/ui"]["ToOpenTypeTag"] = github_com_andlabs_ui.ToOpenTypeTag
Api["github.com/andlabs/ui"]["ProgressBar"] = github_com_andlabs_ui.ProgressBar{}
Api["github.com/andlabs/ui/new"]["ProgressBar"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.ProgressBar)
}
Api["github.com/andlabs/ui"]["NewProgressBar"] = github_com_andlabs_ui.NewProgressBar
Api["github.com/andlabs/ui"]["RadioButtons"] = github_com_andlabs_ui.RadioButtons{}
Api["github.com/andlabs/ui/new"]["RadioButtons"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.RadioButtons)
}
Api["github.com/andlabs/ui"]["NewRadioButtons"] = github_com_andlabs_ui.NewRadioButtons
Api["github.com/andlabs/ui"]["Separator"] = github_com_andlabs_ui.Separator{}
Api["github.com/andlabs/ui/new"]["Separator"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Separator)
}
Api["github.com/andlabs/ui"]["NewHorizontalSeparator"] = github_com_andlabs_ui.NewHorizontalSeparator
Api["github.com/andlabs/ui"]["NewVerticalSeparator"] = github_com_andlabs_ui.NewVerticalSeparator
Api["github.com/andlabs/ui"]["Slider"] = github_com_andlabs_ui.Slider{}
Api["github.com/andlabs/ui/new"]["Slider"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Slider)
}
Api["github.com/andlabs/ui"]["NewSlider"] = github_com_andlabs_ui.NewSlider
Api["github.com/andlabs/ui"]["Spinbox"] = github_com_andlabs_ui.Spinbox{}
Api["github.com/andlabs/ui/new"]["Spinbox"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Spinbox)
}
Api["github.com/andlabs/ui"]["NewSpinbox"] = github_com_andlabs_ui.NewSpinbox
Api["github.com/andlabs/ui"]["Tab"] = github_com_andlabs_ui.Tab{}
Api["github.com/andlabs/ui/new"]["Tab"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Tab)
}
Api["github.com/andlabs/ui"]["NewTab"] = github_com_andlabs_ui.NewTab
Api["github.com/andlabs/ui"]["Table"] = github_com_andlabs_ui.Table{}
Api["github.com/andlabs/ui/new"]["Table"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Table)
}
Api["github.com/andlabs/ui"]["NewTable"] = github_com_andlabs_ui.NewTable
Api["github.com/andlabs/ui"]["TableColor"] = github_com_andlabs_ui.TableColor{}
Api["github.com/andlabs/ui/new"]["TableColor"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TableColor)
}
Api["github.com/andlabs/ui"]["TableImage"] = github_com_andlabs_ui.TableImage{}
Api["github.com/andlabs/ui/new"]["TableImage"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TableImage)
}
Api["github.com/andlabs/ui"]["TableFalse"] = github_com_andlabs_ui.TableFalse
Api["github.com/andlabs/ui"]["TableTrue"] = github_com_andlabs_ui.TableTrue
Api["github.com/andlabs/ui"]["TableModel"] = github_com_andlabs_ui.TableModel{}
Api["github.com/andlabs/ui/new"]["TableModel"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TableModel)
}
Api["github.com/andlabs/ui"]["NewTableModel"] = github_com_andlabs_ui.NewTableModel
Api["github.com/andlabs/ui"]["TableParams"] = github_com_andlabs_ui.TableParams{}
Api["github.com/andlabs/ui/new"]["TableParams"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TableParams)
}
Api["github.com/andlabs/ui"]["TableTextColumnOptionalParams"] = github_com_andlabs_ui.TableTextColumnOptionalParams{}
Api["github.com/andlabs/ui/new"]["TableTextColumnOptionalParams"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TableTextColumnOptionalParams)
}
Api["github.com/andlabs/ui"]["TextBackground"] = github_com_andlabs_ui.TextBackground{}
Api["github.com/andlabs/ui/new"]["TextBackground"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TextBackground)
}
Api["github.com/andlabs/ui"]["TextColor"] = github_com_andlabs_ui.TextColor{}
Api["github.com/andlabs/ui/new"]["TextColor"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.TextColor)
}
Api["github.com/andlabs/ui"]["TextItalicNormal"] = github_com_andlabs_ui.TextItalicNormal
Api["github.com/andlabs/ui"]["TextItalicOblique"] = github_com_andlabs_ui.TextItalicOblique
Api["github.com/andlabs/ui"]["TextItalicItalic"] = github_com_andlabs_ui.TextItalicItalic
Api["github.com/andlabs/ui"]["TextStretchUltraCondensed"] = github_com_andlabs_ui.TextStretchUltraCondensed
Api["github.com/andlabs/ui"]["TextStretchExtraCondensed"] = github_com_andlabs_ui.TextStretchExtraCondensed
Api["github.com/andlabs/ui"]["TextStretchCondensed"] = github_com_andlabs_ui.TextStretchCondensed
Api["github.com/andlabs/ui"]["TextStretchSemiCondensed"] = github_com_andlabs_ui.TextStretchSemiCondensed
Api["github.com/andlabs/ui"]["TextStretchNormal"] = github_com_andlabs_ui.TextStretchNormal
Api["github.com/andlabs/ui"]["TextStretchSemiExpanded"] = github_com_andlabs_ui.TextStretchSemiExpanded
Api["github.com/andlabs/ui"]["TextStretchExpanded"] = github_com_andlabs_ui.TextStretchExpanded
Api["github.com/andlabs/ui"]["TextStretchExtraExpanded"] = github_com_andlabs_ui.TextStretchExtraExpanded
Api["github.com/andlabs/ui"]["TextStretchUltraExpanded"] = github_com_andlabs_ui.TextStretchUltraExpanded
Api["github.com/andlabs/ui"]["TextWeightMinimum"] = github_com_andlabs_ui.TextWeightMinimum
Api["github.com/andlabs/ui"]["TextWeightThin"] = github_com_andlabs_ui.TextWeightThin
Api["github.com/andlabs/ui"]["TextWeightUltraLight"] = github_com_andlabs_ui.TextWeightUltraLight
Api["github.com/andlabs/ui"]["TextWeightLight"] = github_com_andlabs_ui.TextWeightLight
Api["github.com/andlabs/ui"]["TextWeightBook"] = github_com_andlabs_ui.TextWeightBook
Api["github.com/andlabs/ui"]["TextWeightNormal"] = github_com_andlabs_ui.TextWeightNormal
Api["github.com/andlabs/ui"]["TextWeightMedium"] = github_com_andlabs_ui.TextWeightMedium
Api["github.com/andlabs/ui"]["TextWeightSemiBold"] = github_com_andlabs_ui.TextWeightSemiBold
Api["github.com/andlabs/ui"]["TextWeightBold"] = github_com_andlabs_ui.TextWeightBold
Api["github.com/andlabs/ui"]["TextWeightUltraBold"] = github_com_andlabs_ui.TextWeightUltraBold
Api["github.com/andlabs/ui"]["TextWeightHeavy"] = github_com_andlabs_ui.TextWeightHeavy
Api["github.com/andlabs/ui"]["TextWeightUltraHeavy"] = github_com_andlabs_ui.TextWeightUltraHeavy
Api["github.com/andlabs/ui"]["TextWeightMaximum"] = github_com_andlabs_ui.TextWeightMaximum
Api["github.com/andlabs/ui"]["UnderlineNone"] = github_com_andlabs_ui.UnderlineNone
Api["github.com/andlabs/ui"]["UnderlineSingle"] = github_com_andlabs_ui.UnderlineSingle
Api["github.com/andlabs/ui"]["UnderlineDouble"] = github_com_andlabs_ui.UnderlineDouble
Api["github.com/andlabs/ui"]["UnderlineSuggestion"] = github_com_andlabs_ui.UnderlineSuggestion
Api["github.com/andlabs/ui"]["UnderlineColorSpelling"] = github_com_andlabs_ui.UnderlineColorSpelling
Api["github.com/andlabs/ui"]["UnderlineColorGrammar"] = github_com_andlabs_ui.UnderlineColorGrammar
Api["github.com/andlabs/ui"]["UnderlineColorAuxiliary"] = github_com_andlabs_ui.UnderlineColorAuxiliary
Api["github.com/andlabs/ui"]["UnderlineColorCustom"] = github_com_andlabs_ui.UnderlineColorCustom{}
Api["github.com/andlabs/ui/new"]["UnderlineColorCustom"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.UnderlineColorCustom)
}
Api["github.com/andlabs/ui"]["Window"] = github_com_andlabs_ui.Window{}
Api["github.com/andlabs/ui/new"]["Window"] = func(args ...interface{}) interface{} {
					return new(github_com_andlabs_ui.Window)
}
Api["github.com/andlabs/ui"]["NewWindow"] = github_com_andlabs_ui.NewWindow
}
