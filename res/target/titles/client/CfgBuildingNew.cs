using System;
using UnityEngine;

/**
 * C#定义
 */
namespace {MODULE}
{
    /**
     * @author {AUTHOR}
     * Created on {DATE}.
     */
    [Serializable]
    public class CfgBuildingNew
    {
        // 以下为属性声明

        // ID
        [SerializeField] private ushort buildingType;
        // 建筑
        [SerializeField] private string name;
        // 是否可以升级
        [SerializeField] private byte promotion;
        // 门X坐标
        [SerializeField] private float gateX;
        // 门Y坐标
        [SerializeField] private float gateY;
        // 地形
        [SerializeField] private int terrainFlags;
        // 供需人口类型
        [SerializeField] private short supplyPopulationType;
        // 有没有门
        [SerializeField] private sbyte isDoor;
        // 建筑描述
        [SerializeField] private string desc1;
        // 建筑描述
        [SerializeField] private string desc2;
        // 建筑描述
        [SerializeField] private string desc3;
        // 测试1
        [SerializeField] private bool[] f1;
        // 测试2
        [SerializeField] private byte[] f2;
        // 测试3
        [SerializeField] private ushort[] f3;
        // 测试4
        [SerializeField] private uint[] f4;
        // 测试5
        [SerializeField] private sbyte[] f5;
        // 测试6
        [SerializeField] private short[] f6;
        // 测试7
        [SerializeField] private int[] f7;
        // 测试8
        [SerializeField] private float[] f8;
        // 测试9
        [SerializeField] private string[] f9;
        // 测试10
        [SerializeField] private string[] f10;

        // 以下为Get方法

        /**
         * ID
         *
         * @remark 注释列
         * @return ushort
         */
        publicushort BuildingType => buildingType;

        /**
         * 建筑
         *
         * @remark 建筑名称
         * @return string
         */
        publicstring Name => name;

        /**
         * 是否可以升级
         *
         * @remark 0：代表不能升级
1：代表可以升级
         * @return byte
         */
        publicbyte Promotion => promotion;

        /**
         * 门X坐标
         *
         * @remark 坐标从0开始
         * @return float
         */
        publicfloat GateX => gateX;

        /**
         * 门Y坐标
         *
         * @remark 
         * @return float
         */
        publicfloat GateY => gateY;

        /**
         * 地形
         *
         * @remark 4
         * @return int
         */
        publicint TerrainFlags => terrainFlags;

        /**
         * 供需人口类型
         *
         * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

         * @return short
         */
        publicshort SupplyPopulationType => supplyPopulationType;

        /**
         * 有没有门
         *
         * @remark 5
         * @return sbyte
         */
        publicsbyte IsDoor => isDoor;

        /**
         * 建筑描述
         *
         * @remark 建筑描述
string(20)
         * @return string
         */
        publicstring Desc1 => desc1;

        /**
         * 建筑描述
         *
         * @remark 建筑描述
string
         * @return string
         */
        publicstring Desc2 => desc2;

        /**
         * 建筑描述
         *
         * @remark 建筑描述
json
         * @return string
         */
        publicstring Desc3 => desc3;

        /**
         * 测试1
         *
         * @remark 测试1
         * @return bool[]
         */
        publicbool[] F1 => f1;

        /**
         * 测试2
         *
         * @remark 测试2
         * @return byte[]
         */
        publicbyte[] F2 => f2;

        /**
         * 测试3
         *
         * @remark 测试3
         * @return ushort[]
         */
        publicushort[] F3 => f3;

        /**
         * 测试4
         *
         * @remark 测试4
         * @return uint[]
         */
        publicuint[] F4 => f4;

        /**
         * 测试5
         *
         * @remark 测试5
         * @return sbyte[]
         */
        publicsbyte[] F5 => f5;

        /**
         * 测试6
         *
         * @remark 测试6
         * @return short[]
         */
        publicshort[] F6 => f6;

        /**
         * 测试7
         *
         * @remark 测试7
         * @return int[]
         */
        publicint[] F7 => f7;

        /**
         * 测试8
         *
         * @remark 测试8
         * @return float[]
         */
        publicfloat[] F8 => f8;

        /**
         * 测试9
         *
         * @remark 测试9
         * @return string[]
         */
        publicstring[] F9 => f9;

        /**
         * 测试10
         *
         * @remark 测试10
         * @return string[]
         */
        publicstring[] F10 => f10;

        // 以下为解释数据方法

        // Json数据解释
        public void ParseJson(string data)
        {
            JsonUtility.FromJsonOverwrite(data, this);
        }

        // 二进制数据解释
//		public void ParseBinary(proxy: xu.BinaryReaderProxy)
//		{   
            //this.buildingType = proxy.();
            //this.name = proxy.();
            //this.promotion = proxy.();
            //this.gateX = proxy.();
            //this.gateY = proxy.();
            //this.terrainFlags = proxy.();
            //this.supplyPopulationType = proxy.();
            //this.isDoor = proxy.();
            //this.desc1 = proxy.();
            //this.desc2 = proxy.();
            //this.desc3 = proxy.();
            //this.f1 = proxy.();
            //this.f2 = proxy.();
            //this.f3 = proxy.();
            //this.f4 = proxy.();
            //this.f5 = proxy.();
            //this.f6 = proxy.();
            //this.f7 = proxy.();
            //this.f8 = proxy.();
            //this.f9 = proxy.();
            //this.f10 = proxy.();
//		}

        // 静态实例化
        public static ${CLASS} FromJson(string data)
        {
            return JsonUtility.FromJson<${CLASS}>(data);
        }
    }
}