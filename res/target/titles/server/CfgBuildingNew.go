package client

/**
 * @author ${AUTHOR}
 * Created on ${DATE}.
 */
type ICfgBuildingNew interface {

    /**
     * ID
     *
     * @remark 注释列
     * @return uint16
     */
    GetBuildingType() uint16

    /**
     * 建筑
     *
     * @remark 建筑名称
     * @return string
     */
    GetName() string

    /**
     * 规格
     *
     * @remark 
     * @return uint16
     */
    GetLayoutX() uint16

    /**
     * 门X坐标
     *
     * @remark 坐标从0开始
     * @return float32
     */
    GetGateX() float32

    /**
     * 门Y坐标
     *
     * @remark 
     * @return float32
     */
    GetGateY() float32

    /**
     * 地形
     *
     * @remark 4
     * @return int32
     */
    GetTerrainFlags() int32

    /**
     * 供需人口类型
     *
     * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

     * @return int16
     */
    GetSupplyPopulationType() int16

    /**
     * 有没有门
     *
     * @remark 5
     * @return int8
     */
    GetIsDoor() int8

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string(20)
     * @return string
     */
    GetDesc1() string

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string
     * @return string
     */
    GetDesc2() string

    /**
     * 建筑描述
     *
     * @remark 建筑描述
json
     * @return string
     */
    GetDesc3() string

    /**
     * 测试1
     *
     * @remark 测试1
     * @return []bool
     */
    GetF1() []bool

    /**
     * 测试2
     *
     * @remark 测试2
     * @return []uint8
     */
    GetF2() []uint8

    /**
     * 测试3
     *
     * @remark 测试3
     * @return []uint16
     */
    GetF3() []uint16

    /**
     * 测试4
     *
     * @remark 测试4
     * @return []uint32
     */
    GetF4() []uint32

    /**
     * 测试5
     *
     * @remark 测试5
     * @return []int8
     */
    GetF5() []int8

    /**
     * 测试6
     *
     * @remark 测试6
     * @return []int16
     */
    GetF6() []int16

    /**
     * 测试7
     *
     * @remark 测试7
     * @return []int32
     */
    GetF7() []int32

    /**
     * 测试8
     *
     * @remark 测试8
     * @return []float32
     */
    GetF8() []float32

    /**
     * 测试9
     *
     * @remark 测试9
     * @return []string
     */
    GetF9() []string

    /**
     * 测试10
     *
     * @remark 测试10
     * @return []string
     */
    GetF10() []string

}

/**
 * @author ${AUTHOR}
 * Created on ${DATE}.
 */
type CfgBuildingNew struct {

    /**
     * ID
     */
    buildingType uint16

    /**
     * 建筑
     */
    name string

    /**
     * 规格
     */
    layoutX uint16

    /**
     * 门X坐标
     */
    gateX float32

    /**
     * 门Y坐标
     */
    gateY float32

    /**
     * 地形
     */
    terrainFlags int32

    /**
     * 供需人口类型
     */
    supplyPopulationType int16

    /**
     * 有没有门
     */
    isDoor int8

    /**
     * 建筑描述
     */
    desc1 string

    /**
     * 建筑描述
     */
    desc2 string

    /**
     * 建筑描述
     */
    desc3 string

    /**
     * 测试1
     */
    f1 []bool

    /**
     * 测试2
     */
    f2 []uint8

    /**
     * 测试3
     */
    f3 []uint16

    /**
     * 测试4
     */
    f4 []uint32

    /**
     * 测试5
     */
    f5 []int8

    /**
     * 测试6
     */
    f6 []int16

    /**
     * 测试7
     */
    f7 []int32

    /**
     * 测试8
     */
    f8 []float32

    /**
     * 测试9
     */
    f9 []string

    /**
     * 测试10
     */
    f10 []string

}

//属性接口实现-----------------------

/**
 * ID
*/
func (o *CfgBuildingNew) GetBuildingType() uint16 {
	return o.buildingType
}

/**
 * 建筑
*/
func (o *CfgBuildingNew) GetName() string {
	return o.name
}

/**
 * 规格
*/
func (o *CfgBuildingNew) GetLayoutX() uint16 {
	return o.layoutX
}

/**
 * 门X坐标
*/
func (o *CfgBuildingNew) GetGateX() float32 {
	return o.gateX
}

/**
 * 门Y坐标
*/
func (o *CfgBuildingNew) GetGateY() float32 {
	return o.gateY
}

/**
 * 地形
*/
func (o *CfgBuildingNew) GetTerrainFlags() int32 {
	return o.terrainFlags
}

/**
 * 供需人口类型
*/
func (o *CfgBuildingNew) GetSupplyPopulationType() int16 {
	return o.supplyPopulationType
}

/**
 * 有没有门
*/
func (o *CfgBuildingNew) GetIsDoor() int8 {
	return o.isDoor
}

/**
 * 建筑描述
*/
func (o *CfgBuildingNew) GetDesc1() string {
	return o.desc1
}

/**
 * 建筑描述
*/
func (o *CfgBuildingNew) GetDesc2() string {
	return o.desc2
}

/**
 * 建筑描述
*/
func (o *CfgBuildingNew) GetDesc3() string {
	return o.desc3
}

/**
 * 测试1
*/
func (o *CfgBuildingNew) GetF1() []bool {
	return o.f1
}

/**
 * 测试2
*/
func (o *CfgBuildingNew) GetF2() []uint8 {
	return o.f2
}

/**
 * 测试3
*/
func (o *CfgBuildingNew) GetF3() []uint16 {
	return o.f3
}

/**
 * 测试4
*/
func (o *CfgBuildingNew) GetF4() []uint32 {
	return o.f4
}

/**
 * 测试5
*/
func (o *CfgBuildingNew) GetF5() []int8 {
	return o.f5
}

/**
 * 测试6
*/
func (o *CfgBuildingNew) GetF6() []int16 {
	return o.f6
}

/**
 * 测试7
*/
func (o *CfgBuildingNew) GetF7() []int32 {
	return o.f7
}

/**
 * 测试8
*/
func (o *CfgBuildingNew) GetF8() []float32 {
	return o.f8
}

/**
 * 测试9
*/
func (o *CfgBuildingNew) GetF9() []string {
	return o.f9
}

/**
 * 测试10
*/
func (o *CfgBuildingNew) GetF10() []string {
	return o.f10
}


//解释接口实现-----------------------

func (o *CfgBuildingNew) parseJson(proxy IJsonReaderProxy) {
	//以下为从Json数据中解释出字段数据
    o.buildingType = proxy.GetUInt16("building_type")
    o.name = proxy.GetString("name")
    o.layoutX = proxy.GetUInt16("layoutX")
    o.gateX = proxy.GetFloat32("gateX")
    o.gateY = proxy.GetFloat32("gateY")
    o.terrainFlags = proxy.GetInt32("terrain_flags")
    o.supplyPopulationType = proxy.GetInt16("supply_population_type")
    o.isDoor = proxy.GetInt8("isDoor")
    o.desc1 = proxy.GetString("desc1")
    o.desc2 = proxy.GetString("desc2")
    o.desc3 = proxy.GetString("desc3")
    o.f1 = proxy.GetBooleanArray("f1")
    o.f2 = proxy.GetUInt8Array("f2")
    o.f3 = proxy.GetUInt16Array("f3")
    o.f4 = proxy.GetUInt32Array("f4")
    o.f5 = proxy.GetInt8Array("f5")
    o.f6 = proxy.GetInt16Array("f6")
    o.f7 = proxy.GetInt32Array("f7")
    o.f8 = proxy.GetFloat32Array("f8")
    o.f9 = proxy.GetStringArray("f9")
    o.f10 = proxy.GetStringArray("f10")
}

func (o *CfgBuildingNew) parseYaml(proxy IYamlReaderProxy) {
	//以下为从Yaml数据中解释出字段数据
    o.buildingType = proxy.Get("building_type")
    o.name = proxy.Get("name")
    o.layoutX = proxy.Get("layoutX")
    o.gateX = proxy.Get("gateX")
    o.gateY = proxy.Get("gateY")
    o.terrainFlags = proxy.Get("terrain_flags")
    o.supplyPopulationType = proxy.Get("supply_population_type")
    o.isDoor = proxy.Get("isDoor")
    o.desc1 = proxy.Get("desc1")
    o.desc2 = proxy.Get("desc2")
    o.desc3 = proxy.Get("desc3")
    o.f1 = proxy.Get("f1")
    o.f2 = proxy.Get("f2")
    o.f3 = proxy.Get("f3")
    o.f4 = proxy.Get("f4")
    o.f5 = proxy.Get("f5")
    o.f6 = proxy.Get("f6")
    o.f7 = proxy.Get("f7")
    o.f8 = proxy.Get("f8")
    o.f9 = proxy.Get("f9")
    o.f10 = proxy.Get("f10")
}

func (o *CfgBuildingNew) parseBinary(proxy IBinaryReaderProxy) {
	//以下为从二进制数据中解释出字段数据
    o.buildingType = proxy.ReadUInt16()
    o.name = proxy.ReadString()
    o.layoutX = proxy.ReadUInt16()
    o.gateX = proxy.ReadFloat32()
    o.gateY = proxy.ReadFloat32()
    o.terrainFlags = proxy.ReadInt32()
    o.supplyPopulationType = proxy.ReadInt16()
    o.isDoor = proxy.ReadInt8()
    o.desc1 = proxy.ReadString()
    o.desc2 = proxy.ReadString()
    o.desc3 = proxy.ReadString()
    o.f1 = proxy.ReadBooleanArray()
    o.f2 = proxy.ReadUInt8Array()
    o.f3 = proxy.ReadUInt16Array()
    o.f4 = proxy.ReadUInt32Array()
    o.f5 = proxy.ReadInt8Array()
    o.f6 = proxy.ReadInt16Array()
    o.f7 = proxy.ReadInt32Array()
    o.f8 = proxy.ReadFloat32Array()
    o.f9 = proxy.ReadStringArray()
    o.f10 = proxy.ReadStringArray()
}