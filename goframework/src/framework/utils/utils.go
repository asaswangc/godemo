package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"github.com/pborman/uuid"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// UserHomeDir 获取用户目录
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// AbsPath 获取绝对路径
func AbsPath(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = UserHomeDir() + inPath[5:]
	}
	inPath = os.ExpandEnv(inPath)
	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}
	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}

// https://segmentfault.com/a/1190000017346458
// 加解密算法：对称性加密算法、非对称性加密算法、散列算法，其中散列算法不可逆，无法解密，故而只能用于签名校验、身份验证
// 对称性加密算法：DES、3DES、AES
// 非对称性加密算法：RSA、DSA、ECC
// 散列算法：MD5、SHA1、HMAC

// GenerateRSAKey 生成私钥和公钥, bits参数指定证书大小
// 也可以直接通过openssl命令生成：
// 私钥：openssl genrsa -out rsa_private_key.pem 2048
// 公钥：openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func GenerateRSAKey(bits int) error {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer func() {
		_ = privateFile.Close()
	}()

	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	defer func() {
		_ = publicFile.Close()
	}()

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}
	return nil
}

// EncryptWithRSA rsa加密
func EncryptWithRSA(plainText string, publicKeyPath string) (string, error) {
	//打开公钥文件
	keyFile, err := os.Open(publicKeyPath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = keyFile.Close()
	}()

	//读取公钥内容
	info, _ := keyFile.Stat()
	buf := make([]byte, info.Size())
	_, err = keyFile.Read(buf)
	if err != nil {
		return "", err
	}

	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(cipherBytes), nil
}

// DecryptWithRSA rsa解密
func DecryptWithRSA(cipherText string, privateKeyPath string) (string, error) {
	//打开私钥文件
	keyFile, err := os.Open(privateKeyPath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = keyFile.Close()
	}()

	//获取私钥内容
	info, _ := keyFile.Stat()
	buf := make([]byte, info.Size())
	_, err = keyFile.Read(buf)
	if err != nil {
		return "", err
	}

	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//对密文进行解密
	cipherBytes, err := base64.RawURLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	plainBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

// EncryptWithSha256 sha256加密
func EncryptWithSha256(data string) string {
	// 先base64对原始数据进行编码
	tmp := base64.StdEncoding.EncodeToString([]byte(data))

	// 再使用sha256进行两层加密
	h := sha256.New()
	h.Write([]byte(tmp))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func StructToJson(Struct interface{}) (string, error) {
	bytes, err := json.Marshal(Struct)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func GetUuid() string {
	uuidWithHyphen := uuid.NewRandom()
	return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}
