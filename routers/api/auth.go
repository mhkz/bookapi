package api

import (
	"bookapi/pkg/logging"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"bookapi/models"
	"bookapi/pkg/e"
	"bookapi/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

type Email struct {
	Email    string `form:"email" json:"email" valid: "Required; MaxSize(50)"`
	Password string `form:"password" json:"password" valid: "Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	requestData, _ := ioutil.ReadAll(c.Request.Body)
	bodyString := string(requestData)
	var email Email
	var username string
	var password string
	var valid = validation.Validation{}
	if err := json.Unmarshal([]byte(bodyString), &email); err == nil {
		username = email.Email
		password = email.Password
	}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
