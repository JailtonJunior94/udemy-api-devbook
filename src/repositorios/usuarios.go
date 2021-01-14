package repositorios

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JailtonJunior94/udemy-api-devbook/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	criadoEm := time.Now()
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO Usuarios ([Name], [Nick], [Email], [Senha], [CriadoEm]) VALUES (@nome, @nick, @email, @senha, @criadoEm); SELECT SCOPE_IDENTITY()")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	lastInsertId := 0
	erro = statement.QueryRow(sql.Named("nome", usuario.Nome),
		sql.Named("nick", usuario.Nick),
		sql.Named("email", usuario.Email),
		sql.Named("senha", usuario.Senha),
		sql.Named("criadoEm", criadoEm)).Scan(&lastInsertId)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT [Id], [Name], [Nick], [Email], [CriadoEm] FROM Usuarios WHERE [Name] LIKE @nomeOuNick OR [Nick] LIKE @nomeOuNick",
		sql.Named("nomeOuNick", nomeOuNick))

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
