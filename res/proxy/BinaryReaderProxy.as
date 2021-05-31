package xu {
	import flash.utils.ByteArray;

	public class BinaryReaderProxy {
		private var charsetName:String;
		private var data:ByteArray;

		public function BinaryReaderProxy(data:ByteArray, charsetName:String) {
			super();
			this.data=data;
			this.charsetName=charsetName;
		}

		public function readBoolean():Boolean {
			return data.readBoolean();
		}

		public function readInt8():int {
			return data.readByte();
		}

		public function readInt16():int {
			return data.readShort();
		}

		public function readInt32():int {
			return data.readInt();
		}

		public function readUInt8():uint {
			return data.readUnsignedByte();
//			var val:int=data.readByte();
//			if (val < 0) {
//				return uint(val + 256);
//			} else {
//				return uint(val);
//			}
		}

		public function readUInt16():uint {
			return data.readUnsignedShort();
//			var val:int=data.readShort();
//			if (val < 0) {
//				return uint(val + 65536);
//			} else {
//				return uint(val);
//			}
		}

		public function readUInt32():uint {
			return data.readUnsignedInt();
//			var val:int=data.readInt();
//			if (val < 0) {
//				return uint(val + uint.MAX_VALUE);
//			} else {
//				return uint(val);
//			}
		}

		public function readFloat32():Number {
			return data.readFloat();
		}

		public function readString():String {
			const len:uint=data.readUnsignedShort();
			var str:String=data.readMultiByte(len, charsetName);
			return str;
		}
		
		//---------------------------------------------
		
		public function readBooleanArray():Vector.<Boolean> {
			return this.readArray(this.readBoolean);
		}

		public function readInt8Array():Vector.<int> {
			return this.readArray(this.readInt8);
		}

		public function readInt16Array():Vector.<int> {
			return this.readArray(this.readInt16);
		}

		public function readInt32Array():Vector.<int> {
			return this.readArray(this.readInt32);
		}

		public function readUInt8Array():Vector.<uint> {
			return this.readArray(this.readUInt8);
		}

		public function readUInt16Array():Vector.<uint> {
			return this.readArray(this.readUInt16);
		}

		public function readUInt32Array():Vector.<uint> {
			return this.readArray(this.readUInt32);
		}

		public function readFloat32Array():Vector.<Number> {
			return this.readArray(this.readFloat32);
		}
		
		public function readStringArray():Vector.<String> {
			return this.readArray(this.readString);
		}
		
		private function readArray(func:Function): *{
			const len:int = this.data.readUnsignedShort();
			var rs:Vector.<*> = new Vector.<*>(len);
			for (var i = 0; i < len; i++) {
				rs[i] = func.call(this);
			}
			return rs;
		}
	}
}
