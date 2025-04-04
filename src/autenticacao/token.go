package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"usuarioID":  usuarioID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	chaveSecreta := []byte(config.SecretKey)

	return token.SignedString(chaveSecreta)
}

// Validartoken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

    if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return nil
    }

    return errors.New("Token inválido!")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
