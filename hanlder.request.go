package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addRequest(ctx *gin.Context) {
	from_user := ctx.Param("nickname")
	to_user := ctx.PostForm("to_user")
	amount, _ := strconv.ParseFloat(ctx.PostForm("amount"), 64)
	request := new(transaction)
	request.Amount = amount
	request.From_user = from_user
	request.To_user = to_user
	db.Model(request).Insert()
	fmt.Println("Hello")

	processRequest(db)

	ctx.Redirect(http.StatusTemporaryRedirect, "/user_state/Ilya/lk")
	return
}
