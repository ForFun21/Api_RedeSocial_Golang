package autenticacao

import (
	"api/src/config"
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