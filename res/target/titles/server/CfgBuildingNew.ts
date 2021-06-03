namespace {modle} { 
    /**
	 * @author {AUTHOR}
	 * Created on {DATE}.
	 */
	export class CfgBuildingNew {
        //以下为属性声明

        // ID
        private _buildingType: number;

        // 建筑
        private _name: string;

        // 规格
        private _layoutX: number;

        // 门X坐标
        private _gateX: number;

        // 门Y坐标
        private _gateY: number;

        // 地形
        private _terrainFlags: number;

        // 供需人口类型
        private _supplyPopulationType: number;

        // 有没有门
        private _isDoor: number;

        // 建筑描述
        private _desc1: string;

        // 建筑描述
        private _desc2: string;

        // 建筑描述
        private _desc3: any;

        // 测试1
        private _f1: boolean[];

        // 测试2
        private _f2: number[];

        // 测试3
        private _f3: number[];

        // 测试4
        private _f4: number[];

        // 测试5
        private _f5: number[];

        // 测试6
        private _f6: number[];

        // 测试7
        private _f7: number[];

        // 测试8
        private _f8: number[];

        // 测试9
        private _f9: string[];

        // 测试10
        private _f10: string[];

        //以下为Get方法

        /**
         * ID
         *
         * @remark 注释列
         * @return number
         */
        get buildingType(): number {
            return this._buildingType;
        }

        /**
         * 建筑
         *
         * @remark 建筑名称
         * @return string
         */
        get name(): string {
            return this._name;
        }

        /**
         * 规格
         *
         * @remark 
         * @return number
         */
        get layoutX(): number {
            return this._layoutX;
        }

        /**
         * 门X坐标
         *
         * @remark 坐标从0开始
         * @return number
         */
        get gateX(): number {
            return this._gateX;
        }

        /**
         * 门Y坐标
         *
         * @remark 
         * @return number
         */
        get gateY(): number {
            return this._gateY;
        }

        /**
         * 地形
         *
         * @remark 4
         * @return number
         */
        get terrainFlags(): number {
            return this._terrainFlags;
        }

        /**
         * 供需人口类型
         *
         * @remark 平民 1 -> 修改为 233
骑士 2 -> 修改为 234
贵族 3 -> 修改为 235

         * @return number
         */
        get supplyPopulationType(): number {
            return this._supplyPopulationType;
        }

        /**
         * 有没有门
         *
         * @remark 5
         * @return number
         */
        get isDoor(): number {
            return this._isDoor;
        }

        /**
         * 建筑描述
         *
         * @remark 建筑描述
string(20)
         * @return string
         */
        get desc1(): string {
            return this._desc1;
        }

        /**
         * 建筑描述
         *
         * @remark 建筑描述
string
         * @return string
         */
        get desc2(): string {
            return this._desc2;
        }

        /**
         * 建筑描述
         *
         * @remark 建筑描述
json
         * @return any
         */
        get desc3(): any {
            return this._desc3;
        }

        /**
         * 测试1
         *
         * @remark 测试1
         * @return boolean[]
         */
        get f1(): boolean[] {
            return this._f1;
        }

        /**
         * 测试2
         *
         * @remark 测试2
         * @return number[]
         */
        get f2(): number[] {
            return this._f2;
        }

        /**
         * 测试3
         *
         * @remark 测试3
         * @return number[]
         */
        get f3(): number[] {
            return this._f3;
        }

        /**
         * 测试4
         *
         * @remark 测试4
         * @return number[]
         */
        get f4(): number[] {
            return this._f4;
        }

        /**
         * 测试5
         *
         * @remark 测试5
         * @return number[]
         */
        get f5(): number[] {
            return this._f5;
        }

        /**
         * 测试6
         *
         * @remark 测试6
         * @return number[]
         */
        get f6(): number[] {
            return this._f6;
        }

        /**
         * 测试7
         *
         * @remark 测试7
         * @return number[]
         */
        get f7(): number[] {
            return this._f7;
        }

        /**
         * 测试8
         *
         * @remark 测试8
         * @return number[]
         */
        get f8(): number[] {
            return this._f8;
        }

        /**
         * 测试9
         *
         * @remark 测试9
         * @return string[]
         */
        get f9(): string[] {
            return this._f9;
        }

        /**
         * 测试10
         *
         * @remark 测试10
         * @return string[]
         */
        get f10(): string[] {
            return this._f10;
        }

        //以下为解释数据方法

        // Json数据解释
        public parseJson(data: any): void {
            // 以下为从 Json数据代理 中解释出字段数据
            this._buildingType = data.building_type_j;
            this._name = data.name_j;
            this._layoutX = data.layoutX_j;
            this._gateX = data.gateX_j;
            this._gateY = data.gateY_j;
            this._terrainFlags = data.terrain_flags_j;
            this._supplyPopulationType = data.supply_population_type_j;
            this._isDoor = data.isDoor_j;
            this._desc1 = data.desc1_j;
            this._desc2 = data.desc2_j;
            this._desc3 = data.desc3_j;
            this._f1 = data.f1_j;
            this._f2 = data.f2_j;
            this._f3 = data.f3_j;
            this._f4 = data.f4_j;
            this._f5 = data.f5_j;
            this._f6 = data.f6_j;
            this._f7 = data.f7_j;
            this._f8 = data.f8_j;
            this._f9 = data.f9_j;
            this._f10 = data.f10_j;
        }

        // 二进制数据解释
        public parseBinary(proxy: xu.BinaryReaderProxy): void {
            // 以下为从 二进制数据代理 中解释出字段数据
            this._buildingType = proxy.readUInt16();
            this._name = proxy.readString();
            this._layoutX = proxy.readUInt16();
            this._gateX = proxy.readFloat32();
            this._gateY = proxy.readFloat32();
            this._terrainFlags = proxy.readInt32();
            this._supplyPopulationType = proxy.readInt16();
            this._isDoor = proxy.readInt8();
            this._desc1 = proxy.readString();
            this._desc2 = proxy.readString();
            this._desc3 = proxy.readJsonObject();
            this._f1 = proxy.readBooleanArray();
            this._f2 = proxy.readUInt8Array();
            this._f3 = proxy.readUInt16Array();
            this._f4 = proxy.readUInt32Array();
            this._f5 = proxy.readInt8Array();
            this._f6 = proxy.readInt16Array();
            this._f7 = proxy.readInt32Array();
            this._f8 = proxy.readFloat32Array();
            this._f9 = proxy.readStringArray();
            this._f10 = proxy.readStringArray();
        }
	}
}