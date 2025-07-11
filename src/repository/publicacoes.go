package repository

import (
	"api/src/models"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	var id uint64
	erro := repositorio.db.QueryRow(
		`INSERT INTO publicacoes (titulo, conteudo, autor_id)
         VALUES ($1, $2, $3)
         RETURNING id`,
		publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID,
	).Scan(&id)
	if erro != nil {
		return 0, erro
	}
	return id, nil
}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`SELECT p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criado_em AS criadaEm, u.nick
        FROM publicacoes p
        INNER JOIN usuarios u ON u.id = p.autor_id
        WHERE p.id = $1`,
		publicacaoID,
	)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
       SELECT DISTINCT 
           p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criado_em, u.nick
       FROM publicacoes p
       JOIN usuarios u 
         ON u.id = p.autor_id
       LEFT JOIN seguidores s
         ON s.usuario_id = p.autor_id 
         AND s.seguidor_id = $1
       WHERE p.autor_id = $1
          OR s.seguidor_id = $1
       ORDER BY p.id DESC`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE publicacoes
        SET titulo = $1, conteudo = $2
        WHERE id = $3`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`DELETE FROM publicacoes
        WHERE id = $1`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criado_em AS criadaEm, u.nick
        FROM publicacoes p
        JOIN usuarios u ON u.id = p.autor_id
        WHERE p.autor_id = $1`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida a uma publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE publicacoes
        SET curtidas = curtidas + 1
        WHERE id = $1`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Descurtir subtrai uma curtida a uma publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE publicacoes
        SET curtidas = CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END
        WHERE id = $1`,
	)
	if erro != nil {
		return erro
	}
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
