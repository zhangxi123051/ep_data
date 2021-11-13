package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Type          string `json:"type"`
	Status        string `json:"status"`
	AddSystemTime int    `json:"add_system_time"`
	LastLoginTime int    `json:"last_login_time"`
	Remark        string `json:"remark"`
	CreateAt      string `json:"create_at"`
	UpdatedAt     string `json:"update_at"`
	Isbond        string `json:"isbond"`
	Secret        string `json:"secret"`
}

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type ResData struct {
	Uname *UserInfo `json:"userInfo"`
	Tkn   *Token    `json:"token"`
}

type ResMsg struct {
	Data *ResData `json:"data"`
}

func UserLogin(c *gin.Context) {
	user := UserInfo{
		ID:       1,
		Username: "admin",
		Secret:   "Z4B7KBXK5C72SBGT",
	}
	toke := Token{
		AccessToken:      "123456",
		ExpiresIn:        1629701673,
		RefreshExpiresIn: 1629701673,
	}
	rdata := ResData{
		Tkn:   &toke,
		Uname: &user,
	}
	msg := ResMsg{
		Data: &rdata,
	}

	message, _ := json.Marshal(msg)
	log.Println(string(message))

	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Println("%v", &json)
	c.JSON(http.StatusOK,
		gin.H{
			"message": msg,
			"code":    http.StatusOK})
	return
}
