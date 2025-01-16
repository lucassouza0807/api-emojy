# API EMOJI

Este repositório contém uma API desenvolvida em Go para o desafio tecnico. Abaixo estão as instruções para instalar as dependências, configurar o ambiente e rodar a API.

## Requisitos

- Go 1.18+ (ou superior)
- Git
- Dependências que serão instaladas automaticamente via Go Modules

## Instalação das Dependências

1. Clone o repositório para sua máquina local:

    ```bash
    git clone https://github.com/lucassouza0807/api-emojy.git
    cd api-emojy
    ```

2. Instale as dependências necessárias utilizando Go Modules:

    ```bash
    go mod tidy
    ```

Isso irá baixar todas as dependências mencionadas no `go.mod`.

## Como Rodar a API

1. Para rodar a API localmente, execute o seguinte comando:

    ```bash
    go run main.go
    ```

    Isso irá iniciar o servidor na porta padrão `8080`. Você pode acessar a API através de `http://localhost:8080`.

2. Caso queira compilar a API para uma versão executável, rode o seguinte comando:

    ```bash
    go build -o api
    ```

    Isso irá gerar o arquivo executável `api`. Para rodá-lo, basta executar:

    ```bash
    ./api
    ```


# API Endpoints

### 1. **Create User**
   - **Método**: POST
   - **URL**: `http://localhost:8080/api/v1/create-user`
   - **Descrição**: Cria um novo usuário.
   - **Body**:
     ```json
     {
       "name": "Lucas",
       "email": "lucassouza0807@gmail.com",
       "password": "@md11nice"
     }
     ```

### 2. **Login**
   - **Método**: POST
   - **URL**: `http://localhost:8080/api/v1/login`
   - **Descrição**: Faz a autenticação.
   - **Body**:
     ```json
     {
       "email": "lucassouza0807@gmail.com",
       "password": "@md11nice"
     }
     ```

### 3. **Get Phrases**
   - **Método**: GET
   - **URL**: `http://localhost:8080/api/v1/phrases`
   - **Descrição**: Obtém as frases do usuário via token.
   - **Autorização**: Bearer Token
   - **Exemplo de Header Authorization**:
     ```
     Authorization: Bearer <your_token_here>
     ```

### 4. **Edit Phrase**
   - **Método**: PUT
   - **URL**: `http://localhost:8080/api/v1/edit-phrase/{id}`
   - **Descrição**: Edita uma frase existente.
   - **Parâmetros**:
     - `{id}`: ID da frase a ser editada.
   - **Body**:
     ```json
     {
       "original_phrase": "gato dasd",
       "emojified_phrase": "gato teste"
     }
     ```

### 5. **Store Phrase**
   - **Método**: POST
   - **URL**: `http://localhost:8080/api/v1/store-phrase`
   - **Descrição**: Cria uma nova frase na base de dados.
   - **Body**:
     ```json
     {
       "original_phrase": "sdas",
       "emojified_phrase": "adasdsa"
     }
     ```

### 6. **Delete Phrase**
   - **Método**: DELETE
   - **URL**: `http://localhost:8080/api/v1/delete-phrase/{id}`
   - **Descrição**: Deleta uma frase no banco de dados.
   - **Parâmetros**:
     - `{id}`: ID da frase a ser deletada.

### 7. **Search Phrase**
   - **Método**: GET
   - **URL**: `http://localhost:8080/api/v1/search-phrase?query={query}`
   - **Descrição**: Procura uma frase baseada no texto original.
   - **Parâmetros**:
     - `query`: Texto para buscar.
   - **Exemplo de URL**: `http://localhost:8080/api/v1/search-phrase?query=frase`



