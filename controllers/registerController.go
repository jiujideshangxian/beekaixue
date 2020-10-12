package controllers

import (
	db_mysql2 "DataCertPaltPhone/db_mysql"
	"DataCertPaltPhone/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

type ResgiterController struct {
	beego.Controller
}
//该方法用于用户注册的逻辑
func (r *ResgiterController) Post() {
	//解析用户端提交的数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("数据解析失败，请重试。")
		return
	}
	//将解析到的数据保存到数据库里
	row,err := db_mysql2.AddUser(user)
	/*
	将处理结果返回到客户端浏览器，
	如果成功跳转登入页面，nn
	如果失败跳转错误页面
	 */
	if err != nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户信息注册失败")
		r.TplName = "err.html"
		return
	}
	fmt.Println(row)
	md5Hash := md5.New()
	md5Hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(md5Hash.Sum(nil))
	r.TplName = "login.html"
}
