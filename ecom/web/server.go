package web

import (
	//"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
	"fmt"
	"github.com/thanhpk/sutu.shop/ecom/usecase"
)

type Usecases struct {
	BrowseProduct usecase.IBrowseProduct
	Login usecase.ILogin
	Registry usecase.IRegistry
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
//	app.Use(recovery.New())
	app.Get("/login/authbyfacebook/:accesstoken", func(ctx *iris.Context) {
		accesstoken := ctx.Param("accesstoken")
		customer := ucs.Login.AuthByFacebook(accesstoken)
		ctx.JSON(iris.StatusOK, customer)
	}, func(ctx *iris.Context) {
		fmt.Println("Hello, refresh one time more to get panic!")
		ctx.EmitError(iris.StatusInternalServerError)
	});

	app.Post("/login/authbyphone", func(ctx *iris.Context) {
		phone := ctx.FormValueString("phone")
		password := ctx.FormValueString("password")
		customer := ucs.Login.AuthByPhone(phone, password)
		ctx.JSON(iris.StatusOK, customer)
	}, func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	});

	app.Post("/registry/isDuplicated", func(ctx *iris.Context) {
		phone := ctx.FormValueString("phone")
		isDuplicate := ucs.Registry.IsDuplicated(phone)
		ctx.JSON(iris.StatusOK, isDuplicate)
	}, func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	});

	app.Post("/registry/registry", func(ctx *iris.Context) {
		phone := ctx.FormValueString("phone")
		password := ctx.FormValueString("password")

		newCustomerID := ucs.Registry.Registry(phone, password)
		ctx.JSON(iris.StatusOK, newCustomerID)
	}, func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	})
}
