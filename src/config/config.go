package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver PostgreSQL
)

var (
	// StringConexaoBanco é a string de conexão com o Postgres
	StringConexaoBanco string

	// Porta onde a API vai estar rodando
	Porta int

	// SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	// Carrega .env apenas em ambiente de desenvolvimento
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Nenhum arquivo .env carregado (GO_ENV != production)")
		}
	}

	// Determina porta HTTP (PaaS ou variável local)
	if p := os.Getenv("PORT"); p != "" {
		if porta, err := strconv.Atoi(p); err == nil {
			Porta = porta
		}
	} else if p2, err := strconv.Atoi(os.Getenv("API_PORT")); err == nil {
		Porta = p2
	} else {
		Porta = 9000
	}

	// Coleta variáveis do Postgres
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgUser := os.Getenv("DB_USUARIO")
	pgPass := os.Getenv("DB_SENHA")
	pgDB := os.Getenv("DB_BANCO")

	// Monta string de conexão PostgreSQL
	// Substitua todo o bloco if/else por:
	StringConexaoBanco = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPass, pgDB,
	)

	// Define chave secreta
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
