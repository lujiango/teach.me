package qiniu

import (
	"github.com/qiniu/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
	"teach.me/teaching/config"
	"teach.me/teaching/tlog"
)

type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func Upload(filePath string) (string, error) {
	tlog.Debug("filePath : ", filePath)
	//init AK，SK
	conf.ACCESS_KEY = config.Gconfig.Ak
	conf.SECRET_KEY = config.Gconfig.Sk

	//create Client
	c := kodo.New(0, nil)

	//set upload policy
	policy := &kodo.PutPolicy{
		Scope: config.Gconfig.Bucket,
		//set Token's expires
		Expires: 3600,
	}
	//make a token
	token := c.MakeUptoken(policy)
	tlog.Debug("token : ", token)
	//make uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	//call PutFileWithoutKey to upload,named fileName by hash
	err := uploader.PutFileWithoutKey(nil, &ret, token, filePath, nil)
	tlog.Debug(ret)
	if err != nil {
		tlog.Error("io.Put failed : ", err)
		return "", err
	}
	return config.Gconfig.QiNiuDomain + ret.Key, nil
}

func DownloadPrivateUrl(url string) string {
	conf.ACCESS_KEY = config.Gconfig.Ak
	conf.SECRET_KEY = config.Gconfig.Sk
	policy := kodo.GetPolicy{}
	c := kodo.New(0, nil)
	return c.MakePrivateUrl(url, &policy)
}
