# RESTful API com Autenticação JWT em Go

Este projeto demonstra um exemplo básico de um serviço RESTful implementado em Go, utilizando autenticação JWT (JSON Web Tokens). O serviço possui um endpoint protegido, que requer autenticação para ser acessado.

## Funcionalidades

- **Login:** Permite que usuários façam login e recebam um token JWT.
- **Endpoint Protegido:** O endpoint `/home` só pode ser acessado por usuários que possuam um token JWT válido.

## Requisitos

- **Go:** Versão 1.16 ou superior

## Instalação e Execução

1. **Clone o Repositório:**

    ```bash
    git clone https://github.com/seuusuario/rest-jwt-go.git
    cd rest-jwt-go
    ```

2. **Instale as Dependências:**

    ```bash
    go mod tidy
    ```

3. **Execute o Serviço:**

    Inicie o servidor com o comando:

    ```bash
    go run main.go
    ```

    O servidor estará disponível em [http://localhost:8000](http://localhost:8000).

## Como Usar

### 1. Fazer Login e Obter o Token JWT

O primeiro passo é fazer login enviando um POST para o endpoint `/login` com as credenciais de usuário.

**Requisição:**

```bash
curl -X POST -d '{"username":"user","password":"password"}' \
-H "Content-Type: application/json" \
http://localhost:8000/login
