package excelize
import (
github_com_qax_os_excelize "github.com/qax-os/excelize"
)
var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}
func init() {
if _, ok := Api["github.com/qax-os/excelize"]; !ok {
   Api["github.com/qax-os/excelize"] = map[string]interface{}{}
}
if _, ok := Api["github.com/qax-os/excelize/new"]; !ok {
	Api["github.com/qax-os/excelize/new"] = map[string]interface{}{}
}
Api["github.com/qax-os/excelize"]["STCellFormulaTypeArray"] = github_com_qax_os_excelize.STCellFormulaTypeArray
Api["github.com/qax-os/excelize"]["STCellFormulaTypeDataTable"] = github_com_qax_os_excelize.STCellFormulaTypeDataTable
Api["github.com/qax-os/excelize"]["STCellFormulaTypeNormal"] = github_com_qax_os_excelize.STCellFormulaTypeNormal
Api["github.com/qax-os/excelize"]["STCellFormulaTypeShared"] = github_com_qax_os_excelize.STCellFormulaTypeShared
Api["github.com/qax-os/excelize"]["Bar"] = github_com_qax_os_excelize.Bar
Api["github.com/qax-os/excelize"]["BarStacked"] = github_com_qax_os_excelize.BarStacked
Api["github.com/qax-os/excelize"]["BarPercentStacked"] = github_com_qax_os_excelize.BarPercentStacked
Api["github.com/qax-os/excelize"]["Bar3DClustered"] = github_com_qax_os_excelize.Bar3DClustered
Api["github.com/qax-os/excelize"]["Bar3DStacked"] = github_com_qax_os_excelize.Bar3DStacked
Api["github.com/qax-os/excelize"]["Bar3DPercentStacked"] = github_com_qax_os_excelize.Bar3DPercentStacked
Api["github.com/qax-os/excelize"]["Col"] = github_com_qax_os_excelize.Col
Api["github.com/qax-os/excelize"]["ColStacked"] = github_com_qax_os_excelize.ColStacked
Api["github.com/qax-os/excelize"]["ColPercentStacked"] = github_com_qax_os_excelize.ColPercentStacked
Api["github.com/qax-os/excelize"]["Col3DClustered"] = github_com_qax_os_excelize.Col3DClustered
Api["github.com/qax-os/excelize"]["Col3D"] = github_com_qax_os_excelize.Col3D
Api["github.com/qax-os/excelize"]["Col3DStacked"] = github_com_qax_os_excelize.Col3DStacked
Api["github.com/qax-os/excelize"]["Col3DPercentStacked"] = github_com_qax_os_excelize.Col3DPercentStacked
Api["github.com/qax-os/excelize"]["Doughnut"] = github_com_qax_os_excelize.Doughnut
Api["github.com/qax-os/excelize"]["Line"] = github_com_qax_os_excelize.Line
Api["github.com/qax-os/excelize"]["Pie"] = github_com_qax_os_excelize.Pie
Api["github.com/qax-os/excelize"]["Pie3D"] = github_com_qax_os_excelize.Pie3D
Api["github.com/qax-os/excelize"]["Radar"] = github_com_qax_os_excelize.Radar
Api["github.com/qax-os/excelize"]["Scatter"] = github_com_qax_os_excelize.Scatter
Api["github.com/qax-os/excelize"]["DataValidationTypeCustom"] = github_com_qax_os_excelize.DataValidationTypeCustom
Api["github.com/qax-os/excelize"]["DataValidationTypeDate"] = github_com_qax_os_excelize.DataValidationTypeDate
Api["github.com/qax-os/excelize"]["DataValidationTypeDecimal"] = github_com_qax_os_excelize.DataValidationTypeDecimal
Api["github.com/qax-os/excelize"]["DataValidationTypeTextLeng"] = github_com_qax_os_excelize.DataValidationTypeTextLeng
Api["github.com/qax-os/excelize"]["DataValidationTypeTime"] = github_com_qax_os_excelize.DataValidationTypeTime
Api["github.com/qax-os/excelize"]["DataValidationTypeWhole"] = github_com_qax_os_excelize.DataValidationTypeWhole
Api["github.com/qax-os/excelize"]["DataValidationOperatorBetween"] = github_com_qax_os_excelize.DataValidationOperatorBetween
Api["github.com/qax-os/excelize"]["DataValidationOperatorEqual"] = github_com_qax_os_excelize.DataValidationOperatorEqual
Api["github.com/qax-os/excelize"]["DataValidationOperatorGreaterThan"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThan
Api["github.com/qax-os/excelize"]["DataValidationOperatorGreaterThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThanOrEqual
Api["github.com/qax-os/excelize"]["DataValidationOperatorLessThan"] = github_com_qax_os_excelize.DataValidationOperatorLessThan
Api["github.com/qax-os/excelize"]["DataValidationOperatorLessThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorLessThanOrEqual
Api["github.com/qax-os/excelize"]["DataValidationOperatorNotBetween"] = github_com_qax_os_excelize.DataValidationOperatorNotBetween
Api["github.com/qax-os/excelize"]["DataValidationOperatorNotEqual"] = github_com_qax_os_excelize.DataValidationOperatorNotEqual
Api["github.com/qax-os/excelize"]["SourceRelationship"] = github_com_qax_os_excelize.SourceRelationship
Api["github.com/qax-os/excelize"]["SourceRelationshipChart"] = github_com_qax_os_excelize.SourceRelationshipChart
Api["github.com/qax-os/excelize"]["SourceRelationshipComments"] = github_com_qax_os_excelize.SourceRelationshipComments
Api["github.com/qax-os/excelize"]["SourceRelationshipImage"] = github_com_qax_os_excelize.SourceRelationshipImage
Api["github.com/qax-os/excelize"]["SourceRelationshipTable"] = github_com_qax_os_excelize.SourceRelationshipTable
Api["github.com/qax-os/excelize"]["SourceRelationshipDrawingML"] = github_com_qax_os_excelize.SourceRelationshipDrawingML
Api["github.com/qax-os/excelize"]["SourceRelationshipDrawingVML"] = github_com_qax_os_excelize.SourceRelationshipDrawingVML
Api["github.com/qax-os/excelize"]["SourceRelationshipHyperLink"] = github_com_qax_os_excelize.SourceRelationshipHyperLink
Api["github.com/qax-os/excelize"]["SourceRelationshipWorkSheet"] = github_com_qax_os_excelize.SourceRelationshipWorkSheet
Api["github.com/qax-os/excelize"]["SourceRelationshipChart201506"] = github_com_qax_os_excelize.SourceRelationshipChart201506
Api["github.com/qax-os/excelize"]["SourceRelationshipChart20070802"] = github_com_qax_os_excelize.SourceRelationshipChart20070802
Api["github.com/qax-os/excelize"]["SourceRelationshipChart2014"] = github_com_qax_os_excelize.SourceRelationshipChart2014
Api["github.com/qax-os/excelize"]["SourceRelationshipCompatibility"] = github_com_qax_os_excelize.SourceRelationshipCompatibility
Api["github.com/qax-os/excelize"]["NameSpaceDrawingML"] = github_com_qax_os_excelize.NameSpaceDrawingML
Api["github.com/qax-os/excelize"]["NameSpaceDrawingMLChart"] = github_com_qax_os_excelize.NameSpaceDrawingMLChart
Api["github.com/qax-os/excelize"]["NameSpaceDrawingMLSpreadSheet"] = github_com_qax_os_excelize.NameSpaceDrawingMLSpreadSheet
Api["github.com/qax-os/excelize"]["NameSpaceSpreadSheet"] = github_com_qax_os_excelize.NameSpaceSpreadSheet
Api["github.com/qax-os/excelize"]["NameSpaceXML"] = github_com_qax_os_excelize.NameSpaceXML
Api["github.com/qax-os/excelize"]["EMU"] = github_com_qax_os_excelize.EMU
Api["github.com/qax-os/excelize"]["HSLModel"] = github_com_qax_os_excelize.HSLModel
Api["github.com/qax-os/excelize"]["HSLToRGB"] = github_com_qax_os_excelize.HSLToRGB
Api["github.com/qax-os/excelize"]["RGBToHSL"] = github_com_qax_os_excelize.RGBToHSL
Api["github.com/qax-os/excelize"]["ReadZipReader"] = github_com_qax_os_excelize.ReadZipReader
Api["github.com/qax-os/excelize"]["ThemeColor"] = github_com_qax_os_excelize.ThemeColor
Api["github.com/qax-os/excelize"]["TitleToNumber"] = github_com_qax_os_excelize.TitleToNumber
Api["github.com/qax-os/excelize"]["ToAlphaString"] = github_com_qax_os_excelize.ToAlphaString
Api["github.com/qax-os/excelize"]["Comment"] = github_com_qax_os_excelize.Comment{}
Api["github.com/qax-os/excelize/new"]["Comment"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.Comment)
}
Api["github.com/qax-os/excelize"]["DataValidation"] = github_com_qax_os_excelize.DataValidation{}
Api["github.com/qax-os/excelize/new"]["DataValidation"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.DataValidation)
}
Api["github.com/qax-os/excelize"]["NewDataValidation"] = github_com_qax_os_excelize.NewDataValidation
Api["github.com/qax-os/excelize"]["DataValidationErrorStyleStop"] = github_com_qax_os_excelize.DataValidationErrorStyleStop
Api["github.com/qax-os/excelize"]["DataValidationErrorStyleWarning"] = github_com_qax_os_excelize.DataValidationErrorStyleWarning
Api["github.com/qax-os/excelize"]["DataValidationErrorStyleInformation"] = github_com_qax_os_excelize.DataValidationErrorStyleInformation
Api["github.com/qax-os/excelize"]["ErrSheetNotExist"] = github_com_qax_os_excelize.ErrSheetNotExist{}
Api["github.com/qax-os/excelize/new"]["ErrSheetNotExist"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.ErrSheetNotExist)
}
Api["github.com/qax-os/excelize"]["File"] = github_com_qax_os_excelize.File{}
Api["github.com/qax-os/excelize/new"]["File"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.File)
}
Api["github.com/qax-os/excelize"]["NewFile"] = github_com_qax_os_excelize.NewFile
Api["github.com/qax-os/excelize"]["OpenReader"] = github_com_qax_os_excelize.OpenReader
Api["github.com/qax-os/excelize"]["HSL"] = github_com_qax_os_excelize.HSL{}
Api["github.com/qax-os/excelize/new"]["HSL"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.HSL)
}
Api["github.com/qax-os/excelize"]["Rows"] = github_com_qax_os_excelize.Rows{}
Api["github.com/qax-os/excelize/new"]["Rows"] = func(args ...interface{}) interface{} {
					return new(github_com_qax_os_excelize.Rows)
}
}
