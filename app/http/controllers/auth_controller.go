package controllers

import (
	"fmt"
	"github.com/zhaozhentao/goblog/app/models/user"
	requests "github.com/zhaozhentao/goblog/app/requests"
	"github.com/zhaozhentao/goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
		    fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
		} else {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}

	// 3. 表单不通过 —— 重新显示表单
}

// Login 显示登录表单
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{},  "auth.login")
}

// DoLogin 处理登录表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	//
}
