package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"transaction/db"
)

func AddRequest(ctx *gin.Context) {
	from_user := ctx.Param("nickname")
	to_user := ctx.PostForm("to_user")
	amount, _ := strconv.ParseFloat(ctx.PostForm("amount"), 32)
	if amount < 0 {
		ctx.Redirect(http.StatusTemporaryRedirect, "/user_state/"+from_user+"/lk")
		return
	}
	request := new(db.Transaction)
	request.Amount = amount
	request.From_user = from_user
	request.To_user = to_user
	db.DB.Model(request).Insert()
	fmt.Println("Hello")

	db.ProcessRequest(db.DB)

	ctx.Redirect(http.StatusTemporaryRedirect, "/user_state/"+from_user+"/lk")
	return
}
