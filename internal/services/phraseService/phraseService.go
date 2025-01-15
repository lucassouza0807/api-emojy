package phraseService

import (
	connection "api-emoji/config"
	"api-emoji/internal/models"
	"fmt"
	"log"
)

func DestroyPhrase(phraseId string) error {
	db, err := connection.GetDatabaseConnection()

	if err != nil {
		return fmt.Errorf("erro ao excluir frase")
	}

	result := db.Delete(&models.Phrase{}, phraseId)

	if err := result.Error; err != nil {
		return fmt.Errorf(err.Error())
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("frase não encontrada")
	}

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	return nil
}

func EditPhrase(phraseId string, phrase *models.Phrase) error {
	db, err := connection.GetDatabaseConnection()

	if err != nil {
		return fmt.Errorf("erro ao editar frase")
	}

	result := db.Model(&models.Phrase{}).Where("id", phraseId).Updates(&models.Phrase{
		OriginalPhrase:  phrase.OriginalPhrase,
		EmojifiedPhrase: phrase.EmojifiedPhrase,
	})

	if result.RowsAffected == 0 {
		return fmt.Errorf("frase não encontrada")
	}

	if err := result.Error; err != nil {
		return fmt.Errorf("erro: %v", err.Error())
	}

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	return nil

}

func GetPhrases(userId int, page int) ([]models.Phrase, int, error) {
	var pageSize int = 5
	var userPhrases []models.Phrase
	var total int64

	offset := (page - 1) * pageSize

	db, err := connection.GetDatabaseConnection()
	if err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao obter conexão com o banco de dados: %v", err)
	}

	if err := db.Model(&models.Phrase{}).
		Where("user_id = ?", userId).
		Count(&total).Error; err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao contar frases: %v", err)
	}

	if err := db.Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Where("user_id = ?", userId).
		Find(&userPhrases).Error; err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao buscar frases: %v", err)
	}

	lastPage := int((total + int64(pageSize) - 1) / int64(pageSize))

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	return userPhrases, lastPage, nil
}

func SearchForPhrases(userId int, page int, search string) ([]models.Phrase, int, error) {
	var pageSize int = 5
	var userPhrases []models.Phrase
	var total int64

	offset := (page - 1) * pageSize

	db, err := connection.GetDatabaseConnection()
	if err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao obter conexão com o banco de dados: %v", err)
	}

	query := db.Model(&models.Phrase{}).Where("user_id = ?", userId)

	if search != "" {
		query = query.Where("original_phrase LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao contar frases: %v", err)
	}

	if err := query.Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&userPhrases).Error; err != nil {
		return userPhrases, 0, fmt.Errorf("erro ao buscar frases: %v", err)
	}

	lastPage := int((total + int64(pageSize) - 1) / int64(pageSize))

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	return userPhrases, lastPage, nil
}

func StorePhrase(emojifiedPhrase *models.Phrase) error {
	db, err := connection.GetDatabaseConnection()

	if err != nil {
		return fmt.Errorf("erro ao cadastrar frase")
	}

	phrases := models.Phrase{
		UserID:          emojifiedPhrase.UserID,
		OriginalPhrase:  emojifiedPhrase.OriginalPhrase,
		EmojifiedPhrase: emojifiedPhrase.EmojifiedPhrase,
	}

	result := db.Create(&phrases)

	if result.Error != nil {
		log.Printf("Erro ao cadastrar frase: %v", result.Error)
		return fmt.Errorf("erro ao cadastar frase")
	}

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	return nil

}
