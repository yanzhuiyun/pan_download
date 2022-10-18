package controller

import (
	"fmt"
	"net/http"
	"pandownload/logic"
	"pandownload/middleware"
	"pandownload/settings"
	"pandownload/utils"
)

type FileService struct {
	BasicController
}

func (f *FileService) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	_, fpHeader, err := r.FormFile("filename")
	fmt.Println("filename=", fpHeader.Filename)
	responseInfo := ParamBoolean{}
	if err != nil {
		responseInfo.Flag = false
		w.Write(utils.JSONData(responseInfo))
		fmt.Println("err=", err)
	}
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	fpHeader.Header.Set("username", username)
	err = logic.SaveUpload(fpHeader, settings.StorePath())
	if err == nil {
		responseInfo.Flag = true
	}
	fmt.Println(err)
	w.Write(utils.JSONData(responseInfo))
}

func (f *FileService) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	//获取用户文件
	filename := r.FormValue("filename")
	fmt.Println("filename==>", filename)
	data := logic.GetfileData(username, filename)
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Write(data)
}

func (f *FileService) GetUserFilesHandler(w http.ResponseWriter, r *http.Request) {
	//获取username
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	//获取文件列表
	filenames := logic.GetuserFiles(username)
	if filenames == nil {
		return
	}
	reponseInfo := ParamFiles{filenames}
	w.Write(utils.JSONData(reponseInfo))
}

// DeleteFileHandler 删除文件
func (f *FileService) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	fmt.Println("filename==>", filename)
	err = logic.DeleteFile(username, filename)
	reponseInfo := &ParamBoolean{}
	if err == nil {
		reponseInfo.Flag = true
	}
	w.Write(utils.JSONData(reponseInfo))
}

func (f *FileService) OnlineAnalysisFileHandler(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	//获取用户文件
	filename := r.FormValue("filename")
	data := logic.GetfileData(username, filename)
	if data == nil {
		w.Write([]byte("file not existed"))
		return
	}
	format := logic.ConfirmFormat(filename)
	if format == "" {
		format = "application/plain"
	}
	w.Header().Set("Content-Type", format)
	w.Write(data)
}

func (f *FileService) SearchFileHandler(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.JSONData(ParamString{Msg: err.Error()}))
		return
	}
	searchStr := r.FormValue("search")
	fmt.Println("search==>", searchStr)
	files, err := logic.SearchDoc(username, searchStr)
	if err != nil {
		w.Write(nil)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resonseInfo := ParamStringSlice{Slice: files}
	w.Write(utils.JSONData(resonseInfo))
}

func (f *FileService) RemoteDownloadHandler(w http.ResponseWriter, r *http.Request) {
	//获取http或https的链接
	//url := r.FormValue("url")
	////将资源下载至服务器
	//resp, err := http.DefaultClient.Get(url)
	//if err != nil {
	//	w.WriteHeader(http.StatusServiceUnavailable)
	//	return
	//}

}
