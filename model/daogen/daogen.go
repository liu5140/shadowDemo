package daogen

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/utils"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}

var dao_template *template.Template

var dao_ext_template *template.Template

var model_template = `
package model

import (
	"shadowDemo/model/dao"
	"shadowDemo/zframework/datasource"
)

type Model struct {
	%v
}

var model *Model = nil

func ModelInit() {
	model = &Model{}

	db := datasource.DataSourceInstance().Master()

	%v
}


`
var model_en_template = `

package model

import "shadowDemo/model/do"

func GetModel() *Model {
	return model
}

var initialModels []interface{} = []interface{}{
	//new(do.Player),
	%v
}

func GetInitialModels() []interface{} {
	return initialModels
}

`

type ModelFill struct {
	StructLines string
	InitLines   string
	EnLines     string
}

func unescaped(x string) interface{} { return template.HTML(x) }

func DaoGenEntry() {
	input := ""
	output := ""
	templatePath := ""
	modelPath := ""
	enPath := ""
	outputService := ""

	flag.StringVar(&input, "i", "./model/do", "input files")
	flag.StringVar(&modelPath, "m", "./model/model.go", "model.go")
	flag.StringVar(&templatePath, "t", "./model/daogen", "template files")
	flag.StringVar(&output, "o", "./model/dao", "output directory")
	flag.StringVar(&enPath, "e", "./model/initial_model.go", "initial_model.go")
	flag.StringVar(&outputService, "s", "./service", "output directory")

	flag.Parse()

	var fpaths []string

	files, _ := ioutil.ReadDir(input)
	for _, fi := range files {
		if !fi.IsDir() {
			path := input + "/" + fi.Name()
			fpaths = append(fpaths, path)
		}
	}

	// Log.Debugf("fpaths = %v", fpaths)
	err := errors.New("")
	dao_template, err = template.New("dao").Funcs(template.FuncMap{
		"LowerCaseFirstLetter": utils.LowerCaseFirstLetter,
		"unescaped":            unescaped,
	}).ParseFiles(templatePath+"/dao.tmpl", templatePath+"/dao_ext.tmpl", templatePath+"/service.tmpl")
	if err != nil {
		Log.Error(err)
		return
	}

	// templateContent2, err := ioutil.ReadFile(templatePath + "/dao.tmpl")
	// if err != nil {
	// 	Log.Error(err)
	// 	return
	// }

	modelLine1 := ""
	modelLine2 := ""
	modelEn := ""

	for _, fpath := range fpaths {
		base := filepath.Base(fpath)
		if base != "do.go" &&
			base != "session_do.go" &&
			base != "auth_key_do.go" &&
			base != "init.go" {
			Log.Infoln("======fpath==", fpath)
			sti := do2dao(fpath, output, outputService)
			if len(sti.StructName) == 0 {
				continue
			}

			daoName := fmt.Sprintf("%vDao *dao.%vDao\n", sti.StructName, sti.StructName)
			daoNew := fmt.Sprintf("model.%vDao = dao.New%vDao(db)\n", sti.StructName, sti.StructName)
			daoEn := fmt.Sprintf("new(do.%v),\n", sti.StructName)

			modelLine1 = modelLine1 + daoName
			modelLine2 = modelLine2 + daoNew
			modelEn = modelEn + daoEn
		}
	}

	mf := &ModelFill{
		StructLines: modelLine1,
		InitLines:   modelLine2,
		EnLines:     modelEn,
	}
	writeModel(modelPath, mf)

	writeModelEn(enPath, mf)
}
func writeModelEn(enPath string, mf *ModelFill) {
	modelgo := fmt.Sprintf(model_en_template, mf.EnLines)
	ioutil.WriteFile(enPath, []byte(modelgo), 0666)
}

func writeModel(modelPath string, mf *ModelFill) {
	modelgo := fmt.Sprintf(model_template, mf.StructLines, mf.InitLines)
	ioutil.WriteFile(modelPath, []byte(modelgo), 0666)
}

type StructInfo struct {
	StructName string
	FieldNames []string
}

func do2dao(dofile string, outpath string, outputService string) *StructInfo {
	fileContent, err := ioutil.ReadFile(dofile)
	if err != nil {
		Log.Error(err)
		return nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", string(fileContent), parser.AllErrors)

	// spew.Dump(f)

	sti := &StructInfo{}

	ast.Inspect(f, func(n ast.Node) bool {
		//Log.Debugf("n type = %v", reflect.TypeOf(n))
		switch n.(type) {

		case *ast.TypeSpec:
			r := n.(*ast.TypeSpec)
			// Log.Debugf("r spec = %v", r.Name)
			sti.StructName = r.Name.Name
		case *ast.StructType:
			r := n.(*ast.StructType)
			// Log.Debugf("r = %#v, pos = %v, end = %v", r, r.Pos(), r.End())
			// Log.Debugf("r.Fields = %#v", r.Fields)
			for _, v := range r.Fields.List {
				nlen := len(v.Names)
				if nlen > 0 {
					// Log.Debugf("v = %#v", v.Names[0].Name)
					sti.FieldNames = append(sti.FieldNames, v.Names[0].Name)
				}
			}
		}

		return true
	})

	Log.Debugf("====StructName==", sti.StructName)
	// 生成文件
	outFile := outpath + "/" + utils.ToSnakeCase(sti.StructName) + "_dao.go"
	//增加一个判断，如果文件已经存在这不进行生成
	executeTempl(sti, outFile, "dao.tmpl")

	outFileext := outpath + "/" + utils.ToSnakeCase(sti.StructName) + "_dao_ext.go"

	executeTempl(sti, outFileext, "dao_ext.tmpl")

	outFileService := outputService + "/" + utils.ToSnakeCase(sti.StructName) + "_service.go"

	executeTempl(sti, outFileService, "service.tmpl")

	return sti
}

func executeTempl(sti *StructInfo, outFile string, temp string) error {
	Log.Debugf("outFile = %v", outFile)
	if !checkFileIsExist(outFile) {
		ofile, err := os.OpenFile(outFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			Log.Error(err)
			return nil
		}

		err = dao_template.ExecuteTemplate(ofile, temp, sti)
		if err != nil {
			Log.Error(err)
			return nil
		}
		ofile.Close()
	}
	return nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Log.Debug(filename + " not exist")
		exist = false
	}
	return exist
}
