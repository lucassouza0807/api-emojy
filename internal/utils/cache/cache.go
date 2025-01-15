package cache

import (
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

// Simulação de uma estrutura de dados de usuário
type User struct {
	UserID int
	Name   string
	Email  string
}

func main() {
	// Conectar ao Memcache
	mc := memcache.New("localhost:11211")
	if mc == nil {
		log.Fatal("Memcache não está disponível")
	}

	// Simulação de dados de usuário
	user := User{
		UserID: 1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
	}

	// Armazenar dados de usuário no Memcache
	err := mc.Set(&memcache.Item{
		Key:   fmt.Sprintf("user_%d", user.UserID),                 // Usando o user_id como chave
		Value: []byte(fmt.Sprintf("%s,%s", user.Name, user.Email)), // Armazenando os dados como uma string simples
	})
	if err != nil {
		log.Fatal("Erro ao definir item no cache:", err)
	}

	// Pesquisa de dados via user_id
	key := fmt.Sprintf("user_%d", user.UserID)
	item, err := mc.Get(key)
	if err != nil {
		log.Fatal("Erro ao obter item do cache:", err)
	}

	// Exibindo os dados armazenados no Memcache
	data := string(item.Value)
	fmt.Printf("Dados recuperados do Memcache: %s\n", data)
}
