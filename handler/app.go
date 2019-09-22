package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"net/http"
)

var e = createMux()

func main() {

	// Echoインスタンス作成
	e := echo.New()
	http.Handle("/", e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	//e.Use(UseAppEngine)


	// ルーティング
	e.GET("/metal", FetchMetal())

	// サーバー起動
	e.Start(":1323")
	appengine.Main()
}


func createMux() *echo.Echo {
	e := echo.New()

	return e
}