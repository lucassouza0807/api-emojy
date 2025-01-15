package phrasecontroller

import (
	typeHelper "api-emoji/internal/helpers/typeHelper"
	"api-emoji/internal/models"
	jwtService "api-emoji/internal/services/jwtService"
	"api-emoji/internal/services/phraseService"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmojifiedPhrase struct {
	OriginalPhrase  string `json:"original_phrase"`
	EmojifiedPhrase string `json:"emojified_phrase"`
}

func EditPhrase(ctx *gin.Context) {
	var requestBody EmojifiedPhrase

	phrase_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Nenhum campo foi passado na requisição",
		})

		return
	}

	err := phraseService.EditPhrase(phrase_id, &models.Phrase{
		OriginalPhrase:  requestBody.OriginalPhrase,
		EmojifiedPhrase: requestBody.EmojifiedPhrase,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Frase modificada com sucesso",
	})

}

func DeletePhrase(ctx *gin.Context) {
	err := phraseService.DestroyPhrase(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Frase excluída com sucesso",
	})

}
func CreatePhrase(ctx *gin.Context) {
	auth_token := ctx.GetHeader("Authorization")
	token := auth_token[len("Bearer "):]

	decodedToken, err := jwtService.DecodeToken(token)

	if err != nil {
		log.Printf("Erro ao decodar token %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao cadastrar frase",
		})

		return
	}

	var emojifiedPhrase EmojifiedPhrase

	if err := ctx.ShouldBindJSON(&emojifiedPhrase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Nenhum campo foi passado na requisição",
		})

		return
	}

	err = phraseService.StorePhrase(&models.Phrase{
		UserID:          uint(decodedToken.Id),
		OriginalPhrase:  emojifiedPhrase.OriginalPhrase,
		EmojifiedPhrase: emojifiedPhrase.EmojifiedPhrase,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Frase criada com sucesso",
	})

}

func SearchForPhrases(ctx *gin.Context) {
	auth_token := ctx.GetHeader("Authorization")
	token := auth_token[len("Bearer "):]

	decodedToken, err := jwtService.DecodeToken(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao obter fraes",
		})

		return
	}

	page := typeHelper.StringToInt(ctx.Query("page"), 1)

	userPhrases, totalPages, err := phraseService.SearchForPhrases(decodedToken.Id, page, ctx.Query("query"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	current_page := ctx.Query("page")

	if len(ctx.Query("page")) == 0 {
		current_page = "1"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":         userPhrases,
		"last_page":    totalPages,
		"current_page": current_page,
	})

}

func GetUserPhrases(ctx *gin.Context) {
	auth_token := ctx.GetHeader("Authorization")
	token := auth_token[len("Bearer "):]

	decodedToken, err := jwtService.DecodeToken(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao obter fraes",
		})

		return
	}

	page := typeHelper.StringToInt(ctx.Query("page"), 1)

	userPhrases, totalPages, err := phraseService.GetPhrases(decodedToken.Id, page)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	current_page := ctx.Query("page")

	if len(ctx.Query("page")) == 0 {
		current_page = "1"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":         userPhrases,
		"last_page":    totalPages,
		"current_page": current_page,
	})

}
