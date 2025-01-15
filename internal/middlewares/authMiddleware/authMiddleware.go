package authMiddleware

import (
	"api-emoji/internal/services/jwtService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken(ctx *gin.Context) {

	auth_token := ctx.GetHeader("Authorization")

	if auth_token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não authorizado, nenhum token passado",
		})
	}

	token := auth_token[len("Bearer "):]

	err := jwtService.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não autorizado",
		})

		ctx.Abort()
	}

	ctx.Next()
}
