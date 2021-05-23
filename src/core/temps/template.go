package temps

import (
	"io"
	"text/template"
)

type TemplateProxy struct {
	Name     string
	Template *template.Template
}

func (p *TemplateProxy) CloneTemplate() *template.Template {
	if nil != p.Template {
		return nil
	}
	clone, _ := p.Template.Clone()
	return clone
}

func (p *TemplateProxy) Execute(wr io.Writer, data interface{}, clone bool) error {
	if clone {
		return p.CloneTemplate().Execute(wr, data)
	}
	return p.Template.Execute(wr, data)
}

func (p *TemplateProxy) ExecuteTemplate(wr io.Writer, data interface{}, clone bool) error {
	if clone {
		return p.CloneTemplate().ExecuteTemplate(wr, p.Name, data)
	}
	return p.Template.ExecuteTemplate(wr, p.Name, data)
}

func (p *TemplateProxy) ExecuteTemplateByName(wr io.Writer, name string, data interface{}, clone bool) error {
	if clone {
		return p.CloneTemplate().ExecuteTemplate(wr, name, data)
	}
	return p.Template.ExecuteTemplate(wr, name, data)
}
