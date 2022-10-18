package main

import (
	"fmt"
	"log"
	"net/http"
	"pandownload/dao/mysql"
	"pandownload/dao/redis"
	"pandownload/router"
	"pandownload/settings"
	"pandownload/utils/email"
)

func main() {
	err := settings.Init()
	if err != nil {
		fmt.Println("init setting error:", err)
	}
	err = mysql.Init()
	if err != nil {
		fmt.Println("init mysql error:", err)
	}
	err = redis.Init()
	if err != nil {
		fmt.Println("init redis error:", err)
	}
	email.Init()
	router.Init()
	apiConfig := settings.API()
	//	go func() {
	log.Fatal(http.ListenAndServe(":"+apiConfig.Port, nil))
	//	}()
	//	str := fmt.Sprintf("%s:%s start listening %s ", apiConfig.Host, apiConfig.Port, apiConfig.Name)
	//	fmt.Println(str)
	//	quit := make(chan os.Signal)
	//	signal.Notify(quit, os.Interrupt)
	//	<-quit
	//	log.Println("Shutdown Server ...")
	//	//存留五分钟进行处理
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	//	defer cancel()
	//label:
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("service shutdown")
	//			break label
	//		default:
	//			fmt.Println("service is stopping")
	//			time.Sleep(time.Minute)
	//		}
	//	}
}
