package logic

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"pandownload/dao/mysql"
	"pandownload/dao/redis"
	"pandownload/model"
	"pandownload/utils"
	"strings"
)

const (
	// FileSizeLimits 文件超过1G使用分块上传
	FileSizeLimits = 1024 * 1024 * 1024
	FilePart       = 1024 * 1024 * 500
)

func SaveUpload(fpHeader *multipart.FileHeader, dstpath string) (err error) {
	hash := utils.FileMd5(fpHeader)
	//尝试打开文件
	fp, err := os.Open(dstpath + hash)
	if err != nil {
		//打开失败，文件不存在
		//创建文件
		src, _ := fpHeader.Open()
		fp, err = os.Create(dstpath + hash)
		if err != nil {
			return err
		}
		io.Copy(fp, src)
		defer func(src multipart.File) {
			err := src.Close()
			if err != nil {
				return
			}
		}(src)
		//传入文件hash值文件名进行保存到唯一文件表
		//唯一文件与用户文件可能发生冲突
		//存储值mysql
		err = mysql.SaveFile(&model.File{
			Hash:   hash,
			Path:   dstpath,
			Status: 1,
		})
		if err != nil {
			return
		}
	}
	defer closeFile(fp)
	username := fpHeader.Header.Get("username")
	filename := username + "_" + fpHeader.Filename
	err = mysql.SaveuserFile(&model.UserFile{
		Username: username,
		Filename: filename,
		Filehash: hash,
	})
	err = redis.IncrHashId(hash, 1)
	err = redis.CreateInddoc(username, []rune(fpHeader.Filename), fpHeader.Filename)
	return
}

// GetfileData 获取文件数据
func GetfileData(username string, filename string) []byte {
	sqlFilename := username + "_" + filename
	hash := mysql.Gethash(sqlFilename)
	path := mysql.GetPath(hash)
	//打开文件
	fmt.Println(path)
	fp, err := os.Open(path + hash)
	if err != nil {
		return nil
	}
	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func GetuserFiles(username string) []string {
	filenames, err := mysql.Getfiles(username)
	if err != nil {
		return nil
	}
	for i := 0; i < len(filenames); i++ {
		filenames[i] = strings.TrimPrefix(filenames[i], username+"_")
	}
	return filenames
}

func DeleteFile(username string, filename string) error {
	err := mysql.DeleteFile(username + "_" + filename)
	hash := mysql.Gethash(username + "_" + filename)
	err = redis.IncrHashId(hash, -1)
	return err
}

func closeFile(fp *os.File) {
	fp.Close()
}

func ConfirmFormat(filename string) string {
	extend := "content_type." + strings.SplitN(filename, ".", 2)[1]
	format := viper.GetString(extend)
	return format
}

func SearchDoc(username string, searchStr string) ([]string, error) {
	return redis.DocInterStore(searchStr, username)
}
