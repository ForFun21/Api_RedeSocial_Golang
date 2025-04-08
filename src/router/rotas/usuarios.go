package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAltenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BurscarUsuarios,
		RequerAltenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAltenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAltenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAltenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAltenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAltenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguidores",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAltenticacao: true,
	},
}
