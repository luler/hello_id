package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/jwt_helper"
	"go_test/app/helper/log_helper"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// 登录
func Login(c *gin.Context) {
	type Param struct {
		Name     string `validate:"required,max=50" label:"账号"`
		Password string `validate:"required,max=50" label:"密码"`
	}
	var param Param
	request_helper.InputStruct(c, &param)
	var user model.User
	err := db_helper.Db().Where("name=?", param.Name).First(&user)
	if err.Error != nil {
		if param.Name == os.Getenv("ADMIN_NAME") {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
			user.Name = param.Name
			user.Password = string(hashedPassword)
			user.Status = 1
			res := db_helper.Db().Save(&user)
			if res.Error != nil {
				exception_helper.CommonException("系统异常")
			}
			db_helper.Db().Where("name=?", param.Name).First(&user)
		} else {
			exception_helper.CommonException("用户不存在")
		}
	}

	//判断密码是否一致
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if passwordErr != nil {
		exception_helper.CommonException("密码输入有误")
	}
	//是否启用
	if user.Status != 1 {
		exception_helper.CommonException("用户已被禁用")
	}

	res := jwt_helper.IssueToken(gin.H{
		"uid": user.Id,
	})

	log_helper.Info("登录成功", res)

	response_helper.Success(c, "登录成功", res)
}

// CAS登录
func CasLogin(c *gin.Context) {
	type Param struct {
		Code    string `validate:"required" label:"授权码"`
		Open_id string `validate:"required" label:"开放应用用户ID"`
	}
	var param Param
	request_helper.InputStruct(c, &param)

	// 将 JSON 数据编码为字节切片
	jsonData, _ := json.Marshal(map[string]interface{}{
		"code":      param.Code,
		"appid":     os.Getenv("CAS_APPID"),
		"appsecret": os.Getenv("CAS_APPSECRET"),
	})

	// 创建 POST 请求
	req, _ := http.NewRequest("POST", os.Getenv("CAS_HOST")+"/api/auth/checkCode", bytes.NewBuffer(jsonData))
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		exception_helper.CommonException(err.Error())
	}
	defer resp.Body.Close()
	// 读取响应体
	body, _ := ioutil.ReadAll(resp.Body)

	res := make(map[string]interface{})
	json.Unmarshal(body, &res)
	if int(res["code"].(float64)) != 200 {
		exception_helper.CommonException(res["message"])
	}

	var user model.User
	if db_helper.Db().Where("name=?", os.Getenv("ADMIN_NAME")).First(&user).Error != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		user.Name = os.Getenv("ADMIN_NAME")
		user.Password = string(hashedPassword)
		user.Status = 1
		if db_helper.Db().Save(&user).Error != nil {
			exception_helper.CommonException("系统异常")
		}
		db_helper.Db().Where("name=?", os.Getenv("ADMIN_NAME")).First(&user)
	}

	res = jwt_helper.IssueToken(gin.H{
		"uid": user.Id,
	})

	c.Redirect(http.StatusMovedPermanently, "/helloId/idRule?token="+res["token"].(string))
}

// 重新设置密码
func ResetPassword(c *gin.Context) {
	type Param struct {
		Password        string `validate:"required,min=5,max=20" label:"新密码"`
		ConfirmPassword string
	}
	var param Param
	request_helper.InputStruct(c, &param)
	if param.Password != param.ConfirmPassword {
		exception_helper.CommonException("确认密码输入有误")
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", param.Password)
	if !matched {
		exception_helper.CommonException("密码只能包含数字或英文符号")
	}
	uid, _ := c.Get("uid")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	db_helper.Db().Where("id=?", uid).Updates(&model.User{
		Password: string(hashedPassword),
	})

	response_helper.Success(c, "修改成功")
}
