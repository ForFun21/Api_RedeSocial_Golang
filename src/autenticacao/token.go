package autenticacao

import (
    "time"
    "github.com/golang-jwt/jwt/v5" // Use a versão mais nova e mantida
)

func CriarToken(usuarioID uint64) (string, error) {
    permissoes := jwt.MapClaims{
        "authorized": true,
        "exp":        time.Now().Add(time.Hour * 6).Unix(),
        "usuarioID":  usuarioID,
    }

    // Altere para HS256 (HMAC com SHA-256)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

    // Use uma chave secreta forte (em produção, armazene em variável de ambiente!)
    chaveSecreta := []byte("SuaChaveSecretaMuitoLongaESeguraAqui123!@#")

    return token.SignedString(chaveSecreta)
}