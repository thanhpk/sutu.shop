package util

import "encoding/json"
import "net/http"
import "bytes"
import model "github.com/thanhpk/sutu.shop/ecom/model"

type FacebookGraphApi struct {
}

type FbError struct {
	Message string
	Type string
	Code string
}

type modelFbAppInfo model.FbAppInfo
type FbAppInfo struct {
	modelFbAppInfo
	Error FbError
}

type modelFbUserInfo model.FbUserInfo
type FbUserInfo struct {
	modelFbUserInfo
	Error FbError
}

func (fb FacebookGraphApi) GetApp(accesstoken string) *model.FbAppInfo {
	resp, err := http.Get("https://graph.facebook.com/app?access_token=" + accesstoken)
	if err != nil {
		panic (err)
	}
	defer resp.Body.Close()
	
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	pAppInfo := &FbAppInfo{}
  err = json.Unmarshal(buf.Bytes(), &pAppInfo)
	if err != nil {
		panic(err)
	}

	if pAppInfo.Error.Code != "" {
		return nil
	}

	modelFbAppInfo := model.FbAppInfo(pAppInfo.modelFbAppInfo)
	return &modelFbAppInfo
}

func (fb FacebookGraphApi) GetUser(accesstoken string) *model.FbUserInfo {
		resp, err := http.Get("https://graph.facebook.com/app?access_token=" + accesstoken)
	if err != nil {
		panic (err)
	}
	defer resp.Body.Close()
	
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	pUserInfo := &FbUserInfo{}
  err = json.Unmarshal(buf.Bytes(), pUserInfo)
	if err != nil {
		panic(err)
	}
	
	if pUserInfo.Error.Code != "" {
		return nil
	}

	
	modelFbUserInfo := model.FbUserInfo(pUserInfo.modelFbUserInfo)
	return &modelFbUserInfo

}
