package web

import (
	"github.com/kataras/iris"
	//"fmt"
	"github.com/thanhpk/sutu.shop/ecom/usecase"
)

type Usecases struct {
	BrowseProduct usecase.IBrowseProduct
	Login usecase.ILogin
	Registry usecase.IRe
}

type Web struct {
}

func (w *Web) Run(port string, ucs Usecases) {
	server := iris.New()
	server.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Write("Customer 404 not found error page")
	})

	server.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write("")
	})

	route(server, ucs)

	server.Listen(":" + port)
}

func route(app *iris.Framework, ucs Usecases) {
	iris.Post("/login/authbyfacebook/:accesstoken", func(ctx *iris.Context) {
		customer := ucs.Login.AuthByFacebook(ctx.Param("accesstoken"))
		ctx.JSON(iris.StatusOK, customer)
	}, func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	});

	iris.Post("/login/authbyphone", func(ctx *iris.Context) {
		phone := ctx.FormValueString("phone")
		password := ctx.FormValueString("password")
		customer := ucs.Login.AuthByPhone(phone, password)
		ctx.JSON(iris.StatusOK, customer)
	}, func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	});
}
