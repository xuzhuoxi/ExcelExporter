package core

//type ab struct {
//	A string
//	B string
//	C string
//}
//
//var const_temp = `开始啦
//{{$proxy := .}}
//{{$rows := $proxy.DataRows}}
//内容：
//{{range $k, $row := $rows}}
//    {{$k}} = {{$row.ValueAtIndex 0}}, {{$row.ValueAtIndex 1}}{{end}}
//结束啦`
//
//func TestLoadTemplate(t *testing.T) {
//	path := osxu.RunningBaseDir() + "Temp/const.tmp"
//	temp, _ := LoadTemplate(path)
//
//	datas := []ab{{A: "A1", B: "B1", C: "B1"}, {A: "A2", B: "B2"}, {A: "A3", B: "B3"}, {A: "A4", B: "B4"}}
//	//a:={B:"B1"}
//	//a["A"] = "A1"
//	//datas2 := make(map[string]string)
//	//datas2["A1"] = "B1"
//	//datas2["A2"] = "B2"
//	//datas2["A3"] = "B3"
//	//datas2["A4"] = "B4"
//
//	temp.Execute(os.Stdout, datas)
//}
//
//func TestLoadTemplates(t *testing.T) {
//	baseDir := osxu.RunningBaseDir()
//	path1 := baseDir + "Temp/const.tmp"
//	path2 := baseDir + "Temp/const2.tmp"
//	temp, _ := LoadTemplates(path1 + "," + path2)
//
//	datas := []ab{{A: "A1", B: "B1", C: "B1"}, {A: "A2", B: "B2"}, {A: "A3", B: "B3"}, {A: "A4", B: "B4"}}
//	//a:={B:"B1"}
//	//a["A"] = "A1"
//	//datas2 := make(map[string]string)
//	//datas2["A1"] = "B1"
//	//datas2["A2"] = "B2"
//	//datas2["A3"] = "B3"
//	//datas2["A4"] = "B4"
//	fmt.Println(temp, temp.Template)
//	temp.ExecuteTemplate(os.Stdout, datas)
//}
