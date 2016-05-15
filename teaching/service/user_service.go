package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
	"teach.me/teaching/config"
	"teach.me/teaching/http/req"
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
/*
1. check phone is vaild?
2. register phone and pwd
3. if success, then auto login


*/
func UserBaseReg(w http.ResponseWriter, r *http.Request) {
	// 0. parse parameter from request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var user map[string]string
	err = json.Unmarshal(body, &user)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	// 1. check whether phone is vaild?
	checkUrl := fmt.Sprintf("http://%s:%d%s", config.Gconfig.Server, config.Gconfig.Port, req.TEACHING_USER_BASE_PHONE_CHECK)
	res, ex := http.Post(checkUrl, "application/json;charset=utf-8", strings.NewReader("{\"phone\":\""+user["phone"]+"\"}"))
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}
	defer res.Body.Close()
	body, ex = ioutil.ReadAll(res.Body)
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}

	if res.StatusCode != http.StatusOK {
		tlog.Error(res.Body)
		w.WriteHeader(res.StatusCode)
		io.WriteString(w, string(body))
		return
	}
	var vaild map[string]int

	ex = json.Unmarshal(body, &vaild)
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}

	if vaild["vaild"] != 1 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, ret.USER_IS_EXIST)
		return
	}

	// 2. register phone and pwd
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
	// 3. if success, then auto login
	url := fmt.Sprintf("http://%s:%d%s", config.Gconfig.Server, config.Gconfig.Port, req.TEACHING_USER_BASE_LOGIN)
	res, ex = http.Post(url, "application/json;charset=utf-8", strings.NewReader("{\"phone\":\""+user["phone"]+"\", \"pwd\":\""+user["pwd"]+"\"}"))
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}
	defer res.Body.Close()
	body, ex = ioutil.ReadAll(res.Body)
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}

	if res.StatusCode != http.StatusOK {
		tlog.Error(res.Body)
		w.WriteHeader(res.StatusCode)
		io.WriteString(w, string(body))
		return
	}
	w.WriteHeader(res.StatusCode)
	io.WriteString(w, string(body))
}

/**
login  by phone and pwd
*/
func UserBaseLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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
	delete(result, "_id")
	resp, ex := json.Marshal(result)
	if ex != nil {
		tlog.Error(ex)
		ret.HanderError(w, ex)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(resp))
}

/*
whether phone is registed?

*/
func UserBasePhoneCheck(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var u map[string]string
	err = json.Unmarshal(body, &u)
	if err != nil {
		tlog.Error(err)
		ret.HanderError(w, err)
		return
	}
	var result bson.M
	mongo.GetCollection(mongo.USER_COLL).Find(bson.M{"phone": u["phone"]}).One(&result)
	var vaild = 1
	if len(result) > 0 {
		vaild = 0
	}
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "{\"vaild\" : "+strconv.Itoa(vaild)+"}")
}
