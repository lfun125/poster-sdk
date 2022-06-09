
/**
 * 给出的KEY是PKCS1 格式的需要转成PKCS8
 * openssl pkcs8 -topk8 -inform PEM -in {pkcs1.key PKCS1的私钥匙} -outform PEM -nocrypt -out {out.pem pkcs8的私钥}
 */
import java.nio.charset.StandardCharsets;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.Signature;
import java.security.spec.PKCS8EncodedKeySpec;
import java.util.Base64;

public class Main {

    // 这是PKCS8的私钥
    public static String key = "MIIBVwIBADANBgkqhkiG9w0BAQEFAASCAUEwggE9AgEAAkEAsFC2PBKPEHwg8GzLMpT2wAez5nbJs4W9u/36y4y29Jz0e3IrcK8mXxdtxSMrA7cBPyOMT7PkBrcwNp1n225KWwIDAQABAkEArbE1dYSK76B5CqECpHffhVmRPm6zUWllerc/xBqBegMqZAyUkbSX4S+h76X6COwrAGf9rXayt3mTxrxbi0fXYQIhAMxKzfrYd2cA3c9scOKuUSveXQOpllfCdaWCgfWoIbFTAiEA3PEhlQ4nSy5kHObrfEos7vePOiP1+YXEf7PwrTIeudkCIQCb/obYg5BxQ7Ub3Sc5wHfU8p+92zIk4yUoc8Y+ydqoZwIhANh6dvIJ5Rw3vyXGaFLmhrI5458O3xJ2K7sIPrgkVJC5AiEAzAiLCp4XVsbrJLz93C0OSNuVXrhnUqEeda4tcsOrxx0=";

    public static void main(String[] args) throws Exception {

        String data = "你好中国";
        String sign = generateSignature(data);
        // 输出字符串
        System.out.println(sign);
    }

    /**
     * 对字符串生成签名
     * 
     * @param data 需要签名的字符串
     * @return
     * @throws Exception
     */
    public static String generateSignature(final String data) throws Exception {
        PrivateKey privateKey = loadPrivateKey(key);
        Signature privateSignature = Signature.getInstance("SHA256withRSA");
        privateSignature.initSign(privateKey);
        privateSignature.update(data.getBytes(StandardCharsets.UTF_8));
        byte[] s = privateSignature.sign();
        return convertByteToHexadecimal(s);
    }

    /**
     * 加载私钥
     * 
     * @param privateKeyStr PKCS8私钥字符串
     * @return
     * @throws Exception
     */
    public static PrivateKey loadPrivateKey(String privateKeyStr) throws Exception {
        String string = privateKeyStr
                .replace("-----BEGIN RSA PRIVATE KEY-----", "")
                .replace("-----END RSA PRIVATE KEY-----", "")
                .replace("-----BEGIN PRIVATE KEY-----", "")
                .replace("-----END PRIVATE KEY-----", "")
                .replaceAll("\\s", "");
        byte[] buffer = base64Decode(string);
        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(buffer);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");
        return keyFactory.generatePrivate(keySpec);
    }

    public static String base64Encode(byte[] data) {
        return Base64.getEncoder().encodeToString(data);
    }

    public static byte[] base64Decode(String data) {
        return Base64.getDecoder().decode(data);
    }

    public static String convertByteToHexadecimal(byte[] byteArray) {
        String hex = "";
        // Iterating through each byte in the array
        for (byte i : byteArray) {
            hex += String.format("%02X", i);
        }
        return hex.toLowerCase();
    }

}