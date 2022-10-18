package router

import (
	"net/http"
	"pandownload/controller"
)

//type router struct {
//	//路由转发器
//	mux map[string]controller.HTTPService
//}
//
//func (rr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	path := strings.Split(r.URL.Path, "/")[1]
//	rr.mux[path].ServeHTTP(w,r)
//
//}
//
////服务注册
//func (rr *router) registerHTTPService(path string, service controller.HTTPService) {
//	prefix := strings.Split(path, "/")
//	rr.mux[prefix[1]] = service
//}

func Init() {
	userClient := &controller.UserService{Handler: make(map[string]func(http.ResponseWriter, *http.Request))}
	fileClient := &controller.FileService{BasicController: controller.BasicController{Handler: make(map[string]func(http.ResponseWriter, *http.Request))}}
	//进行接口装载
	userService := controller.HTTPService(userClient)
	fileService := controller.HTTPService(fileClient)
	userService.RegisterHandler("/user/login", userClient.LoginHandler)
	userService.RegisterHandler("/user/signup", userClient.SignUpHandler)
	userService.RegisterHandler("/user/sendemail", userClient.SendemailHandler)
	fileService.RegisterHandler("/file/upload", fileClient.UploadFileHandler)
	fileService.RegisterHandler("/file/download", fileClient.DownloadHandler)
	fileService.RegisterHandler("/file/getfiles", fileClient.GetUserFilesHandler)
	fileService.RegisterHandler("/file/delete", fileClient.DeleteFileHandler)
	fileService.RegisterHandler("/file/search", fileClient.SearchFileHandler)
	fileService.RegisterHandler("/file/online_analysis", fileClient.OnlineAnalysisFileHandler)
	userService.RegisterHandler("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PONG"))
	})
}
