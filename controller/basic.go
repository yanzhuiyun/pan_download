package controller

import (
	"log"
	"net/http"
	"pandownload/middleware"
)

type HTTPService interface {
	RegisterHandler(relativepath string, handler func(http.ResponseWriter, *http.Request))
	http.Handler
}

// BasicController 基本控制器
type BasicController struct {
	Handler map[string]func(w http.ResponseWriter, r *http.Request)
}

func (b *BasicController) RegisterHandler(relativepath string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.Handle(relativepath, b)
	b.Handler[relativepath] = handler
}

func (b *BasicController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//跨域中间件
	middleware.Cors(w, r)
	//err := middleware.AccessToken(w, r)
	//if err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write(utils.JSONData(ParamString{Msg: logic.InvalidToken}))
	//	return
	//}
	b.Handler[r.URL.Path](w, r)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic info is: %v", err)
		}
	}()
}
