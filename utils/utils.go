package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
)

//1. 用于对用户密码等私密信息进行加密
//2. 用于计算文件的唯一hash值

var (
	sh1 = sha1.New()
	md  = md5.New()
)

const (
	Sha1Secret = "sha1.acat.pan-download.com"
	Md5Secret  = "md5.acat.pan-download.com"
)

func init() {
	sh1.Write([]byte(Sha1Secret))
	md.Write([]byte(Md5Secret))
}

func SHA1(data []byte) string {
	return hex.EncodeToString(sh1.Sum(data))
}

func MD5(data []byte) string {
	return hex.EncodeToString(md.Sum(data))
}

func FileMd5(file *multipart.FileHeader) string {
	fp, err := file.Open()
	if err != nil {
		return ""
	}
	defer fp.Close()
	io.Copy(md, fp)
	return hex.EncodeToString(md.Sum(nil))
}

func FileSha1(file *multipart.FileHeader) string {
	fp, err := file.Open()
	if err != nil {
		return ""
	}
	defer fp.Close()
	io.Copy(sh1, fp)
	return hex.EncodeToString(md.Sum(nil))
}

// ParseBody 解析body的数据
func ParseBody(reader io.Reader) map[string]string {
	info := make(map[string]string)
	decode := json.NewDecoder(reader)
	decode.Decode(&info)
	return info
}

func JSONData(info any) []byte {
	data, err := json.Marshal(info)
	if err != nil {
		return nil
	}
	return data
}
