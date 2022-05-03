// handlers.user_info.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCurrentState(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	if userQuery, err := getUserFromDB(db, nickname); err == nil {
		ctx.HTML(
			http.StatusOK,
			"user_info.html",
			gin.H{
				"payload":      userQuery,
				"transactions": userQuery.transactions,
			},
		)
	} else {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
}

func getUsers(ctx *gin.Context) {
	var users []struct {
		Nickname string
		Balance  float64
	}
	db.Model((*user)(nil)).Column("nickname", "balance").Select(&users)
	ctx.HTML(http.StatusOK,
		"index.html",
		gin.H{"users": users})
}
