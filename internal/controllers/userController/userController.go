package userController

import (
	connection "api-emoji/config"
	"api-emoji/internal/models"
	"api-emoji/internal/services/jwtService"
	useHash "api-emoji/internal/utils/hash"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserType struct {
	User_id  *int    `json:"user_id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func CreateUser(ctx *gin.Context) {
	db, err := connection.GetDatabaseConnection()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	var requestBody UserType

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Nunhum campo foi passo no corpo da requisição",
		})

		return
	}

	hashedPassowrd, _ := useHash.HashPassword(*requestBody.Password)

	user := models.User{
		Name:     *requestBody.Name,
		Email:    *requestBody.Email,
		Password: hashedPassowrd,
	}

	result := db.Create(&user)

	if result.Error != nil {
		if result.Error.Error() == fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'users.uni_users_email'", *requestBody.Email) {
			ctx.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": fmt.Sprintf("O email '%s' já está em uso.", *requestBody.Email),
			})

			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Houve um erro ao inserir no banco de dados tente mais tarde.",
		})

		return
	}

	token, err := jwtService.CreateToken(&jwtService.UserToken{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	})

	//Case the cant be created
	if err != nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Usúario criado com sucesso",
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token":   token,
		"exp":     int64((time.Minute * 60 * 48).Seconds()),
		"success": true,
		"message": "Usúario criado com sucesso",
	})

}
