package main

import (
	"bytes"
	"testing"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	TestHttp(nil)
}

func marshalBson_test(t *testing.T) {
	data, err := bson.Marshal(&Person{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", data)


}

func TestHttp(t *testing.T) {
	resp, err := http.Get("https://graph.facebook.com/app?access_token=EAAR77iecZCjUBAJ8khGWPxn2ZCuczvNU4lPxqhlSZASxjUptwTOZCGyRqgxGVlyyROCd9yxHwWuHDXEZCRgWGb9SZBkgK1988cqSbcEPV137ZAj6MVNZAhZCy3HsZC7yiY5SfNBbjsFbnlKrwoPNHGUgnQJ9Y4DhWIn94ZD")
	if err != nil {
		panic (err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)


	appInfo := model.FbAppInfo{}
  json.Unmarshal(buf.Bytes(), &appInfo)
	
	fmt.Println(appInfo.Name)
}
