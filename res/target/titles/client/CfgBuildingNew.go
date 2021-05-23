package ${MODULE}

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
    Getbuilding_type() uint16

    /**
     * 建筑
     *
     * @remark 建筑名称
     * @return string
     */
    Getname() string

    /**
     * 是否可以升级
     *
     * @remark 0：代表不能升级
1：代表可以升级
     * @return uint8
     */
    Getpromotion() uint8

    /**
     * 门X坐标
     *
     * @remark 坐标从0开始
     * @return float32
     */
    GetgateX() float32

    /**
     * 门Y坐标
     *
     * @remark 
     * @return float32
     */
    GetgateY() float32

    /**
     * 地形
     *
     * @remark 4
     * @return int32
     */
    Getterrain_flags() int32

    /**
     * 供需人口类型
     *
     * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

     * @return int16
     */
    Getsupply_population_type() int16

    /**
     * 有没有门
     *
     * @remark 5
     * @return int8
     */
    GetisDoor() int8

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string(20)
     * @return 
     */
    Getdesc1() 

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string
     * @return string
     */
    Getdesc2() string

    /**
     * 建筑描述
     *
     * @remark 建筑描述
json
     * @return string
     */
    Getdesc3() string

    /**
     * 测试1
     *
     * @remark 测试1
     * @return []bool
     */
    Getf1() []bool

    /**
     * 测试2
     *
     * @remark 测试2
     * @return []uint8
     */
    Getf2() []uint8

    /**
     * 测试3
     *
     * @remark 测试3
     * @return []uint16
     */
    Getf3() []uint16

    /**
     * 测试4
     *
     * @remark 测试4
     * @return []uint32
     */
    Getf4() []uint32

    /**
     * 测试5
     *
     * @remark 测试5
     * @return []int8
     */
    Getf5() []int8

    /**
     * 测试6
     *
     * @remark 测试6
     * @return []int16
     */
    Getf6() []int16

    /**
     * 测试7
     *
     * @remark 测试7
     * @return []int32
     */
    Getf7() []int32

    /**
     * 测试8
     *
     * @remark 测试8
     * @return []float32
     */
    Getf8() []float32

    /**
     * 测试9
     *
     * @remark 测试9
     * @return []string
     */
    Getf9() []string

    /**
     * 测试10
     *
     * @remark 测试10
     * @return 
     */
    Getf10() 

}

/**
 * @author ${AUTHOR}
 * Created on ${DATE}.
 */
type CfgBuildingNew struct {
	//以下为字段定义------------

    /**
     * ID
     *
     * @remark 注释列
     * @return uint16
     */
    Getbuilding_type() uint16

    /**
     * 建筑
     *
     * @remark 建筑名称
     * @return string
     */
    Getname() string

    /**
     * 是否可以升级
     *
     * @remark 0：代表不能升级
1：代表可以升级
     * @return uint8
     */
    Getpromotion() uint8

    /**
     * 门X坐标
     *
     * @remark 坐标从0开始
     * @return float32
     */
    GetgateX() float32

    /**
     * 门Y坐标
     *
     * @remark 
     * @return float32
     */
    GetgateY() float32

    /**
     * 地形
     *
     * @remark 4
     * @return int32
     */
    Getterrain_flags() int32

    /**
     * 供需人口类型
     *
     * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

     * @return int16
     */
    Getsupply_population_type() int16

    /**
     * 有没有门
     *
     * @remark 5
     * @return int8
     */
    GetisDoor() int8

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string(20)
     * @return 
     */
    Getdesc1() 

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string
     * @return string
     */
    Getdesc2() string

    /**
     * 建筑描述
     *
     * @remark 建筑描述
json
     * @return string
     */
    Getdesc3() string

    /**
     * 测试1
     *
     * @remark 测试1
     * @return []bool
     */
    Getf1() []bool

    /**
     * 测试2
     *
     * @remark 测试2
     * @return []uint8
     */
    Getf2() []uint8

    /**
     * 测试3
     *
     * @remark 测试3
     * @return []uint16
     */
    Getf3() []uint16

    /**
     * 测试4
     *
     * @remark 测试4
     * @return []uint32
     */
    Getf4() []uint32

    /**
     * 测试5
     *
     * @remark 测试5
     * @return []int8
     */
    Getf5() []int8

    /**
     * 测试6
     *
     * @remark 测试6
     * @return []int16
     */
    Getf6() []int16

    /**
     * 测试7
     *
     * @remark 测试7
     * @return []int32
     */
    Getf7() []int32

    /**
     * 测试8
     *
     * @remark 测试8
     * @return []float32
     */
    Getf8() []float32

    /**
     * 测试9
     *
     * @remark 测试9
     * @return []string
     */
    Getf9() []string

    /**
     * 测试10
     *
     * @remark 测试10
     * @return 
     */
    Getf10() 

}
//属性接口实现-----------------------

${PROPERTY_FIELD_GET}

//解释接口实现-----------------------

func (o *${CLASS}) parseJson(proxy IJsonReaderProxy) {
	//以下为从Json数据中解释出字段数据
${PARSE_JSON_CONTENT}
}

func (o *${CLASS}) parseBinary(proxy IBinaryReaderProxy) {
	//以下为从二进制数据中解释出字段数据
${PARSE_BINARY_CONTENT}
}