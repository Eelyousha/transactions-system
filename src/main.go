/*
Нужно реализовать систему транзакций.
Как происходит транзакция:
Идет запрос на сервер от клиента, по клиенту выстраивается очередь на вывод.
Важно :
1) у каждого клиента есть своя очередь;
2) при нехватке денег, нужно блокировать запрос
Что нужно реализовавть :
бд на postgresql, где будет схема с клиентами и их балансами
сервер, которые проверяет все условия(хватает ли денег, если сервер упадет, то история, которая идет на вывод не должна пропасть) и делает изменение баланса(на + или -)

*/

// main.go

package main

import (
	"transaction/db"
	"transaction/web"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB(":5432", "postgres", "password", "postgres")

	err := db.ProcessRequest(db.DB)
	for err == nil {
		err = db.ProcessRequest(db.DB)
	}
	// Set the router as the default one provided by Gin
	web.Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	web.Router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	web.InitializeRoutes()

	// Start serving the application
	web.Router.Run()

}
