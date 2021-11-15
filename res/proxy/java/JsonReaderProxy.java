package xu;

import java.util.List;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

public class JsonReaderProxy {

	private JSONObject object;

	public void setObject(JSONObject object) {
		this.object = object;
	}

	public boolean getBoolean(String key) throws JSONException {
		return object.getBoolean(key);
	}

	public byte getInt8(String key) throws JSONException {
		return (byte) object.getInt(key);
	}

	public short getInt16(String key) throws JSONException {
		return (short) object.getInt(key);
	}

	public int getInt32(String key) throws JSONException {
		return object.getInt(key);
	}

	public short getUInt8(String key) throws JSONException {
		return (short) object.getInt(key);
	}

	public int getUInt16(String key) throws JSONException {
		return object.getInt(key);
	}

	public long getUInt32(String key) throws JSONException {
		return object.getLong(key);
	}

	public float getFloat32(String key) throws JSONException {
		return (float) object.getDouble(key);
	}

	public String getString(String key) throws JSONException {
		return object.getString(key);
	}

	public JSONObject getJsonObject(String key) throws JSONException {
		return object.getJSONObject(key);
	}

	// -----------------------------------------------------

	public boolean[] getBooleanArray(String key) throws JSONException {
		return getBoolArray(key);
	}

	public byte[] getInt8Array(String key) throws JSONException {
		int[] arr = getIntArray(key);
		byte[] rs = new byte[arr.length];
		for (int i = 0; i < arr.length; i++) {
			rs[i] = (byte) arr[i];
		}
		return rs;
	}

	public short[] getInt16Array(String key) throws JSONException {
		int[] arr = getIntArray(key);
		short[] rs = new short[arr.length];
		for (int i = 0; i < arr.length; i++) {
			rs[i] = (short) arr[i];
		}
		return rs;
	}

	public int[] getInt32Array(String key) throws JSONException {
		return getIntArray(key);
	}

	public short[] getUInt8Array(String key) throws JSONException {
		int[] arr = getIntArray(key);
		short[] rs = new short[arr.length];
		for (int i = 0; i < arr.length; i++) {
			rs[i] = (short) arr[i];
		}
		return rs;
	}

	public int[] getUInt16Array(String key) throws JSONException {
		return getIntArray(key);
	}

	public long[] getUInt32Array(String key) throws JSONException {
		return getLongArray(key);
	}

	public float[] getFloat32Array(String key) throws JSONException {
		double[] arr = getDoubleArray(key);
		float[] rs = new float[arr.length];
		for (int i = 0; i < arr.length; i++) {
			rs[i] = (float) arr[i];
		}
		return rs;
	}

	public String[] getStringArray(String key) throws JSONException {
		return getStrArray(key);
	}

	public JSONArray getJsonArray(String key) throws JSONException {
		return object.getJSONArray(key);
	}

	// --------------------------------------

	private final boolean[] getBoolArray(String key) throws JSONException {
		JSONArray ja = object.getJSONArray(key);
		boolean[] rs = new boolean[ja.length()];
		for (int i = 0; i < rs.length; i++) {
			rs[i] = ja.getBoolean(i);
		}
		return rs;
	}

	private final int[] getIntArray(String key) throws JSONException {
		JSONArray ja = object.getJSONArray(key);
		int[] rs = new int[ja.length()];
		for (int i = 0; i < rs.length; i++) {
			rs[i] = ja.getInt(i);
		}
		return rs;
	}

	private final long[] getLongArray(String key) throws JSONException {
		JSONArray ja = object.getJSONArray(key);
		long[] rs = new long[ja.length()];
		for (int i = 0; i < rs.length; i++) {
			rs[i] = ja.getLong(i);
		}
		return rs;
	}

	private final double[] getDoubleArray(String key) throws JSONException {
		JSONArray ja = object.getJSONArray(key);
		double[] rs = new double[ja.length()];
		for (int i = 0; i < rs.length; i++) {
			rs[i] = ja.getDouble(i);
		}
		return rs;
	}

	private final String[] getStrArray(String key) throws JSONException {
		List<Object> list = object.getJSONArray(key).toList();
		String[] rs = new String[list.size()];
		list.toArray(rs);
		return rs;
	}

}
