package main

import (
	mgo "gopkg.in/mgo.v2"
	"github.com/creamdog/gonfig"
	"os"
	"fmt"
	"time"
	"github.com/thanhpk/sutu.shop/ecom/web"
	"github.com/thanhpk/sutu.shop/ecom/usecase"
	"github.com/thanhpk/sutu.shop/ecom/model"
	"github.com/thanhpk/sutu.shop/ecom/db"
	"github.com/thanhpk/sutu.shop/ecom/util"
)


func ConnectDb(hostname string, username string, password string, port string) *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{hostname + ":" + port},
		Timeout:  20 * time.Second,
//		Database: AuthDatabase,
//		Username: username,
//		Password: password,
	}
	
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic (err)
	}

	return session
}

func main() {
	f, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	config, err := gonfig.FromJson(f)

	if err != nil {
		panic(err)
	}

	port, _ := config.GetString("endpoint/port", "8081")
	fmt.Println(port)
	
	fbappid, err := config.GetString("facebook/appid", "1262162703875637")
	mgodbname, err := config.GetString("database/database", "sutu.shop")
	mgousername, err := config.GetString("database/username", "")
	mgoport, err := config.GetString("database/port", "27017")
	mgohost, err := config.GetString("database/hostname", "127.0.0.1")
	mgopassword, err := config.GetString("database/password", "")
	
	mgosession := ConnectDb(mgohost, mgousername, mgopassword, mgoport)

	customerMgt :=  MakeCustomerMgt(mgodbname, mgosession, fbappid)
	usecases := web.Usecases{
		Login: usecase.Login{
			CustomerMgt: customerMgt,
		},
		Registry: usecase.Registry{
			CustomerMgt: customerMgt,
		},
	}
	web := web.Web{}
	web.Run(port, usecases)	
}

func MakeCustomerMgt(dbname string, session *mgo.Session, fbappid string) *model.CustomerMgt {
	customerMgt := model.CustomerMgt{}
	customerMgt.FbAppId = fbappid
	customerMgt.Repo = db.NewMongoCustomerRepository(dbname, session, "customer")
	customerMgt.Fb = util.FacebookGraphApi{}
	return &customerMgt
}

