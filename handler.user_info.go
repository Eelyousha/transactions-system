// handlers.user_info.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func getCurrentState(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	connectDB(":5432", "postgres", "password", "postgres")

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
