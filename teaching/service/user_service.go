package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
	"teach.me/teaching/config"
	"teach.me/teaching/http/ret"
	"teach.me/teaching/mongo"
	"teach.me/teaching/tlog"
	"teach.me/teaching/utils"
)

const (
	TOKEN_NULL    = 1
	TOKEN_EXPIRE  = 2
	TOKEN_ILLEGAL = 3
)

type User struct {
	Name     string `json:name`
	NickName string `json:nickName`
	Phone    string `json:phone`
	Avatar   string `json:avatar`
	Sex      int    `json:sex`
	City     string `json:city`
	Token    string `json:token`
}

type TokenError int

type Session struct {
	Token  string
	Expire int64
}

// token mapping session
var sc = make(map[string]Session, 100)

//func CheckToken(token string) (bool, int) {
//	if token == nil || len(token) == 0 {
//		return false, 1
//	}
//	session := sc[token]

//	if session == nil {
//		return false, 2
//	}
//	if session.Expire < time.Now().Unix() {
//		return false, 3
//	}
//	return true, 0

//}

func UserBaseReg(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var user map[string]string
	err = json.Unmarshal(bs, &user)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	user["name"] = user["phone"]
	user["nickName"] = user["phone"]
	user["avatar"] = ""
	user["sex"] = "-1"
	user["city"] = ""
	err = mongo.GetCollection(mongo.USER_COLL).Insert(user)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	//	url := fmt.Sprintf("http://%s:%d%s", config.Gconfig.Server, config.Gconfig.Port, req.TEACHING_USER_BASE_LOGIN)
	//	http.Post(url, "application/json;charset=utf-8")
	UserBaseLogin(w, r)
}

func UserBaseLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	fmt.Println(string(body))
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var temp map[string]string
	err = json.Unmarshal(body, &temp)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var result bson.M
	mongo.GetCollection(mongo.USER_COLL).Find(bson.M{"phone": temp["phone"], "pwd": temp["pwd"]}).One(&result)
	fmt.Println(result)
	if len(result) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, ret.USER_CHECK_FAILED)
		return
	}

	var expire = time.Now().Unix() + config.Gconfig.ExpireToken
	var token = fmt.Sprintf("%s-%d", utils.Rand().Hex(), expire)
	session := Session{Token: token, Expire: expire}
	sc[token] = session
	result["token"] = token
	resp, ex := json.Marshal(result)
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(resp))
}
func UserBasePhoneCheck(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
	}
	var u map[string]string
	err = json.Unmarshal(body, &u)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
	}
	var result bson.M
	mongo.GetCollection(mongo.USER_COLL).Find(bson.M{"phone": u["phone"]}).One(&result)
	var vaild = 1
	if len(result) > 0 {
		vaild = 0
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{vaild:"+strconv.Itoa(vaild)+"}")
}
