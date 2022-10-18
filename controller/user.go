package controller

import (
	"fmt"
	"log"
	"net/http"
	"pandownload/logic"
	"pandownload/middleware"
	"pandownload/utils"
)

type UserService struct {
	Handler map[string]func(w http.ResponseWriter, r *http.Request)
}

func (u *UserService) RegisterHandler(relativepath string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.Handle(relativepath, u)
	u.Handler[relativepath] = handler
}

func (u *UserService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Cors(w, r)
	//调用对应处理器
	u.Handler[r.URL.Path](w, r)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic info is: %v", err)
		}
	}()
}

func (u *UserService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//获取json数据
	userinfo := utils.ParseBody(r.Body)
	fmt.Println(userinfo)
	userinfo["password"] = utils.MD5([]byte(userinfo["password"]))
	err := logic.CheckuserInfo(userinfo)
	responseInfo := ParamLogin{}
	if err != nil {
		responseInfo.Flag = false
	} else {

		responseInfo.Flag = true
		token, _ := logic.GenerateToken(userinfo["username"])
		w.Header().Set("Authorization", token)
	}
	w.Write(utils.JSONData(responseInfo))
}

func (u *UserService) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	userinfo := utils.ParseBody(r.Body)
	err := logic.CheckuserInfo(userinfo)
	responseInfo := ParamSignUp{}
	userinfo["password"] = utils.MD5([]byte(userinfo["password"]))
	email := userinfo["email"]
	authcode := userinfo["authcode"]
	fmt.Println(userinfo)
	//用户已存在或验证码错误
	if err.Error() == logic.ErrPassword || err == nil {
		responseInfo.Set(logic.ErruserExisted, false)
	} else if errs := logic.VerifyauthCode(email, authcode); errs != nil {
		if errs.Error() == logic.ErrauthCode {
			responseInfo.Set(logic.ErrauthCode, false)
		}
	} else if err.Error() == logic.ErrUsername {
		//用户不存在，前往注册
		err = logic.SignupUser(userinfo["username"], userinfo["password"], email)
		if err != nil {
			responseInfo.Set(logic.ErrServicebusy, false)
		} else {
			responseInfo.Set("", true)
		}
	}
	w.Write(utils.JSONData(responseInfo))
}

func (u *UserService) SendemailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	fmt.Println("email=   ", email)
	flag := ParamBoolean{}
	if email == "" {
		flag.Flag = false
		return
	} else {
		authCode, _ := logic.GenerateauthCode(email)
		err := logic.SendmailTo(email, "验证码为"+authCode)
		if err != nil {
			flag.Flag = false
		} else {
			flag.Flag = true
		}
	}
	w.Write(utils.JSONData(flag))
}
