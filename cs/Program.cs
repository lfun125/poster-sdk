class Program
{
    static void Main()
    {
        //SetUser();
        string token = Login();
        token = Decrp(token, privateKey);
        string url = $"https://m.daishuhaibao.com/#/pages/account/login/login?token={token}";
        Process.Start(url);
        Console.ReadKey();
    }
    private const string baseUrl = "https://api.daishuhaibao.com";
    private const string mid = "2276907093";//迅德服务商账号
    private const string privateKey = @"<RSAKeyValue><Modulus>uHn7fV3+50qFiWnTtAjCRhxa+bx56d+CRYcwtL2nJZ02ofs4qCKc/W5eV1Rcljebg/fPzpNn8ObaDIVJUno9qQ==</Modulus><Exponent>AQAB</Exponent><P>8QolinOY1G2QReOnRjFLUaAjdMS1xg5NVaGB/gZ4mSs=</P><Q>w+0cJTdvOTf3+NQN/JH2dzIJJYdyrJWfsnUWXYY68ns=</Q><DP>NcVpsoUvJtR6Rt0OR95fSwpKXpYzZsdyARRSh3SyGpk=</DP><DQ>GBAaAdQiG+Ps16mrohaHL7J8fxh4lAu4VmGdApWGzds=</DQ><InverseQ>TfvOOKOy2/lr9moT1vPh7nBhMUxFsIzj7rDTqXPLozA=</InverseQ><D>TwSxHLoufZQGLx4NeIhn7vAmZ+K7tdnvSyjlTAYQmqNJn39plkDSNax0z0KvJRX28gFBImP+PsKjlNuv3Y+GoQ==</D></RSAKeyValue>";
    private const string myid = "44815127401";//SetUser()获取到的用户编号
    private static void SetUser()
    {
        string url = baseUrl + "/merchant/set_user";
        string nickname = "迅德研发测试";
        string openid = "M3aLsz4tPB_0218";
        JObject jo = new JObject();
        jo["mid"] = mid;
        jo["nickname"] = nickname;
        jo["open_id"] = openid;
        jo["sign"] = (Signature(nickname + openid, privateKey));
        string result = Post(url, jo.ToString());
        Console.WriteLine(result);
    }
    private static string Login()
    {
        string url = baseUrl + "/merchant/login";
        string openid = "M3aLsz4tPB_0218";
        JObject jo = new JObject();
        jo["mid"] = mid;
        jo["open_id"] = openid;
        jo["sign"] = (Signature(openid, privateKey));
        string result = Post(url, jo.ToString());
        jo = JObject.Parse(result);
        return jo["token"].ToString();
    }

    private static string Post(string url, string postDataStr)
    {
        byte[] byteArray = Encoding.UTF8.GetBytes(postDataStr);
        HttpWebRequest objWebRequest = (HttpWebRequest)WebRequest.Create(url);
        objWebRequest.Method = "POST";
        objWebRequest.ContentType = "application/json";
        objWebRequest.ContentLength = byteArray.Length;
        using (Stream newStream = objWebRequest.GetRequestStream())
        {
            newStream.Write(byteArray, 0, byteArray.Length); //写入参数 
        }
        HttpWebResponse response = (HttpWebResponse)objWebRequest.GetResponse();
        using (StreamReader sr = new StreamReader(response.GetResponseStream(), Encoding.UTF8))
        {
            return sr.ReadToEnd(); // 返回的数据
        }
    }
    private static string Signature(string originalText, string privateKey)
    {
        byte[] byteData = Encoding.UTF8.GetBytes(originalText);
        RSACryptoServiceProvider provider = new RSACryptoServiceProvider();
        provider.FromXmlString(privateKey);
        byteData = provider.SignData(byteData, new SHA256CryptoServiceProvider());
        //return Convert.ToBase64String(byteData);//Base64格式
        return BitConverter.ToString(byteData, 0).Replace("-", string.Empty).ToLower();//HEX格式
    }
    private static string Decrp(string originalText, string privateKey)
    {
        byte[] byteData = ConvertHexStringToBytes(originalText);
        RSACryptoServiceProvider provider = new RSACryptoServiceProvider();
        provider.FromXmlString(privateKey);
        byteData = provider.Decrypt(byteData, false);
        return Encoding.UTF8.GetString(byteData);
        //return Convert.ToBase64String(byteData);//Base64格式
        //return BitConverter.ToString(byteData, 0).Replace("-", string.Empty).ToLower();//HEX格式
    }
    public static byte[] ConvertHexStringToBytes(string hexString)
    {
        hexString = hexString.Replace(" ", "");
        byte[] returnBytes = new byte[hexString.Length / 2];
        for (int i = 0; i < returnBytes.Length; i++)
        {
            returnBytes[i] = Convert.ToByte(hexString.Substring(i * 2, 2), 16);
        }

        return returnBytes;
    }
}