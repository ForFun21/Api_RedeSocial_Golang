package repository

import (
	"api/src/models"
	"database/sql"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (u Usuarios) Criar(usuaio models.Usuario) (uint64, error) {
	return 0, nil
}