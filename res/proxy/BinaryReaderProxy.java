package xu;

import java.io.DataInputStream;
import java.nio.charset.Charset;

import org.json.JSONArray;
import org.json.JSONObject;

public class BinaryReaderProxy {

	private DataInputStream dis;
	private Charset charset;
	private StringBuilder sb = new StringBuilder();

	public BinaryReaderProxy(DataInputStream dis, String charsetName) {
		super();
		this.dis = dis;
		this.charset = Charset.forName(charsetName);
	}

	public boolean readBoolean() throws Exception {
		return dis.readBoolean();
	}

	public byte readInt8() throws Exception {
		return dis.readByte();
	}

	public short readInt16() throws Exception {
		return dis.readShort();
	}

	public int readInt32() throws Exception {
		return dis.readInt();
	}

	public short readUInt8() throws Exception {
		byte val = dis.readByte();
		if (val < 0) {
			return (short) (val + ((short) Byte.MAX_VALUE + 1) * 2);
		} else {
			return (short) val;
		}
	}

	public int readUInt16() throws Exception {
		short val = dis.readShort();
		if (val < 0) {
			return (int) (val + ((int) Short.MAX_VALUE + 1) * 2);
		} else {
			return (int) val;
		}
	}

	public long readUInt32() throws Exception {
		int val = dis.readInt();
		if (val < 0) {
			return (long) (val + ((long) Long.MAX_VALUE + 1) * 2);
		} else {
			return (long) val;
		}
	}

	public float readFloat32() throws Exception {
		float value = dis.readFloat();
		return value;
	}

	public String readString() throws Exception {
		int len = dis.readUnsignedShort();
		byte[] strByte = new byte[len];
		dis.read(strByte);
		return new String(strByte, charset).toString();
	}

	public JSONObject readJsonObject() throws Exception {
		String jsonStr = readString();
		JSONObject obj = new JSONObject(jsonStr);
		return obj;
	}

	// -----------------------------------------------------

	public boolean[] readBooleanArray() throws Exception {
		int len = dis.readUnsignedShort();
		boolean[] rs = new boolean[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readBoolean();
		}
		return rs;
	}

	public short[] readUInt8Array() throws Exception {
		int len = dis.readUnsignedShort();
		short[] rs = new short[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readUInt8();
		}
		return rs;
	}

	public int[] readUInt16Array() throws Exception {
		int len = dis.readUnsignedShort();
		int[] rs = new int[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readUInt16();
		}
		return rs;
	}

	public long[] readUInt32Array() throws Exception {
		int len = dis.readUnsignedShort();
		long[] rs = new long[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readUInt32();
		}
		return rs;
	}

	public byte[] readInt8Array() throws Exception {
		int len = dis.readUnsignedShort();
		byte[] rs = new byte[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readInt8();
		}
		return rs;
	}

	public short[] readInt16Array() throws Exception {
		int len = dis.readUnsignedShort();
		short[] rs = new short[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readInt16();
		}
		return rs;
	}

	public int[] readInt32Array() throws Exception {
		int len = dis.readUnsignedShort();
		int[] rs = new int[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readInt32();
		}
		return rs;
	}

	public float[] readFloat32Array() throws Exception {
		int len = dis.readUnsignedShort();
		float[] rs = new float[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readFloat32();
		}
		return rs;
	}

	public String[] readStringArray() throws Exception {
		int len = dis.readUnsignedShort();
		String[] rs = new String[len];
		for (int i = 0; i < len; i++) {
			rs[i] = this.readString();
		}
		return rs;
	}

	public JSONArray readJsonArray() throws Exception {
		String[] jsonContextArray = readStringArray();
		sb.setLength(0);
		sb.append("[");
		for (String string : jsonContextArray) {
			sb.append("\"" + string + "\",");
		}
		sb.setCharAt(sb.length() - 1, ']');
		JSONArray a = new JSONArray(sb.toString());
		return a;
	}
}
