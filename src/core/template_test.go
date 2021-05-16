package core

import (
	"os"
	"testing"
)

type ab struct {
	A string
	B string
	C string
}

var (
	const_temp = `开始啦
{{$abArr := .}}
内容：{{range $index, $ab := $abArr}}
    {{$index}}={{$ab.A}},{{$ab.B}},{{$ab.C}} {{end}}
结束啦`

	datas = []ab{{A: "A1", B: "B1", C: "B1"}, {A: "A2", B: "B2"}, {A: "A3", B: "B3"}, {A: "A4", B: "B4"}}
)

func TestNewTemplate(t *testing.T) {
	temp, err := NewTemplate("Test", const_temp)
	if nil != err {
		t.Fatal(err)
	}
	temp.Execute(os.Stdout, datas)
}

//func TestLoadTemplate(t *testing.T) {
//	path := filex.Combine(osxu.GetRunningDir(), "test/const.tmp")
//	temp, err := LoadTemplate(path)
//	if nil != err {
//		t.Fatal(err)
//	}
//	temp.Execute(os.Stdout, datas)
//}
//
//func TestLoadTemplates(t *testing.T) {
//	baseDir := osxu.GetRunningDir()
//	path1 := filex.Combine(baseDir, "test/const.tmp")
//	path2 := filex.Combine(baseDir, "test/const2.tmp")
//	temp, _ := LoadTemplates(path1 + "," + path2)
//
//	fmt.Println(temp, temp.Template)
//	temp.ExecuteTemplate(os.Stdout, datas)
//}
