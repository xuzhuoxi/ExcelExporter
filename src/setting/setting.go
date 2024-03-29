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

	EnvPath     = RootPath
	SystemPath  = filex.Combine(EnvPath, "system.yaml")
	ProjectPath = filex.Combine(EnvPath, "project.yaml")
	ExcelPath   = filex.Combine(EnvPath, "excel.yaml")
)

type Settings struct {
	System  *SystemSettings // 工具系统配置
	Project *ProjectSetting // 项目配置
	Excel   *ExcelSetting   // Excel表格相关配置
}

func (s *Settings) Init(envPath string) {
	s.initPath(envPath)
	s.loadSystemSetting()
	s.loadProjectSetting()
	s.loadExcelSetting()
}

func (s *Settings) initPath(envPath string) {
	if len(envPath) > 0 {
		EnvPath = envPath
		SystemPath = filex.Combine(EnvPath, "system.yaml")
		ProjectPath = filex.Combine(EnvPath, "project.yaml")
		ExcelPath = filex.Combine(EnvPath, "excel.yaml")
	}
}

func (s *Settings) InitLangSetting(lang string) error {
	ref, ok := s.System.FindProgramLanguage(lang)
	if !ok {
		return errors.New(fmt.Sprintf("Lang(%s) is not supported!", lang))
	}
	langSetting := &LangSetting{}
	UnmarshalData(ref.RefPath, langSetting)

	//// 以下打印用于测试语言配置的初始化信息
	//fmt.Println("InitLangSetting", lang, langSetting)

	ref.Setting = langSetting
	return nil
}

func (s *Settings) loadSystemSetting() error {
	system := &SystemSettings{}
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
