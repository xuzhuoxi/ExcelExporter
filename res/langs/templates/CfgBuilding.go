package xu

/**
 * @author xuzhuoxi
 * Created on 2021-05-23.
 */
type ICfgBuilding interface {
	//以下为接口定义------------

	/**
	 * ID
	 *
	 * @remark 注释列
	 * @return {uint16}
	 */
	getBuilding_type() uint16
	/**
	 * 建筑
	 *
	 * @remark 建筑名称
	 * @return {string}
	 */
	getName() string
	/**
	 * 是否可以升级
	 *
	 * @remark 0：代表不能升级
1：代表可以升级
	 * @return {uint8}
	 */
	getPromotion() uint8
	/**
	 * 规格
	 *
	 * @remark 
	 * @return {uint16}
	 */
	getLayoutX() uint16
	/**
	 * 建筑key
	 *
	 * @remark 建筑key指的是对应类型的1级建筑的ID
	 * @return {uint32}
	 */
	getType_idx() uint32
	/**
	 * 门X坐标
	 *
	 * @remark 坐标从0开始
	 * @return {float32}
	 */
	getGateX() float32
	/**
	 * 门Y坐标
	 *
	 * @remark 
	 * @return {float32}
	 */
	getGateY() float32
	/**
	 * 地形
	 *
	 * @remark 4
	 * @return {int32}
	 */
	getTerrain_flags() int32
	/**
	 * 供需人口类型
	 *
	 * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

	 * @return {int16}
	 */
	getSupply_population_type() int16
	/**
	 * 有没有门
	 *
	 * @remark 5.0
	 * @return {int8}
	 */
	getIsDoor() int8
	/**
	 * 建筑描述
	 *
	 * @remark 建筑描述
string(20)
	 * @return {string}
	 */
	getDesc() string
}

/**
 * @author xuzhuoxi
 * Created on 2021-05-23.
 */
type CfgBuilding struct {
	//以下为字段定义------------

	//ID
	_building_type uint16 `json:"building_type,omitempty"`
	//建筑
	_name string `json:"name,omitempty"`
	//是否可以升级
	_promotion uint8 `json:"promotion,omitempty"`
	//规格
	_layoutX uint16 `json:"layoutX,omitempty"`
	//建筑key
	_type_idx uint32 `json:"type_idx,omitempty"`
	//门X坐标
	_gateX float32 `json:"gateX,omitempty"`
	//门Y坐标
	_gateY float32 `json:"gateY,omitempty"`
	//地形
	_terrain_flags int32 `json:"terrain_flags,omitempty"`
	//供需人口类型
	_supply_population_type int16 `json:"supply_population_type,omitempty"`
	//有没有门
	_isDoor int8 `json:"isDoor,omitempty"`
	//建筑描述
	_desc string `json:"desc,omitempty"`
}

//属性接口实现-----------------------

func (o CfgBuilding) getBuilding_type() uint16 {
	return o._building_type
}

func (o CfgBuilding) getName() string {
	return o._name
}

func (o CfgBuilding) getPromotion() uint8 {
	return o._promotion
}

func (o CfgBuilding) getLayoutX() uint16 {
	return o._layoutX
}

func (o CfgBuilding) getType_idx() uint32 {
	return o._type_idx
}

func (o CfgBuilding) getGateX() float32 {
	return o._gateX
}

func (o CfgBuilding) getGateY() float32 {
	return o._gateY
}

func (o CfgBuilding) getTerrain_flags() int32 {
	return o._terrain_flags
}

func (o CfgBuilding) getSupply_population_type() int16 {
	return o._supply_population_type
}

func (o CfgBuilding) getIsDoor() int8 {
	return o._isDoor
}

func (o CfgBuilding) getDesc() string {
	return o._desc
}

//解释接口实现-----------------------

func (o *CfgBuilding) parseJson(proxy IJsonReaderProxy) {
	//以下为从Json数据中解释出字段数据
	o._building_type = proxy.getUInt16("building_type")
	o._name = proxy.getString("name")
	o._promotion = proxy.getUInt8("promotion")
	o._layoutX = proxy.getUInt16("layoutX")
	o._type_idx = proxy.getUInt32("type_idx")
	o._gateX = proxy.getFloat32("gateX")
	o._gateY = proxy.getFloat32("gateY")
	o._terrain_flags = proxy.getInt32("terrain_flags")
	o._supply_population_type = proxy.getInt16("supply_population_type")
	o._isDoor = proxy.getInt8("isDoor")
	o._desc = proxy.getString("desc")

}

func (o *CfgBuilding) parseBinary(proxy IBinaryReaderProxy) {
	//以下为从二进制数据中解释出字段数据
	o._building_type = proxy.readUInt16()
	o._name = proxy.readString()
	o._promotion = proxy.readUInt8()
	o._layoutX = proxy.readUInt16()
	o._type_idx = proxy.readUInt32()
	o._gateX = proxy.readFloat32()
	o._gateY = proxy.readFloat32()
	o._terrain_flags = proxy.readInt32()
	o._supply_population_type = proxy.readInt16()
	o._isDoor = proxy.readInt8()
	o._desc = proxy.readString()

}
