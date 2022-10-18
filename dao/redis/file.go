package redis

import (
	"context"
	"fmt"
	"time"
)

const (
	//文件存有者数量的key
	fileHashKeys = "fileHashs:"
)

func IncrHashId(hash string, increment float64) (err error) {
	err = client.ZIncrBy(context.Background(), fileHashKeys, increment, hash).Err()
	return
}

// DeleteFileStatusByZero 随主函数开启而开启
func DeleteFileStatusByZero(done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
			//每天进行一次删除
		default:
			client.ZRemRangeByScore(context.Background(), fileHashKeys, "0", "0")
			time.Sleep(time.Hour * 24)
		}
	}
}

func generateIndKey(username string, ch rune) string {
	key := fmt.Sprintf("%s:ind:%v", username, ch)
	return key
}

func CreateInddoc(username string, filename []rune, doc string) (err error) {
	tx := client.TxPipeline()
	for _, ch := range filename {
		key := generateIndKey(username, ch)
		err = tx.SAdd(context.Background(), key, doc).Err()
		if err != nil {
			err = tx.Discard()
			fmt.Println("文件索引创建失败:", filename)
			if err != nil {
				return err
			}
		}
	}
	tx.Exec(context.Background())
	return err
}

// ParseSearch 返回一个key的集合
func parseSearch(searchStr string, username string) (all []string) {
	str := []rune(searchStr)
	//获取一个key的集合返回
	for _, v := range str {
		all = append(all, generateIndKey(username, v))
	}
	return
}

// DocInterStore 获取一个文档的集合
func DocInterStore(searchStr string, username string) ([]string, error) {
	//获取key
	keys := parseSearch(searchStr, username)
	//持有key去寻炸
	docs, err := client.SInter(context.Background(), keys...).Result()
	return docs, err
}
