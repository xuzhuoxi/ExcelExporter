package {module};

import org.json.JSONObject;
import xu.BinaryReaderProxy;
import xu.JsonReaderProxy;

/**
 * @author {AUTHOR}
 * Created on {DATE}.
 */
 public class CfgBuildingNew {

    // ID
     private int buildingType;

    // 建筑
     private String name;

    // 建筑key
     private long typeIdx;

    // 门X坐标
     private float gateX;

    // 门Y坐标
     private float gateY;

    // 地形
     private int terrainFlags;

    // 供需人口类型
     private int supplyPopulationType;

    // 有没有门
     private int isDoor;

    // 建筑描述
     private String desc1;

    // 建筑描述
     private String desc2;

    // 建筑描述
     private JSONObject desc3;

    // 测试1
     private boolean[] f1;

    // 测试2
     private short[] f2;

    // 测试3
     private int[] f3;

    // 测试4
     private long[] f4;

    // 测试5
     private byte[] f5;

    // 测试6
     private short[] f6;

    // 测试7
     private int[] f7;

    // 测试8
     private float[] f8;

    // 测试9
     private String[] f9;

    // 测试10
     private String[] f10;


    /**
     * ID
     *
     * @remark 注释列
     * @return int
     */
    publicint getBuildingType() {
        return this.buildingType;
    }

    /**
     * 建筑
     *
     * @remark 建筑名称
     * @return String
     */
    publicString getName() {
        return this.name;
    }

    /**
     * 建筑key
     *
     * @remark 建筑key指的是对应类型的1级建筑的ID
     * @return long
     */
    publiclong getTypeIdx() {
        return this.typeIdx;
    }

    /**
     * 门X坐标
     *
     * @remark 坐标从0开始
     * @return float
     */
    publicfloat getGateX() {
        return this.gateX;
    }

    /**
     * 门Y坐标
     *
     * @remark 
     * @return float
     */
    publicfloat getGateY() {
        return this.gateY;
    }

    /**
     * 地形
     *
     * @remark 4
     * @return int
     */
    publicint getTerrainFlags() {
        return this.terrainFlags;
    }

    /**
     * 供需人口类型
     *
     * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

     * @return int
     */
    publicint getSupplyPopulationType() {
        return this.supplyPopulationType;
    }

    /**
     * 有没有门
     *
     * @remark 5
     * @return int
     */
    publicint getIsDoor() {
        return this.isDoor;
    }

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string(20)
     * @return String
     */
    publicString getDesc1() {
        return this.desc1;
    }

    /**
     * 建筑描述
     *
     * @remark 建筑描述
string
     * @return String
     */
    publicString getDesc2() {
        return this.desc2;
    }

    /**
     * 建筑描述
     *
     * @remark 建筑描述
json
     * @return JSONObject
     */
    publicJSONObject getDesc3() {
        return this.desc3;
    }

    /**
     * 测试1
     *
     * @remark 测试1
     * @return boolean[]
     */
    publicboolean[] getF1() {
        return this.f1;
    }

    /**
     * 测试2
     *
     * @remark 测试2
     * @return short[]
     */
    publicshort[] getF2() {
        return this.f2;
    }

    /**
     * 测试3
     *
     * @remark 测试3
     * @return int[]
     */
    publicint[] getF3() {
        return this.f3;
    }

    /**
     * 测试4
     *
     * @remark 测试4
     * @return long[]
     */
    publiclong[] getF4() {
        return this.f4;
    }

    /**
     * 测试5
     *
     * @remark 测试5
     * @return byte[]
     */
    publicbyte[] getF5() {
        return this.f5;
    }

    /**
     * 测试6
     *
     * @remark 测试6
     * @return short[]
     */
    publicshort[] getF6() {
        return this.f6;
    }

    /**
     * 测试7
     *
     * @remark 测试7
     * @return int[]
     */
    publicint[] getF7() {
        return this.f7;
    }

    /**
     * 测试8
     *
     * @remark 测试8
     * @return float[]
     */
    publicfloat[] getF8() {
        return this.f8;
    }

    /**
     * 测试9
     *
     * @remark 测试9
     * @return String[]
     */
    publicString[] getF9() {
        return this.f9;
    }

    /**
     * 测试10
     *
     * @remark 测试10
     * @return String[]
     */
    publicString[] getF10() {
        return this.f10;
    }


    // Json数据解释
    public void parseJson(JsonReaderProxy proxy) {
        // 以下为从 Json数据代理 中解释出字段数据
        o.buildingType = proxy.getUInt16("building_type_j")
        o.name = proxy.getString("name_j")
        o.typeIdx = proxy.getUInt32("type_idx_j")
        o.gateX = proxy.getDouble("gateX_j")
        o.gateY = proxy.getDouble("gateY_j")
        o.terrainFlags = proxy.getInt32("terrain_flags_j")
        o.supplyPopulationType = proxy.getInt16("supply_population_type_j")
        o.isDoor = proxy.getInt8("isDoor_j")
        o.desc1 = proxy.getString("desc1_j")
        o.desc2 = proxy.getString("desc2_j")
        o.desc3 = proxy.getJsonObject("desc3_j")
        o.f1 = proxy.getBooleanArray("f1_j")
        o.f2 = proxy.getUInt8Array("f2_j")
        o.f3 = proxy.getUInt16Array("f3_j")
        o.f4 = proxy.getUInt32Array("f4_j")
        o.f5 = proxy.getInt8Array("f5_j")
        o.f6 = proxy.getInt16Array("f6_j")
        o.f7 = proxy.getInt32Array("f7_j")
        o.f8 = proxy.getFloat32Array("f8_j")
        o.f9 = proxy.getStringArray("f9_j")
        o.f10 = proxy.getStringArray("f10_j")
    }

    // 二进制数据解释
    public void parseBinary(BinaryReaderProxy proxy) {
        // 以下为从 二进制数据代理 中解释出字段数据
        o.buildingType = proxy.readUInt16()
        o.name = proxy.readString()
        o.typeIdx = proxy.readUInt32()
        o.gateX = proxy.readFloat32()
        o.gateY = proxy.readFloat32()
        o.terrainFlags = proxy.readInt32()
        o.supplyPopulationType = proxy.readInt16()
        o.isDoor = proxy.readInt8()
        o.desc1 = proxy.readString()
        o.desc2 = proxy.readString()
        o.desc3 = proxy.readJsonObject()
        o.f1 = proxy.readBooleanArray()
        o.f2 = proxy.readUInt8Array()
        o.f3 = proxy.readUInt16Array()
        o.f4 = proxy.readUInt32Array()
        o.f5 = proxy.readInt8Array()
        o.f6 = proxy.readInt16Array()
        o.f7 = proxy.readInt32Array()
        o.f8 = proxy.readFloat32Array()
        o.f9 = proxy.readStringArray()
        o.f10 = proxy.readStringArray()
    }
 }