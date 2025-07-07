# API Rede Social em Go

Esta documentação descreve o funcionamento da API de Rede Social desenvolvida em Go, incluindo como configurá-la e utilizá-la localmente.

---

## 1. Visão Geral e Pré‑requisitos

Esta é uma API de rede social escrita em Go (>= 1.23). Ela oferece endpoints para:

* **Usuários**: cadastro, busca, atualização, exclusão, seguir/parar de seguir, lista de seguidores e seguindo, atualização de senha.
* **Autenticação**: login via JWT.
* **Publicações**: criar, listar, buscar, editar, excluir, curtir/descurtir e listar por usuário.

**Pré‑requisitos**

* Go 1.23+
* MySQL
* `git` ([https://git-scm.com/](https://git-scm.com/))
* Postman ou similar para testar endpoints

---

## 2. Como rodar localmente

1. **Clone o repositório**

   ```bash
   git clone https://github.com/ForFun21/Api_RedeSocial_Golang.git
   cd Api_RedeSocial_Golang
   ```

2. **Instale as dependências**

   ```bash
   go mod download
   ```

3. **Configure as variáveis de ambiente**

   Copie o arquivo `.env.example` como `.env` e ajuste conforme explicado na seção 4.

4. **Crie o banco de dados & tabelas**

   * No MySQL, crie um banco com o nome definido em `DB_BANCO`.
   * Execute os scripts SQL em `sql/` (por exemplo, `tabelas.sql`) para criar as tabelas necessárias.

5. **Execute a aplicação**

   ```bash
   go run main.go
   ```

   Por padrão, ela irá “escutar” na porta definida em `API_PORT` (ex.: `5000`).
   Você verá no console:

   ```
   Escutando na porta 5000
   ```

---

## 3. Configuração do Banco de Dados

No diretório `sql/` estão os scripts para criar as tabelas. Exemplo de tabelas:

```sql
CREATE TABLE usuarios (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  senha VARCHAR(255) NOT NULL,
  criado_em DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE seguidores (
  usuario_id INT,
  seguidor_id INT,
  criado_em DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (usuario_id, seguidor_id),
  FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
  FOREIGN KEY (seguidor_id) REFERENCES usuarios(id)
);

CREATE TABLE publicacoes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  usuario_id INT,
  conteudo TEXT NOT NULL,
  criado_em DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

CREATE TABLE curtidas (
  publicacao_id INT,
  usuario_id INT,
  criado_em DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (publicacao_id, usuario_id),
  FOREIGN KEY (publicacao_id) REFERENCES publicacoes(id),
  FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
```

---

## 4. Variáveis de Ambiente

No arquivo `.env`, configure:

```dotenv
DB_USUARIO=Seu usuário do mysql
DB_SENHA=Senha do seu usuário do mysql
DB_BANCO=Nome do banco 
API_PORT=5000
SECRET_KEY=<sua_chave_secreta_para_JWT>
```

* **DB\_USUARIO**, **DB\_SENHA**, **DB\_BANCO**: credenciais do MySQL.
* **API\_PORT**: porta em que o servidor HTTP irá rodar.
* **SECRET\_KEY**: chave usada para assinar tokens JWT.

---

## 5. Estrutura de Pastas

```
Api_RedeSocial_Golang/
├── main.go             # inicialização do app (carrega config e router)
├── go.mod, go.sum      # dependências
├── .env                # variáveis de ambiente
├── sql/                # scripts SQL para criação de tabelas
└── src/
    ├── config/         # carregamento de env e conexão
    ├── router/
    │   ├── router.go   # gera *mux.Router
    │   └── rotas/      # definição de todas as rotas
    └── controllers/    # lógica de cada endpoint
```

---

## 6. Rotas e Exemplos de Uso

### 6.1 Usuários

```http
POST   /usuarios                             # Criar usuário (sem token)
GET    /usuarios                             # Listar usuários (token)
GET    /usuarios/{usuarioId}                 # Buscar usuário por ID (token)
PUT    /usuarios/{usuarioId}                 # Atualizar usuário (token)
DELETE /usuarios/{usuarioId}                 # Excluir usuário (token)
POST   /usuarios/{usuarioId}/seguir          # Seguir usuário (token)
POST   /usuarios/{usuarioId}/parar-de-seguir # Parar de seguir (token)
GET    /usuarios/{usuarioId}/seguidores      # Listar seguidores (token)
GET    /usuarios/{usuarioId}/seguindo       # Listar seguindo (token)
POST   /usuarios/{usuarioId}/atualizar-senha # Atualizar senha (token)
```

### 6.2 Autenticação

```http
POST /login   # Recebe JSON { email, senha } e retorna { token } sem precisar de token prévio
```

### 6.3 Publicações

```http
POST   /publicacoes                          # Criar publicação (token)
GET    /publicacoes                          # Listar todas as publicações (token)
GET    /publicacoes/{publicacaoId}           # Buscar publicação por ID (token)
PUT    /publicacoes/{publicacaoId}           # Atualizar publicação (token)
DELETE /publicacoes/{publicacaoId}           # Excluir publicação (token)
GET    /usuarios/{usuarioId}/publicacoes     # Listar publicações de um usuário (token)
POST   /publicacoes/{publicacaoId}/curtir    # Curtir publicação (token)
POST   /publicacoes/{publicacaoId}/descurtir # Descurtir publicação (token)
```

---

## Exemplos de Requisição

### Criar Usuário

```bash
curl -X POST http://localhost:5000/usuarios \
  -H "Content-Type: application/json" \
  -d '{"nome":"Jhon","email":"jhon@ex.com","senha":"minhaSenha"}'
```

#### Exemplo no Postman

1. Abra o Postman e crie uma nova requisição.
2. Selecione **POST** e insira a URL: `http://localhost:5000/usuarios`.
3. Em **Headers**, adicione:
   - `Content-Type: application/json`
4. Em **Body**, selecione **raw** e **JSON**, e insira:
   ```json
   {
     "nome": "Jhon",
     "email": "jhon@ex.com",
     "senha": "minhaSenha"
   }
   ```
5. Clique em **Send** para enviar a requisição.

---

### Login

```bash
curl -X POST http://localhost:5000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"jhon@ex.com","senha":"minhaSenha"}'
# → { "token": "eyJhbGciOi..." }
```

#### Exemplo no Postman

1. Crie uma nova requisição no Postman.
2. Selecione **POST** e informe a URL: `http://localhost:5000/login`.
3. Em **Headers**, adicione:
   - `Content-Type: application/json`
4. Em **Body**, raw JSON:
   ```json
   {
     "email": "jhon@ex.com",
     "senha": "minhaSenha"
   }
   ```
5. Clique em **Send** e verifique o corpo da resposta com o token JWT.

---

### Listar Publicações (com token)

```bash
curl http://localhost:5000/publicacoes \
  -H "Authorization: Bearer <seu_token_aqui>"
```

#### Exemplo no Postman

1. No Postman, crie uma requisição **GET** com a URL `http://localhost:5000/publicacoes`.
2. Em **Headers**, adicione:
   - `Authorization: Bearer <seu_token_aqui>`
3. Clique em **Send** para obter a lista de publicações.
