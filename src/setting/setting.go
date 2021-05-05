package setting

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	RootPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res`
	//RootPath    = osxu.GetRunningDir()

	SystemPath  = filex.Combine(RootPath, "system.yaml")
	ProjectPath = filex.Combine(RootPath, "prject.yaml")
	ExcelPath   = filex.Combine(RootPath, "excel.yaml")
	LangGoPath  = filex.Combine(RootPath, "./langs/go.yaml")
)

type Settings struct {
	System  *SystemSetting
	Project *ProjectSetting
	Excel   *ExcelSetting
	LangMap map[string]*LangSetting
}

func (s *Settings) Init() {
	s.loadSystemSetting()
	s.loadProjectSetting()
	s.loadExcelSetting()
}

func (s *Settings) loadSystemSetting() error {
	system := &SystemSetting{}
	err := s.initSetting(SystemPath, system)
	if nil != err {
		return err
	}
	s.System = system
	return nil
}

func (s *Settings) loadProjectSetting() error {
	project := &ProjectSetting{}
	err := s.initSetting(ProjectPath, project)
	if nil != err {
		return err
	}
	s.Project = project
	return nil
}

func (s *Settings) loadExcelSetting() error {
	excel := &ExcelSetting{}
	err := s.initSetting(ExcelPath, excel)
	if nil != err {
		return err
	}
	s.Excel = excel
	return nil
}

func (s *Settings) InitLangSetting(lang string) error {
	ok, ref := s.System.FindLangRef(lang)
	if !ok {
		return errors.New(fmt.Sprintf("Lang(%s) is not supported!", lang))
	}
	path := filex.Combine(RootPath, ref.Ref)
	langSetting := &LangSetting{}
	s.initSetting(path, langSetting)
	s.LangMap[lang] = langSetting
	return nil
}

func (s *Settings) initSetting(path string, settingRef interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if nil != err {
		return err
	}
	err = yaml.Unmarshal(bs, settingRef)
	if nil != err {
		return err
	}
	return nil
}
