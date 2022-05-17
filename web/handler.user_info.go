// handlers.user_info.go

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"transaction/db"
)

func GetCurrentState(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	if userQuery, err := db.GetUserFromDB(db.DB, nickname); err == nil {
		ctx.HTML(
			http.StatusOK,
			"user_info.html",
			gin.H{
				"payload":      userQuery,
				"transactions": userQuery.Transactions,
			},
		)
	} else {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
}

func GetUsers(ctx *gin.Context) {
	var users []struct {
		Nickname string
		Balance  float64
	}
	db.DB.Model((*db.User)(nil)).Column("nickname", "balance").Select(&users)
	ctx.HTML(http.StatusOK,
		"index.html",
		gin.H{"users": users})
}
