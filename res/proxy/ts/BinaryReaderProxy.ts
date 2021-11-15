/**
 * Created by Administrator on 2017/8/28.
 * @author xuzhuoxi
 */
namespace xu {
    export class BinaryReaderProxy {
        private byteArray: egret.ByteArray;
        private charsetNamed: string;

        constructor(byteArray: egret.ByteArray, charsetNamed: string) {
            this.byteArray = byteArray;
            this.charsetNamed = charsetNamed;
        }

        public readBoolean(): boolean {
            return this.byteArray.readBoolean();
        }

        public readInt8(): number {
            return this.byteArray.readByte();
        }

        public readInt16(): number {
            return this.byteArray.readShort();
        }

        public readInt32(): number {
            return this.byteArray.readInt();
        }

        public readUInt8(): number {
            return this.byteArray.readUnsignedByte();
        }

        public readUInt16(): number {
            return this.byteArray.readUnsignedShort();
        }

        public readUInt32(): number {
            return this.byteArray.readUnsignedInt();
        }

        public readString(): string {
            let len: number = this.byteArray.readUnsignedShort();
            return this.byteArray.readUTFBytes(len);
        }
		
		public readJsonObject(): any {
			let jsonStr: string = this.readString();
			let rs = JSON.parse(jsonStr);
			return rs;
		}

		//---------------------------------------------
		
        public readBooleanArray(): boolean[] {
			return this.readArray(this.readBoolean);
        }

        public readInt8Array(): number[] {
			return this.readArray(this.readInt8);
        }

        public readInt16Array(): number[] {
            return this.readArray(this.readInt16);
        }

        public readInt32Array(): number[] {
            return this.readArray(this.readInt32);
        }

        public readUInt8Array(): number[] {
            return this.readArray(this.readUInt8);
        }

        public readUInt16Array(): number[] {
            return this.readArray(this.readUInt16);
        }

        public readUInt32Array(): number[] {
            return this.readArray(this.readUInt32);
        }

        public readStringArray(): string[] {
            return this.readArray(this.readString);
        }
		
		public readJsonArray(): any[] {
			return this.readArray(this.readJsonObject);
		}
		
		private readArray(func:Function):any{
			let thisObj = this;
			let len: number = thisObj.byteArray.readUnsignedShort();
			let rs:any[] = [];
			for (let i = 0; i < len; i++) {
				rs.push(func.call(thisObj));
			}
            return rs;
		}
    }
}