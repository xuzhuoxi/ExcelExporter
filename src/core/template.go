package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"io"
	"io/ioutil"
	"strings"
	"text/template"
)

type TemplateProxy struct {
	Name     string
	Template *template.Template
}

func (temp *TemplateProxy) CloneTemplate() *template.Template {
	if nil != temp.Template {
		return nil
	}
	clone, _ := temp.Template.Clone()
	return clone
}

func (temp *TemplateProxy) Execute(wr io.Writer, data interface{}) error {
	return temp.Template.Execute(wr, data)
}

func (temp *TemplateProxy) ExecuteTemplate(wr io.Writer, data interface{}) error {
	return temp.Template.ExecuteTemplate(wr, temp.Name, data)
}

func (temp *TemplateProxy) ExecuteTemplateByName(wr io.Writer, name string, data interface{}) error {
	return temp.Template.ExecuteTemplate(wr, name, data)
}

func NewTemplate(name string, text string) (*TemplateProxy, error) {
	temp, err := template.New(name).Parse(text)
	if nil != err {
		return nil, err
	}
	rs := &TemplateProxy{Name: name, Template: temp}
	return rs, nil
}

func LoadTemplate(tempFile string) (*TemplateProxy, error) {
	//if temp, ok := templateMap[tempFile]; ok {
	//	return temp, nil
	//}
	if !filex.IsExist(tempFile) {
		return nil, errors.New(fmt.Sprintf("Templete File Not Found: \"%s\"", tempFile))
	}
	body, err := ioutil.ReadFile(tempFile)
	if nil != err {
		return nil, err
	}
	text := string(body)
	temp, err := template.New(tempFile).Parse(text)
	if nil != err {
		return nil, err
	}
	_, name := filex.Split(tempFile)
	rs := &TemplateProxy{Name: name, Template: temp}
	//templateMap[tempFile] = rs
	return rs, nil
}

func LoadTemplates(tempFiles string) (*TemplateProxy, error) {
	files := strings.Split(tempFiles, ",")
	temp, err := template.ParseFiles(files...)
	if nil != err {
		return nil, err
	}
	_, name := filex.Split(files[0])
	rs := &TemplateProxy{Name: name, Template: temp}
	return rs, nil
}
