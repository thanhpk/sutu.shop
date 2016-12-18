package main

import (
	"github.com/creamdog/gonfig"
	"os"
	"fmt"
	"github.com/thanhpk/sutu.shop/ecom/web"
	"github.com/thanhpk/sutu.shop/ecom/usecase"
)

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

	usecases := web.Usecases{
		Login: usecase.Login{},
	}
	web := web.Web{}
	web.Run(port, usecases)	
}
