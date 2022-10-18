package middleware

import (
	"errors"
	"log"
	"net/http"
	"pandownload/logic"
)

// Cors 同源策略
func Cors(w http.ResponseWriter, r *http.Request) {
	log.Printf("visited %s\n", r.URL.Path)
	//跨域中间件
	method := r.Method
	origin := r.Header.Get("Origin") //请求头部
	if origin != "" {
		//接收客户端发送的origin （重要！）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,Authorization")
		//设置缓存时间
		w.Header().Set("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	//允许类型校验
	if method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
	}
}

// AccessToken token验证
func AccessToken(w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("Authorization")
	_, err := logic.ParseToken(token)
	if err != nil {
		return err
	}
	return nil
}

func GetUsernameToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	username := logic.GetUsername(token)
	if username == "" {
		return "", errors.New(logic.InvalidToken)
	}
	return username, nil
}

func CostTimeTotal(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
}
