package logincontroller

import (
	connection "api-emoji/config"
	models "api-emoji/internal/models"
	jwtService "api-emoji/internal/services/jwtService"
	useHash "api-emoji/internal/utils/hash"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func Login(ctx *gin.Context) {
	var user models.User //Model to scan query results

	db, err := connection.GetDatabaseConnection()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Nunhum campo foi passo no corpo da requisição",
		})

		return
	}

	result := db.Where("email", credentials.Email).First(&user)

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "O e-mail fornecido não existe.",
		})

		return
	}

	passwordMatches := useHash.CheckPasswordHash(*credentials.Password, user.Password)

	if !passwordMatches {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credenciais inválidas",
		})

		return
	}

	token, err := jwtService.CreateToken(&jwtService.UserToken{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Houve um erro tentar fazer login",
		})

		return
	}

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"exp":   int64((time.Minute * 60 * 48).Seconds()),
		"user":  user,
	})

}
