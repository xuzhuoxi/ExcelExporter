package setting

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	//RootPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res`
	RootPath = osxu.GetRunningDir()

	SystemPath  = filex.Combine(RootPath, "system.yaml")
	ProjectPath = filex.Combine(RootPath, "project.yaml")
	ExcelPath   = filex.Combine(RootPath, "excel.yaml")
	LangGoPath  = filex.Combine(RootPath, "./langs/go.yaml")
)

type Settings struct {
	System  *SystemSetting
	Project *ProjectSetting
	Excel   *ExcelSetting
}

func (s *Settings) Init() {
	s.loadSystemSetting()
	s.loadProjectSetting()
	s.loadExcelSetting()
}

func (s *Settings) InitLangSetting(lang string) error {
	ref, ok := s.System.FindProgramLanguage(lang)
	if !ok {
		return errors.New(fmt.Sprintf("Lang(%s) is not supported!", lang))
	}
	langSetting := &LangSetting{}
	UnmarshalData(ref.RefPath, langSetting)
	ref.Setting = langSetting
	return nil
}

func (s *Settings) loadSystemSetting() error {
	system := &SystemSetting{}
	err := UnmarshalData(SystemPath, system)
	if nil != err {
		return err
	}
	s.System = system
	return nil
}

func (s *Settings) loadProjectSetting() error {
	project := &ProjectSetting{}
	err := UnmarshalData(ProjectPath, project)
	if nil != err {
		return err
	}
	s.Project = project
	return nil
}

func (s *Settings) loadExcelSetting() error {
	excel := &ExcelSetting{}
	err := UnmarshalData(ExcelPath, excel)
	if nil != err {
		return err
	}
	s.Excel = excel
	return nil
}

//------------------------

func UnmarshalData(dataPath string, dataRef interface{}) error {
	bs, err := ioutil.ReadFile(dataPath)
	if nil != err {
		return err
	}
	err = yaml.Unmarshal(bs, dataRef)
	if nil != err {
		return err
	}
	return nil
}
